package catalogo

import "testing"

func TestRegistrarYBuscar(t *testing.T) {
	original := []Titulo{{ID: "1", Nombre: "El viaje", Genero: "Aventura"}}
	nueva, err := Registrar(original, Titulo{ID: "2", Nombre: "Viaje al mar"})
	if err != nil || len(nueva) != 2 || len(original) != 1 {
		t.Fatal("registro incorrecto", err)
	}
	if _, err := Registrar(nueva, Titulo{ID: "1", Nombre: "Duplicado"}); err == nil {
		t.Fatal("debe rechazar ID duplicado")
	}
	if got := BuscarPorNombre(nueva, "VIAJE"); len(got) != 2 {
		t.Fatalf("búsqueda: %v", got)
	}
}
