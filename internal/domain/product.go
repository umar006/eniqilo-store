package domain

import (
	"time"

	"github.com/google/uuid"
)

var (
	ProductCategoryClothing    = "Clothing"
	ProductCategoryAccessories = "Accessories"
	ProductCategoryFootware    = "Footware"
	ProductCategoryBeverages   = "Beverages"
)

var ProductCategory = []string{
	ProductCategoryClothing,
	ProductCategoryAccessories,
	ProductCategoryFootware,
	ProductCategoryBeverages,
}

type Product struct {
	ID          string    `db:"id"`
	CreatedAt   time.Time `db:"created_at"`
	Name        string    `db:"name"`
	Sku         string    `db:"sku"`
	Category    string    `db:"category"`
	ImageUrls   []string  `db:"image_urls"`
	Notes       string    `db:"notes"`
	Price       int       `db:"price"`
	Stock       int       `db:"stock"`
	Location    string    `db:"stock"`
	IsAvailable bool      `db:"is_available"`
}

type ProductRequest struct {
	Name        string   `json:"name"`
	Sku         string   `json:"sku"`
	Category    string   `json:"category"`
	ImageUrls   []string `json:"imageUrls"`
	Notes       string   `json:"notes"`
	Price       int      `json:"price"`
	Stock       int      `json:"stock"`
	Location    string   `json:"location"`
	IsAvailable bool     `json:"isAvailable"`
}

type CreateProductResponse struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
}

type ProductResponse struct {
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	Name        string    `json:"name"`
	Sku         string    `json:"sku"`
	Category    string    `json:"category"`
	ImageUrls   []string  `json:"imageUrls"`
	Notes       string    `json:"notes"`
	Price       int       `json:"price"`
	Stock       int       `json:"stock"`
	Location    string    `json:"location"`
	IsAvailable bool      `json:"isAvailable"`
}

type UpdateProductResponse struct {
	ID        string    `json:"id"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (pr *ProductRequest) NewProduct() Product {
	id := uuid.New()
	rawCreatedAt := time.Now().Format(time.RFC3339)
	createdAt, _ := time.Parse(time.RFC3339, rawCreatedAt)

	return Product{
		ID:          id.String(),
		CreatedAt:   createdAt,
		Name:        pr.Name,
		Sku:         pr.Sku,
		Category:    pr.Category,
		ImageUrls:   pr.ImageUrls,
		Notes:       pr.Notes,
		Price:       pr.Price,
		Stock:       pr.Stock,
		Location:    pr.Location,
		IsAvailable: pr.IsAvailable,
	}
}