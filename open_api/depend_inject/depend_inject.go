package depend_inject

import "errors"

type SimpleRepository struct {
	Error bool
}

func NewSimpleRepository(isError bool) *SimpleRepository {
	return &SimpleRepository{
		Error: isError,
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
