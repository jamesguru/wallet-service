package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/jamesguru/dev-api/models"
	"github.com/labstack/echo/v4"
)

var wallets []models.Wallet

func AddWallet(c echo.Context) error {
	var wallet models.Wallet
	err := c.Bind(&wallet)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	if wallet.Name == "" {
		return c.JSON(http.StatusPreconditionFailed, "name cannot be empty")
	}
	if wallet.Author == "" {
		return c.JSON(http.StatusPreconditionFailed, "author cannot be empty")
	}
	log.Printf("name: %s", wallet.Name)
	log.Printf("Author: %s", wallet.Author)
	wallets = append(wallets, wallet)
	return c.JSON(http.StatusOK, wallets)
}
func GetWallet(c echo.Context) error {
	id := c.Param("id")
	idx, _ := strconv.ParseInt(id, 10, 64)
	return c.JSON(http.StatusOK, wallets[idx])
}

func GetWallets(c echo.Context) error {
	return c.JSON(http.StatusOK, "This is wallet service")
}
