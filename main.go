package main

import (
	"log"
	"net/http"
	"training/goproject/config"
	"training/goproject/infrastructure/db/sqlc"
	"training/goproject/infrastructure/router"
	"training/goproject/registry"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

func main() {
	config.ReadConfig()
	conn := db.OpenDB()
	defer db.OpenDB().Close()

	r := registry.NewRegistry(conn)
	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

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
