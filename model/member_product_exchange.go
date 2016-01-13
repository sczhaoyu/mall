package model

import (
	"errors"
	"fmt"
	"time"
)

type MemberProductExchange struct {
	Id             int64     `json:"id"`
	Seller         int       `json:"seller"`         //所属商家
	OrderId        int       `json:"orderId"`        //订单ID
	OrderProductId int       `json:"orderProductId"` //网单ID
	ProductId      int       `json:"productId"`      //商品ID
	MemberId       int       `json:"memberId"`       //用户ID
	MemberName     string    `json:"memberName"`     //用户Name
	ProvinceId2    int       `json:"provinceId2"`    //Province
	CityId2        int       `json:"cityId2"`        //city
	AreaId2        int       `json:"areaId2"`        //area
	AddressAll2    string    `json:"addressAll2"`    //省市区组合
	AddressInfo2   string    `json:"addressInfo2"`   //详细地址
	Phone2         string    `json:"phone2"`
	ChangeName2    string    `json:"changeName2"` //换货收货人
	ProvinceId     int       `json:"provinceId"`  //Province
	CityId         int       `json:"cityId"`      //city
	AreaId         int       `json:"areaId"`      //area
	AddressAll     string    `json:"addressAll"`  //省市区组合
	AddressInfo    string    `json:"addressInfo"` //详细地址
	ChangeName     string    `json:"changeName"`  //收货人地址
	Phone          string    `json:"phone"`
	Email          string    `json:"email"`
	ZipCode        string    `json:"zipCode"`  //邮编
	Question       string    `json:"question"` //问题描述
	Image          string    `json:"image"`    //图片
	Name           string    `json:"name"`     //联系人姓名
	State          int       `json:"state"`    //换货状态：1、未处理；2、审核通过待收货；3、已经收货；4、发货处理完成；5、不予处理原件退还；6、不处理
	CreateTime     time.Time `json:"createTime"`
	UpdateTime     time.Time `json:"updateTime"`
	OptId          int       `json:"optId"`         //处理人
	Remark         string    `json:"remark"`        //备注
	AddressResult  string    `json:"addressResult"` //退货地址
}

func (m *MemberProductExchange) Save() error {
	_, err := ShopDB.Insert(m)
	return err
}

func (m *MemberProductExchange) Delete() error {
	_, err := ShopDB.Where("id=?", m.Id).Delete(MemberProductExchange{})
	return err
}
func GetMemberProductExchangeById(id int64) (*MemberProductExchange, error) {
	var ret MemberProductExchange
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("MemberProductExchange Not Found Value: %v", id))
	}
	return &ret, nil
}

func (m *MemberProductExchange) Update() error {
	_, err := ShopDB.Where("id=?", m.Id).Update(m)
	return err
}
