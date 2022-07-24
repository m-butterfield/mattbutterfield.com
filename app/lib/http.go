package lib

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func BaseRouter() (*gin.Engine, error) {
	var r *gin.Engine
	if gin.Mode() == gin.ReleaseMode {
		r = gin.New()
	} else {
		r = gin.Default()
	}
	if err := r.SetTrustedProxies(nil); err != nil {
		return nil, err
	}
	return r, nil
}

func InternalError(err error, c *gin.Context) {
	log.Println(err)
	c.AbortWithStatus(http.StatusInternalServerError)
}
