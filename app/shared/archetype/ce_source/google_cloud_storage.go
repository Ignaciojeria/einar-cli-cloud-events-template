package ce_source

type GoogleCloudStorageBucketSource string

type GoogleCloudStorageObjectIntermediate struct {
	source GoogleCloudStorageBucketSource
}

// NewGoogleCloudStorageBucketSource inicia la construcción con el nombre del bucket.
func NewGoogleCloudStorageBucketSource(bucketName string) GoogleCloudStorageBucketSource {
	return GoogleCloudStorageBucketSource("gs://" + bucketName)
}

// WithObject agrega un objeto (archivo) al bucket y retorna un GoogleCloudStorageObjectIntermediate.
func (bs GoogleCloudStorageBucketSource) WithObject(objectName string) GoogleCloudStorageObjectIntermediate {
	updatedSource := GoogleCloudStorageBucketSource(string(bs) + "/" + objectName)
	return GoogleCloudStorageObjectIntermediate{source: updatedSource}
}

// Build finaliza la construcción y solo está disponible en GoogleCloudStorageObjectIntermediate.
func (oi GoogleCloudStorageObjectIntermediate) Build() string {
	return string(oi.source)
}
