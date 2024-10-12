package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

// Replace with your credentials file path
const credentialsFile = "auth.json" // <-- UPDATE THIS!

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading env %v", err)

	}

	backupFile := os.Getenv("BACKUP_PATH")
	ctx := context.Background()
	srv, err := drive.NewService(ctx, option.WithCredentialsFile(credentialsFile))
	if err != nil {
		log.Fatalf("Unable to create Drive service: %v", err)
	}

	uploadFile(srv, backupFile)
}

func uploadFile(srv *drive.Service, filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Unable to open file: %v", err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatalf("Unable to get file info: %v", err)
	}

	fileMetadata := &drive.File{
		Name: fileInfo.Name(),
	}

	res, err := srv.Files.Create(fileMetadata).Media(file).Do()
	if err != nil {
		log.Fatalf("Unable to upload file: %v", err)
	}

	fmt.Printf("File ID: %s\n", res.Id)
	fmt.Printf("Uploaded file '%s' successfully.\n", filePath)
}
