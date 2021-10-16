package main

import(
	// "gomvc/models"
	"fmt"
	"os"
	"reflect"
)
func main(){
	modelName := os.Args[1]
	StubStorage := map[string]interface{}{
		"funcA": modelName,
	}
	f := StubStorage["funcA"]
	fmt.Println(reflect.TypeOf(f))
	// models.$modelName()

	
}