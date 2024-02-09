package domain

import (
	"time"

	"gorm.io/gorm"
)

type Entrada struct {
	gorm.Model
	Tipo              string `json:"tipo"` //Normal o especial
	InmuebleID        int
	Inmueble          Inmueble  `json:"inmueble_id"` //pertenece o has many?
	Fecha             time.Time `json:"fecha"`
	Origen            string    `json:"origen"`
	ProgramaID        int
	Programa          Programa `json:"programa_id"` //pertenece o has many?
	AreaResponsable   string   //`json:"recibe"`
	Entrega           string   //`json:"recibe"`
	PuestoEntrega     string   //`json:"recibe"`
	EmpresaEntrega    string   //`json:"recibe"`
	Autoriza          string   //`json:"autoriza"`
	TipoVehiculo      string   //`json:"tipo_vehiculo"`
	Capacidad         string   //`json:"capacidad"`
	Motivo            string
	Placas            string `json:"placas"`
	NombreConductor   string `json:"nombre_conductor"`
	Observaciones     string `json:"observaciones"`
	Folio             string `json:"folio" gorm:"unique"`
	Entradas_Producto []Entradas_Producto
}
