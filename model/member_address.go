package model

import (
	"errors"
	"fmt"
	"time"
)

type MemberAddress struct {
	Id          int64     `json:"id"`
	MemberId    int       `json:"memberId"`
	MemberName  string    `json:"memberName"`  //收货人
	ProvinceId  int       `json:"provinceId"`  //省ID
	CityId      int       `json:"cityId"`      //市ID
	AreaId      int       `json:"areaId"`      //地区ID
	AddAll      string    `json:"addAll"`      //省市区组合
	AddressInfo string    `json:"addressInfo"` //详细地址
	Mobile      string    `json:"mobile"`      //手机
	Phone       string    `json:"phone"`       //电话
	Email       string    `json:"email"`       //邮箱
	ZipCode     string    `json:"zipCode"`     //邮编
	State       int       `json:"state"`       //1、默认；2、不是默认
	CreateTime  time.Time `json:"createTime"`
	UpdateTime  time.Time `json:"updateTime"`
}

func (m *MemberAddress) Save() error {
	_, err := ShopDB.Insert(m)
	return err
}

func (m *MemberAddress) Delete() error {
	_, err := ShopDB.Where("id=?", m.Id).Delete(MemberAddress{})
	return err
}
func GetMemberAddressById(id int64) (*MemberAddress, error) {
	var ret MemberAddress
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("MemberAddress Not Found Value: %v", id))
	}
	return &ret, nil
}

func (m *MemberAddress) Update() error {
	_, err := ShopDB.Where("id=?", m.Id).Update(m)
	return err
}
