package usecases

type UseCaseInterface interface {
	CreateUser(name, email string) (int64, error)
	UpdateUser(id int64, name, email string) (int64, error)
}
