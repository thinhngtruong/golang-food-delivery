package restaurantmodel

import (
	"errors"
	"learn-api/common"
	"strings"
)

const EntityName = "restaurant"

type Restaurant struct {
	common.SQLModel `json:",inline"`
	Name            string         `json:"name" gorm:"column:name;"`
	Addr            string         `json:"address" gorm:"column:addr;"`
	Logo            *common.Image  `json:"logo" form:"column:logo"`
	Cover           *common.Images `json:"cover" form:"column:cover"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantUpdate struct {
	Name  *string        `json:"name" gorm:"column:name;"`
	Addr  *string        `json:"address" gorm:"column:addr;"`
	Logo  *common.Image  `json:"logo" form:"column:logo"`
	Cover *common.Images `json:"cover" form:"column:cover"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

type RestaurantCreate struct {
	Id    int            `json:"id" gorm:"column:id;"`
	Name  string         `json:"name" gorm:"column:name;"`
	Addr  string         `json:"address" gorm:"column:addr;"`
	Logo  *common.Image  `json:"logo" form:"column:logo"`
	Cover *common.Images `json:"cover" form:"column:cover"`
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

func (res *RestaurantCreate) Validate() error {
	res.Name = strings.TrimSpace(res.Name)

	if len(res.Name) == 0 {
		return errors.New("restaurant name can't be blank")
	}

	return nil
}
