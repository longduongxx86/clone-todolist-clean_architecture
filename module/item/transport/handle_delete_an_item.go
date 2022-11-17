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
// @Sumary      Delete a Todo
// @Description delete a todo
// @Tags        todos
// @Accept      json
// @Produce     json
// @Param       id  path     string true "id"
// @Success     200 {object} object
// @Failed      400 string "Bad request"
// @Failed      404 string "Not found"
// @Router      /todos/{id} [delete]
func HandleDeleteAnItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		storage := todostorage.NewMySQLStorage(db)
		biz := todobiz.NewDeleteToDoItemBiz(storage)

		if err := biz.DeleteItem(c.Request.Context(), map[string]interface{}{"id": id}); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
