package v1

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/hongdagen/Go-000/Week02/Error-Homework/service"
	pkgerror "github.com/pkg/errors"
	"os"
)

func DoApiv1() {
	_, err := service.DoService()
	if err != nil {
		// 最顶层应该处理
		if errors.Is(pkgerror.Cause(err), sql.ErrNoRows) {
			// 打印堆栈
			fmt.Printf("发生了sql.ErrNoRows错误\n")
			fmt.Printf("原始错误发生信息：%T %v\n", pkgerror.Cause(err), pkgerror.Cause(err))
			fmt.Printf("堆栈信息：\n%+v\n", err)
			os.Exit(1)
		}
	}

	fmt.Println("查询结束")
}
