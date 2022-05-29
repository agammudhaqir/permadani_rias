package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func main() {

	// 1. Config Viper conf.json
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.SetConfigName("app.conf")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	// 2. Connect database mysql
	username := viper.GetString("database.user")
	password := viper.GetString("database.password")
	database := viper.GetString("database.name")
	host := viper.GetString("database.host")
	port := viper.GetInt("database.port")

	psqlInfo := fmt.Sprintf("%s:%s@(%s:%d)/%s", username, password, host, port, database)

	db, err := sql.Open("mysql", psqlInfo)
	if err != nil {
		fmt.Println("hlo")
		log.Fatalln(err)
	}

	defer db.Close()

	maxLifetime, _ := time.ParseDuration(viper.GetString("database.max_lifetime_connection") + "s")
	db.SetMaxIdleConns(viper.GetInt("database.max_idle_dbection"))
	db.SetConnMaxLifetime(maxLifetime)

	// 3. Routing
	router := gin.New()
	router.Use(cors.Default())
	sessionstore := memstore.NewStore([]byte("SuperDuperRahasia1234567890!@#$%"))
	sessionstore.Options(sessions.Options{
		Path:   "/",     // default /, atau bisa kosongin aja sama aja
		MaxAge: 60 * 30, // dalam satuan detik
	})
	router.Use(sessions.Sessions("PERMADANI RIAS", sessionstore))

	// //---- HANDLING FORCE ERROR ----
	router.Use(gin.Recovery())

	Routing(router)

	tmphttpreadheadertimeout, _ := time.ParseDuration(viper.GetString("server.readheadertimeout") + "s")
	tmphttpreadtimeout, _ := time.ParseDuration(viper.GetString("server.readtimeout") + "s")
	tmphttpwritetimeout, _ := time.ParseDuration(viper.GetString("server.writetimeout") + "s")
	tmphttpidletimeout, _ := time.ParseDuration(viper.GetString("server.idletimeout") + "s")
	//---- RUNNING SERVER WITH PORT -----
	s := &http.Server{
		Addr:              ":" + viper.GetString("server.port"),
		Handler:           router,
		ReadHeaderTimeout: tmphttpreadheadertimeout,
		ReadTimeout:       tmphttpreadtimeout,
		WriteTimeout:      tmphttpwritetimeout,
		IdleTimeout:       tmphttpidletimeout,
		//MaxHeaderBytes:    1 << 20,
	}
	fmt.Println("Server running on port:", viper.GetString("server.port"))
	s.ListenAndServe()

}

func Routing(router *gin.Engine) {
	// Root static
	router.Static("/img-storage", "./asset/img")
}
