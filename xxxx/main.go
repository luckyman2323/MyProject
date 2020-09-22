package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

const (
	CreatProductFunc		="createProduct"
	GetProductFunc			="getProduct"
	GetProductStorageFunc	="getProductStorage"

	CollectionProductName			="collectionProduct"
	CollectionProductStorageName	="collectionProductStorage"
)
// 自定义结构体，名称可随意设置
type FristChainCode struct {
}

//实现Chaincode接口的Init()方法
func (t *FristChainCode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("---------初始化----------")
	return shim.Success(nil)
}

//实现Chaincode接口的Invoke()方法
func (t *FristChainCode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {


	//获取要执行的函数的函数名（首字母小写）和参数
	fun, args := stub.GetFunctionAndParameters()
	fmt.Println("---------"+fun+"方法调用中--------")
	//通过switch来确定要调用的函数
	switch fun {
	case CreatProductFunc:
		fmt.Println("--------创建商品--------")
		return createProduct(stub, args)
	case GetProductFunc:
		fmt.Println("--------获取商品信息--------")
		return getProduct(stub, args)
	case GetProductStorageFunc:
		fmt.Println("--------获取库存商品信息--------")
		return getProductStorage(stub, args)
	default:
		return shim.Error("方法名有误")
	}

	return shim.Success(nil)
}

type product struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

type productStorage struct {
	Id string `json:"id"`
	Storage int `json:"storage"`
}

func createProduct(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("参数只能有三个")
	}
	id := args[0]
	name := args[1]
	storage, err := strconv.Atoi(args[2])
	if err != nil{
		fmt.Println("数据转换错误"+err.Error())
		return shim.Error(err.Error())
	}
	var product = product{
		Id: id,
		Name: name,
	}

	var productstorage  = productStorage{
		Id: id,
		Storage: storage,
	}

	productData, err := json.Marshal(product)
	if err != nil{
		fmt.Println("序列化失败"+err.Error())
		return shim.Error(err.Error())
	}

	productstorageData, err := json.Marshal(productstorage)
	if err != nil{
		fmt.Println("序列化失败"+err.Error())
		return shim.Error(err.Error())
	}

	if err = stub.PutPrivateData(CollectionProductName,id, productData);err != nil{
		fmt.Println("Product写入数据失败"+err.Error())
		return shim.Error(err.Error())
	}
	if err = stub.PutPrivateData(CollectionProductStorageName,id, productstorageData);err != nil{
		fmt.Println("ProductStorage写入数据失败"+err.Error())
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}


func getProduct(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//判断参数个数是否为1
	if len(args) != 1 {
		return shim.Error("参数只能有一个")
	}


	data, err := stub.GetPrivateData(CollectionProductName,args[0])
	if err != nil{
		fmt.Println("获取Product失败"+err.Error())
		return shim.Error(err.Error())
	}

	fmt.Println(string(data))
	//返回数据给链码调用者
	return shim.Success(data)
}

func getProductStorage(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//判断参数个数是否为1
	if len(args) != 1 {
		return shim.Error("参数只能有一个")
	}


	data, err := stub.GetPrivateData(CollectionProductStorageName,args[0])
	if err != nil{
		fmt.Println("获取ProductStorage失败"+err.Error())
		return shim.Error(err.Error())
	}

	fmt.Println(string(data))
	//返回数据给链码调用者
	return shim.Success(data)
}


//函数入口
func main() {
	//启动链码：将继承了Chaincode接口的FristChainCode结构体作为入参，此处使用了多态
	err := shim.Start(new(FristChainCode))

	if err != nil {
		fmt.Printf("链码执行失败:%s", err)
	}
}
