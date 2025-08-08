package product

import (
	"oms-test/database"
	"oms-test/models"

	"gorm.io/gorm"
)

type ProductRepo struct{}

func (repo ProductRepo) GetProduct(id string) (*models.Product, error) {
	var product models.Product
	result := database.DB.First(&product, id)

	if result.Error != nil {
		return &models.Product{}, result.Error
	}
	return &product, nil
}

func (repo ProductRepo) SearchProducts(nameOrCategory string) (*[]models.Product, error) {
	var products []models.Product
	likeQuery := "%" + nameOrCategory + "%"
	result := database.DB.Where("name ILIKE ? OR category ILIKE ?", likeQuery, likeQuery).
		Find(&products)

	if result.Error != nil {
		return &[]models.Product{}, result.Error
	}
	return &products, nil
}

func (repo ProductRepo) CreateProduct(product *models.Product) {
	database.DB.Create(&product)
}

func (repo ProductRepo) CheckProductUniquenessExcludingSkuId(skuId *string, id *uint) bool {
	var product models.Product
	result := database.DB.Where("sku_id = ? AND id != ?", skuId, id).Find(&product)
	return result.Error != nil
}

func (repo ProductRepo) UpdateProduct(newProduct, oldProduct *models.Product) {
	database.DB.Model(oldProduct).Updates(*newProduct)
}

func (repo ProductRepo) DeleteProduct(id uint) {
	database.DB.Delete(&models.Product{}, id)
}

func (repo ProductRepo) InflowProduct(product *models.Product, quantity uint) error {
	result := database.DB.Model(&models.Product{}).
		Where("id = ?", product.ID).
		Update("quantity_available", gorm.Expr("quantity_available + ?", quantity))

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo ProductRepo) OutflowProduct(product *models.Product, quantity uint) error {
	result := database.DB.Model(&models.Product{}).
		Where("id = ? and quantity_available >= ?", product.ID, quantity).
		Update("quantity_available", gorm.Expr("quantity_available - ?", quantity))

	if result.Error != nil {
		return result.Error
	}
	return nil
}
