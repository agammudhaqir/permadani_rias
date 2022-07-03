package control

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	core "permadani_rias/core"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func HomePage(c *gin.Context, db *sql.DB) {
	c.HTML(http.StatusOK, "home.html", gin.H{
		"headerClient": template.HTML(core.HeaderTemplateClient()),
		"footerClient": template.HTML(core.FooterTemplateClient()),
		"base_url":     viper.GetString("base_url"),
	})
}

func GetCatalog(c *gin.Context, db *sql.DB) {
	SQL := "SELECT id from tbl_barang where id = 1"
	// fmt.Println(SQL)
	res, err := db.Query(SQL)
	if err != nil {
        // http.Redirect(response, request, "/error", 302)
        // return, so no else is needed
        return
    }

    if err != nil {
        panic(err)
    }
    defer res.Close()
    for res.Next() {
        var (
            id int
        )
        if err := res.Scan(&id); err != nil {
            panic(err)
        }
        fmt.Printf("%d\n", id)
    }
    if err := res.Err(); err != nil {
        panic(err)
    }
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}
