package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/rs/cors"
)

var db *sql.DB

func handlerHello(w http.ResponseWriter, r *http.Request) {
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

	instanceID := aws.StringValue(runResult.Instances[0].InstanceId)

	_, err = db.Exec("INSERT INTO ec2_instances (id) VALUES ($1)", instanceID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(instanceID)
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

	_, err = db.Exec("DELETE FROM ec2_instances WHERE id = $1", data.InstanceId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Instance terminated successfully!"))
}

func main() {
	var err error
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", dbUser, dbPass, dbName, dbHost)
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)}
	mux := http.NewServeMux()

	mux.HandleFunc("/api/hello", handlerHello)
	mux.HandleFunc("/api/ec2/create", handlerCreateEC2Instance)
	mux.HandleFunc("/api/ec2/terminate", handlerTerminateEC2Instance)

	handler := cors.Default().Handler(mux)

	log.Println("Starting server on port 80...")
	err = http.ListenAndServe(":80", handler)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}