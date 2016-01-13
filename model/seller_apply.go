package model

import (
	"errors"
	"fmt"
	"time"
)

type SellerApply struct {
	Id                    int64     `json:"id"`
	UserId                int       `json:"userId"`                //用户id
	Name                  string    `json:"name"`                  //用户名
	Password              string    `json:"password"`              //密码
	Company               string    `json:"company"`               //公司名称
	BussinessLicense      string    `json:"bussinessLicense"`      //营业执照注册号
	TaxLicense            string    `json:"taxLicense"`            //税务登记号
	CompanyProvince       string    `json:"companyProvince"`       //企业注册省
	CompanyCity           string    `json:"companyCity"`           //企业注册市
	CompanyStartTime      time.Time `json:"companyStartTime"`      //营业开始日期
	CompanyEndTime        time.Time `json:"companyEndTime"`        //营业结束日期
	CompanyAdd            string    `json:"companyAdd"`            //常用地址
	Telephone             string    `json:"telephone"`             //联系电话
	Fax                   string    `json:"fax"`                   //传真
	BussinessLicenseImage string    `json:"bussinessLicenseImage"` //营业执照扫描件
	Organization          string    `json:"organization"`          //组织机构代码
	BankUser              string    `json:"bankUser"`              //开户行账户名称
	BankName              string    `json:"bankName"`              //开户行
	BankNameBranch        string    `json:"bankNameBranch"`        //开户行支行名称
	BrandNameCode         string    `json:"brandNameCode"`         //开户行支行联行号
	BankCode              string    `json:"bankCode"`              //银行账号
	BankProvince          string    `json:"bankProvince"`          //开户行省
	BankCity              string    `json:"bankCity"`              //开户行市
	LegalPerson           string    `json:"legalPerson"`           //法定代表人
	LegalPersonCard       string    `json:"legalPersonCard"`       //法定代表人身份证
	PersonCardUp          string    `json:"personCardUp"`          //身份证正面扫描件
	PersonCardDown        string    `json:"personCardDown"`        //身份证背面扫描件
	PersonPhone           string    `json:"personPhone"`           //手机
	Email                 string    `json:"email"`                 //email
	State                 int       `json:"state"`                 //1、提交申请；2、审核通过；3、缴纳保证金；4、审核失败
	Type                  int       `json:"type"`                  //1、平台自营；2、商家入驻
	OptId                 int       `json:"optId"`                 //审核人ID
	Bond                  int       `json:"bond"`                  //保证金额度
	CreateTime            time.Time `json:"createTime"`            //创建时间
	UpdateTime            time.Time `json:"updateTime"`            //更新时间
}

func (s *SellerApply) Save() error {
	_, err := ShopDB.Insert(s)
	return err
}

func (s *SellerApply) Delete() error {
	_, err := ShopDB.Where("id=?", s.Id).Delete(SellerApply{})
	return err
}
func GetSellerApplyById(id int64) (*SellerApply, error) {
	var ret SellerApply
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("SellerApply Not Found Value: %v", id))
	}
	return &ret, nil
}

func (s *SellerApply) Update() error {
	_, err := ShopDB.Where("id=?", s.Id).Update(s)
	return err
}
