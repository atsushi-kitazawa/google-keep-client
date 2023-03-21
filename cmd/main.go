package main

import (
	"context"
	"log"

	"google.golang.org/api/keep/v1"
)

func main() {
	do()
}

func do() {
	log.Println("start")

	ctx := context.Background()
	service, err := keep.NewService(ctx)
	if err != nil {
		log.Fatalf("failed to create service. err = %v", err)
	}

	res, err := service.Notes.List().Do()
	if err != nil {
		log.Fatalf("failed to call list(). err = %v", err)
	}

	log.Println(res.HTTPStatusCode)
	log.Println(len(res.Notes))
}
