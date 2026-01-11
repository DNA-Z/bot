package product

type ProductService struct{}

func NewProductService() *ProductService {
	return &ProductService{}
}

func (s *ProductService) List() []Product {
	return allProducts
}
