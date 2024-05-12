package main

import (
	"chaincode/api"       //
	"chaincode/model"     //模型
	"chaincode/pkg/utils" //工具包
	"fmt"                 //标准库中的fmt和time
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type BlockChainRealEstate struct { //定义智能合约结构
}

// Init 链码初始化
func (t *BlockChainRealEstate) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("链码初始化")
	//初始化默认数据
	var accountIds = [6]string{
		"5feceb66ffc8",
		"6b86b273ff34",
		"d4735e3a265e",
		"4e07408562be",
		"4b227777d4dd",
		"ef2d127de37b",
	}
	var userNames = [6]string{"管理员", "①号业主", "②号业主", "③号业主", "④号业主", "⑤号业主"}
	var balances = [6]float64{0, 5000000, 5000000, 5000000, 5000000, 5000000}
	//初始化账号数据账号，余额
	for i, val := range accountIds {
		account := &model.Account{
			AccountId: val,
			UserName:  userNames[i],
			Balance:   balances[i],
		}
		// 写入账本
		if err := utils.WriteLedger(account, stub, model.AccountKey, []string{val}); err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
	}
	return shim.Success(nil)
}

// Invoke 实现Invoke接口调用智能合约
func (t *BlockChainRealEstate) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	funcName, args := stub.GetFunctionAndParameters()
	switch funcName {
	case "hello": //简单问候
		return api.Hello(stub, args)
	case "queryAccountList": //查询账户列表
		return api.QueryAccountList(stub, args)
	case "createRealEstate": //创建房地产记录
		return api.CreateRealEstate(stub, args)
	case "queryRealEstateList": //查询房地产列表
		return api.QueryRealEstateList(stub, args)
	case "createSelling": //创建销售记录
		return api.CreateSelling(stub, args)
	case "createSellingByBuy": //通过购买创建销售记录
		return api.CreateSellingByBuy(stub, args)
	case "querySellingList": //查询销售列表
		return api.QuerySellingList(stub, args)
	case "querySellingListByBuyer": //根据购买者查询销售列表
		return api.QuerySellingListByBuyer(stub, args)
	case "updateSelling": //更新销售记录
		return api.UpdateSelling(stub, args)
	case "createDonating": //更新捐赠记录
		return api.CreateDonating(stub, args)
	case "queryDonatingList": //查询捐赠列表
		return api.QueryDonatingList(stub, args)
	case "queryDonatingListByGrantee": //根据受赠人查询捐赠列表
		return api.QueryDonatingListByGrantee(stub, args)
	case "updateDonating": //更新捐赠记录
		return api.UpdateDonating(stub, args)
	default: //错误操作
		return shim.Error(fmt.Sprintf("没有该功能: %s", funcName))
	}
}

func main() {
	timeLocal, err := time.LoadLocation("Asia/Shanghai") //设置当前时间
	if err != nil {
		panic(err)
	}
	time.Local = timeLocal
	err = shim.Start(new(BlockChainRealEstate))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
