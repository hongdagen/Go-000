package service

import "github.com/hongdagen/Go-000/Week02/Error-Homework/dao"

func DoService()(int, error){
	// 调用包内或项目内其他函数，直接返回err
	return dao.DoDao()
}