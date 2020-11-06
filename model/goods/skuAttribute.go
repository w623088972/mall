package goods

import (
	"myself/mall/model"
)

type SkuAttribute struct {
	Id          int    `form:"id"           json:"id"           gorm:"primary_key" binding:"required"`
	GoodsId     int    `form:"goods_id"     json:"goods_id"     gorm:"column:goods_id"`
	AttributeId int    `form:"attribute_id" json:"attribute_id" gorm:"column:attribute_id"`
	CreatedAt   string `form:"created_at"   json:"created_at"   gorm:"column:created_at"`
	UpdatedAt   string `form:"updated_at"   json:"updated_at"   gorm:"column:updated_at"`
}

func (sa *SkuAttribute) TableName() string {
	return "sku_attribute"
}

type SkuAttributeRes struct {
	SpecId        int    `json:"spec_id"`
	SpecName      string `json:"spec_name"`
	AttributeId   int    `json:"attribute_id"`
	AttributeName string `json:"attribute_name"`
}

func SkuAttributeList(skuId int) ([]*SkuAttributeRes, error) {
	attributes := make([]*SkuAttributeRes, 0)
	err := model.DB.Self.Table("sku_attribute").Select("sku_attribute.attribute_id,attribute.name attribute_name,attribute.spec_id,spec.name spec_name").
		Joins("LEFT JOIN attribute ON attribute.id = sku_attribute.attribute_id").
		Joins("LEFT JOIN spec ON spec.id = attribute.spec_id").
		Where("sku_attribute.sku_id = ?", skuId).
		Find(&attributes).Error
	if err != nil {
		return nil, err
	}

	return attributes, nil
}
