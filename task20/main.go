package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ec2"
)

func handlerHello(w http.ResponseWriter, r *http.Request) {
	// Разрешаем запросы с указанных доменов
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Обрабатываем предварительную проверку CORS запросов
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var data struct {
		Name string `json:"name"`
	}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	message := fmt.Sprintf("Hello, %s!", data.Name)
	fmt.Fprintln(w, message)
}
func handlerCreateEC2Instance(w http.ResponseWriter, r *http.Request) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-2"),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	svc := ec2.New(sess)
	runResult, err := svc.RunInstances(&ec2.RunInstancesInput{
		ImageId:      aws.String("ami-0c9c942bd7bf113a2"),
		InstanceType: aws.String("t2.micro"),
		MinCount:     aws.Int64(1),
		MaxCount:     aws.Int64(1),
		SecurityGroupIds: aws.StringSlice([]string{"sg-04882473cf1ebd81d",}),
		KeyName: aws.String("id_rsa"),
		TagSpecifications: []*ec2.TagSpecification{
			{
				ResourceType: aws.String("instance"),
				Tags: []*ec2.Tag{
					{
						Key:   aws.String("Name"),
						Value: aws.String("UbuntuGo"),
					},
				},
			},
		},
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(runResult.Instances[0].InstanceId)
}

func handlerTerminateEC2Instance(w http.ResponseWriter, r *http.Request) {
	var data struct {
		InstanceId string `json:"instanceId"`
	}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-2"),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	svc := ec2.New(sess)
	_, err = svc.TerminateInstances(&ec2.TerminateInstancesInput{
		InstanceIds: []*string{aws.String(data.InstanceId)},
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Instance terminated successfully!"))
}

func main() {
	http.HandleFunc("/api/hello", handlerHello)
	http.HandleFunc("/api/ec2/create", handlerCreateEC2Instance)
	http.HandleFunc("/api/ec2/terminate", handlerTerminateEC2Instance)

	log.Println("Starting server on port 80...")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}