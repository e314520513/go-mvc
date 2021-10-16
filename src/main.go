package main

import(
	
	// "gomvc/controllers"
	"gomvc/routers"
	// "gomvc/helpers"
	// "gomvc/models"
	"gomvc/connection"
)
func main(){
	connection.ConnectDatabase()
	routers.Router()
}