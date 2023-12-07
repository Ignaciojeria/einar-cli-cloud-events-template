package ce_source

import (
	"testing"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

func TestGoogleCloudStorageSourceBuilder(t *testing.T) {
	// Configuración del nombre del bucket y del objeto.
	bucketName := "my-bucket"
	objectName := "my-object.txt"

	// Creación de una nueva fuente de Google Cloud Storage.
	source := NewGoogleCloudStorageBucketSource(bucketName).
		WithObject(objectName).
		Build()

	// Creación de un evento CloudEvents.
	event := cloudevents.NewEvent()
	event.SetSource(source)
	event.SetType("entityType")

	// Comprobación de que la fuente se construyó correctamente.
	expectedSource := "gs://my-bucket/my-object.txt"
	if event.Source() != expectedSource {
		t.Errorf("La fuente de Google Cloud Storage no se construyó correctamente, obtenido: %s, esperado: %s", source, expectedSource)
	}
}
