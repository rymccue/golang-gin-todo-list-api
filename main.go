package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rymccue/golang-gin-todo-list-api/controllers"
	"github.com/rymccue/golang-gin-todo-list-api/routes"
	"github.com/rymccue/golang-gin-todo-list-api/utils/database"
)

func main() {
	db, err := database.Connect(os.Getenv("PGUSER"), os.Getenv("PGPASS"), os.Getenv("PGDB"), os.Getenv("PGHOST"), os.Getenv("PGPORT"))
	if err != nil {
		log.Fatal("err", err)
	}
	todoController := controllers.NewTodoController(db)
	r := gin.Default()
	routes.CreateRoutes(r, todoController)
	r.Run()
}
