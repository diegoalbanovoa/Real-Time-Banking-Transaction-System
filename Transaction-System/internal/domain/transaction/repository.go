package transaction

// Repository define las operaciones que un repositorio de transacciones debe implementar.
// Este patrón de diseño, conocido como "Repository Pattern", desacopla la lógica de negocio
// de la capa de persistencia, facilitando la mantenibilidad y el testeo.
type Repository interface {
	// Save guarda una transacción en el repositorio.
	// Recibe una transacción (Transaction) como argumento.
	// Retorna un error si ocurre algún problema al guardar la transacción.
	Save(t *Transaction) error
}
