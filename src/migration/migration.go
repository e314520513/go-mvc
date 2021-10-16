package main

import(
	"fmt"
	"gomvc/connection"
)
func main(){
	connection.ConnectDatabase()
	fmt.Println("connected")
}