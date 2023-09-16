package main

import (
	"flag"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func runHttp() {
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello, World!\n")
	})

	http.ListenAndServe(":8080", nil)
}

func runGin() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.SetTrustedProxies([]string{})
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!\n")
	})
	r.Run()
}

func main() {
	useGin := flag.Bool("gin", false, "use gin framework")
	flag.Parse()

	if *useGin {
		runGin()
	}

	runHttp()
}
