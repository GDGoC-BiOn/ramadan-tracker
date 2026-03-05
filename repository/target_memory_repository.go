package repository

import (
	"errors"
	"ramadan-tracker-bts/models"
	"sync"
)

// TargetMemoryRepository adalah implementasi in-memory dari TargetRepository
type TargetMemoryRepository struct {
	data []models.Target
	mu   sync.RWMutex // Untuk thread-safety
}

// NewTargetMemoryRepository membuat instance baru TargetMemoryRepository
func NewTargetMemoryRepository() *TargetMemoryRepository {
	return &TargetMemoryRepository{
		data: []models.Target{
			{ID: "1", Ibadah: "Tarawih 30 Hari", Status: "Proses"},
			{ID: "2", Ibadah: "Baca Al-Quran 1 Juz", Status: "Pending"},
		},
	}
}

// FindAll mengembalikan semua target
func (r *TargetMemoryRepository) FindAll() ([]models.Target, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.data, nil
}

// FindByID mencari target berdasarkan ID
func (r *TargetMemoryRepository) FindByID(id string) (*models.Target, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, t := range r.data {
		if t.ID == id {
			return &t, nil
		}
	}
	return nil, errors.New("target tidak ditemukan")
}

// Create menambahkan target baru
func (r *TargetMemoryRepository) Create(target models.Target) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Cek apakah ID sudah ada
	for _, t := range r.data {
		if t.ID == target.ID {
			return errors.New("ID sudah digunakan")
		}
	}

	r.data = append(r.data, target)
	return nil
}

// Update memperbarui target yang ada
func (r *TargetMemoryRepository) Update(id string, target models.Target) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, t := range r.data {
		if t.ID == id {
			r.data[i].Ibadah = target.Ibadah
			r.data[i].Status = target.Status
			return nil
		}
	}
	return errors.New("target tidak ditemukan")
}

// Delete menghapus target
func (r *TargetMemoryRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, t := range r.data {
		if t.ID == id {
			r.data = append(r.data[:i], r.data[i+1:]...)
			return nil
		}
	}
	return errors.New("target tidak ditemukan")
}
