package web

import (
	"Cerrifix/tutorial"
	"database/sql"
	"log"
	"net/http"
	"os"
)

func CerriHome(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
	var buscar sql.NullString

	buscar.String = r.FormValue("buscar")
	if buscar.String != "" {
		buscar.Valid = true
	}

	db, err := sql.Open("sqlite3", os.Getenv("BLUEPRINT_DB_URL"))

	if err != nil {
		log.Fatal(err)
	}

	queries := tutorial.New(db)

	reparacion, err := queries.GetRepByTelParcial(r.Context(), buscar)
	if err != nil {
		log.Fatal(err)
	}

	idBuscador := reparacion[0].ID

	reparacion2, err := queries.GetRepByID(r.Context(), idBuscador)

	if err != nil {
		log.Fatal(err)
	}

	err = Buscador(reparacion2).Render(r.Context(), w)
	if err != nil {
		log.Fatal(err)
	}
}
