package goods

import (
	"myself/mall/model"
)

type GoodsSpec struct {
	Id        int    `form:"id"         json:"id"         gorm:"primary_key" binding:"required"`
	GoodsId   int    `form:"goods_id"   json:"goods_id"   gorm:"column:goods_id"`
	SpecId    int    `form:"spec_id"    json:"spec_id"    gorm:"column:spec_id"`
	CreatedAt string `form:"created_at" json:"created_at" gorm:"column:created_at"`
	UpdatedAt string `form:"updated_at" json:"updated_at" gorm:"column:updated_at"`
}

func (gs *GoodsSpec) TableName() string {
	return "goods_spec"
}

type GoodsSpecRes struct {
	GoodsSpec
	Name            string               `json:"name"`
	GoodsAttributes []*GoodsAttributeRes `json:"goods_attributes"`
}

func GoodsSpecList(goodsId int) ([]*GoodsSpecRes, error) {
	specs := make([]*GoodsSpecRes, 0)
	err := model.DB.Self.Table("goods_spec").Select("goods_spec.*,spec.name").
		Joins("LEFT JOIN spec ON spec.id = goods_spec.spec_id").
		Where("goods_spec.goods_id = ?", goodsId).Find(&specs).Error
	if err != nil {
		return nil, err
	}

	return specs, nil
}
