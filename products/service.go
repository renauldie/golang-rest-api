package products

type Service interface {
	GetProducts(limit int, page int) ([]Product, int64, error)
	CreateProduct(request ProductRequest) (Product, error)
	GetProductById(id int) (Product, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) GetProducts(page int, limit int) ([]Product, int64, error) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}
	offset := (page - 1) * limit
	return s.repository.FindAll(limit, offset)
}

func (s *service) CreateProduct(request ProductRequest) (Product, error) {
	var product Product
	product.ID = request.ID
	product.Name = request.Name
	product.Price = request.Price
	product.Stock = request.Stock
	return s.repository.CreateProduct(product)
}

func (s *service) GetProductById(id int) (Product, error) {
	product, err := s.repository.FindById(id)
	if err != nil {
		return Product{}, err
	}

	return product, nil
}
