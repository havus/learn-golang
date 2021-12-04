package depend_inject

type FooRepository struct {}

func NewFooRepo() *FooRepository {
	return &FooRepository{}
}

type FooService struct {
	FooRepository *FooRepository
}

func NewFooService(repo *FooRepository) *FooService {
	return &FooService{
		FooRepository: repo,
	}
}