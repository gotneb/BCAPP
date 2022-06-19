package main

import (
	"net/http"
	"os"

	"math/big"

	"github.com/gin-gonic/gin"
	"github.com/gotneb/bcapp/nibblelinx"
)

func main() {
	nibblelinx.Init()
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Helo there!")
	})

	r.GET("/modp/:n/:p1", func(c *gin.Context) {
		_n, _p1 := c.Param("n"), c.Param("p1")
		c.String(http.StatusOK, "Given values, n=%s and p1=%s\n\n", _n, _p1)

		n, _ := new(big.Int).SetString(_n, 10)
		p1, _ := new(big.Int).SetString(_p1, 10)
		result := nibblelinx.ModP(n, p1)
		c.JSON(http.StatusOK, gin.H{
			"result": result,
		})
	})

	// Handle calls to nibblelinx.DoubleP()
	r.GET("/doublep/:x/:y", func(c *gin.Context) {
		_x, _y := c.Param("x"), c.Param("y")
		c.String(http.StatusOK, "Given values, x=%s and y=%s\n\n", _x, _y)

		x, _ := new(big.Int).SetString(_x, 10)
		y, _ := new(big.Int).SetString(_y, 10)
		x.Set(nibblelinx.DoubleP(x, y)[0])
		y.Set(nibblelinx.DoubleP(x, y)[1])
		c.JSON(http.StatusOK, gin.H{
			"x": x,
			"y": y,
		})
	})

	// Handle calls to nibblelinx.Inverse()
	r.GET("/inverse/:r/:p", func(c *gin.Context) {
		_r, _p := c.Param("r"), c.Param("p")
		c.String(http.StatusOK, "Given values, r=%s and p=%s\n\n", _r, _p)

		r, _ := new(big.Int).SetString(_r, 10)
		p, _ := new(big.Int).SetString(_p, 10)
		result := nibblelinx.Inverse(r, p)
		c.JSON(http.StatusOK, gin.H{
			"result": result,
		})
	})

	port := os.Getenv("PORT")
	r.Run(":" + port)
}
