package model

import (
	"errors"
	"fmt"
	"time"
)

type ProductType struct {
	Id              int64     `json:"id"`
	Name            string    `json:"name"`            //名称
	Sort            int       `json:"sort"`            //排序
	ProductNormIds  string    `json:"productNormIds"`  //关联规格ID集合
	ProductBrandIds string    `json:"productBrandIds"` //关联品牌ID集合
	CreateId        int       `json:"createId"`        //创建人
	CreateTime      time.Time `json:"createTime"`      //创建时间
}

func (p *ProductType) Save() error {
	_, err := ShopDB.Insert(p)
	return err
}

func (p *ProductType) Delete() error {
	_, err := ShopDB.Where("id=?", p.Id).Delete(ProductType{})
	return err
}
func GetProductTypeById(id int64) (*ProductType, error) {
	var ret ProductType
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("ProductType Not Found Value: %v", id))
	}
	return &ret, nil
}

func (p *ProductType) Update() error {
	_, err := ShopDB.Where("id=?", p.Id).Update(p)
	return err
}
