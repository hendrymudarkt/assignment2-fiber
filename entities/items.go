package entities

type Items struct {
	Item_id     int    `gorm:"primaryKey;" json:"lineItemId"`
	Item_code   string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	Order_id    int    `json:"-"`
}