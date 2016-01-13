package model

import (
	"errors"
	"fmt"
	"time"
)

type SearchKeyword struct {
	Id         int64     `json:"id"`
	Keyword    string    `json:"keyword"`    //关键字
	CreateTime time.Time `json:"createTime"` //添加时间
	CreateId   int       `json:"createId"`
	CreateName string    `json:"createName"`
	UpdateTime time.Time `json:"updateTime"`
}

func (s *SearchKeyword) Save() error {
	_, err := ShopDB.Insert(s)
	return err
}

func (s *SearchKeyword) Delete() error {
	_, err := ShopDB.Where("id=?", s.Id).Delete(SearchKeyword{})
	return err
}
func GetSearchKeywordById(id int64) (*SearchKeyword, error) {
	var ret SearchKeyword
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("SearchKeyword Not Found Value: %v", id))
	}
	return &ret, nil
}

func (s *SearchKeyword) Update() error {
	_, err := ShopDB.Where("id=?", s.Id).Update(s)
	return err
}
