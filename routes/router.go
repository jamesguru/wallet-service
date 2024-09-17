package routes

import (
	"github.com/jamesguru/dev-api/controllers"
	"github.com/labstack/echo/v4"
)

func Setup() {
	e := echo.New()
	e.POST("/wallet", controllers.AddWallet)
	e.GET("/wallet/:id", controllers.GetWallet)
	e.GET("/wallets", controllers.GetWallets)
	e.Start(":3002")
}
