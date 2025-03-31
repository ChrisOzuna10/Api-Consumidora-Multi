package main

import (
	"Api-Consumer-Multi/src/Websoket"
	mqtt "Api-Consumer-Multi/src/consumer"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	go mqtt.StartMQTTConsumer()

	r := gin.Default()

	r.GET("/ws", Websoket.HandleWebSocket)

	log.Println("Servidor WebSocket corriendo en http://localhost:4000/ws")
	r.Run(":8080")
}
