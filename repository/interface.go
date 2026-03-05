package repository

import "ramadan-tracker-bts/models"

// TargetRepository adalah interface untuk akses data Target
// Interface ini bisa diimplementasikan oleh berbagai storage:
// - In-memory (untuk demo/testing)
// - PostgreSQL (production)
// - MongoDB, Redis, dll.
type TargetRepository interface {
	FindAll() ([]models.Target, error)
	FindByID(id string) (*models.Target, error)
	Create(target models.Target) error
	Update(id string, target models.Target) error
	Delete(id string) error
}
