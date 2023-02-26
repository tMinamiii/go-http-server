package form

// User form
type User struct {
	Name   string  `json:"name"`
	Gender string  `json:"gender"`
	Height float64 `json:"height"`
	Weight float64 `json:"weight"`
}
