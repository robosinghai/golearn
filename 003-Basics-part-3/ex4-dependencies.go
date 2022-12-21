package main

import (
	"fmt"
	"log"
	"os"

	Utils "golearn/utils" // NOTE: LOCAL PACKAGE IMPORTED

	"github.com/joho/godotenv" // NOTE: ONLINE PACKAGE IMPORTED
)

func main() {

	// APP CONFIG DEMO
	fmt.Printf("APP CONFIG DEMO (.env file)\n")
	err := godotenv.Load() // Loading .env file which contains config
	if err != nil {
		log.Fatal("Error loading .env file") // Error notif
	} else {
		log.Print(".env File loaded... reading env variables")
		log.Print("-----------------------------------------------")
	}

	s3Bucket := os.Getenv("S3_BUCKET")   // Read env variable 1
	secretKey := os.Getenv("SECRET_KEY") // Read env variable 2

	// now do something with s3 or whatever
	log.Printf("S3 Bucket: %s\tKey: %s\n", s3Bucket, secretKey)
	log.Println("-----------------------------------------------")

	// FUNCTIONS DEMO - CALLING A UTILS FUNCTION FROM A LOCAL PACKAGE
	fmt.Printf("FUNCTION CALL / LOCAL PACKAGE IMPORT DEMO\n")
	Utils.DemoFunction() // See next slide for details on this func
}
