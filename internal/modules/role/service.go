package role

type Service interface {
	GetByID(id int) (Role, error)
	Store(request CreateRole) (Role, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s service) GetByID(id int) (Role, error) {
	role, err := s.repository.GetByID(id)
	if err != nil {
		return Role{}, err
	}

	return role, nil
}

func (s service) Store(request CreateRole) (Role, error) {
	var role Role
	role.ID = request.ID
	role.Name = request.Name

	role, err := s.repository.Store(role)
	if err != nil {
		return Role{}, err
	}

	return role, nil
}
