package main

import (
	"fmt"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/stretchr/testify/assert"
)

//用于初始化链码
func checkInit(t *testing.T,stub *shim.MockStub, args [][]byte){
	res := stub.MockInit("1",args)//uuid string args [][]byte
	//shime.OK: 200  res.Status: 200
	assert.Equal(t,int32(shim.OK),res.Status,"Init failed: %v",string(res.Message))//返回true
}

//用于调用链码
func checkInvoke(t *testing.T,stub *shim.MockStub, args[][]byte)(data []byte){
	res := stub.MockInvoke("1",args)
	if assert.Equal(t,int32(shim.OK), res.Status, "Invoke failded:%s",string(res.Message)){
	}

	return res.Payload
}

func Test(t *testing.T){

	cc := new(FristChainCode)
	//通过FristChainCode结构体获取Mock对象
	stub := shim.NewMockStub("FristChainCode",cc)
	//通过Mock对象进行Init操作
	checkInit(t, stub, [][]byte{[]byte("init")})


	var argsCraete [][]byte
	//定义参数User参数,[]byte("方法名"), []byte("参数列表"), []byte("参数列表")
	argsCraete = append(argsCraete,[]byte(CreatProductFunc),[]byte("1001"),[]byte("锤子手机"),[]byte("6666"))
	//test调用链码
	checkInvoke(t,stub,argsCraete)


	var getProductArgs [][]byte
	//进行Id查找的test
	getProductArgs = append(getProductArgs,[]byte(GetProductFunc),[]byte("1001"))
	data := checkInvoke(t,stub,getProductArgs)
	fmt.Println(string(data))

	var getProductStorageArgs [][]byte
	//进行Id查找的test
	getProductStorageArgs = append(getProductStorageArgs,[]byte(GetProductStorageFunc),[]byte("1001"))
	data = checkInvoke(t,stub,getProductStorageArgs)
	fmt.Println(string(data))

	return
}

