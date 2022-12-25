package main

import (
	"fmt"
	"net/http"
	"time"

	controllers "permadani_rias/controller"
	helper "permadani_rias/helper"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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

	//psqlInfo := fmt.Sprintf("%s:%s@(%s:%d)/%s", username, password, host, port, database)
	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",host, port, username, password, database)
	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, database)

	// db, err := sql.Open("postgres", psqlInfo)
	// if err != nil {
	// 	fmt.Println("Error connect database, please check!!!")
	// 	log.Fatalln(err)
	// }
	// defer db.Close()

	// maxLifetime, _ := time.ParseDuration(viper.GetString("database.max_lifetime_connection") + "s")
	// db.SetMaxIdleConns(viper.GetInt("database.max_idle_dbection"))
	// db.SetConnMaxLifetime(maxLifetime)

	//connectionString := fmt.Sprintf("%s:%s@(%s:%d)/%s", username, password, host, port, database)
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, database)
	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		fmt.Println("Connection to DB Error : ", err)
	}
	defer db.Close()
	db.SetConnMaxLifetime(30 * time.Second)
	db.SetConnMaxIdleTime(30)

	dbs := helper.DBStruct{Dbx: db}

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
	// fmt.Println(db)
	// return
	router.Use(gin.Recovery())

	Routing(router, dbs)

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

func Routing(router *gin.Engine, dbs helper.DBStruct) {
	// Root static
	router.Static("/img-storage", "./asset/img")
	router.Static("/asset", "./views/static")
	// Load HTML
	router.LoadHTMLGlob("views/html/*/*.html")

	// Root group

	router.GET("", func(c *gin.Context) { controllers.LoginPage(c) })
	router.GET("/password", func(c *gin.Context) { controllers.ForgetPassword(c) })

	routeAuth := router.Group("auth")
	{
		routeAuth.POST("login", func(c *gin.Context) { controllers.Login(c) })
	}
	// CLIENT
	routeClient := router.Group("client")
	{
		routeClient.GET("", func(c *gin.Context) { controllers.HomePage(c, dbs) })
		routeClient.GET("catalog", func(c *gin.Context) { controllers.CatalogPage(c, dbs) })
		routeClient.GET("jadwal", func(c *gin.Context) { controllers.ClientJadwal(c) })
		routeClient.GET("status", func(c *gin.Context) { controllers.ClientStatus(c) })
		routeClient.GET("transaksi", func(c *gin.Context) { controllers.ClientTransaksi(c) })
	}
	routeHome := router.Group("home/api")
	{
		routeHome.GET("get-catalog", func(c *gin.Context) { controllers.GetCatalog(c, dbs) })
	}


	// ADMIN
	routeAdmin := router.Group("admin")
	{
		routeAdmin.GET("", func(c *gin.Context) { controllers.AdminDashboard(c) })
		routeAdmin.GET("data-barang", func(c *gin.Context) { controllers.DataBarang(c) })
		routeAdmin.GET("data-penyewaan", func(c *gin.Context) { controllers.DataPenyewaan(c) })
		routeAdmin.GET("data-jadwal", func(c *gin.Context) { controllers.DataJadwal(c) })
		routeAdmin.GET("pengeluaran", func(c *gin.Context) { controllers.Pengeluaran(c) })
		routeAdmin.GET("pemasukan", func(c *gin.Context) { controllers.Pemasukan(c) })
		routeAdmin.GET("transaksi", func(c *gin.Context) { controllers.DataTransaksi(c) })
	}
}
