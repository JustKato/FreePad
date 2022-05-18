package controllers

import (
	"time"
)

func TaskManager() {

	// Run the migrations function
	go handleMigrations()

	for range time.Tick(time.Second * 5) {
		// fmt.Printf("%s\n", time.Now().Format("02/01/2006 03:04 PM"))
	}

}

func handleMigrations() {
	time.AfterFunc(time.Second*30, func() {
		// Run Migrations
	})
}
