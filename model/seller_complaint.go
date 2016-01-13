package model

import (
	"errors"
	"fmt"
	"time"
)

type SellerComplaint struct {
	Id                  int64     `json:"id"`
	UserId              int       `json:"userId"`              //投诉人ID
	UserName            string    `json:"userName"`            //投诉人账户
	OrderId             int       `json:"orderId"`             //订单ID
	OrderProductId      int       `json:"orderProductId"`      //网单ID
	State               int       `json:"state"`               //状态：1、买家投诉待审核；2、买家投诉不通过；3、买家投诉通过；4、卖家申诉待审核；5、卖家申诉不通过；6、卖家申诉通过；
	ProductBackId       int       `json:"productBackId"`       //退货管理id，如没有置为0
	ProductExchangeId   int       `json:"productExchangeId"`   //换货管理id，如没有置为0
	Content             string    `json:"content"`             //投诉内容
	Image               string    `json:"image"`               //投诉图片
	ComplaintTime       time.Time `json:"complaintTime"`       //投诉时间
	SellerId            int       `json:"sellerId"`            //投诉商家
	SellerComplaintTime time.Time `json:"sellerComplaintTime"` //商家申诉时间
	SellerCompContent   string    `json:"sellerCompContent"`   //商家申诉内容
	SellerCompImage     string    `json:"sellerCompImage"`     //商家申诉图片
	OptId               int       `json:"optId"`               //平台处理人ID
	OptContent          string    `json:"optContent"`          //平台处理结果
	UserContent         string    `json:"userContent"`         //客户反馈意见
	CreateTime          time.Time `json:"createTime"`          //创建时间
	OptTime             time.Time `json:"optTime"`             //处理时间
}

func (s *SellerComplaint) Save() error {
	_, err := ShopDB.Insert(s)
	return err
}

func (s *SellerComplaint) Delete() error {
	_, err := ShopDB.Where("id=?", s.Id).Delete(SellerComplaint{})
	return err
}
func GetSellerComplaintById(id int64) (*SellerComplaint, error) {
	var ret SellerComplaint
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("SellerComplaint Not Found Value: %v", id))
	}
	return &ret, nil
}

func (s *SellerComplaint) Update() error {
	_, err := ShopDB.Where("id=?", s.Id).Update(s)
	return err
}
