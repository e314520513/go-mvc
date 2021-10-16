package main

import(
	"fmt"
	"gomvc/models"
	"os"
)
func main(){
	modelName := os.Args[1]
	migration.{modelName}()

	
}