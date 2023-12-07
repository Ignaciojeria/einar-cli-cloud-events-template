package ce_source

import (
	"testing"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

// TestCloudEventWithACLSource verifica si el source del ACL se establece correctamente en un evento CloudEvents.
func TestCloudEventWithACLSource(t *testing.T) {
	// Configuración del ACL, entidad de origen, entidad de transformación y entidad de origen.
	aclIdentifier := "miACL"
	originalSource := "sistemaOriginal"
	sourceEntity := "miEntidadOrigen"
	transformationEntity := "miEntidadTransformacion"

	// Construcción del source.
	source := NewAntiCorruptionLayerSource(aclIdentifier).
		SetOriginalSource(originalSource).
		SetSourceEntity(sourceEntity).
		SetTransformationEntity(transformationEntity).
		Build()

	// Creación de un evento CloudEvents.
	event := cloudevents.NewEvent()
	event.SetSource(source)
	event.SetType("entityType")

	// Comprobación de que el source del evento es correcto.
	if event.Source() != source {
		t.Errorf("El source del evento CloudEvents no se estableció correctamente, obtenido: %s, esperado: %s", event.Source(), source)
	}
}
