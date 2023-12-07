package ce_source

type FirestoreSource string

type FirestoreIntermediate struct {
	source FirestoreSource
}

// NewFirestoreSource inicia la construcción de la URL.
func NewFirestoreSource(projectID string) FirestoreSource {
	return FirestoreSource("https://firestore.googleapis.com/v1/projects/" + projectID + "/databases/(default)/documents")
}

// WithCollection agrega una colección y retorna un FirestoreIntermediate.
func (fs FirestoreSource) WithCollection(collection string) FirestoreIntermediate {
	updatedSource := FirestoreSource(string(fs) + "/" + collection)
	return FirestoreIntermediate{source: updatedSource}
}

// WithDocument agrega un documento y retorna un FirestoreSource (el tipo final que puede llamar a Build).
func (fi FirestoreIntermediate) WithDocument(document string) FirestoreSource {
	return FirestoreSource(string(fi.source) + "/" + document)
}

// Build solo está disponible en FirestoreSource (el tipo final).
func (fs FirestoreSource) Build() string {
	return string(fs)
}
