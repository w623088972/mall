package goods

import (
	"myself/mall/model"
)

type GoodsAttribute struct {
	Id          int    `form:"id"           json:"id"           gorm:"primary_key" binding:"required"`
	GoodsId     int    `form:"goods_id"     json:"goods_id"     gorm:"column:goods_id"`
	AttributeId int    `form:"attribute_id" json:"attribute_id" gorm:"column:attribute_id"`
	CreatedAt   string `form:"created_at"   json:"created_at"   gorm:"column:created_at"`
	UpdatedAt   string `form:"updated_at"   json:"updated_at"   gorm:"column:updated_at"`
}

func (ga *GoodsAttribute) TableName() string {
	return "goods_attribute"
}

type GoodsAttributeRes struct {
	GoodsAttribute
	Name string `json:"name"`
}

func GoodsAttributeList(goodsId, specId int) ([]*GoodsAttributeRes, error) {
	attributes := make([]*GoodsAttributeRes, 0)
	err := model.DB.Self.Table("goods_attribute").Select("goods_attribute.*,attribute.name").
		Joins("LEFT JOIN attribute ON attribute.id = goods_attribute.attribute_id").
		Where("goods_attribute.goods_id = ?", goodsId).
		Where("attribute.spec_id = ?", specId).
		Find(&attributes).Error
	if err != nil {
		return nil, err
	}

	return attributes, nil
}
