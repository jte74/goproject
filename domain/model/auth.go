package model

type Auth struct {
	// Name
	// in: string
	Name string `json:"name,omitempty"`
	// Password
	// in: string
	Password string `json:"password,omitempty"`
}
