/**
  @Author : hanxiaodong
*/

package web

import (
	"net/http"
	"fmt"
	"education/web/controller"
)


// 啟動Web服務並指定路由信息
func WebStart(app controller.Application)  {

	fs:= http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// 指定路由信息(匹配請求)
	http.HandleFunc("/", app.LoginView)
	http.HandleFunc("/login", app.Login)
	http.HandleFunc("/loginout", app.LoginOut)

	http.HandleFunc("/index", app.Index)
	http.HandleFunc("/help", app.Help)

	http.HandleFunc("/addEduInfo", app.AddEduShow)	// 顯示添加信息頁面
	http.HandleFunc("/addEdu", app.AddEdu)	// 提交信息請求

	http.HandleFunc("/queryPage", app.QueryPage)	// 轉至根據社員編號與姓名查詢信息頁面
	http.HandleFunc("/query", app.FindCertByNoAndName)	// 根據社員編號與姓名查詢信息

	http.HandleFunc("/queryPage2", app.QueryPage2)	// 轉至根據學號查詢信息頁面
	http.HandleFunc("/query2", app.FindByID)	// 根據學號查詢信息


	http.HandleFunc("/modifyPage", app.ModifyShow)	// 修改信息頁面
	http.HandleFunc("/modify", app.Modify)	//  修改信息

	http.HandleFunc("/upload", app.UploadFile)

	fmt.Println("啟動Web服務, 監聽端口號為: 9000")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("Web服務啟動失敗: %v", err)
	}

}



