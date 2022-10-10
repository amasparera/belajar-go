package teslayer

type CategoryRepository interface {
	FindById(Id string) (*Category, error)
}
