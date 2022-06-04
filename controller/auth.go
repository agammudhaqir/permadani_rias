package control

import (
	"html/template"
	"net/http"
	core "permadani_rias/core"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"headerAuth": template.HTML(core.HeaderAuth()),
		"footerAuth": template.HTML(core.FooterAuth()),
		"base_url": viper.GetString("base_url"),
		"ciri": "login",
	})
}
func ForgetPassword(c *gin.Context) {
	c.HTML(http.StatusOK, "password.html", gin.H{
		"headerAuth": template.HTML(core.HeaderAuth()),
		"footerAuth": template.HTML(core.FooterAuth()),
		"base_url": viper.GetString("base_url"),
		"ciri": "fore",
	})
}
func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "auth/login.html", gin.H{
		"headerAuth": template.HTML(core.HeaderAuth()),
		"footerAuth": template.HTML(core.FooterAuth()),
	})
}
