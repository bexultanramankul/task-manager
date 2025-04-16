package middleware

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AccessMiddleware struct {
	jwtSecret string
}

func NewAccessMiddleware(jwtSecret string) *AccessMiddleware {
	return &AccessMiddleware{jwtSecret: jwtSecret}
}

// AdminOnly проверяет, что пользователь является администратором
func (m *AccessMiddleware) AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("userRole")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authentication required"})
			return
		}

		if userRole != "admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "admin access required"})
			return
		}

		c.Next()
	}
}

// BoardOwner проверяет, что пользователь является владельцем доски
// Для этого нам нужно получить boardID из параметров запроса и сравнить с userID в токене
func (m *AccessMiddleware) BoardOwner() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("userID")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authentication required"})
			return
		}

		boardID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid board ID"})
			return
		}

		// Здесь должна быть логика проверки принадлежности доски пользователю
		// Так как мы не обращаемся к БД, то нужно либо:
		// 1. Передавать owner_id в параметрах запроса (небезопасно)
		// 2. Использовать кеш (Redis)
		// 3. Оставить проверку в обработчике
		// В данном примере просто пропускаем - реальная проверка должна быть в обработчике
		_ = boardID
		_ = userID

		c.Next()
	}
}

// TaskOwnerOrBoardOwner проверяет, что пользователь является либо создателем задачи, либо владельцем доски
// Аналогично BoardOwner, без доступа к БД мы не можем проверить это
func (m *AccessMiddleware) TaskOwnerOrBoardOwner() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("userID")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authentication required"})
			return
		}

		taskID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
			return
		}

		// Аналогично BoardOwner, без доступа к БД проверка невозможна
		_ = taskID
		_ = userID

		c.Next()
	}
}

// UserOwner проверяет, что пользователь работает со своим собственным ресурсом
func (m *AccessMiddleware) UserOwner() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("userID")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authentication required"})
			return
		}

		requestedUserID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
			return
		}

		if uint(requestedUserID) != userID.(uint) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "you can only access your own resources"})
			return
		}

		c.Next()
	}
}
