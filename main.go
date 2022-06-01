package main

import (
	"log"
	"training/goproject/config"
	"training/goproject/infrastructure/db/sqlc"
	"training/goproject/infrastructure/router"
	"training/goproject/registry"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	_ "training/goproject/docs"
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

	if err := e.Start("localhost:8081"); err != nil {
		log.Fatalln(err)
	}

	// users, err := queries.ListUsers(context.Background())
	// if err != nil {
	// 	log.Fatal("ListUsers error:", err)
	// }
	// fmt.Println(users)

	//  erro:= queries.CreateUser(context.Background(), db.CreateUserParams{
	// 	Name:      "fdsfs",
	// 	Firstname: "fsdfds",
	// 	Age:       12,
	// })
	// if erro != nil {
	// 	log.Fatal("CreateUser erro:", erro)
	// }
	// fmt.Println(insertedUser)

	// fetchedUser, err := queries.GetUser(context.Background(), insertedUser.Id)
	// if err != nil {
	// 	log.Fatal("GetUser error:", err)
	// }
	// fmt.Println(fetchedUser)

	// err = queries.DeleteUser(context.Background(), insertedUser.Id)
	// if err != nil {
	// 	log.Fatal("DeleteUser error:", err)
	// }
	// config.ReadConfig()
}
