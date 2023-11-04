

package controller

import (
	"net/http"
	"path/filepath"
	"html/template"
	"fmt"
)

func ShowView(w http.ResponseWriter, r *http.Request, templateName string, data interface{})  {

	// 指定視圖所在路徑
	pagePath := filepath.Join("web", "tpl", templateName)

	resultTemplate, err := template.ParseFiles(pagePath)
	if err != nil {
		fmt.Printf("創建模板實例錯誤: %v", err)
		return
	}

	err = resultTemplate.Execute(w, data)
	if err != nil {
		fmt.Printf("在模板中融合數據時發生錯誤: %v", err)
		//fmt.Fprintf(w, "顯示在客戶端瀏覽器中的錯誤信息")
		return
	}

}
