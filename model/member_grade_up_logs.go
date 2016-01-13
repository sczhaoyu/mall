package model

import (
	"errors"
	"fmt"
	"time"
)

type MemberGradeUpLogs struct {
	Id          int64     `json:"id"`
	MemberId    int       `json:"memberId"`    //会员ID
	MemberName  string    `json:"memberName"`  //会员名称
	BeforeExper int       `json:"beforeExper"` //升级之前的经验值
	AfterExper  int       `json:"afterExper"`  //升级之后的经验值
	BeforeGrade int       `json:"beforeGrade"` //升级之前的等级
	AfterGrade  int       `json:"afterGrade"`  //升级之后的等级
	CreateTime  time.Time `json:"createTime"`  //创建时间
}

func (m *MemberGradeUpLogs) Save() error {
	_, err := ShopDB.Insert(m)
	return err
}

func (m *MemberGradeUpLogs) Delete() error {
	_, err := ShopDB.Where("id=?", m.Id).Delete(MemberGradeUpLogs{})
	return err
}
func GetMemberGradeUpLogsById(id int64) (*MemberGradeUpLogs, error) {
	var ret MemberGradeUpLogs
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("MemberGradeUpLogs Not Found Value: %v", id))
	}
	return &ret, nil
}

func (m *MemberGradeUpLogs) Update() error {
	_, err := ShopDB.Where("id=?", m.Id).Update(m)
	return err
}
