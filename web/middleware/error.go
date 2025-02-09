package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kohge2/upsdct-server/utils"
	"github.com/kohge2/upsdct-server/web/response"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if err := c.Errors.Last(); err != nil {
			if appErr, ok := err.Err.(*utils.AppErr); ok {
				c.JSON(
					appErr.Code,
					response.NewErrorResponse(
						appErr.Type, appErr.Message, appErr.Code),
				)
				return
			}

			// アプリケーション側で定義したerror以外は500
			c.JSON(
				http.StatusInternalServerError,
				response.NewErrorResponse(
					utils.ErrTypeInternalServer, utils.ErrMsgInternalServer, http.StatusInternalServerError),
			)
			return
		}
	}
}
