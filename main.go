package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)


func main() {
    ctx := context.Background()

    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        go consumes(ctx)
        produces(ctx)
        c.JSON(http.StatusOK, gin.H{"message": "gin tonic with lime is awesome"})
    })

    r.Run()
}
