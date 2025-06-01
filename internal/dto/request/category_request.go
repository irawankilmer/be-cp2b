package request

type CategoryRequest struct {
	Name         string `json:"name" binding:"required"`
	Type         string `json:"type" binding:"required,oneof=pemasukan pengeluaran pindah"`
	Descriptions string `json:"descriptions" binding:"omitempty"`
}
