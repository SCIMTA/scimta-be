package responses

type UserResponse struct {
	User struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
}
