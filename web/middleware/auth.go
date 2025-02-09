package middleware

import "github.com/gin-gonic/gin"

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO セッション情報からユーザー情報を取得
		c.Set("userID", "test1")
		c.Next()
	}
}
