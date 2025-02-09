package router

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kohge2/upsdct-server/database"
	"github.com/kohge2/upsdct-server/web/middleware"
)

func NewRouter(db *database.DB, txn *database.Transaction) *gin.Engine {
	r := gin.Default()

	// TODO ここでDIする
	log.Println(db, txn)

	authorized := r.Group("/api")
	authorized.Use(middleware.Auth())
	authorized.Use(middleware.ErrorMiddleware())
	{
		authorized.GET("/invoices", nil)
		authorized.POST("/invoices", nil)
	}

	return r
}
