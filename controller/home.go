package control

import (
	"html/template"
	"net/http"
	core "permadani_rias/core"
	"permadani_rias/helper"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func HomePage(c *gin.Context, dbs helper.DBStruct) {
	c.HTML(http.StatusOK, "home.html", gin.H{
		"headerClient": template.HTML(core.HeaderTemplateClient()),
		"footerClient": template.HTML(core.FooterTemplateClient()),
		"base_url":     viper.GetString("base_url"),
	})
}

func GetCatalog(c *gin.Context, dbs helper.DBStruct) {

	// get data dari database
	query := "SELECT * FROM tbl_barang where is_deleted = false"

	sql := dbs.DatabaseQueryRows(query) 

	c.JSON(http.StatusOK, gin.H{
		"status": "OK BRO",
		"data":   sql,
	})
}

