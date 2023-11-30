package cloudfunction_todo

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"

	todo "github.com/mytodolist1/be_p3/modul"
)

func init() {
	functions.HTTP("MytodolistInsertTodo", MytodolistInsertTodo)
}

func MytodolistInsertTodo(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		// w.Header().Set("Access-Control-Allow-Origin", "https://mytodolist1.github.io")
		w.Header().Set("Access-Control-Allow-Origin", "https://mytodolist.my.id")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	// Set CORS headers for the main request.
	// w.Header().Set("Access-Control-Allow-Origin", "https://mytodolist1.github.io")
	w.Header().Set("Access-Control-Allow-Origin", "https://mytodolist.my.id")
	fmt.Fprintf(w, todo.GCFHandlerInsertTodo("PASETOPUBLICKEY", "MONGOSTRING", "mytodolist", "user", r))
}
