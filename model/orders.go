package model

import (
	"errors"
	"fmt"
	"time"
)

type Orders struct {
	Id                  int64     `json:"id"`
	OrderSn             string    `json:"orderSn"`             //订单号
	RelationOrderSn     string    `json:"relationOrderSn"`     //关联订单编号
	OrderType           int       `json:"orderType"`           //订单类型：1、合并订单 2、子订单
	SellerId            int       `json:"sellerId"`            //商家ID
	MemberId            int       `json:"memberId"`            //会员ID
	MemberName          string    `json:"memberName"`          //会员name
	OrderState          int       `json:"orderState"`          //订单状态：1、未付款的订单；2、待确认的订单；3、待发货的订单；4、已发货的订单；5、已完成的订单；6、取消的订单
	PayTime             time.Time `json:"payTime"`             //付款时间
	PaymentStatus       int       `json:"paymentStatus"`       //付款状态：0 买家未付款 1 买家已付款
	InvoiceStatus       int       `json:"invoiceStatus"`       //发票状态0、不要发票；1、单位；2个人
	InvoiceTitle        string    `json:"invoiceTitle"`        //发票抬头
	InvoiceType         string    `json:"invoiceType"`         //发票类型（内容）
	MoneyProduct        float64   `json:"moneyProduct"`        //商品金额，等于订单中所有的商品的单价乘以数量之和
	MoneyLogistics      float64   `json:"moneyLogistics"`      //物流费用
	MoneyOrder          float64   `json:"moneyOrder"`          //订单总金额，等于商品总金额＋运费-优惠金额
	MoneyPaidBalance    float64   `json:"moneyPaidBalance"`    //余额账户支付总金额
	MoneyPaidReality    float64   `json:"moneyPaidReality"`    //现金支付金额
	MoneyGiftcardAmount float64   `json:"moneyGiftcardAmount"` //礼品卡抵用金额
	MoneyCouponcode     float64   `json:"moneyCouponcode"`     //优惠码优惠金额
	MoneyLowerAmount    float64   `json:"moneyLowerAmount"`    //下单立减金额
	MoneyBack           float64   `json:"moneyBack"`           //退款的金额，订单没有退款则为0
	MoneyIntegral       float64   `json:"moneyIntegral"`       //积分换算金额
	Integral            int       `json:"integral"`            //订单使用的积分数量
	LowerId             int       `json:"lowerId"`             //下单立减id
	GiftcardId          int       `json:"giftcardId"`          //礼品卡ID
	CouponcodeId        int       `json:"couponcodeId"`        //优惠码ID
	Ip                  string    `json:"ip"`                  //ip
	PaymentName         string    `json:"paymentName"`         //支付方式名称
	PaymentCode         string    `json:"paymentCode"`         //支付方式code
	Name                string    `json:"name"`                //收货人
	ProvinceId          int       `json:"provinceId"`          //Province
	CityId              int       `json:"cityId"`              //city
	AreaId              int       `json:"areaId"`              //area
	AddressAll          string    `json:"addressAll"`          //省市区组合
	AddressInfo         string    `json:"addressInfo"`         //详细地址
	Mobile              string    `json:"mobile"`
	Email               string    `json:"email"`
	ZipCode             string    `json:"zipCode"`          //邮编
	Remark              string    `json:"remark"`           //订单备注
	DeliverTime         time.Time `json:"deliverTime"`      //发货时间
	FinishTime          time.Time `json:"finishTime"`       //订单完成时间
	TradeSn             string    `json:"tradeSn"`          //在线支付交易流水号
	Source              int       `json:"source"`           //会员来源1、pc；2、H5；3、Android；4、IOS
	LogisticsId         int       `json:"logisticsId"`      //物流公司ID
	LogisticsName       string    `json:"logisticsName"`    //物流公司
	LogisticsNumber     string    `json:"logisticsNumber"`  //发票快递单号
	IsGiftCardOrder     int       `json:"isGiftCardOrder"`  //是否为礼品卡订单0、不是；1、是
	IsCodconfim         int       `json:"isCodconfim"`      //是否货到付款订单0、不是；1、是
	CodconfirmId        int       `json:"codconfirmId"`     //货到付款确认人
	CodconfirmName      string    `json:"codconfirmName"`   //货到付款确认Name
	CodconfirmTime      time.Time `json:"codconfirmTime"`   //货到付款确认时间
	CodconfirmRemark    string    `json:"codconfirmRemark"` //货到付款确认备注
	CodconfirmState     int       `json:"codconfirmState"`  //货到付款状态 0、非货到付款；1、待确认；2、确认通过可以发货；3、订单取消
	CreateTime          time.Time `json:"createTime"`
	UpdateTime          time.Time `json:"updateTime"`
}

func (o *Orders) Save() error {
	_, err := ShopDB.Insert(o)
	return err
}

func (o *Orders) Delete() error {
	_, err := ShopDB.Where("id=?", o.Id).Delete(Orders{})
	return err
}
func GetOrdersById(id int64) (*Orders, error) {
	var ret Orders
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("Orders Not Found Value: %v", id))
	}
	return &ret, nil
}

func (o *Orders) Update() error {
	_, err := ShopDB.Where("id=?", o.Id).Update(o)
	return err
}
