package book

import "fmt"

// why interface?

// this is contract
// type Repository interface{
// 	FindAll() 				([]Book, error)
// 	FindById(id int) 	(Book, error)
// 	Create(Book) 			(Book, error)
// }

// ================ IMPL ================
type RepositoryFileImpl struct {
}

func NewRepoFile() *RepositoryFileImpl {
	return &RepositoryFileImpl{}
}

func (repo *RepositoryFileImpl) FindAll() ([]Book, error) {
	var books []Book

	fmt.Println("FindAll")

	return books, nil
}

func (repo *RepositoryFileImpl) FindById(id int) (Book, error) {
	var book Book

	fmt.Println("FindById")

	return book, nil
}

func (repo *RepositoryFileImpl) Create(book Book) (Book, error) {
	fmt.Println("Create")

	return book, nil
}
