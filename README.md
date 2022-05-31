# goproject

# GolangRestSwag

## Go

go mod init xxx/xxx (launch go project)
go mod tidy  
go build  
go run .  

## Go release

To do the command in git bash  
https://goreleaser.com/quick-start/

## Run docker

Build image with => docker build . -t golangrestswag  
Run image in a container => docker run --publish 8080:8080  golangrestswag  

## Go routine

**Sample :** 

func main() {  
  go print1()  
  fmt.Println("0")  
  time.Sleep(2 * time.Second)  
}  

func print1() {  
**2nd thread** go print2()  
**Main thread**	fmt.Println("1")  
  }  
  
  func print2() {  
	time.Sleep(time.Second)  
	fmt.Println("2")  
  }  
  
  **OUTPUT :**  
  0  
  1  
  2  

Une goroutine ne sera pas terminée, même la fonction/goroutine appelante est déjà terminée.  
Une goroutine ne sera terminée que si la fonction principale est déjà terminée ou si la goroutine a déjà fait son travail.  
Comme ce qui se passe dans l'exemple ci-dessus, la fonction print1 qui est l'appelant de print2 est terminée avant print2, mais le message de print2 est toujours imprimé.

https://medium.com/@manus.can/golang-tutorial-goroutines-and-channels-c2cd491f77ab

**SQLC**

https://conroy.org/introducing-sqlc
docker run --rm -v C:/Users/OWTPF3EZC3A/Training/Golang/goProject:/src -w /src kjconroy/sqlc generate
**Swagger**

Installation :
$ go get -d github.com/swaggo/swag/cmd/swag
$ go install github.com/swaggo/swag/cmd/swag@latest
$ swag init
$ go get -u github.com/swaggo/echo-swagger
$ go get github.com/labstack/echo/v4

For add a specific parameter => 

@Param <name> body <model> <required> <comment>
e.g.
@Param pet body Pet true "Pet to create"

After each header modifications run this command => swag init (to update swagger.json/yaml)

Sample :
***Main***
import "github.com/swaggo/echo-swagger" // IMPORTANT
import "github.com/labstack/echo/v4" // IMPORTANT v4 !!!!
**_ "training/goproject/docs"** // IMPORTANT


func main() {
	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":1323"))
}

***Model***

package model

type User struct {
	// Id
	// in: int32
	Id        int32       `json:"id"`
	// Name
	// in: string
	Name      string     `json:"name"`
	// Firstname
	// in: string
	Firstname  string    `json:"firstname"`
	// Age
	// in: int32
	Age       int32     `json:"age"`
}

func (User) TableName() string { return "users" }

***controller***
// Return All Users godoc
// @Summary Return All Users
// @Description Return All Users
// @Tags id
// @Accept  json
// @Produce  json
// @Success 200 
// @Router /users [get]
func (uc *userController) GetUsers(c Context) error {
	fmt.Println("Endpoint Hit: GetUsers")
	var u []*model.User

	u, err := uc.userInteractor.Get(u)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}

***Link***

http://localhost:1323/swagger/index.html 

**Clean Architecture**

https://manakuro.medium.com/clean-architecture-with-go-bce409427d31
