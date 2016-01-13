package model

import (
	"errors"
	"fmt"
	"time"
)

type OrderInventedProduct struct {
	Id                      int64     `json:"id"`
	OrderInventedId         int       `json:"orderInventedId"`         //订单ID
	OrderInventedSn         string    `json:"orderInventedSn"`         //订单号
	SellerId                int       `json:"sellerId"`                //商家ID
	ProductId               int       `json:"productId"`               //商品id
	ProductName             string    `json:"productName"`             //商品名称
	ProductSku              string    `json:"productSku"`              //抽象商品sku
	PackageGroupsId         int       `json:"packageGroupsId"`         //促销套装0、不是促销套装；大于0，促销套装ID，价格是促销套装的ID
	MallGroupsId            int       `json:"mallGroupsId"`            //商城套装0，不是上次套装；大于0商城套装ID
	GiftId                  int       `json:"giftId"`                  //赠品ID0、没有赠品；大于0、不是赠品；大于0赠品的ID
	IsGift                  int       `json:"isGift"`                  //是否是赠品，0、不是；1、是；
	MoneyPrice              float64   `json:"moneyPrice"`              //商品单价
	Number                  int       `json:"number"`                  //商品数量
	MoneyAmount             float64   `json:"moneyAmount"`             //网单金额
	ActivityId              int       `json:"activityId"`              //团购ID，为0正常购买
	FlashsalesId            int       `json:"flashsalesId"`            //限时抢购ID，为0正常购买
	LogisticsId             int       `json:"logisticsId"`             //物流
	LogisticsName           string    `json:"logisticsName"`           //物流name
	LogisticsNumber         string    `json:"logisticsNumber"`         //发票快递单号
	ShippingTime            time.Time `json:"shippingTime"`            //发货时间
	CloseTime               time.Time `json:"closeTime"`               //网单完成关闭或取消关闭时间
	SystemRemark            string    `json:"systemRemark"`            //系统备注，不给用户显示
	MemberProductBackId     int       `json:"memberProductBackId"`     //退货ID，默认为0
	MemberProductExchangeId int       `json:"memberProductExchangeId"` //换货ID，默认为0
	CreateTime              time.Time `json:"createTime"`
	UpdateTime              time.Time `json:"updateTime"`
}

func (o *OrderInventedProduct) Save() error {
	_, err := ShopDB.Insert(o)
	return err
}

func (o *OrderInventedProduct) Delete() error {
	_, err := ShopDB.Where("id=?", o.Id).Delete(OrderInventedProduct{})
	return err
}
func GetOrderInventedProductById(id int64) (*OrderInventedProduct, error) {
	var ret OrderInventedProduct
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("OrderInventedProduct Not Found Value: %v", id))
	}
	return &ret, nil
}

func (o *OrderInventedProduct) Update() error {
	_, err := ShopDB.Where("id=?", o.Id).Update(o)
	return err
}
