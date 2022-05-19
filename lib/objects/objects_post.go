package objects

import (
	"fmt"
	"os"
)

type Post struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

func getStorageDirectory() string {

	baseStoragePath, exists := os.LookupEnv("PAD_STORAGE_PATH")
	if !exists {
		baseStoragePath = "/tmp/"
	}

	// Check if the base storage path exists
	if _, err := os.Stat(baseStoragePath); os.IsNotExist(err) {
		// Looks like the base storage path was NOT set, create the dir
		err = os.Mkdir(baseStoragePath, 0777)
		// Check for errors
		if err != nil {
			// No way this sends an error unless it goes horribly wrong.
			panic(err)
		}
	}

	// Return the base storage path
	return baseStoragePath
}

func GetPost(fileName string) Post {
	// Get the base storage directory and make sure it exists
	storageDir := getStorageDirectory()

	// Generate the file path
	filePath := fmt.Sprintf("%s%s", storageDir, fileName)

	fmt.Println("Reading: ", filePath)

	p := Post{
		Name:    fileName,
		Content: "",
	}

	// Check if the file exits
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		fmt.Println("Found ", filePath)
		// File does exist, read it and set the content
		data, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Println("Error:", err)
		}

		// Get the content of the file and put it in the response
		p.Content = string(data)
		fmt.Println("Loaded content for ", filePath)
	}

	return p
}

func WritePost(p Post) error {

	return nil
}
