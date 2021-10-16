package main

import(
	"fmt"
	"os"
	"reflect"
	"gomvc/models"
)
func main(){
	modelName := os.Args[1]
	
	
	models.funcs[modelName]

	
}