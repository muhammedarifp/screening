package repo

type DataBaseInterface interface {
	CreateNewUser(name, email string) (int64, error)
	UpdateUser(id int64, name, email string) (int64, error)
}
