package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rymccue/golang-gin-todo-list-api/controllers"
)

func CreateRoutes(r *gin.Engine, tc *controllers.TodoController) {
	r.GET("/items", tc.Items)
	r.POST("/item", tc.Create)
	r.GET("/item/:itemID", tc.Get)
	r.PUT("/item/:itemID", tc.Update)
	r.DELETE("/item/:itemID", tc.Delete)
}
