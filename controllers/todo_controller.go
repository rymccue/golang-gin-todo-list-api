package controllers

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rymccue/golang-gin-todo-list-api/repositories"
)

func NewTodoController(db *sql.DB) *TodoController {
	return &TodoController{DB: db}
}

type TodoController struct {
	DB *sql.DB
}

func (tc *TodoController) Create(c *gin.Context) {
	title := c.PostForm("title")
	description := c.PostForm("description")
	_, err := repositories.CreateItem(tc.DB, title, description)
	if err != nil {
		log.Println("err", err)
		c.String(http.StatusInternalServerError, "")
		return
	}
	c.String(http.StatusCreated, "")
}

func (tc *TodoController) Items(c *gin.Context) {
	allStr := c.Query("all")
	all, err := strconv.ParseBool(allStr)
	if err != nil {
		all = true
	}
	items, err := repositories.GetItems(tc.DB, all)
	if err != nil {
		log.Println("err", err)
		c.String(http.StatusInternalServerError, "")
		return
	}
	c.JSON(http.StatusOK, items)
}

func (tc *TodoController) Get(c *gin.Context) {
	itemIDStr := c.Params.ByName("itemID")
	itemID, err := strconv.Atoi(itemIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Bad item id",
		})
		return
	}
	item, err := repositories.GetItem(tc.DB, itemID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.String(http.StatusNotFound, "")
			return
		}
		log.Println("err", err)
		c.String(http.StatusInternalServerError, "")
		return
	}
	c.JSON(http.StatusOK, item)
}

func (tc *TodoController) Update(c *gin.Context) {
	itemIDStr := c.Params.ByName("itemID")
	itemID, err := strconv.Atoi(itemIDStr)
	if err != nil {
		log.Println("err", err)
		c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Bad item id",
		})
		return
	}
	title := c.PostForm("title")
	description := c.PostForm("description")
	completedStr := c.PostForm("completed")
	completed, err := strconv.ParseBool(completedStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Bad completed input",
		})
		return
	}
	err = repositories.UpdateItem(tc.DB, itemID, title, description, completed)
	if err != nil {
		log.Println("err", err)
		c.String(http.StatusInternalServerError, "")
		return
	}
	c.String(http.StatusOK, "")
}

func (tc *TodoController) Delete(c *gin.Context) {
	itemIDStr := c.Param("itemID")
	itemID, err := strconv.Atoi(itemIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Bad item id",
		})
		return
	}
	err = repositories.DeleteItem(tc.DB, itemID)
	if err != nil {
		log.Println("err", err)
		c.String(http.StatusInternalServerError, "")
		return
	}
	c.String(http.StatusOK, "")
}
