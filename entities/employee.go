package entities

type Employee struct {
	ID         int    `json:"-"` // Struct tags
	Name       string `json:"name"`
	Department string `json:"speciality"`
	ProjectID  int    `json:"project"`
}
