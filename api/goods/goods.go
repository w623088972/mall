package goods

import (
	"github.com/gin-gonic/gin"
	"myself/mall/conf"
	goodsM "myself/mall/model/goods"
	"strconv"
)

func GoodsInfo(c *gin.Context) {
	goodsId, _ := strconv.Atoi(c.Query("goods_id"))

	//获取商品详情
	goods, err := goodsM.GoodsInfo(goodsId)
	if err != nil {
		conf.SendResponse(c, nil, "GoodsInfo goodsM.GoodsInfo failed. err is "+err.Error(), nil, "chs")
		return
	}

	//获取商品规格
	goodsSpecs, err := goodsM.GoodsSpecList(goodsId)
	if err != nil {
		conf.SendResponse(c, nil, "GoodsInfo goodsM.GoodsSpecList failed. err is "+err.Error(), nil, "chs")
		return
	}
	for _, val := range goodsSpecs {
		//获取商品属性
		goodsAttributes, err := goodsM.GoodsAttributeList(goodsId, val.SpecId)
		if err != nil {
			conf.SendResponse(c, nil, "GoodsInfo goodsM.GoodsAttributeList failed. err is "+err.Error(), nil, "chs")
			return
		}
		val.GoodsAttributes = goodsAttributes
	}

	//获取商品sku
	sku, err := goodsM.SkuList(goodsId)
	if err != nil {
		conf.SendResponse(c, nil, "GoodsInfo goodsM.GoodsSpecList failed. err is "+err.Error(), nil, "chs")
		return
	}
	for _, val := range sku {
		//获取商品sku属性
		skuAttributes, err := goodsM.SkuAttributeList(val.Id)
		if err != nil {
			conf.SendResponse(c, nil, "GoodsInfo goodsM.SkuAttributeList failed. err is "+err.Error(), nil, "chs")
			return
		}
		val.Combines = skuAttributes
	}

	data := make(map[string]interface{})
	data["goods_detail"] = goods
	data["goods_specs"] = goodsSpecs
	data["goods_skus"] = sku

	conf.SendResponse(c, nil, "", data, "chs")
}
