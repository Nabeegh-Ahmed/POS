package models

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Collection struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Items       []Item    `json:"items" gorm:"foreignKey:CollectionId"`
	StoredAt    time.Time `json:"storedAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (collection *Collection) Validate() error {
	// Check all fields are filled
	if collection.Name == "" || collection.Description == "" {
		return errors.New("missing required fields")
	}
	return nil
}

func (collection *Collection) BeforeCreate(db *gorm.DB) error {
	collection.ID = uuid.New()
	collection.StoredAt = time.Now()
	return nil
}

// BeforeSave hook runs automatically before saving the exchange collection
func (collection *Collection) BeforeSave(db *gorm.DB) error {
	collection.UpdatedAt = time.Now()
	return nil
}

func (collection *Collection) CreateCollection(db *gorm.DB) (*Collection, error) {
	err := collection.Validate()
	if err != nil {
		return nil, err
	}

	err = db.Debug().Create(&collection).Error
	return collection, err
}

func (collection *Collection) GetCollections(offset, limit int, db *gorm.DB) ([]Collection, error) {
	var collections []Collection
	err := db.Debug().Preload("Items").Find(&collections).Offset(offset).Limit(limit).Error
	if err != nil {
		return nil, err
	}
	return collections, nil
}

func (collection *Collection) GetCollection(id string, db *gorm.DB) (*Collection, error) {
	err := db.Debug().Model(&Collection{}).Preload("Items").Where("id = ?", id).Take(&collection).Error
	return collection, err
}

func (collection *Collection) UpdateCollection(collectionData Collection, db *gorm.DB) (*Collection, error) {
	updateCollection := make(map[string]string)
	if collectionData.Name != "" {
		updateCollection["name"] = collectionData.Name
	}
	if collectionData.Description != "" {
		updateCollection["description"] = collectionData.Description
	}

	err := db.Debug().Model(&Collection{}).Where("id = ?", collection.ID).Updates(updateCollection).Error
	return collection, err
}
