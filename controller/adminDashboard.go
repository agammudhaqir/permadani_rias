package control

import (
	"html/template"
	"net/http"
	core "permadani_rias/core"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func AdminDashboard(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"headerAdmin": template.HTML(core.HeaderTemplateAdmin()),
		"footerAdmin": template.HTML(core.FooterTemplateAdmin()),
		"base_url":     viper.GetString("base_url"),
		"title_menu": "Dashboard",
	})
}
