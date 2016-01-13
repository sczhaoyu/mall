package model

import (
	"errors"
	"fmt"
)

type Config struct {
	Id            int64 `json:"id"`
	IntegralScale int   `json:"integralScale"` //积分换算比例，多少积分换算成1元人民币
}

func (c *Config) Save() error {
	_, err := ShopDB.Insert(c)
	return err
}

func (c *Config) Delete() error {
	_, err := ShopDB.Where("id=?", c.Id).Delete(Config{})
	return err
}
func GetConfigById(id int64) (*Config, error) {
	var ret Config
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("Config Not Found Value: %v", id))
	}
	return &ret, nil
}

func (c *Config) Update() error {
	_, err := ShopDB.Where("id=?", c.Id).Update(c)
	return err
}
