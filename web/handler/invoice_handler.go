package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kohge2/upsdct-server/usecase"
	"github.com/kohge2/upsdct-server/web/request"
	"github.com/kohge2/upsdct-server/web/response"
)

type InvoiceHandler struct {
	invoiceUseCase usecase.InvoiceUseCase
}

func NewInvoiceHandler(invoiceUseCase usecase.InvoiceUseCase) *InvoiceHandler {
	return &InvoiceHandler{invoiceUseCase: invoiceUseCase}
}

// CreateInvoice upsdct-server godoc
//
// @Summary 請求書 登録
// @Description
// @Description ⚫︎パラメータについて: <br> 「paidDueDate」支払い期日 フォーマットはRFC3339の文字列(例:日本時間22時の場合 "2025-05-31T22:00:00+09:00") <br> 「partnerCompnayID」: 取引先ID <br> 「paidAmount」: 支払い金額
// @Description ⚫︎説明: <br> ログイン中のユーザー情報を取得し、そのユーザーが所属する企業の取引先企業への支払いについての請求書を登録するAPI
// @Produce json
// @Security Token || DebugUser
// @Param       request body request.CreateInvoiceRequest true " "
// @Tags invoice
// @Success     200  {object} response.PostResponse ""
// @Router      /api/invoices [post]
func (h *InvoiceHandler) CreateInvoice(c *gin.Context) {
	req := new(request.CreateInvoiceRequest)
	if err := req.Bind(c); err != nil {
		c.Error(err)
		return
	}
	paidDueDate, err := req.GetPaidDueDate()
	if err != nil {
		c.Error(err)
		return
	}
	// ログイン中のユーザー情報取得
	userID := c.GetString("userID")

	if err := h.invoiceUseCase.CreateInvoice(req.PaidAmount, req.PartnerCompanyID, userID, *paidDueDate, time.Now().UTC()); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.NewPostResponse())

}
