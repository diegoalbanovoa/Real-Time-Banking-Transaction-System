package account

// Repository define las operaciones que un repositorio de cuentas debe implementar.
// Este patrón de diseño se llama "Repository Pattern" y permite desacoplar la lógica de negocio
// de la capa de persistencia (por ejemplo, una base de datos).
type Repository interface {
	// Save guarda una cuenta (Account) en el repositorio.
	// Retorna un error si no se puede realizar la operación.
	Save(a *Account) error

	// FindByID busca una cuenta por su ID único.
	// Retorna un puntero a la cuenta (Account) y un error si no se encuentra.
	FindByID(id int) (*Account, error)
}
