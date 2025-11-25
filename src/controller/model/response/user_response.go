package response

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Age   int8   `json:"age"`
	Email string `json:"email"`
}