package handler

import (
	"ramadan-tracker-bts/models"
	"ramadan-tracker-bts/service"

	"github.com/gofiber/fiber/v2"
)

// TargetHandler menangani HTTP requests untuk Target
type TargetHandler struct {
	service service.TargetServiceInterface
}

// NewTargetHandler membuat instance baru TargetHandler
func NewTargetHandler(service service.TargetServiceInterface) *TargetHandler {
	return &TargetHandler{service: service}
}

// GetAll menangani GET /api/targets
func (h *TargetHandler) GetAll(c *fiber.Ctx) error {
	targets, err := h.service.GetAll()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":  "Gagal mengambil data",
			"detail": err.Error(),
		})
	}
	return c.JSON(fiber.Map{"data": targets})
}

// GetByID menangani GET /api/targets/:id
func (h *TargetHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")

	target, err := h.service.GetByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error":  "Data tidak ditemukan",
			"detail": err.Error(),
		})
	}

	return c.JSON(fiber.Map{"data": target})
}

// Create menangani POST /api/targets
func (h *TargetHandler) Create(c *fiber.Ctx) error {
	var newTarget models.Target

	// Parse JSON body
	if err := c.BodyParser(&newTarget); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":  "Format JSON tidak valid",
			"detail": err.Error(),
		})
	}

	// Validasi data
	if validationErrors := newTarget.Validate(); len(validationErrors) > 0 {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Validasi gagal",
			"details": validationErrors,
		})
	}

	// Simpan ke repository
	if err := h.service.Create(newTarget); err != nil {
		return c.Status(409).JSON(fiber.Map{
			"error":  "Gagal menyimpan data",
			"detail": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Target berhasil ditambahkan",
		"data":    newTarget,
	})
}

// Update menangani PUT /api/targets/:id
func (h *TargetHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")

	var updatedTarget models.Target

	if err := c.BodyParser(&updatedTarget); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":  "Format JSON tidak valid",
			"detail": err.Error(),
		})
	}

	// Validasi data (kecuali ID)
	updatedTarget.ID = id
	if validationErrors := updatedTarget.Validate(); len(validationErrors) > 0 {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Validasi gagal",
			"details": validationErrors,
		})
	}

	if err := h.service.Update(id, updatedTarget); err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error":  "Data tidak ditemukan",
			"detail": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Target berhasil diupdate",
		"data":    updatedTarget,
	})
}

// Delete menangani DELETE /api/targets/:id
func (h *TargetHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.service.Delete(id); err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error":  "Data tidak ditemukan",
			"detail": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Target berhasil dihapus",
	})
}
