// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package tutorial

import (
	"database/sql"
	"time"
)

type Reparacione struct {
	ID            int64
	Apellido      string
	Telefono      string
	Equipo        string
	Cable         int64
	Pata          int64
	Soporte       int64
	Control       int64
	Estado        string
	Falla         string
	FechaIngreso  time.Time
	FechaEntrega  sql.NullTime
	Costo         sql.NullInt64
	Observaciones sql.NullString
}
