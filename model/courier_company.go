package model

import (
	"errors"
	"fmt"
	"time"
)

type CourierCompany struct {
	Id          int64     `json:"id"`
	State       int       `json:"state"`       //状态 1-可用 2-不可用
	CompanyMark string    `json:"companyMark"` //快递代码
	CompanyName string    `json:"companyName"` //快递公司名称
	Seq         int       `json:"seq"`         //排序号
	CompanyType string    `json:"companyType"` //快递类型 1-平邮 2-EMS 3-快递
	SellerId    int       `json:"sellerId"`    //商家id
	CreateTime  time.Time `json:"createTime"`
	UpdateTime  time.Time `json:"updateTime"`
	ImagePath   string    `json:"imagePath"` //快递模板路径
	Content     string    `json:"content"`   //打印模板生成的js文件
}

func (c *CourierCompany) Save() error {
	_, err := ShopDB.Insert(c)
	return err
}

func (c *CourierCompany) Delete() error {
	_, err := ShopDB.Where("id=?", c.Id).Delete(CourierCompany{})
	return err
}
func GetCourierCompanyById(id int64) (*CourierCompany, error) {
	var ret CourierCompany
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("CourierCompany Not Found Value: %v", id))
	}
	return &ret, nil
}

func (c *CourierCompany) Update() error {
	_, err := ShopDB.Where("id=?", c.Id).Update(c)
	return err
}
