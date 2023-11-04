
package service

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"time"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"education/sdkInit"
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

type ServiceSetup struct {
	ChaincodeID	string
	Client	*channel.Client
}

func regitserEvent(client *channel.Client, chaincodeID, eventID string) (fab.Registration, <-chan *fab.CCEvent) {

	reg, notifier, err := client.RegisterChaincodeEvent(chaincodeID, eventID)
	if err != nil {
		fmt.Println("注冊鏈碼事件失敗: %s", err)
	}
	return reg, notifier
}

func eventResult(notifier <-chan *fab.CCEvent, eventID string) error {
	select {
	case ccEvent := <-notifier:
		fmt.Printf("接收到鏈碼事件: %v\n", ccEvent)
	case <-time.After(time.Second * 20):
		return fmt.Errorf("不能根據指定的事件ID接收到相應的鏈碼事件(%s)", eventID)
	}
	return nil
}

func InitService(chaincodeID, channelID string, org *sdkInit.OrgInfo, sdk *fabsdk.FabricSDK) (*ServiceSetup, error) {
	handler := &ServiceSetup{
		ChaincodeID:chaincodeID,
	}
	//prepare channel client context using client context
	clientChannelContext := sdk.ChannelContext(channelID, fabsdk.WithUser(org.OrgUser), fabsdk.WithOrg(org.OrgName))
	// Channel client is used to query and execute transactions (Org1 is default org)
	client, err := channel.New(clientChannelContext)
	if err != nil {
		return nil, fmt.Errorf("Failed to create new channel client: %s", err)
	}
	handler.Client = client
	return handler, nil
}
