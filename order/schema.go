package order

import "time"

type Order struct {
	ID         int64     `gorm:"primary_key" json:"id"`
	Country    string    `gorm:"not null;check:country IN ('lt', 'us', 'de')" json:"country"`
	ProxyCount int64     `gorm:"not null;check:proxy_count <= 100" json:"proxy_count"`
	Name       string    `gorm:"not null;type:varchar(30)" json:"name"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}
