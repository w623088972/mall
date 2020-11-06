package goods

import (
	"github.com/shopspring/decimal"
	"myself/mall/model"
)

type Sku struct {
	Id        int             `form:"id"         json:"id"         gorm:"primary_key" binding:"required"`
	GoodsId   int             `form:"goods_id"   json:"goods_id"   gorm:"column:goods_id"`
	Name      string          `form:"name"       json:"name"       gorm:"column:name"`
	Price     decimal.Decimal `form:"price"      json:"price"      gorm:"column:price"`
	Stock     int             `form:"stock"      json:"stock"      gorm:"column:stock"`
	CreatedAt string          `form:"created_at" json:"created_at" gorm:"column:created_at"`
	UpdatedAt string          `form:"updated_at" json:"updated_at" gorm:"column:updated_at"`
}

func (s *Sku) TableName() string {
	return "sku"
}

type SkuRes struct {
	Sku
	Combines []*SkuAttributeRes `json:"combines"`
}

func SkuList(goodsId int) ([]*SkuRes, error) {
	sku := make([]*SkuRes, 0)
	err := model.DB.Self.Table("sku").Where("goods_id = ?", goodsId).Find(&sku).Error
	if err != nil {
		return nil, err
	}

	return sku, nil
}
