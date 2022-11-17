package todotrpt

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	todobiz "first-app/module/item/business"
	todostorage "first-app/module/item/storage"
)

// GetATodo godoc
//
// @Sumary      Get a Todo
// @Description Get a todo
// @Tags        todos
// @Accept      json
// @Produce     json
// @Param       id  query    string true "id" minlength(8) maxlength(100)
// @Success     200 {object} todomodel.ToDoItem
// @Failed      400 string "Bad request"
// @Failed      404 string "Not found"
// @Router      /todos/{id} [get]
func HandleFindAnItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		storage := todostorage.NewMySQLStorage(db)
		biz := todobiz.NewFindToDoItemBiz(storage)

		data, err := biz.FindAnItem(c.Request.Context(), map[string]interface{}{"id": id})

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}
