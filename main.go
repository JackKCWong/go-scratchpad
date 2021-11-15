package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	scratchpad "github.com/JackKCWong/go-scratchpad/internal/core"
	"github.com/gin-gonic/gin"
)

func main() {
	unixsock := flag.String("unixsock", "", "path to unix socket")
	flag.Parse()

	fmt.Printf("%+v", os.Environ())
	
	r := gin.Default()

	var listener net.Listener
	var err error
	if len(*unixsock) > 0 {
		listener, err = net.Listen("unix", *unixsock)
	} else {
		listener, err = net.Listen("tcp", ":8080")
	}

	if err != nil {
		panic(err)
	}

	setupRoutes(r)

	r.RunListener(listener)
}

func setupRoutes(r *gin.Engine) {
	r.Static("/", "./web")

	r.POST("/wasm", func(c *gin.Context) {
		snippet, ok := c.GetPostForm("snippet")
		if !ok {
			c.JSON(400, gin.H{"error": "missing snippet"})
			return
		}

		bincode, err := scratchpad.CompileSnippet(snippet)

		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.Data(200, "application/wasm", bincode)
	})
}
