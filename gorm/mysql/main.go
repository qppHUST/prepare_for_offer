package main

import (
	"database/sql"
	"fmt"

	"github.com/qppHUST/prepareForOffer/gorm/mysql/connection"
	"github.com/qppHUST/prepareForOffer/gorm/mysql/model"
)

func main() {
	db := connection.GetConnection()
	db.Begin(&sql.TxOptions{})
	var food model.Food
	res := db.Debug().Where("id = ?", "1").First(&food)
	db.Commit()
	fmt.Println(res.Error)
	fmt.Println(food)
}
