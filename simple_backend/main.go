package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type book struct {
	ID           string `json:"ID"`
	Tittle       string `json:"Tittle"`
	Author       string `json:"Author"`
	Already_Read bool   `json:"Already_Read"`
}

var books = []book{
	{ID: "1", Tittle: "Góry Pindos", Author: "Andrzej Murawski", Already_Read: true},
	{ID: "2", Tittle: "Powojnie", Author: "Tony Judt", Already_Read: true},
	{ID: "3", Tittle: "Stary Człowiek i Morze", Author: "Ernest Hemingway", Already_Read: true},
	{ID: "4", Tittle: "Raspberry PI Receptury", Author: "Simon Monk", Already_Read: false},
	{ID: "5", Tittle: "Monte Cassino 1944", Author: "Zbigniew Wawer", Already_Read: false},
	{ID: "6", Tittle: "Katanga 1960-1963", Author: "Daniel Kowalczuk", Already_Read: false},
}

func getBooks(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, books)
}

func getMainPage(context *gin.Context) {
	context.File("../simple_frontend/index.html")
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.Static("/simple_frontend", "../simple_frontend")
	router.GET("/", getMainPage)

	router.GET("/books", getBooks)

	router.Run(":8080")
}
