package main

import (
	"github.com/Sas-Kirakosyan/my-go-project/db"
	"github.com/Sas-Kirakosyan/my-go-project/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")

}

// var all = models.GetAllEvents()
// context.JSON(http.StatusCreated, gin.H{"massage": "event created", "event": all})
