package main

import (
	"log"
	"training/goproject/config"
	_ "training/goproject/docs"
	db "training/goproject/infrastructure/db/sqlc"
	"training/goproject/infrastructure/router"
	"training/goproject/registry"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

// @title Article API
// @version 1.0
// @description This is a sample service
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email julien.tellier@openwt.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8081
// @BasePath /
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	config.ReadConfig()
	conn := db.OpenDB()
	defer db.OpenDB().Close()

	r := registry.NewRegistry(conn)
	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	if err := e.Start(config.C.App.Addr); err != nil {
		log.Fatalln(err)
	}
}
