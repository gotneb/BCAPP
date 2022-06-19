package main

import (
	"net/http"
	"os"

	"math/big"

	"github.com/gin-gonic/gin"
	"github.com/gotneb/bcapp/nibblelinx"
)

func initHandlers() {
	r := gin.Default()

	// Handle startup page
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello there!")
	})

	// Handle calls to nibblelinx.ModP()
	r.GET("/modp/:n/:p1", func(c *gin.Context) {
		_n, _p1 := c.Param("n"), c.Param("p1")

		n, _ := new(big.Int).SetString(_n, 10)
		p1, _ := new(big.Int).SetString(_p1, 10)
		result := nibblelinx.ModP(n, p1)

		c.JSON(http.StatusOK, gin.H{
			"result": result.String(),
		})
	})

	// Handle calls to nibblelinx.Inverse()
	r.GET("/inverse/:r/:p", func(c *gin.Context) {
		_r, _p := c.Param("r"), c.Param("p")

		r, _ := new(big.Int).SetString(_r, 10)
		p, _ := new(big.Int).SetString(_p, 10)
		result := nibblelinx.Inverse(r, p)

		c.JSON(http.StatusOK, gin.H{
			"result": result.String(),
		})
	})

	// Handle calls to nibblelinx.DoubleP()
	r.GET("/doublep/:x/:y", func(c *gin.Context) {
		_x, _y := c.Param("x"), c.Param("y")

		x, _ := new(big.Int).SetString(_x, 10)
		y, _ := new(big.Int).SetString(_y, 10)

		c.JSON(http.StatusOK, gin.H{
			"result": [2]string{nibblelinx.DoubleP(x, y)[0].String(), nibblelinx.DoubleP(x, y)[1].String()},
		})
	})

	port := os.Getenv("PORT")
	r.Run(":" + port)
}

func main() {
	nibblelinx.Init()
	initHandlers()
}
