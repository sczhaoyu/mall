package model

import (
	"errors"
	"fmt"
	"time"
)

type CodeMaster struct {
	CodeDiv      string    `json:"codeDiv"`      //code组id
	CodeCd       string    `json:"codeCd"`       //codeId
	CodeText     string    `json:"codeText"`     //code中文名称
	SortOrder    int       `json:"sortOrder"`    //组内顺序为
	UseYn        int       `json:"useYn"`        //是否使用
	CreateUserId int       `json:"createUserId"` //创建人Id
	CreateUser   string    `json:"createUser"`   //创建人
	CreateTime   time.Time `json:"createTime"`   //创建时间
	UpdateUserId int       `json:"updateUserId"` //修改人Id
	UpdateUser   string    `json:"updateUser"`   //修改人
	UpdateTime   time.Time `json:"updateTime"`   //修改时间
}

func (c *CodeMaster) Save() error {
	_, err := ShopDB.Insert(c)
	return err
}

func (c *CodeMaster) Delete() error {
	_, err := ShopDB.Where("code_div=?", c.CodeDiv).Delete(CodeMaster{})
	return err
}
func GetCodeMasterByCodeDiv(codeDiv string) (*CodeMaster, error) {
	var ret CodeMaster
	b, err := ShopDB.Where("code_div=?", codeDiv).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("CodeMaster Not Found Value: %v", codeDiv))
	}
	return &ret, nil
}

func (c *CodeMaster) Update() error {
	_, err := ShopDB.Where("code_div=?", c.CodeDiv).Update(c)
	return err
}
