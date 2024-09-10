package account

// Repository define las operaciones que un repositorio de cuentas debe implementar
type Repository interface {
	Save(a *Account) error
	FindByID(id int) (*Account, error)
}
