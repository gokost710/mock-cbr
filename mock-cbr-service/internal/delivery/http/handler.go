package httpdelivery

import (
	"github.com/gin-gonic/gin"
)

type RatesUsecase interface {
	GetDailyXML(dateReq string) (xml []byte, status int, err error)
}

type Handler struct {
	usecase RatesUsecase
}

func NewHandler(usecase RatesUsecase) *Handler {
	return &Handler{usecase: usecase}
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	router.GET("/scripts/XML_daily.asp", h.getDailyXML)
}

func (h *Handler) getDailyXML(c *gin.Context) {
	dateReq := c.Query("date_req")
	xml, status, err := h.usecase.GetDailyXML(dateReq)
	if err != nil {
		c.String(status, err.Error())
		return
	}
	c.Data(status, "application/xml; charset=utf-8", xml)
}
