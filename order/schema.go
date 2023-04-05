package order

type Order struct {
	ID         int64  `gorm:"primary_key" json:"id"`
	ProxyCount int64  `gorm:"not null;check:proxy_count <= 100" json:"proxy_count"`
	Name       string `gorm:"type:varchar(30)" json:"name"`
}
