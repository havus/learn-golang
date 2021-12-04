package depend_inject

import "errors"

type SimpleRepository struct {
	Error bool
}

func NewSimpleRepository() *SimpleRepository {
	return &SimpleRepository{
		Error: true,
	}
}

type SimpleService struct {
	*SimpleRepository
}

func NewSimpleService(repo *SimpleRepository) (*SimpleService, error) {
	if repo.Error {
		return nil, errors.New("Failed create service")
	} else {
		return &SimpleService{
			SimpleRepository: repo,
		}, nil
	}
}
