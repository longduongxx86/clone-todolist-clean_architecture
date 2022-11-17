package todotrpt

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	uuid "github.com/google/uuid"

	todobiz "first-app/module/item/business"
	todomodel "first-app/module/item/model"
	todostorage "first-app/module/item/storage"
)

// CreateTodo godoc
//
// @Sumary      Create  Todo
// @Description create  todo
// @Tags        todos
// @Accept      json
// @Produce     json
// @Param       dataItem body     todomodel.ToDoItem true "newTodo"
// @Success     200      {object} todomodel.ToDoItem
// @Failed      400 string "Bad request"
// @Failed      404 string "Not found"
// @Router      /todos [post]
func HandleCreateItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var dataItem todomodel.ToDoItem

		if err := c.ShouldBind(&dataItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		dataItem.Title = strings.TrimSpace(dataItem.Title)
		dataItem.Id = uuid.New().String()

		storage := todostorage.NewMySQLStorage(db)
		biz := todobiz.NewCreateToDoItemBiz(storage)

		if err := biz.CreateNewItem(c.Request.Context(), &dataItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": dataItem.Id})
	}
}
