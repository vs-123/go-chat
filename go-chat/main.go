package main

import (
	"net/http" // For serving our index.html file

	"github.com/gin-gonic/gin" // For creating the server
	"gopkg.in/olahol/melody.v1" // For sending and receiving messages (a websocket framework)
)

func main() {
	r := gin.Default() // Setting up gin
	m := melody.New() // Setting up melody

	r.GET("/", func(c *gin.Context) { // Setting a route for the main URL where the index.html will be served
		http.ServeFile(c.Writer, c.Request, "index.html") // Serving the index.html file
	})

	r.GET("/ws", func(c *gin.Context) { // Setting up the route for melody to handle requests
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) { // Handling messages that gets sent to the /ws route
		m.Broadcast(msg) // Broadcasting or sending the message to all clients listening to the /ws route
	})

	r.Run(":8080") // Start the server on port 8080
}