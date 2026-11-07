// Package catalogo contiene las primeras funciones del prototipo StreamUIDE.
package catalogo

import (
	"errors"
	"strings"
)

// Titulo representa un contenido audiovisual con datos ficticios.
type Titulo struct {
	ID              string
	Nombre          string
	Genero          string
	Anio            int
	DuracionMinutos int
	Disponible      bool
}

// Registrar devuelve una nueva lista; no modifica el catálogo recibido.
func Registrar(catalogo []Titulo, nuevo Titulo) ([]Titulo, error) {
	if strings.TrimSpace(nuevo.ID) == "" || strings.TrimSpace(nuevo.Nombre) == "" {
		return nil, errors.New("el identificador y el nombre son obligatorios")
	}
	for _, existente := range catalogo {
		if existente.ID == nuevo.ID {
			return nil, errors.New("el identificador ya existe")
		}
	}
	resultado := append([]Titulo(nil), catalogo...)
	return append(resultado, nuevo), nil
}

// BuscarPorNombre selecciona títulos cuyo nombre contiene el texto buscado.
func BuscarPorNombre(catalogo []Titulo, texto string) []Titulo {
	consulta := strings.ToLower(strings.TrimSpace(texto))
	var resultados []Titulo
	for _, titulo := range catalogo {
		if strings.Contains(strings.ToLower(titulo.Nombre), consulta) {
			resultados = append(resultados, titulo)
		}
	}
	return resultados
}
