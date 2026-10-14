// Program entry point: assembly and startup only, no business logic.
package main

import (
	"log"
	"net/http"

	"github.com/ravenmk2/baodou/internal/server"
	"github.com/ravenmk2/baodou/internal/store"
)

var version = "dev" // default when not injected; release builds inject via -ldflags "-X main.version="

func main() {
	db, err := store.Open(store.DefaultDSN)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("close database: %v", err)
		}
	}()

	handler := server.New(db)
	log.Printf("listening on :8080 (version %s)", version)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
