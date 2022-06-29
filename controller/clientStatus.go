package control


import (
	"html/template"
	"net/http"
	core "permadani_rias/core"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func ClientStatus(c *gin.Context) {
	c.HTML(http.StatusOK, "status.html", gin.H{
		"headerClient": template.HTML(core.HeaderTemplateClient()),
		"footerClient": template.HTML(core.FooterTemplateClient()),
		"base_url":     viper.GetString("base_url"),
		"title_menu": "Status",
	})
}
