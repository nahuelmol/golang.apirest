import (
	"github.com/gorilla/mux"
)

router := mux.NewRouter().StrictSlash(true)


func main (){
	router.HandleFunc("/", homePage)
}

