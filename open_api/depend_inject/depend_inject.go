package depend_inject

type SimpleRepository struct {

}

func NewSimpleRepository() *SimpleRepository {
	return &SimpleRepository{}
}

type SimpleService struct {
	*SimpleRepository
}

func NewSimpleService(repo *SimpleRepository) *SimpleService {
	return &SimpleService{
		SimpleRepository: repo,
	}
}
