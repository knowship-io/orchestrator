/*
	The orchestrator acts as the target for Coeus Fireworq queues to 
	init workload processes such as initiating brigade.sh pipelines
*/
package main

import (
	"github.com/gofiber/fiber"
	"github.com/monaco-io/request"
	"log"
)

func main() {
	// Fiber instance
	app := fiber.New()

	// Routes
	app.Get("/", hello)

	// start server
	app.Listen(8082)
}

// Handler
func hello(c *fiber.Ctx) {
	client := request.Client{
		URL:    "https://simpleevents/v1/ax68t7ax321f/secret/",
		Method: "POST",
		Params: map[string]string{"input": "s3://coeus/videos/vid_01.mp4"},
	}
	resp, err := client.Do()

	log.Println(resp.Code, string(resp.Data), err)
	c.Send("Success")
}
