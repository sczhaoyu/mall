package model

import (
	"errors"
	"fmt"
	"time"
)

type SystemGoodsFloor struct {
	Id         int64     `json:"id"`
	Createtime time.Time `json:"createtime"` //创建时间
	Display    int       `json:"display"`    //是否显示 0-不显示 1-显示
	Name       string    `json:"name"`       //楼层名称
	Sort       int       `json:"sort"`       //楼层排序号
	LeftAdv    int       `json:"leftAdv"`    //左侧广告
	BrandList  string    `json:"brandList"`  //品牌 json格式：id-品牌id name-品牌名称
}

func (s *SystemGoodsFloor) Save() error {
	_, err := ShopDB.Insert(s)
	return err
}

func (s *SystemGoodsFloor) Delete() error {
	_, err := ShopDB.Where("id=?", s.Id).Delete(SystemGoodsFloor{})
	return err
}
func GetSystemGoodsFloorById(id int64) (*SystemGoodsFloor, error) {
	var ret SystemGoodsFloor
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("SystemGoodsFloor Not Found Value: %v", id))
	}
	return &ret, nil
}

func (s *SystemGoodsFloor) Update() error {
	_, err := ShopDB.Where("id=?", s.Id).Update(s)
	return err
}
