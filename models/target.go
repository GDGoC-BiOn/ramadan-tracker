package models

// Target merepresentasikan satu target ibadah harian
type Target struct {
	ID     string `json:"id"`
	Ibadah string `json:"ibadah"`
	Status string `json:"status"` // "Proses", "Selesai", "Pending"
}

// ValidStatuses adalah daftar status yang valid
var ValidStatuses = []string{"Proses", "Selesai", "Pending"}

// IsValidStatus memeriksa apakah status valid
func IsValidStatus(status string) bool {
	for _, v := range ValidStatuses {
		if v == status {
			return true
		}
	}
	return false
}

// Validate melakukan validasi data Target
func (t *Target) Validate() map[string]string {
	errors := make(map[string]string)

	if t.ID == "" {
		errors["id"] = "ID tidak boleh kosong"
	}

	if t.Ibadah == "" {
		errors["ibadah"] = "Ibadah tidak boleh kosong"
	}

	if t.Status == "" {
		errors["status"] = "Status tidak boleh kosong"
	} else if !IsValidStatus(t.Status) {
		errors["status"] = "Status harus: Proses, Selesai, atau Pending"
	}

	return errors
}
