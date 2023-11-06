package cloudfunction_wegotour

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"

	ticket "github.com/wegotour/be_p3/modul"
)

func init() {
	functions.HTTP("WgtGetUser", WegotourGetUser)
}

func WegotourGetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "https://wegotour.github.io")
	fmt.Fprintf(w, ticket.GCFHandler("MONGOSTRING", "wegotour", "user"))
}
