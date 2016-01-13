package model

import (
	"errors"
	"fmt"
	"time"
)

type OrderInvented struct {
	Id                  int64     `json:"id"`
	OrderSn             string    `json:"orderSn"`             //订单号
	SellerId            int       `json:"sellerId"`            //商家ID
	MemberId            int       `json:"memberId"`            //会员ID
	MemberName          string    `json:"memberName"`          //会员name
	OrderState          int       `json:"orderState"`          //虚拟订单状态：1、未付款的订单；2、待处理的订单；3、已完成的订单；
	PayTime             time.Time `json:"payTime"`             //在线付款时间
	PaymentStatus       int       `json:"paymentStatus"`       //付款状态：0 买家未付款 1 买家已付款
	InvoiceStatus       int       `json:"invoiceStatus"`       //发票状态0、不要发票；1、单位；2个人
	InvoiceTitle        string    `json:"invoiceTitle"`        //发票抬头
	InvioceState        int       `json:"invioceState"`        //发票抬头类型1、公司；2、个人
	InvoiceType         string    `json:"invoiceType"`         //发票类型
	MoneyProduct        float64   `json:"moneyProduct"`        //商品金额，等于订单中所有的商品的单价乘以数量之和
	MoneyOrder          float64   `json:"moneyOrder"`          //订单总金额，等于商品总金额＋运费-优惠金额
	MoneyPaidBalance    float64   `json:"moneyPaidBalance"`    //余额账户支付总金额
	MoneyPaidReality    float64   `json:"moneyPaidReality"`    //现金支付金额
	MoneyGiftcardAmount float64   `json:"moneyGiftcardAmount"` //礼品卡抵用金额
	MoneyCouponcode     float64   `json:"moneyCouponcode"`     //优惠码优惠金额
	MoneyLowerAmount    float64   `json:"moneyLowerAmount"`    //下单立减金额
	MoneyBack           float64   `json:"moneyBack"`           //退款的金额，订单没有退款则为0
	MoneyIntegral       int       `json:"moneyIntegral"`       //订单消耗的积分数量
	LowerId             int       `json:"lowerId"`             //下单立减id
	GiftcardId          int       `json:"giftcardId"`          //礼品卡ID
	CouponcodeId        int       `json:"couponcodeId"`        //优惠码ID
	Ip                  string    `json:"ip"`                  //ip
	PaymentName         string    `json:"paymentName"`         //支付方式名称
	PaymentCode         string    `json:"paymentCode"`         //支付方式code
	Mobile              string    `json:"mobile"`              //接受手机号码
	CodeExchange        string    `json:"codeExchange"`        //电子兑换码
	IsSuccess           int       `json:"isSuccess"`           //短信是否发生成功0、未发生；2、发生成功；2、发送失败
	Email               string    `json:"email"`
	ZipCode             string    `json:"zipCode"`           //邮编
	Remark              string    `json:"remark"`            //订单备注
	FinishTime          time.Time `json:"finishTime"`        //订单完成时间
	TradeSn             string    `json:"tradeSn"`           //在线支付交易流水号
	Source              string    `json:"source"`            //订单来源
	IsGiftCardOrder     int       `json:"isGiftCardOrder"`   //是否为礼品卡订单0、不是；1、是
	ProductCommentsId   int       `json:"productCommentsId"` //评价ID，如果为0，没有评价
	CreateTime          time.Time `json:"createTime"`
	UpdateTime          time.Time `json:"updateTime"`
}

func (o *OrderInvented) Save() error {
	_, err := ShopDB.Insert(o)
	return err
}

func (o *OrderInvented) Delete() error {
	_, err := ShopDB.Where("id=?", o.Id).Delete(OrderInvented{})
	return err
}
func GetOrderInventedById(id int64) (*OrderInvented, error) {
	var ret OrderInvented
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("OrderInvented Not Found Value: %v", id))
	}
	return &ret, nil
}

func (o *OrderInvented) Update() error {
	_, err := ShopDB.Where("id=?", o.Id).Update(o)
	return err
}
