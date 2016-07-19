package main

import (
	"log"
	"net/http"

	"golang.org/x/net/context"

	"github.com/gnhuy91/stringsvc"
)

func main() {
	ctx := context.Background()

	var svc stringsvc.Service
	svc = stringsvc.NewStringService()

	h := stringsvc.MakeHTTPHandler(ctx, svc)
	log.Fatal(http.ListenAndServe(":8080", h))
}
