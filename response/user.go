package response

// User response
type User struct {
	ID     int64   `json:"id"`
	Name   string  `json:"name"`
	Gender string  `json:"gender"`
	Height float64 `json:"height"`
	Weight float64 `json:"weight"`
}

// Users response
type Users []User
