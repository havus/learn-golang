package book

import (
	"gorm.io/gorm"
)

type Repository interface{
	FindAll() 				([]Book, error)
	FindById(id int) 	(Book, error)
	Create(book Book)	(Book, error)
	Update(book Book)	(Book, error)
	Delete(book Book)	(Book, error)
}

// ================ IMPL ================
type RepositoryImpl struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *RepositoryImpl {
	return &RepositoryImpl{db}
}

func (repo *RepositoryImpl) FindAll() ([]Book, error) {
	var books []Book

	err := repo.db.Find(&books).Error

	return books, err
}

func (repo *RepositoryImpl) FindById(id int) (Book, error) {
	var book Book

	err := repo.db.First(&book, id).Error

	return book, err
}

func (repo *RepositoryImpl) Create(book Book) (Book, error) {
	err := repo.db.Create(&book).Error

	return book, err
}

func (repo *RepositoryImpl) Update(book Book) (Book, error) {
	err := repo.db.Save(&book).Error

	return book, err
}

func (repo *RepositoryImpl) Delete(book Book) (Book, error) {
	err := repo.db.Delete(&book).Error

	return book, err
}
