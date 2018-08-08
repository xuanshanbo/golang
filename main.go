package main
import (
	"fmt"
	my "./server.go"
)

func main(){
	fmt.Println("ces")
	my.StartWebServer("6767")
}