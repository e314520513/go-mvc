package routers
import(
	"net/http"
	"log"
	"gomvc/controllers"
)
func Router(){
	
	http.HandleFunc("/",controllers.Home)
	log.Println("Listening on localhost:9090")
	if err:= http.ListenAndServe(":9090",nil); err!=nil{
		log.Fatal("ListenAndServe: ",err)
	}
}