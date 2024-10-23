package web

import (
	"Cerrifix/tutorial"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func onToBool(key string) int {
	if key == "on" {
		return 1
	}
	return 0
}

func CerrifixWebHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	coso := r.Form
	tiempo := time.Now()
	fmt.Println(coso)

	db, err := sql.Open("sqlite3", os.Getenv("BLUEPRINT_DB_URL"))

	if err != nil {
		log.Fatal(err)
	}

	queries := tutorial.New(db)

	err = queries.CreateRep(r.Context(), tutorial.CreateRepParams{
		Apellido:     r.FormValue("apellido"),
		Telefono:     r.FormValue("telefono"),
		Equipo:       r.FormValue("equipo"),
		Cable:        int64(onToBool(r.FormValue("cable"))),
		Pata:         int64(onToBool(r.FormValue("pata"))),
		Soporte:      int64(onToBool(r.FormValue("soporte"))),
		Control:      int64(onToBool(r.FormValue("control"))),
		Estado:       r.FormValue("estado"),
		Falla:        r.FormValue("falla"),
		FechaIngreso: tiempo,
	})
	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "localhost:8080/cerrifix", http.StatusSeeOther)
	defer db.Close()

}
