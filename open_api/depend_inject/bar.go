package depend_inject

type BarRepository struct {}

func NewBarRepo() *BarRepository {
	return &BarRepository{}
}

type BarService struct {
	*BarRepository
}

func NewBarService(repo *BarRepository) *BarService {
	return &BarService{
		BarRepository: repo,
	}
}