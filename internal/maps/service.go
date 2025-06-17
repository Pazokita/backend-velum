package maps

type Service interface {
	GetMaps() ([]MapMetadata, error)
	GetMap(id int) (*MapMetadata, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetMaps() ([]MapMetadata, error) {
	return s.repo.GetAllMaps()
}

func (s *service) GetMap(id int) (*MapMetadata, error) {
	return s.repo.GetMapByID(id)
}
