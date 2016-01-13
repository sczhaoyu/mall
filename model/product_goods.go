package model

import (
	"errors"
	"fmt"
)

type ProductGoods struct {
	Id                  int64   `json:"id"`
	ProductId           int     `json:"productId"`           //商品ID
	NormAttrId          string  `json:"normAttrId"`          //规格属性值ID，用逗号分隔
	NormName            string  `json:"normName"`            //规格值，用逗号分隔
	MallPcPrice         float64 `json:"mallPcPrice"`         //PC价格
	MallMobilePrice     float64 `json:"mallMobilePrice"`     //Mobile价格
	ProductStock        int     `json:"productStock"`        //库存
	ProductStockWarning int     `json:"productStockWarning"` //库存预警值
	ActualSales         int     `json:"actualSales"`         //所有规格销量相加等于商品总销量
	Sku                 string  `json:"sku"`
	Images              string  `json:"images"` //规格图片（规格图片，用逗号隔开，和规格值对应，如果没有图片，那为空）
}

func (p *ProductGoods) Save() error {
	_, err := ShopDB.Insert(p)
	return err
}

func (p *ProductGoods) Delete() error {
	_, err := ShopDB.Where("id=?", p.Id).Delete(ProductGoods{})
	return err
}
func GetProductGoodsById(id int64) (*ProductGoods, error) {
	var ret ProductGoods
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("ProductGoods Not Found Value: %v", id))
	}
	return &ret, nil
}

func (p *ProductGoods) Update() error {
	_, err := ShopDB.Where("id=?", p.Id).Update(p)
	return err
}
