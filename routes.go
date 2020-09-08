package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {

  c.HTML(
    http.StatusOK,
    "index.html",
    gin.H{
      "title": "Home Page",
    },
  )
}

func initializeRoutes(r *gin.Engine) {
  r.GET("/", showIndexPage)
}
