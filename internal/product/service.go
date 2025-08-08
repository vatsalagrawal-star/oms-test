package product

import (
	"errors"

	"oms-test/models"
)

type ProductService struct {
	repo ProductRepo
}

func (ser *ProductService) getProduct(id string) (*models.Product, error) {
	product, err := ser.repo.GetProduct(id)
	if err != nil {
		return &models.Product{}, err
	}
	return product, nil
}

func (ser *ProductService) searchProducts(nameOrCategory string) (*[]models.Product, error) {
	products, err := ser.repo.SearchProducts(nameOrCategory)
	if err != nil {
		return &[]models.Product{}, err
	}
	return products, nil
}

func (ser *ProductService) createProduct(product *models.Product) (*models.Product, error) {
	if product.Name == "" || product.Category == "" || product.SkuID == "" {
		return &models.Product{}, errors.New("product name, category, and SKU ID cannot be empty")
	}
	ser.repo.CreateProduct(product)
	return product, nil
}

func (ser *ProductService) updateProduct(old, new *models.Product) error {
	if new.SkuID == "" {
		return errors.New("product SKU ID cannot be empty")
	}

	is_duplicate := ser.repo.CheckProductUniquenessExcludingSkuId(&new.SkuID, &old.ID)
	if is_duplicate {
		return errors.New("product with the same SKU ID already exists")
	}

	ser.repo.UpdateProduct(new, old)
	return nil
}

func (ser *ProductService) deleteProduct(idUint uint) {
	ser.repo.DeleteProduct(idUint)
}

func (ser *ProductService) inflowProduct(product *models.Product, quantity uint) error {
	err := ser.repo.InflowProduct(product, quantity)
	return err
}

func (ser *ProductService) outflowProduct(product *models.Product, quantity uint) error {
	if product.QuantityAvailable < quantity {
		return errors.New("not enough quantity available for outflow")
	}

	return ser.repo.OutflowProduct(product, quantity)
}

func (ser *ProductService) getAllProducts() (*[]models.Product, error) {
	products, err := ser.repo.GetAllProducts()
	if err != nil {
		return &[]models.Product{}, err
	}
	return products, nil
}
