package model

import (
	"errors"
	"fmt"
	"time"
)

type ProductTypeAttr struct {
	Id            int64     `json:"id"`
	Name          string    `json:"name"`          //属性名称
	Value         string    `json:"value"`         //属性值，用逗号隔开
	ProductTypeId int       `json:"productTypeId"` //属性分类id
	Type          int       `json:"type"`          //1、检索属性；2、自定义属性
	CreateId      int       `json:"createId"`      //创建人ID
	CreateTime    time.Time `json:"createTime"`    //创建时间
	Sort          int       `json:"sort"`          //排序0到255，越大越靠前展示
}

func (p *ProductTypeAttr) Save() error {
	_, err := ShopDB.Insert(p)
	return err
}

func (p *ProductTypeAttr) Delete() error {
	_, err := ShopDB.Where("id=?", p.Id).Delete(ProductTypeAttr{})
	return err
}
func GetProductTypeAttrById(id int64) (*ProductTypeAttr, error) {
	var ret ProductTypeAttr
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("ProductTypeAttr Not Found Value: %v", id))
	}
	return &ret, nil
}

func (p *ProductTypeAttr) Update() error {
	_, err := ShopDB.Where("id=?", p.Id).Update(p)
	return err
}
