package response

type CategoryResponse struct {
	ID           uint   `json:"ID"`
	Name         string `json:"Name"`
	Type         string `json:"Type"`
	Descriptions string `json:"Descriptions"`
}
