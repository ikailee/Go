package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Home: ", os.Getenv("HOME"))

	shell, ok := os.LookupEnv("SHELL")
	if !ok {
		fmt.Printf("env var SHELL is not set")
	} else {
		fmt.Println("SHELL: ", shell)
	}

	err := os.Setenv("NAME", "Kevin")
	if err != nil {
		fmt.Printf("could not set NAME")
	}

	fmt.Printf("NAME: %s", os.Getenv("NAME"))

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s3Bucket := os.Getenv("S3_BUCKET")
	secretKey := os.Getenv("SECRET_KEY")

	envMap, envErr := godotenv.Read(".env")
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	s3Bucket := envMap["S3_BUCKET"]
	secretKey := envMap["SECRET_KEY"]
}
