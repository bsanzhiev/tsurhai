package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Pong
// мы защитим эту конечную точку, чтобы только запросы, имеющие
// действительный JWT в заголовке запроса, могли получить к ней доступ.
// Необходимо разместить эту проверку где-то глобально и сделать ее
// пригодной для использования всеми конечными точками, которые нам нужно защитить.
func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ping-pong!"})
}
