package model

import (
	"errors"
	"fmt"
)

type SystemLogs struct {
	Id           int64  `json:"id"`
	OptId        int    `json:"optId"`        //操作人ID
	OptName      string `json:"optName"`      //操作人姓名
	OptContent   string `json:"optContent"`   //操作内容
	CreateTime   int    `json:"createTime"`   //创建时间
	MemberRuleId int    `json:"memberRuleId"` //member_rule
	OptObject    string `json:"optObject"`    //操作表
	OptObjectId  int    `json:"optObjectId"`  //操作表ID
}

func (s *SystemLogs) Save() error {
	_, err := ShopDB.Insert(s)
	return err
}

func (s *SystemLogs) Delete() error {
	_, err := ShopDB.Where("id=?", s.Id).Delete(SystemLogs{})
	return err
}
func GetSystemLogsById(id int64) (*SystemLogs, error) {
	var ret SystemLogs
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("SystemLogs Not Found Value: %v", id))
	}
	return &ret, nil
}

func (s *SystemLogs) Update() error {
	_, err := ShopDB.Where("id=?", s.Id).Update(s)
	return err
}
