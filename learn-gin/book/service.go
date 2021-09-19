package book

type Service interface {
	FindAll() 												([]Book, error)
	FindById(id int) 									(Book, error)
	Create(book BookRequest)					(Book, error)
	Update(id int, book BookRequest)	(Book, error)
	Delete(id int)										(Book, error)
}

type ServiceImpl struct {
	repository Repository
}

func NewService(repository Repository) *ServiceImpl {
	return &ServiceImpl{repository}
}

func (service *ServiceImpl) FindAll() ([]Book, error) {
	return service.repository.FindAll()
}

func (service *ServiceImpl) FindById(id int) (Book, error) {
	return service.repository.FindById(id)
	// book, err := service.repository.FindById(id)

	// return book, err
}

func (service *ServiceImpl) Create(bookRequest BookRequest) (Book, error) {
	price, _ 	:= bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()

	book := Book{
		Title: bookRequest.Title,
		Price: int(price),
		Rating: int(rating),
		Description: bookRequest.Description,
	}

	return service.repository.Create(book)
}

func (service *ServiceImpl) Update(id int, bookRequest BookRequest) (Book, error) {
	book, _ 	:= service.repository.FindById(id)
	price, _ 	:= bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()

	book.Title 				= bookRequest.Title
	book.Price 				= int(price)
	book.Rating 			= int(rating)
	book.Description 	= bookRequest.Description

	return service.repository.Update(book)
}

func (service *ServiceImpl) Delete(id int) (Book, error) {
	book, _ := service.repository.FindById(id)

	return service.repository.Delete(book)
}
