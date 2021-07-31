package controllers

import (
	"bytes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTemplateRenderer(r renderer) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.AbortWithStatus(http.StatusBadRequest)
		}
		buf := new(bytes.Buffer)
		buf.ReadFrom(c.Request.Body)
		yml, err := r.Render(id, buf.Bytes())
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
		}
		c.YAML(http.StatusOK, yml)
	}
}
