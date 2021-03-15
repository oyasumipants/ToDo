package main

import (
	"fmt"

	"./app/controllers"
	"./app/models"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()
}
