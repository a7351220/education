package main

import (
	"encoding/json"
	"fmt"
	"education/sdkInit"
	"education/service"
	"education/web"
	"education/web/controller"
	"os"

)

const (
	cc_name = "simplecc"
	cc_version = "1.0.0"
)
func main() {
	// init orgs information
	orgs := []*sdkInit.OrgInfo{
		{
			OrgAdminUser:  "Admin",
			OrgName:       "Org1",
			OrgMspId:      "Org1MSP",
			OrgUser:       "User1",
			OrgPeerNum:    1,
			OrgAnchorFile: os.Getenv("GOPATH") + "/src/education/fixtures/channel-artifacts/Org1MSPanchors.tx",
		},

	}

	// init sdk env info
	info := sdkInit.SdkEnvInfo{
		ChannelID:        "mychannel",
		ChannelConfig:    os.Getenv("GOPATH") + "/src/education/fixtures/channel-artifacts/channel.tx",
		Orgs:             orgs,
		OrdererAdminUser: "Admin",
		OrdererOrgName:   "OrdererOrg",
		OrdererEndpoint:  "orderer.example.com",
		ChaincodeID:      cc_name,
		ChaincodePath:    os.Getenv("GOPATH")+"/src/education/chaincode/",
		ChaincodeVersion: cc_version,
	}

	// sdk setup
	sdk, err := sdkInit.Setup("config.yaml", &info)
	if err != nil {
		fmt.Println(">> SDK setup error:", err)
		os.Exit(-1)
	}

	// create channel and join
	if err := sdkInit.CreateAndJoinChannel(&info); err != nil {
		fmt.Println(">> Create channel and join error:", err)
		os.Exit(-1)
	}

	// create chaincode lifecycle
	if err := sdkInit.CreateCCLifecycle(&info, 1, false, sdk); err != nil {
		fmt.Println(">> create chaincode lifecycle error: %v", err)
		os.Exit(-1)
	}

	// invoke chaincode set status
	fmt.Println(">> 通過鏈碼外部服務設置鏈碼狀態......")

	edu := service.Education{
		Name: "明華",
		Gender: "男",
		Nation: "半獸人",
		EntityID: "101",
		Place: "彰化",
		BirthDay: "2000年01月03日",
		EnrollDate: "2023年11月4日",
		GraduationDate: "2022年7月",
		SchoolName: "中正大學",
		Major: "加密貨幣",
		QuaType: "金融科技",
		Length: "碩二",
		Mode: "戰鬥模式",
		Level: "高級會員",
		Graduation: "在學",
		CertNo: "000",
		Photo: "/static/photo/11.png",
	}

	serviceSetup, err := service.InitService(info.ChaincodeID, info.ChannelID, info.Orgs[0], sdk)
	if err!=nil{
		fmt.Println()
		os.Exit(-1)
	}
	msg, err := serviceSetup.SaveEdu(edu)
	if err != nil {
		fmt.Println(err.Error())
	}else {
		fmt.Println("信息發布成功, 交易編號: " + msg)
	}

	result, err := serviceSetup.FindEduInfoByEntityID("101")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		var edu service.Education
		json.Unmarshal(result, &edu)
		fmt.Println("根據學號查詢信息成功：")
		fmt.Println(edu)
	}


	app := controller.Application{
		Setup: serviceSetup,
	}
	web.WebStart(app)
}
