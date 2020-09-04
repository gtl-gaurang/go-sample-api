package models

// Users ... user table
type Users struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// TableName ... return tablename
func (b *Users) TableName() string {
	return "Users"
}
