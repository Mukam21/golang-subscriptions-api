package repository

import (
	"time"

	"golang-subscriptions-api/internal/model"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(sub *model.Subscription) error {
	return r.db.Create(sub).Error
}

func (r *Repository) GetAll() ([]model.Subscription, error) {
	var subs []model.Subscription
	err := r.db.Find(&subs).Error
	return subs, err
}

func (r *Repository) GetByID(id string) (*model.Subscription, error) {
	var sub model.Subscription
	err := r.db.First(&sub, "id = ?", id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &sub, err
}

func (r *Repository) Update(sub *model.Subscription) error {
	return r.db.Save(sub).Error
}

func (r *Repository) Delete(id string) error {
	return r.db.Delete(&model.Subscription{}, "id = ?", id).Error
}

// SumPrice calculates total price by filters
func (r *Repository) SumPrice(userID, serviceName string, from, to *time.Time) (int64, error) {
	query := r.db.Model(&model.Subscription{}).Select("COALESCE(SUM(price),0)")

	if userID != "" {
		query = query.Where("user_id = ?", userID)
	}
	if serviceName != "" {
		query = query.Where("service_name = ?", serviceName)
	}
	if from != nil {
		query = query.Where("start_date >= ?", *from)
	}
	if to != nil {
		query = query.Where("start_date <= ?", *to)
	}

	var total int64
	err := query.Scan(&total).Error
	return total, err
}
