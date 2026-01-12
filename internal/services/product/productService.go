package product

type ProductService struct{}

func NewProductService() *ProductService {
	return &ProductService{}
}

func (s *ProductService) List() []Product {
	return allProducts
}

func (s *ProductService) Get(idx int) (*Product, error) {
	return &allProducts[idx], nil
}
