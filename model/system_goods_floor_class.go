package model

import (
	"errors"
	"fmt"
)

type SystemGoodsFloorClass struct {
	Id          int64  `json:"id"`          //商品信息 json格式：pc_id-商品类型id pro_id-商品id
	ProductInfo string `json:"productInfo"` //
	FloorId     int    `json:"floorId"`     //楼层id
	Name        string `json:"name"`        //分类名称
}

func (s *SystemGoodsFloorClass) Save() error {
	_, err := ShopDB.Insert(s)
	return err
}

func (s *SystemGoodsFloorClass) Delete() error {
	_, err := ShopDB.Where("id=?", s.Id).Delete(SystemGoodsFloorClass{})
	return err
}
func GetSystemGoodsFloorClassById(id int64) (*SystemGoodsFloorClass, error) {
	var ret SystemGoodsFloorClass
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("SystemGoodsFloorClass Not Found Value: %v", id))
	}
	return &ret, nil
}

func (s *SystemGoodsFloorClass) Update() error {
	_, err := ShopDB.Where("id=?", s.Id).Update(s)
	return err
}
