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

		result := nibblelinx.DoubleP(x, y)
		c.JSON(http.StatusOK, gin.H{
			"result": [2]string{
				result[0].String(),
				result[1].String(),
			},
		})
	})

	// Handle calls to nibblelinx.AddP()
	r.GET("/addp/:x1/:y1/:x2/:y2", func(c *gin.Context) {
		_x1, _y1, _x2, _y2 := c.Param("x1"), c.Param("y1"), c.Param("x2"), c.Param("y2")

		x1, _ := new(big.Int).SetString(_x1, 10)
		y1, _ := new(big.Int).SetString(_y1, 10)
		x2, _ := new(big.Int).SetString(_x2, 10)
		y2, _ := new(big.Int).SetString(_y2, 10)

		result := nibblelinx.AddP(x1, y1, x2, y2)
		//fmt.Printf("%s\n%s\n", result[0], result[1])
		c.JSON(http.StatusOK, gin.H{
			"result": [2]string{
				result[0].String(),
				result[1].String(),
			},
		})
	})

	port := os.Getenv("PORT")
	r.Run(":" + port)
}

func main() {
	nibblelinx.Init()
	initHandlers()
}
