package cloudfunction_todo

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"

	todo "github.com/mytodolist1/be_p3/modul"
)

func init() {
	functions.HTTP("MytodolistUser", MytodolistUser)
}

func MytodolistUser(w http.ResponseWriter, r *http.Request) {
	allowedOrigins := []string{"https://mytodolist.my.id", "http://127.0.0.1:5500", "http://127.0.0.1:5501"}
	origin := r.Header.Get("Origin")

	for _, allowedOrigin := range allowedOrigins {
		if allowedOrigin == origin {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			break
		}
	}
	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		// w.Header().Set("Access-Control-Allow-Origin", "https://mytodolist.my.id")
		w.Header().Set("Access-Control-Allow-Methods", "GET,PUT,DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	// Set CORS headers for the main request.
	// w.Header().Set("Access-Control-Allow-Origin", "https://mytodolist.my.id")
	id := r.URL.Query().Get("_id")
	username := r.URL.Query().Get("username")

	if r.Method == http.MethodDelete {
		fmt.Fprintf(w, todo.GCFHandlerDeleteUser("PASETOPUBLICKEY", "MONGOSTRING", "mytodolist", "user", r))
		return
	}
	if r.Method == http.MethodPut {
		if id != "" {
			fmt.Fprintf(w, todo.GCFHandlerUpdateUser("PASETOPUBLICKEY", "MONGOSTRING", "mytodolist", "user", r))
			return
		}
		if username != "" {
			fmt.Fprintf(w, todo.GCFHandlerChangePassword("PASETOPUBLICKEY", "MONGOSTRING", "mytodolist", "user", r))
			return
		}
	}
	if r.Method == http.MethodGet {
		if id != "" {
			fmt.Fprintf(w, todo.GCFHandlerGetUserByID("MONGOSTRING", "mytodolist", "user", r))
			return
		}
		if username != "" {
			fmt.Fprintf(w, todo.GCFHandlerGetUserByUsername("MONGOSTRING", "mytodolist", "user", r))
			return
		}
		fmt.Fprintf(w, todo.GCFHandlerGetUserFromToken("PASETOPUBLICKEY", "MONGOSTRING", "mytodolist", "user", r))
		return

	}
}
