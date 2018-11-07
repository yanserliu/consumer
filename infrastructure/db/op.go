package db

import (
	"errors"

	"github.com/astaxie/beego/logs"
)

// GetBatchCountViaToken Get Batch via Token
func GetBatchCountViaToken(id int64) int64 {
	n, _ := Engine.Where("id = ?", id).Count(&User{Id: id})
	logs.Error(n)
	return n
}

func GetName() (error, string) {
	var name string
	result, err := Engine.Query("select email from user where id=1 ")
	logs.Error("err:", err)
	if err != nil {
		return errors.New("Error: unused 'order' fields"), ""
	}
	for _, val := range result {
		name = string(val["email"])

	}
	logs.Error("name:", name)
	return nil, name

}

func Insert(beans ...interface{}) (int64, error) {
	return Engine.Insert(beans...)
}

func Update(bean interface{}, condiBeans ...interface{}) (int64, error) {
	return Engine.Update(bean, condiBeans...)
}

func Delete(bean interface{}) (int64, error) {
	return Engine.Delete(bean)
}

func GetIsExit(bean interface{}) (bool, error) {
	return Engine.Get(bean)
}
