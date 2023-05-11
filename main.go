package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"tripatra-api/db"
	"tripatra-api/router"
)

func main() {
	gormDB, err := db.Open()
	if err != nil {
		log.Fatal(err)
	}

	err = router.Router(gormDB)
	if err != nil {
		log.Fatal(err)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	log.Println("Shutting down...")
}
