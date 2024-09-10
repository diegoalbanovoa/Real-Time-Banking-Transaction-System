package transaction

// Repository define las operaciones que un repositorio de transacciones debe implementar
type Repository interface {
	Save(t *Transaction) error
}
