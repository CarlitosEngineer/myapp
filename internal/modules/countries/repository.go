package countries

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, c *Country) error
	FindAll(ctx context.Context, offset, limit int) ([]Country, error)
	FindByID(ctx context.Context, id uint) (*Country, error)
	Update(ctx context.Context, c *Country) error
	Delete(ctx context.Context, id uint) error
}

type gormRepo struct{ db *gorm.DB }

func NewRepository(db *gorm.DB) Repository { return &gormRepo{db: db} }

func (r *gormRepo) Create(ctx context.Context, c *Country) error {
	return r.db.WithContext(ctx).Create(c).Error
}
func (r *gormRepo) FindAll(ctx context.Context, offset, limit int) ([]Country, error) {
	var out []Country
	err := r.db.WithContext(ctx).Offset(offset).Limit(limit).Order("id asc").Find(&out).Error
	return out, err
}
func (r *gormRepo) FindByID(ctx context.Context, id uint) (*Country, error) {
	var c Country
	if err := r.db.WithContext(ctx).First(&c, id).Error; err != nil {
		return nil, err
	}
	return &c, nil
}
func (r *gormRepo) Update(ctx context.Context, c *Country) error {
	return r.db.WithContext(ctx).Save(c).Error
}
func (r *gormRepo) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&Country{}, id).Error
}
