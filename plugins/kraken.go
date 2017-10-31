package plugins

import (
	"log"
	"net/http"

	"github.com/TobiEiss/kraken-go"
	"github.com/labstack/echo"
)

const Pair = "pair"

func KrakenTickerInfo(context echo.Context) error {
	pair := context.Request().Header.Get(Pair)
	if pair == "" {
		return context.JSON(http.StatusBadRequest, "set header pair")
	}

	// fetch tickerinfo
	session := krakenGo.CreateKrakenSession()
	tickerInfo, err := session.GetTickerInfo(pair)
	if err != nil {
		log.Println(err)
		return context.JSON(http.StatusInternalServerError, err)
	}
	return context.JSON(http.StatusOK, tickerInfo[pair])
}
