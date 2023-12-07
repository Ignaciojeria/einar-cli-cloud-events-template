package ce_source

import (
	"testing"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

func TestPostgresSourceBuilder(t *testing.T) {
	// Configuración del nombre de la base de datos y de la tabla.
	databaseName := "mydb"
	tableName := "mytable"

	// Creación de una nueva fuente de PostgreSQL.
	source := NewPostgresSource(databaseName).
		WithTable(tableName).
		Build()

	// Comprobación de que la fuente se construyó correctamente.
	expectedSource := "postgresql://mydb/mytable"

	// Creación de un evento CloudEvents.
	event := cloudevents.NewEvent()
	event.SetSource(source)
	event.SetType("entityType")

	if event.Source() != expectedSource {
		t.Errorf("La fuente de PostgreSQL no se construyó correctamente, obtenido: %s, esperado: %s", source, expectedSource)
	}
}
