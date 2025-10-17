package app

import (
	httpdelivery "mock-cbr-service/internal/delivery/http"
	ratesrepo "mock-cbr-service/internal/repository/rates"
	ratesusecase "mock-cbr-service/internal/usecase/rates"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	repo := ratesrepo.New()
	usecase := ratesusecase.New(repo)

	r := gin.Default()
	h := httpdelivery.NewHandler(usecase)
	h.RegisterRoutes(r)

	return r
}
