package todotrpt

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	todobiz "first-app/module/item/business"
	todomodel "first-app/module/item/model"
	todostorage "first-app/module/item/storage"
)

// UpdateTodo 	godoc
//
// @Sumary      Update  Todo
// @Description update  todo
// @Tags        todos
// @Accept      json
// @Produce     json
// @Param       id       path                        string true "id"
// @Param       dataItem body     todomodel.ToDoItem true   "updateTodo"
// @Success     200      {object} todomodel.ToDoItem
// @Failed      400 string "Bad request"
// @Failed      404 string "Not found"
// @Router      /todos/{id} [put]
func HandleUpdateAnItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var dataItem todomodel.ToDoItem

		if err := c.ShouldBind(&dataItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return	
		}

		storage := todostorage.NewMySQLStorage(db)
		biz := todobiz.NewUpdateToDoItemBiz(storage)

		if err := biz.UpdateItem(c.Request.Context(), map[string]interface{}{"id": id}, &dataItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
