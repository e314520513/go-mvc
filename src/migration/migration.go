package main

import(
	"fmt"
	"gomvc/migration"
	"os"
)
func main(){
	modelName := os.Args[1]
	migration.{modelName}()

	
}