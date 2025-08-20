package countries

import (
	"context"
	"errors"
	"strings"
)

type Service interface {
	Create(ctx context.Context, in CreateCountryDTO) (*Country, error)
	List(ctx context.Context, page, pageSize int) ([]Country, error)
	Get(ctx context.Context, id uint) (*Country, error)
	Update(ctx context.Context, id uint, in UpdateCountryDTO) (*Country, error)
	Delete(ctx context.Context, id uint) error
}

type service struct{ repo Repository }

func NewService(r Repository) Service { return &service{repo: r} }

func (s *service) Create(ctx context.Context, in CreateCountryDTO) (*Country, error) {
	in.Name = strings.TrimSpace(in.Name)
	in.Code = strings.ToUpper(strings.TrimSpace(in.Code))
	if in.Name == "" || len(in.Code) != 2 {
		return nil, errors.New("invalid payload")
	}
	c := &Country{Name: in.Name, Code: in.Code}
	if err := s.repo.Create(ctx, c); err != nil {
		return nil, err
	}
	return c, nil
}

func (s *service) List(ctx context.Context, page, pageSize int) ([]Country, error) {
	if page < 1 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize
	return s.repo.FindAll(ctx, offset, pageSize)
}

func (s *service) Get(ctx context.Context, id uint) (*Country, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *service) Update(ctx context.Context, id uint, in UpdateCountryDTO) (*Country, error) {
	c, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if n := strings.TrimSpace(in.Name); n != "" {
		c.Name = n
	}
	if code := strings.TrimSpace(in.Code); code != "" {
		if len(code) != 2 {
			return nil, errors.New("code must be 2 letters")
		}
		c.Code = strings.ToUpper(code)
	}
	if err := s.repo.Update(ctx, c); err != nil {
		return nil, err
	}
	return c, nil
}

func (s *service) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}
