package products

type ProductService struct {
	repo *ProductRepository
}

func NewProductService(repo *ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) Create(product *Product) error {
	return s.repo.Create(product)
}

func (s *ProductService) GetAll() ([]Product, error) {
	return s.repo.FindAll()
}

func (s *ProductService) GetByID(id uint) (*Product, error) {
	return s.repo.FindByID(id)
}

func (s *ProductService) Update(product *Product) error {
	return s.repo.Update(product)
}

func (s *ProductService) Delete(id uint) error {
	return s.repo.Delete(id)
}
