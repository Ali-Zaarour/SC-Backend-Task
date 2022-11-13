package payload

type NewStudent struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
