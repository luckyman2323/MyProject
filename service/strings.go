package service

import (
	"fmt"
	"strings"
)

var test  []string

//golang字符串操作
func StringsTest() {
	// s := "hello world hello world"
	// str := "wo"
	test = []string{"init","asset4","{\\\"docType\\\":\\\"couchdb\\\"","\\\"ID\\\":\\\"1yjcaygwr2\\\"",
	"\\\"color\\\":\\\"red\\\"",
	"\\\"size\\\":101",
	"\\\"owner\\\":\\\"小红\\\"",
	"\\\"appraisedValue\\\":1000}"}

	test2 := "init asset4 {\"docType\":\"couchdb\",\"ID\":\"1yjcaygwr2\",\"color\":\"red\",\"size\":101,\"owner\":\"小红\",\"appraisedValue\":1000}"

	//返回字符串str中的任何一个字符在字符串s中第一次出现的位置。
	//如果找不到或str为空则返回-1
	// index := strings.IndexAny(s, str)
	// fmt.Println(index) //4
	// index = strings.Index(s, str)
	// fmt.Println(index) //4

	res := strings.Fields(test2)
	// testString := strings.Join(test,",")
	fmt.Println(len(res))
}

func StringsTest2() {
	test1 := "abcd"
	test2 := "一二三四"

	fmt.Println("test1=", test1, "len=", len(test1))
	fmt.Println("test2=", test2, "len=", len(test2))

	test1Byte := []byte(test1)
	test2Byte := []byte(test2)


	fmt.Println("test1=", test1Byte, "len=", len(test1Byte), "cap", cap(test1Byte))
	fmt.Println("test2=", test2Byte, "len=", len(test2Byte), "cap", cap(test2Byte))
}


