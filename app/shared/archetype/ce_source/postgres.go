package ce_source

type PostgresSource string

type PostgresIntermediate struct {
	source PostgresSource
}

// NewPostgresSource inicia la construcción de la URL.
func NewPostgresSource(databaseName string) PostgresSource {
	return PostgresSource("postgresql://" + databaseName)
}

// WithTable agrega una tabla y retorna un PostgresIntermediate.
func (ps PostgresSource) WithTable(tableName string) PostgresIntermediate {
	updatedSource := PostgresSource(string(ps) + "/" + tableName)
	return PostgresIntermediate{source: updatedSource}
}

// Build solo está disponible en PostgresIntermediate (el tipo final).
func (pi PostgresIntermediate) Build() string {
	return string(pi.source)
}
