package ce_source

import (
	"testing"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

func TestFirestoreSourceBuilder(t *testing.T) {
	// Configuración del proyecto y nombres de colección y documento.
	projectID := "my-project"
	collectionName := "my-collection"
	documentName := "my-document"

	// Creación de un nuevo FirestoreSource.
	source := NewFirestoreSource(projectID).
		WithCollection(collectionName).
		WithDocument(documentName).
		Build()

	// Creación de un evento CloudEvents.
	event := cloudevents.NewEvent()
	event.SetSource(source)
	event.SetType("entityType")

	// Comprobación de que el source se construyó correctamente.
	expectedSource := "https://firestore.googleapis.com/v1/projects/my-project/databases/(default)/documents/my-collection/my-document"
	if event.Source() != expectedSource {
		t.Errorf("El source de Firestore no se construyó correctamente, obtenido: %s, esperado: %s", source, expectedSource)
	}
}
