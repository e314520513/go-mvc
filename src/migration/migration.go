package main

import(
	"gomvc/models"
	"os"
)
func main(){
	modelName := os.Args[1]
	StubStorage := map[string]interface{}{
		"funcA": modelName,
	}
	fmt.Println(StubStorage["funcA"])
	// models.$modelName()

	
}