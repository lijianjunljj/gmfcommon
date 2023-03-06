package router

import "github.com/gin-gonic/gin"

type AbstractRouter interface {
	InAuthentic(*gin.RouterGroup)
	Authentic(*gin.RouterGroup)
}
