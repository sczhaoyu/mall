package model

import (
	"errors"
	"fmt"
)

type ProductLookLog struct {
	Id         int64  `json:"id"`
	SiteCookie string `json:"siteCookie"` //cookie埋点
	MemberId   int    `json:"memberId"`   //用户ID，没有登录ID为0
	ProductId  int    `json:"productId"`  //商品ID
	CreateTime string `json:"createTime"` //访问时间
}

func (p *ProductLookLog) Save() error {
	_, err := analysisDB.Insert(p)
	return err
}

func (p *ProductLookLog) Delete() error {
	_, err := analysisDB.Where("id=?", p.Id).Delete(ProductLookLog{})
	return err
}
func GetProductLookLogById(id int64) (*ProductLookLog, error) {
	var ret ProductLookLog
	b, err := analysisDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("ProductLookLog Not Found Value: %v", id))
	}
	return &ret, nil
}

func (p *ProductLookLog) Update() error {
	_, err := analysisDB.Where("id=?", p.Id).Update(p)
	return err
}
