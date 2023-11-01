package cloudfunction_wegotour

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"

	tour "github.com/wegotour/be_p3/modul"
)

func init() {
	functions.HTTP("Wegotour", WegotourGet)
}

func WegotourGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "https://wegotour.github.io")
	fmt.Fprintf(w, tour.GCFHandler("MONGOSTRING", "wegotour", "user"))
}
