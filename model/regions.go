package model

import (
	"errors"
	"fmt"
)

type Regions struct {
	Id          int64  `json:"id"`
	ParentId    int    `json:"parentId"`
	RegionName  string `json:"regionName"`
	ParentPath  string `json:"parentPath"`
	FirstLetter string `json:"firstLetter"`
	RegionType  int    `json:"regionType"`
	AgencyId    int    `json:"agencyId"`
	ShippingId  int    `json:"shippingId"` //市场级别
	Visible     int    `json:"visible"`    //是否可见
	RowId       string `json:"rowId"`      //HP地址库映射字段
}

func (r *Regions) Save() error {
	_, err := ShopDB.Insert(r)
	return err
}

func (r *Regions) Delete() error {
	_, err := ShopDB.Where("id=?", r.Id).Delete(Regions{})
	return err
}
func GetRegionsById(id int64) (*Regions, error) {
	var ret Regions
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("Regions Not Found Value: %v", id))
	}
	return &ret, nil
}

func (r *Regions) Update() error {
	_, err := ShopDB.Where("id=?", r.Id).Update(r)
	return err
}
