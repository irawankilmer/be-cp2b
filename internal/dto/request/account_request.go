package request

type AccountRequest struct {
	Name         string `json:"name" binding:"required"`
	Descriptions string `json:"descriptions" binding:"omitempty"`
}
