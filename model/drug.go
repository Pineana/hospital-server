package model

type Drug struct {
	Id          int    `gorm:"column:id;primary_key" json:"id"`                // 主键
	Name        string `gorm:"column:name;NOT NULL" json:"name"`               // 药品名称
	Stock       string `gorm:"column:stock;NOT NULL" json:"stock"`             // 库存
	Price       string `gorm:"column:price;NOT NULL" json:"price"`             // 价格
	Type        string `gorm:"column:type;NOT NULL" json:"type"`               // 药品类型
	Description string `gorm:"column:description;NOT NULL" json:"description"` // 药品说明
}
