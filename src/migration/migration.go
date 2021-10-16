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
	f := reflect.ValueOf(StubStorage["funcA"])
	fmt.Println(f)
	// models.$modelName()

	
}