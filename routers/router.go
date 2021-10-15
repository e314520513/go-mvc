package routers
import(
	"net/http"
	"log"
	"gomvc/controllers"
)
func Router(){

	http.HandleFunc("/",controllers.Home)

	if err:= http.ListenAndServe(":9090",nil); err!=nil{
		log.Fatal("ListenAndServe: ",err)
	}
}