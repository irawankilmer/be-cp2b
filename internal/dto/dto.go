package dto

type AccountDTO struct {
	ID   uint   `json:"ID"`
	Name string `json:"Name"`
}

type CategoryDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type UserDTO struct {
	ID   uint   `json:"ID"`
	Name string `json:"Name"`
}
