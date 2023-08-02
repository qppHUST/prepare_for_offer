package model

type Food struct {
	Id     int     //表字段名为：id
	Name   string  //表字段名为：name
	Price  float64 //表字段名为：price
	TypeId int     //表字段名为：type_id

	//字段定义后面使用两个反引号``包裹起来的字符串部分叫做标签定义，这个是golang的基础语法，不同的库会定义不同的标签，有不同的含义
	CreateTime int64 `gorm:"column:createtime"` //表字段名为：createtime
}

func (food Food) TableName() string {
	return "food"
}
