package entities

import "time"

type Orders struct {
	Order_id      int   	`gorm:"primaryKey;" json:"-"`
	Customer_name string 	`json:"customerName"`
	Ordered_at    time.Time `json:"orderedAt"`
	Items	  	  []Items 	`gorm:"foreignKey:Order_id;references:Order_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"items"`
}