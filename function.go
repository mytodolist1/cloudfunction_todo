package cfp3

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"

	todo "github.com/mytodolist1/be_p3"
)

func init() {
	functions.HTTP("Mytodolist", MytodolistPost)
}

func MytodolistPost(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "https://mytodolist1.github.io/fe_p3/")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	// Set CORS headers for the main request.
	w.Header().Set("Access-Control-Allow-Origin", "https://mytodolist1.github.io/fe_p3/")
	fmt.Fprintf(w, todo.GCFPostHandler("PASETOPRIVATEKEY", "MONGOSTRING", "mytodolist", "user", r))

}
