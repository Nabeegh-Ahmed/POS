package models

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Item struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Price        float64   `json:"price"`
	Cost         float64   `json:"cost"`
	Barcode      string    `json:"barcode"`
	CollectionId uuid.UUID `json:"collection"`
	StoredAt     time.Time `json:"storedAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

func (item *Item) Validate() error {
	// Check all fields are filled
	if item.Barcode == "" || item.Name == "" || item.Price == 0 || item.CollectionId == uuid.Nil {
		return errors.New("missing required fields")
	}
	return nil
}

// BeforeSave hook runs automatically before saving the exchange order
func (item *Item) BeforeSave(db *gorm.DB) error {
	item.UpdatedAt = time.Now()
	return nil
}

func (item *Item) BeforeCreate(db *gorm.DB) error {
	item.ID = uuid.New()
	item.StoredAt = time.Now()
	item.UpdatedAt = time.Now()
	return nil
}

func (item *Item) CreateItem(db *gorm.DB) (*Item, error) {
	err := item.Validate()
	if err != nil {
		return nil, err
	}
	err = db.Debug().Create(&item).Error
	return item, err
}

func (item *Item) GetItems(offset, limit int, db *gorm.DB) ([]Item, error) {
	var items []Item
	err := db.Debug().Find(&items).Offset(offset).Limit(limit).Error
	return items, err
}

func (item *Item) GetItem(id string, db *gorm.DB) (*Item, error) {
	err := db.Debug().Model(&Item{}).Where("id = ?", id).Take(&item).Error
	return item, err
}

func (item *Item) GetItemsByName(name string, db *gorm.DB) ([]Item, error) {
	var items []Item
	err := db.Debug().Model(&Item{}).Find(&items).Where("name like ?", "%"+name+"%").Error
	return items, err
}

func (item *Item) GetItemByBarcode(barcode string, db *gorm.DB) (*Item, error) {
	err := db.Debug().Model(&Item{}).Where("barcode = ?", barcode).Take(&item).Error
	return item, err
}

func (item *Item) UpdateItem(db *gorm.DB) (*Item, error) {
	updateItem := make(map[string]interface{})
	if item.Name != "" {
		updateItem["Name"] = item.Name
	}
	if item.Price != 0 {
		updateItem["Price"] = item.Price
	}
	if item.Cost != 0 {
		updateItem["Cost"] = item.Cost
	}
	if item.Barcode != "" {
		updateItem["Barcode"] = item.Barcode
	}

	err := db.Debug().Model(&Item{}).Where("id = ?", item.ID).Updates(item).Take(&item).Error
	return item, err
}
