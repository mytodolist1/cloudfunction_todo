package cloudfunction_todo

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"

	todo "github.com/mytodolist1/be_p3/modul"
)

func init() {
	functions.HTTP("MytodolistGetUser", MytodolistGetUser)
}

func MytodolistGetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "https://mytodolist1.github.io")
	fmt.Fprintf(w, todo.GCFHandler("MONGOSTRING", "mytodolist", "user", "username"))
}
