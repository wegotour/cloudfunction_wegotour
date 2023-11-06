package cloudfunction_wegotour

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"

	ticket "github.com/wegotour/be_p3/modul"
)

func init() {
	functions.HTTP("WegotourLogin", WegotourLogin)
}

func WegotourLogin(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "https://wegotour.github.io")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	// Set CORS headers for the main request.
	w.Header().Set("Access-Control-Allow-Origin", "https://wegotour.github.io")
	fmt.Fprintf(w, ticket.GCFHandlerLogIn("PASETOPRIVATEKEY", "MONGOSTRING", "wegotour", "user", r))

}
