package models

type Product struct {
	AbstractCreateUpdateModel
	Name              string  `json:"username"           binding:"required"`
	Category          string  `json:"email"              binding:"required"`
	Price             float32 `json:"name"               binding:"required"`
	QuantityAvailable uint    `json:"quantity_available" binding:"required"`
	SkuID             string  `json:"sku_id"             binding:"required" gorm:"unique"`
}
