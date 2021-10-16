package main

import(
	"gomvc/models"
	"os"
)
func main(){
	modelName := os.Args[1]
	models.$modelName()

	
}