package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kohge2/upsdct-server/adapter"
	"github.com/kohge2/upsdct-server/database"
	"github.com/kohge2/upsdct-server/usecase"
	"github.com/kohge2/upsdct-server/web/handler"
	"github.com/kohge2/upsdct-server/web/middleware"

	"github.com/kohge2/upsdct-server/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(db *database.DB, txn *database.Transaction) *gin.Engine {
	r := gin.Default()

	invoiceUseCase := usecase.NewInvoiceUseCase(
		adapter.NewInvoiceRepository(db), adapter.NewUserRepository(db),
		adapter.NewCompanyRepository(db), adapter.NewPartnerCompanyRepository(db),
		*txn,
	)
	invoiceHandler := handler.NewInvoiceHandler(invoiceUseCase)

	authorized := r.Group("/api")
	authorized.Use(middleware.Auth())
	authorized.Use(middleware.ErrorMiddleware())
	{
		authorized.GET("/invoices", invoiceHandler.GetInvoices)
		authorized.POST("/invoices", invoiceHandler.CreateInvoice)
	}

	doc := r.Group("/docs")
	{
		docs.SwaggerInfo.BasePath = ""
		doc.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}

	return r
}
