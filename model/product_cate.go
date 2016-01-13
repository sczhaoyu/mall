package model

import (
	"errors"
	"fmt"
	"time"
)

type ProductCate struct {
	Id            int64     `json:"id"`
	ProductTypeId int       `json:"productTypeId"` //类型ID
	Pid           int       `json:"pid"`           //父类ID
	Name          string    `json:"name"`          //分类名称
	Path          string    `json:"path"`          //上级分类路径
	Scaling       float64   `json:"scaling"`       //分佣比例(商家给平台的分佣比例，填写1到100的数字)
	CreateId      int       `json:"createId"`      //创建人
	UpdateId      int       `json:"updateId"`      //更新人
	CreateTime    time.Time `json:"createTime"`    //创建时间
	UpdateTime    time.Time `json:"updateTime"`    //更新时间
	Sort          int       `json:"sort"`          //排序
	Status        int       `json:"status"`        //0、默认；1、提交审核；2、审核通过；3、审核失败；4、停用
}

func (p *ProductCate) Save() error {
	_, err := ShopDB.Insert(p)
	return err
}

func (p *ProductCate) Delete() error {
	_, err := ShopDB.Where("id=?", p.Id).Delete(ProductCate{})
	return err
}
func GetProductCateById(id int64) (*ProductCate, error) {
	var ret ProductCate
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("ProductCate Not Found Value: %v", id))
	}
	return &ret, nil
}

func (p *ProductCate) Update() error {
	_, err := ShopDB.Where("id=?", p.Id).Update(p)
	return err
}
