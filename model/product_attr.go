package model

import (
	"errors"
	"fmt"
)

type ProductAttr struct {
	Id                int64  `json:"id"`
	ProductId         int    `json:"productId"`         //商品ID
	ProductTypeAttrId int    `json:"productTypeAttrId"` //属性ID，可以根据属性ID来进行检索
	Name              string `json:"name"`              //属性名称
	Value             string `json:"value"`             //属性值
	State             int    `json:"state"`             //1、检索属性；2、自定义属性；3、商品自定义属性（product_tyep_attr_id=0）
}

func (p *ProductAttr) Save() error {
	_, err := ShopDB.Insert(p)
	return err
}

func (p *ProductAttr) Delete() error {
	_, err := ShopDB.Where("id=?", p.Id).Delete(ProductAttr{})
	return err
}
func GetProductAttrById(id int64) (*ProductAttr, error) {
	var ret ProductAttr
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("ProductAttr Not Found Value: %v", id))
	}
	return &ret, nil
}

func (p *ProductAttr) Update() error {
	_, err := ShopDB.Where("id=?", p.Id).Update(p)
	return err
}
