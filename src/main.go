package main

import (
	"fmt"

	"./app/models"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()

}
