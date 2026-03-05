package service

import "ramadan-tracker-bts/models"

type TargetServiceInterface interface {
	GetAll() ([]models.Target, error)
	GetByID(id string) (*models.Target, error)
	Create(target models.Target) error
	Update(id string, target models.Target) error
	Delete(id string) error
}
