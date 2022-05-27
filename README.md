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

**Clean Architecture**

https://manakuro.medium.com/clean-architecture-with-go-bce409427d31
