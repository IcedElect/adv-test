package routes

import (
	"icedelect/avito-test-adv/controllers"

	"github.com/gin-gonic/gin"
)

type Adv struct {
	advController controllers.Adv
}

func NewAdv(advController controllers.Adv) Adv {
	return Adv{advController}
}

func (a *Adv) AdvRoute(rg *gin.RouterGroup) {
	router := rg.Group("advs")
	router.GET("/", a.advController.All)
	router.POST("/", a.advController.Create)
	router.GET("/:id", a.advController.Find)
}
