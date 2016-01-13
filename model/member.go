package model

import (
	"errors"
	"fmt"
	"time"
)

type Member struct {
	Id              int64     `json:"id"`
	Name            string    `json:"name"`            //用户名
	Password        string    `json:"password"`        //密码
	Grade           int       `json:"grade"`           //会员等级
	GradeValue      int       `json:"gradeValue"`      //会员经验值
	Integral        int       `json:"integral"`        //会员积分
	RegisterTime    time.Time `json:"registerTime"`    //注册时间
	LastLoginTime   time.Time `json:"lastLoginTime"`   //最后登录时间
	LastLoginIp     string    `json:"lastLoginIp"`     //最后登录IP
	LoginNumber     int       `json:"loginNumber"`     //登陆次数
	LastAddressId   int       `json:"lastAddressId"`   //上次使用的地址
	LastPaymentCode int       `json:"lastPaymentCode"` //上次使用的支付方式
	Gender          int       `json:"gender"`          //性别0、保密；1、男；2、女
	Birthday        time.Time `json:"birthday"`        //生日
	Email           string    `json:"email"`           //邮箱
	Qq              string    `json:"qq"`              //qq
	Mobile          string    `json:"mobile"`          //mobile
	Phone           string    `json:"phone"`           //电话
	PwdErrCount     int       `json:"pwdErrCount"`     //密码输入错误次数
	Source          int       `json:"source"`          //会员来源1、pc；2、H5；3、Android；4、IOS ;5、微信商城
	Balance         float64   `json:"balance"`         //账户余额
	BalancePwd      string    `json:"balancePwd"`      //账户支付密码
	IsEmailVerify   int       `json:"isEmailVerify"`   //是否验证邮箱0、未验证；1、验证
	IsSmsVerify     int       `json:"isSmsVerify"`     //是否验证手机0、未验证；1、验证
	SmsVerifyCode   string    `json:"smsVerifyCode"`   //短信验证码
	EmailVerifyCode string    `json:"emailVerifyCode"` //邮件验证码
	CanReceiveSms   int       `json:"canReceiveSms"`   //是否接受短信0、不接受；1、接受
	CanReceiveEmail int       `json:"canReceiveEmail"` //是否接收邮件
	Status          int       `json:"status"`          //1、正常使用；2、停用
	UpdateTime      time.Time `json:"updateTime"`
}

func (m *Member) Save() error {
	_, err := ShopDB.Insert(m)
	return err
}

func (m *Member) Delete() error {
	_, err := ShopDB.Where("id=?", m.Id).Delete(Member{})
	return err
}
func GetMemberById(id int64) (*Member, error) {
	var ret Member
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("Member Not Found Value: %v", id))
	}
	return &ret, nil
}

func (m *Member) Update() error {
	_, err := ShopDB.Where("id=?", m.Id).Update(m)
	return err
}
