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

	fmt.Println("[Task::Cleanup]: Task registered")
	for range time.Tick(time.Minute * 5) {
		objects.CleanupPosts(cleanupInterval)
	}

}
