package model

import (
	"errors"
	"fmt"
	"time"
)

type MemberProductBack struct {
	Id             int64     `json:"id"`
	SellerId       int       `json:"sellerId"`       //所属商家
	OrderId        int       `json:"orderId"`        //订单ID
	OrderProductId int       `json:"orderProductId"` //网单ID
	ProductId      int       `json:"productId"`      //商品ID
	MemberId       int       `json:"memberId"`       //用户ID
	MemberName     string    `json:"memberName"`     //用户Name
	ProvinceId     int       `json:"provinceId"`     //Province
	CityId         int       `json:"cityId"`         //city
	AreaId         int       `json:"areaId"`         //area
	AddressAll     string    `json:"addressAll"`     //省市区组合
	AddressInfo    string    `json:"addressInfo"`    //详细地址
	Phone          string    `json:"phone"`          //退货人手机
	ReturnName     string    `json:"returnName"`     //联系人姓名
	ZipCode        string    `json:"zipCode"`        //邮编
	Question       string    `json:"question"`       //问题描述
	Image          string    `json:"image"`          //图片
	StateReturn    int       `json:"stateReturn"`    //退货状态：1、未处理；2、审核通过待收货；3、已经收货；4、不予处理
	StateMoney     int       `json:"stateMoney"`     //退款状态：1、未退款；2、退款到账户；3、退款到银行
	BackMoney      float64   `json:"backMoney"`      //退款金额
	BackMoneyTime  time.Time `json:"backMoneyTime"`  //退款时间
	CreateTime     time.Time `json:"createTime"`
	UpdateTime     time.Time `json:"updateTime"`
	OptId          int       `json:"optId"`  //处理人
	Remark         string    `json:"remark"` //备注
}

func (m *MemberProductBack) Save() error {
	_, err := ShopDB.Insert(m)
	return err
}

func (m *MemberProductBack) Delete() error {
	_, err := ShopDB.Where("id=?", m.Id).Delete(MemberProductBack{})
	return err
}
func GetMemberProductBackById(id int64) (*MemberProductBack, error) {
	var ret MemberProductBack
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("MemberProductBack Not Found Value: %v", id))
	}
	return &ret, nil
}

func (m *MemberProductBack) Update() error {
	_, err := ShopDB.Where("id=?", m.Id).Update(m)
	return err
}
