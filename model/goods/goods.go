package goods

import (
	"myself/mall/model"
)

type Goods struct {
	Id         int    `form:"id"          json:"id"          gorm:"primary_key" binding:"required" title:"序号"`
	Name       string `form:"name"        json:"name"        gorm:"column:name"        title:"商品名称"`
	CategoryId int    `form:"category_id" json:"category_id" gorm:"column:category_id" title:"分类id"`
	CreatedAt  string `form:"created_at"  json:"created_at"  gorm:"column:created_at"  title:"创建时间"`
	UpdatedAt  string `form:"updated_at"  json:"updated_at"  gorm:"column:updated_at"  title:"修改时间"`
}

func (g *Goods) TableName() string {
	return "goods"
}

func GoodsInfo(goodsId int) (*Goods, error) {
	goods := &Goods{}
	err := model.DB.Self.Table("goods").Where("id = ?", goodsId).First(&goods).Error
	if err != nil {
		return nil, err
	}

	return goods, nil
}
