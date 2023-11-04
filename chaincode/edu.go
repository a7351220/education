package main

import (
"github.com/hyperledger/fabric-chaincode-go/shim"
"github.com/hyperledger/fabric-protos-go/peer"
"fmt"
"encoding/json"
"bytes"

)


type Education struct {
	ObjectType	string	`json:"docType"`
	Name	string	`json:"Name"`		// 姓名
	Gender	string	`json:"Gender"`		// 性别
	Nation	string	`json:"Nation"`		// 種族
	EntityID	string	`json:"EntityID"`		// 學號
	Place	string	`json:"Place"`		// 地區
	BirthDay	string	`json:"BirthDay"`		// 出生日期

	EnrollDate	string	`json:"EnrollDate"`		// 註冊日期
	GraduationDate	string	`json:"GraduationDate"`	// 入學日期
	SchoolName	string	`json:"SchoolName"`	// 校名
	Major	string	`json:"Major"`	// 專業
	QuaType	string	`json:"QuaType"`	// 科系
	Length	string	`json:"Length"`	// 年級
	Mode	string	`json:"Mode"`	// 模式
	Level	string	`json:"Level"`	// 社員等級
	Graduation	string	`json:"Graduation"`	// 是否畢業
	CertNo	string	`json:"CertNo"`	// 社員編號

	Photo	string	`json:"Photo"`	// 照片

	Historys	[]HistoryItem	// 當前edu的歷史紀錄
}

type HistoryItem struct {
	TxId	string
	Education	Education
}

type EducationChaincode struct {

}

func (t *EducationChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response{
	fmt.Println(" ==== Init ====")

	return shim.Success(nil)
}

func (t *EducationChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response{

	fun, args := stub.GetFunctionAndParameters()

	if fun == "addEdu"{
		return t.addEdu(stub, args)		// 添加信息
	}else if fun == "queryEduByCertNoAndName" {
		return t.queryEduByCertNoAndName(stub, args)		// 根據社員編號及姓名查詢信息
	}else if fun == "queryEduInfoByEntityID" {
		return t.queryEduInfoByEntityID(stub, args)	// 根據學號及姓名查詢詳情
	}else if fun == "updateEdu" 
		return t.updateEdu(stub, args)		// 根據社員編號更新訊息
	}else if fun == "delEdu"{
		return t.delEdu(stub, args)	// 根據社員編號刪除訊息
	}

	return shim.Error("指定的函數名稱錯誤")

}


const DOC_TYPE = "eduObj"

// 保存edu
// args: education
func PutEdu(stub shim.ChaincodeStubInterface, edu Education) ([]byte, bool) {

	edu.ObjectType = DOC_TYPE

	b, err := json.Marshal(edu)
	if err != nil {
		return nil, false
	}

	// 保存edu狀態
	err = stub.PutState(edu.EntityID, b)
	if err != nil {
		return nil, false
	}

	return b, true
}

// 根據學號查詢訊息狀態
// args: entityID
func GetEduInfo(stub shim.ChaincodeStubInterface, entityID string) (Education, bool)  {
	var edu Education
	// 根據學號查詢訊息狀態
	b, err := stub.GetState(entityID)
	if err != nil {
		return edu, false
	}

	if b == nil {
		return edu, false
	}

	//對查詢到的狀態進行反序列化
	err = json.Unmarshal(b, &edu)
	if err != nil {
		return edu, false
	}

	// 返回结果
	return edu, true
}

// 根據指定的查詢字符串實現富查詢
func getEduByQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer  resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}

		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		bArrayMemberAlreadyWritten = true
	}

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil

}

// 添加信息
// args: educationObject
// 學號為 key, Education 為 value
func (t *EducationChaincode) addEdu(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 2{
		return shim.Error("給定的參數個數不符合要求")
	}

	var edu Education
	err := json.Unmarshal([]byte(args[0]), &edu)
	if err != nil {
		return shim.Error("反序列化信息時發生錯誤")
	}

	// 查重: 身份證號碼必須唯一
	_, exist := GetEduInfo(stub, edu.EntityID)
	if exist {
		return shim.Error("要添加的身份證號碼已存在")
	}

	_, bl := PutEdu(stub, edu)
	if !bl {
		return shim.Error("保存信息時發生錯誤")
	}

	err = stub.SetEvent(args[1], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("信息添加成功"))
}

// 根據社員編號及姓名查詢信息
// args: CertNo, name
func (t *EducationChaincode) queryEduByCertNoAndName(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 2 {
		return shim.Error("給定的參數個數不符合要求")
	}
	CertNo := args[0]
	name := args[1]

	// 拼裝CouchDB所需要的查詢字符串(是標準的一個JSON串)
	// queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"eduObj\", \"CertNo\":\"%s\"}}", CertNo)
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"%s\", \"CertNo\":\"%s\", \"Name\":\"%s\"}}", DOC_TYPE, CertNo, name)

	// 查詢數據
	result, err := getEduByQueryString(stub, queryString)
	if err != nil {
		return shim.Error("根據社員編號及姓名查詢信息時發生錯誤")
	}
	if result == nil {
		return shim.Error("根據指定的社員編號及姓名沒有查詢到相關的信息")
	}
	return shim.Success(result)
}

// 根據身份證號碼查詢詳情（溯源）
// args: entityID
func (t *EducationChaincode) queryEduInfoByEntityID(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("給定的參數個數不符合要求")
	}

	// 根據學號查詢edu狀態
	b, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("根據學號查詢信息失敗")
	}

	if b == nil {
		return shim.Error("根據學號沒有查詢到相關的信息")
	}

	// 對查詢到的狀態進行反序列化
	var edu Education
	err = json.Unmarshal(b, &edu)
	if err != nil {
		return  shim.Error("反序列化edu信息失败")
	}

	// 獲取歷史變更數據
	iterator, err := stub.GetHistoryForKey(edu.EntityID)
	if err != nil {
		return shim.Error("根據指定的學號查詢對應的歷史變更數據失敗")
	}
	defer iterator.Close()

	// 叠代處理
	var historys []HistoryItem
	var hisEdu Education
	for iterator.HasNext() {
		hisData, err := iterator.Next()
		if err != nil {
			return shim.Error("獲取edu的歷史變更數據失敗")
		}

		var historyItem HistoryItem
		historyItem.TxId = hisData.TxId
		json.Unmarshal(hisData.Value, &hisEdu)

		if hisData.Value == nil {
			var empty Education
			historyItem.Education = empty
		}else {
			historyItem.Education = hisEdu
		}

		historys = append(historys, historyItem)

	}

	edu.Historys = historys

	// 返回
	result, err := json.Marshal(edu)
	if err != nil {
		return shim.Error("序列化edu信息時發生錯誤")
	}
	return shim.Success(result)
}

// 根據學號更新信息
// args: educationObject
func (t *EducationChaincode) updateEdu(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2{
		return shim.Error("給定的參數個數不符合要求")
	}

	var info Education
	err := json.Unmarshal([]byte(args[0]), &info)
	if err != nil {
		return  shim.Error("反序列化edu信息失敗")
	}

	// 根據學號查詢信息
	result, bl := GetEduInfo(stub, info.EntityID)
	if !bl{
		return shim.Error("根據學號查詢信息時發生錯誤")
	}

	result.Name = info.Name
	result.BirthDay = info.BirthDay
	result.Nation = info.Nation
	result.Gender = info.Gender
	result.Place = info.Place
	result.EntityID = info.EntityID
	result.Photo = info.Photo


	result.EnrollDate = info.EnrollDate
	result.GraduationDate = info.GraduationDate
	result.SchoolName = info.SchoolName
	result.Major = info.Major
	result.QuaType = info.QuaType
	result.Length = info.Length
	result.Mode = info.Mode
	result.Level = info.Level
	result.Graduation = info.Graduation
	result.CertNo = info.CertNo;

	_, bl = PutEdu(stub, result)
	if !bl {
		return shim.Error("保存信息信息時發生錯誤")
	}

	err = stub.SetEvent(args[1], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("信息更新成功"))
}

// 根據學號刪除信息（暂不提供）
// args: entityID
func (t *EducationChaincode) delEdu(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2{
		return shim.Error("給定的參數個數不符合要求")
	}

	/*var edu Education
	result, bl := GetEduInfo(stub, info.EntityID)
	err := json.Unmarshal(result, &edu)
	if err != nil {
		return shim.Error("反序列化信息時發生錯誤")
	}*/

	err := stub.DelState(args[0])
	if err != nil {
		return shim.Error("刪除信息時發生錯誤")
	}

	err = stub.SetEvent(args[1], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("信息删除成功"))
}

func main(){
	err := shim.Start(new(EducationChaincode))
	if err != nil{
		fmt.Printf("啟動EducationChaincode時發生錯誤: %s", err)
	}
}

