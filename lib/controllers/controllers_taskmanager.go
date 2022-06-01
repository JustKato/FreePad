package controllers

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/JustKato/FreePad/lib/objects"
)

func TaskManager() {

	// Get the cleanup interval
	cleanupIntervalString, exists := os.LookupEnv("CLEANUP_MAX_AGE")
	if !exists {
		cleanupIntervalString = "-1"
	}

	if cleanupIntervalString == "-1" {
		// Do not cleanup
		return
	}

	// Try and parse the string as an int
	cleanupInterval, err := strconv.Atoi(cleanupIntervalString)
	if err != nil {
		cleanupInterval = 1
	}

	// Run all handlers
	go cleanupHandler(cleanupInterval)
	go savePostHandler()

}

func savePostHandler() {
	// Save the views cache
	fmt.Println("[Task::Save]: File save registered")
	for range time.NewTicker(time.Minute * 1).C {
		objects.SavePostViewsCache()
	}
}

func cleanupHandler(cleanupInterval int) {
	fmt.Println("[Task::Cleanup]: Cleanup task registered")
	for range time.NewTicker(time.Minute * 5).C {
		objects.CleanupPosts(cleanupInterval)
	}
}
