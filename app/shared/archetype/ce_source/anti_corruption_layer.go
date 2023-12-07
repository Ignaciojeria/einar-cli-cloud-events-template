package ce_source

import (
	"fmt"
)

type AntiCorruptionLayerSource struct {
	aclIdentifier        string
	originalSource       string
	sourceEntity         string
	transformationEntity string
}

// NewAntiCorruptionLayerSource inicia la construcción del source para un ACL.
func NewAntiCorruptionLayerSource(aclIdentifier string) *AntiCorruptionLayerSource {
	return &AntiCorruptionLayerSource{
		aclIdentifier: aclIdentifier,
	}
}

// SetOriginalSource establece el origen original del evento.
func (b *AntiCorruptionLayerSource) SetOriginalSource(originalSource string) *AntiCorruptionLayerSource {
	b.originalSource = originalSource
	return b
}

// SetSourceEntity establece la entidad de origen del evento.
func (b *AntiCorruptionLayerSource) SetSourceEntity(sourceEntity string) *AntiCorruptionLayerSource {
	b.sourceEntity = sourceEntity
	return b
}

// SetTransformationEntity establece la entidad de transformación del evento.
func (b *AntiCorruptionLayerSource) SetTransformationEntity(transformationEntity string) *AntiCorruptionLayerSource {
	b.transformationEntity = transformationEntity
	return b
}

// Build construye y retorna el source final del evento.
func (b *AntiCorruptionLayerSource) Build() string {
	// Formato del source: [Original Source]/[Source Entity]/[ACL Identifier]/[Transformation Entity]
	return fmt.Sprintf("%s/%s/%s/%s", b.originalSource, b.sourceEntity, b.aclIdentifier, b.transformationEntity)
}
