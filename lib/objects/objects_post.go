package objects

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"sync"
	"time"

	"github.com/JustKato/FreePad/lib/helper"
)

// Initialize the views cache
var ViewsCache map[string]uint32 = make(map[string]uint32)

// Mutex lock for the ViewsCache
var viewersLock sync.Mutex

type Post struct {
	Name         string `json:"name"`
	LastModified string `json:"last_modified"`
	Content      string `json:"content"`
	Views        uint32 `json:"views"`
}

// Get the path to the views JSON
func getViewsFilePath() (string, error) {
	// Get the path to the storage then append the const name for the storage file
	filePath := path.Join(getStorageDirectory(), "views_storage.json")

	// Check if the file exists
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		// Create the file
		err := os.WriteFile(filePath, []byte(""), 0640)
		if err != nil {
			return ``, err
		}
	}

	// Return the file path
	return filePath, nil
}

// Load the views cache from file
func LoadViewsCache() error {
	// Get the views file path
	viewsFilePath, err := getViewsFilePath()
	if err != nil {
		// This is now a problem for the upstairs
		return err
	}

	// Read the contents of the file as a map
	f, err := os.ReadFile(viewsFilePath)
	if err != nil {
		// This is now a problem for the upstairs
		return err
	}

	// Check if the contents are valid
	if len(f) <= 0 && len(string(f)) <= 0 {
		// The file is completely empty!
		return nil
	}

	var parsedData map[string]uint32

	// Parse the data
	err = json.Unmarshal(f, &parsedData)
	if err != nil {
		// This is now a problem for the function caller :D
		return err
	}

	storageDir := getStorageDirectory()
	// Loop through all of the mapped files
	for fileName := range parsedData {
		// Grab the path to the file
		// TODO: Create a generic function that checks if a post exists by name to use here adn in the GetPost method.
		filePath := path.Join(storageDir, fileName)
		// Check if the file exists
		if _, err := os.Stat(filePath); !os.IsNotExist(err) {
			// Looks like the file does not exist anymore, remove it from the map
			delete(parsedData, fileName)
		}
	}

	// Update the current cache
	ViewsCache = parsedData

	return nil
}

func AddViewToPost(postName string) uint32 {
	// Lock the viewers mapping
	viewersLock.Lock()

	// Check if the map has any value set to this elem
	if _, ok := ViewsCache[postName]; !ok {
		// Set the map value
		ViewsCache[postName] = 0
	}

	// Add to the counter
	ViewsCache[postName]++

	// Unlock
	viewersLock.Unlock()

	// Return the value
	return ViewsCache[postName]
}

func SavePostViewsCache() error {

	data, err := json.Marshal(ViewsCache)
	if err != nil {
		return err
	}

	viewsFilePath, err := getViewsFilePath()
	if err != nil {
		return err
	}

	f, err := os.OpenFile(viewsFilePath, os.O_WRONLY|os.O_CREATE, 0640)
	if err != nil {
		return err
	}
	// Actually close the file
	defer f.Close()

	// Delete all past content
	err = f.Truncate(0)
	if err != nil {
		return err
	}

	// Reset pointer
	_, err = f.Seek(0, 0)
	if err != nil {
		return err
	}

	// Write the json into storage
	_, err = f.Write(data)

	return err
}

// Get the path to the storage directory
func getStorageDirectory() string {

	baseStoragePath, exists := os.LookupEnv("PAD_STORAGE_PATH")
	if !exists {
		baseStoragePath = "/tmp/"
	}

	// Check if the base storage path exists
	if _, err := os.Stat(baseStoragePath); os.IsNotExist(err) {
		// Looks like the base storage path was NOT set, create the dir
		err = os.Mkdir(baseStoragePath, 0640)
		// Check for errors
		if err != nil {
			// No way this sends an error unless it goes horribly wrong.
			panic(err)
		}
	}

	// Return the base storage path
	return baseStoragePath
}

// Get a post from the file system
func GetPost(fileName string) Post {
	// Get the base storage directory and make sure it exists
	storageDir := getStorageDirectory()

	// Generate the file path
	filePath := fmt.Sprintf("%s%s", storageDir, fileName)

	// Get the post views and add 1 to them
	postViews := AddViewToPost(fileName)

	p := Post{
		Name:         fileName,
		Content:      "",
		Views:        postViews,
		LastModified: "Never Before",
	}

	// Check if the file exits
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		// File does exist, read it and set the content
		data, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Println("Error:", err)
		}

		// Get the content of the file and put it in the response
		p.Content = string(data)

		// Get last modified date
		fileData, err := os.Stat(filePath)
		if err == nil {
			p.LastModified = fileData.ModTime().Format("02/01/2006 03:04:05 PM")
		} else {
			fmt.Println(err)
		}
	}

	return p
}

// Write a post to the file system
func WritePost(p Post) error {

	maximumPadSize := helper.GetMaximumPadSize()
	if len(p.Content) > maximumPadSize {
		return errors.New("The pad is too big, please limit to the maximum of " + fmt.Sprint(maximumPadSize) + " characters")
	}

	// Get the base storage directory and make sure it exists
	storageDir := getStorageDirectory()

	// Generate the file path
	filePath := fmt.Sprintf("%s%s", storageDir, p.Name)

	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}

	// Write the contnets
	_, err = f.WriteString(p.Content)
	if err != nil {
		return err
	}

	return f.Close()
}

// Cleanup all of the older posts based on the environment settings
func CleanupPosts(age int) {
	// Initialize the files buffer
	var files []string

	// Get the base path
	root := getStorageDirectory()

	// Scan the directory
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		// Check that the file is not a dir
		if !info.IsDir() {
			files = append(files, path)
		}

		return nil
	})

	// Check if we had any errors
	if err != nil {
		fmt.Println("[Task::Cleanup]: Error", err)
	}

	// The timestamp where the file should be deleted
	tooOldTime := time.Now()

	// Go through all files and process them
	for _, filePath := range files {

		// Get last modified date
		fileData, err := os.Stat(filePath)
		if err == nil {
			fileAge := fileData.ModTime()

			// Check if the file is too old
			if fileAge.Add(time.Duration(age)*time.Minute).Unix() < tooOldTime.Unix() {
				fmt.Println("[Task::Cleanup]: Removing File", filePath)

				// Remove the file
				err = os.Remove(filePath)
				if err != nil {
					fmt.Println("[Task::Cleanup]: Failed to remove file", filePath)
				}
			}

		} else {
			fmt.Println("[Task::Cleanup]: Error", err)
		}

	}
}

func GetAllPosts() []Post {
	// Initialize the list of posts
	postList := []Post{}

	// Get the posts storage directory
	storageDir := getStorageDirectory()

	// Read the directory listing
	files, err := os.ReadDir(storageDir)
	// Check if thereh as been an issues with reading the directory contents
	if err != nil {
		// Log the error
		fmt.Println("Error::GetAllPosts:", err)
		// Return an empty list to have a clean fallback
		return []Post{}
	}

	// Go through all of the files
	for _, v := range files {
		// Process the file into a pad
		postList = append(postList, GetPost(v.Name()))
	}

	// Return the post list
	return postList
}
