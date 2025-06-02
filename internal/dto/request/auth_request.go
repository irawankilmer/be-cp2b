package request

type AuthRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (a *AuthRequest) Sanitize() map[string]any {
	return map[string]any{
		"email": a.Email,
	}
}
