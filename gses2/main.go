package main

import (
	"./handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Обробник для отримання актуального курсу
	r.GET("/rate", handlers.GetCurrentPrice)

	// Обробник для підписки на електронну адресу
	r.POST("/subscribe", handlers.SubscribeHandler)

	// Обробник для розсилки електронних листів
	r.POST("/sendEmails", handlers.SendEmailsHandler)

	r.Run(":8888")
}
