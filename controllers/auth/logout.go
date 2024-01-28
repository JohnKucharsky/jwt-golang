package controllers

import (
	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	c.SetCookie("Authorization", "", 0, "", "", false, true)
}
