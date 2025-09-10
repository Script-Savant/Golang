package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

var client *ethclient.Client

func main() {
	var err error
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.Static("/static", "./static") // for htmx and bootstrap

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	r.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.POST("/api/register", func(c *gin.Context) {
		username := c.PostForm("username")
		email := c.PostForm("email")

		c.String(200, "Registered user %s with email %s", username, email)
	})

	r.GET("/block-number", func(c *gin.Context) {
		block, err := client.BlockByNumber(context.Background(), nil)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error %s", err.Error())
			return
		}
		c.String(http.StatusOK, fmt.Sprintf("Latest block: %d", block.Number().Uint64()))
	})

	r.GET("/latest-blocks", func(c *gin.Context) {
		head, err := client.BlockByNumber(context.Background(), nil)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error: %s", err.Error())
			return
		}

		html := "<ul class='list-group'>"
		for i := 0; i < 5; i++ {
			num := head.Number().Uint64() - uint64(i)
			blk, _ := client.BlockByNumber(context.Background(), big.NewInt(int64(num)))
			html += fmt.Sprintf("<li class='list-group-item'>Block %d -> %s</li>", blk.Number().Uint64(), blk.Hash().Hex())
		}
		html += "</ul>"

		c.Header("Content-Type", "text/html")
		c.String(http.StatusOK, html)
	})

	r.POST("/balance", func(c *gin.Context) {
		address := c.PostForm("address")
		acc := common.HexToAddress(address)

		balance, err := client.BalanceAt(context.Background(), acc, nil)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error: %s", err.Error())
			return
		}

		c.String(http.StatusOK, fmt.Sprintf("Balance: %s Wei", balance.String()))
	})

	r.GET("/watch-balance/:address", func(c *gin.Context) {
		addr := c.Param("address")
		acc := common.HexToAddress(addr)

		balance, err := client.BalanceAt(context.Background(), acc, nil)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error: %s", err.Error())
			return
		}
		c.String(http.StatusOK, fmt.Sprintf("Balance: %s", balance.String()))
	})

	r.Run(":8080")
}
