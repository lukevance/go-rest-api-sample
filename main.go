package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type pizza struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func main() {
	router := gin.Default()
	router.GET("/", getRoot)
	router.GET("/pizzas", getPizzas)
	router.POST("/pizzas", postPizzas)

	router.Run(":8080")
}

// seed data for pizzas
var pizzas = []pizza{
	{ID: "1", Name: "Cheese", Price: 9.99},
	{ID: "2", Name: "Pepperoni", Price: 11.99},
}

// basic greeting at root
func getRoot(c *gin.Context) {
	m := make(map[string]string)
	m["message"] = "Hello Pizza!"
	c.IndentedJSON(http.StatusOK, m)
}

// getPizzas responds with the list of all pizzas
func getPizzas(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, pizzas)
}

// postPizzas adds a pizza to the menu from the request
func postPizzas(c *gin.Context) {
	var newPizza pizza

	//bind the received json to newPizza
	if err := c.BindJSON(&newPizza); err != nil {
		return
	}

	//add pizza to our in memory object
	pizzas = append(pizzas, newPizza)
	c.IndentedJSON(http.StatusCreated, newPizza)
}
