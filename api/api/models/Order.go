package models

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Order struct {
	ID        uuid.UUID  `json:"id"`
	LineItems []LineItem `json:"line_items" gorm:"foreignKey:OrderId"`
	Price     float64    `json:"price"`
	StoredAt  time.Time  `json:"stored_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	Discount  float64    `json:"discount"`
	Tax       float64    `json:"tax"`
}

type LineItem struct {
	OrderId      uuid.UUID `json:"order"`
	ID           uuid.UUID `json:"id"`
	ItemId       uuid.UUID `json:"item_id"`
	Name         string    `json:"name"`
	Price        float64   `json:"price"`
	Cost         float64   `json:"cost"`
	Discount     float64   `json:"discount"`
	Quantity     uint      `json:"quantity"`
	Barcode      string    `json:"barcode"`
	CollectionId uuid.UUID `json:"collection"`
	StoredAt     time.Time `json:"storedAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

func (order *Order) Validate() error {
	// Check all fields are filled
	if len(order.LineItems) == 0 {
		return errors.New("missing required fields")
	}
	// check the total price is equal to the sum of the price of each product
	var totalPrice float64
	for _, item := range order.LineItems {
		totalPrice += item.Price * float64(item.Quantity)
	}
	if totalPrice != order.Price {
		order.Price = totalPrice
	}
	return nil
}

// BeforeSave hook runs automatically before saving the exchange order
func (order *Order) BeforeSave(db *gorm.DB) error {
	order.UpdatedAt = time.Now()
	return nil
}

func (order *Order) BeforeCreate(db *gorm.DB) error {
	order.StoredAt = time.Now()
	order.UpdatedAt = time.Now()
	order.ID = uuid.New()
	return nil
}

func (lineItem *LineItem) BeforeCreate(db *gorm.DB) error {
	lineItem.ID = uuid.New()
	return nil
}

func (order *Order) CreateOrder(db *gorm.DB) (*Order, error) {
	err := order.Validate()
	if err != nil {
		return nil, err
	}
	err = db.Debug().Create(&order).Error
	return order, err
}

func (order *Order) GetOrders(offset, limit int, db *gorm.DB) ([]Order, error) {
	var orders []Order
	err := db.Debug().Preload("LineItems").Find(&orders).Offset(offset).Limit(limit).Error
	return orders, err
}

func (order *Order) GetOrder(id string, db *gorm.DB) (*Order, error) {
	err := db.Debug().Model(&Order{}).Preload("LineItems").Where("id = ?", id).Take(&order).Error
	return order, err
}
