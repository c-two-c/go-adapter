package services

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	bunnystorage "git.sr.ht/~jamesponddotco/bunnystorage-go"
)

func main() {
	config := &bunnystorage.Config{
		Key:         "7fed1593-5800-41ff-b517d6d6cd7a-6d67-403d",
		StorageZone: "sg.storage.bunnycdn.com",
	}

	client, err := bunnystorage.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create BunnyCDN Storage client: %v", err)
	}

	dir := "/path/to/directory"

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatalf("Failed to read directory: %v", err)
	}

	for _, file := range files {
		filePath := filepath.Join(dir, file.Name())

		f, err := os.Open(filePath)
		if err != nil {
			log.Fatalf("Failed to open file: %v", err)
		}
		defer f.Close()

		_, err = client.Upload(context.Background(), "/path/on/bunnycdn/"+file.Name(), "", "", f)
		if err != nil {
			log.Fatalf("Failed to upload file: %v", err)
		}

		log.Printf("Successfully uploaded file: %s", filePath)
	}
}
