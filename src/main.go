package main

import (
	"./app/models"
)

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// log.Println("test")

	// fmt.Println(models.Db)
	// u := &models.User{}
	// u.Name = "test"
	// u.Email = "test@example.com"
	// u.PassWord = "testtest"
	// fmt.Println(u)

	// u.CreateUser()

	// u, _ := models.GetUser(1)
	// fmt.Println(u)

	// u.Name = "Test"
	// u.Email = "test@example.com"
	// // ? なんでこれで渡せる？
	// // uが更新されている
	// u.UpdateUser()

	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// u.DeleteUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// user, _ := models.GetUser(1)
	// user.CreateTodo("First Todo")

	t, _ := models.GetTodo(1)
	t.DeleteTodo()

	// 関数とメソッドで、modelsを対象とするかobjectを対象とするかが変わる
	// user1, _ := models.GetUser(1)
	// todos, _ := user1.GetTodosByUser()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }
}
