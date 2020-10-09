package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
)

type Stu struct {
	Name string `json:"name"`
}
type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func main() {
	//实例化一个数据结构，用于生成json字符串
	stu := Stu{
		Name: "张三",
	}
	//Marshal

	jsonStu, err := json.Marshal(stu)
	if err != nil {
		fmt.Println("生成json字符串错误")
	}

	//jsonStu是[]byte类型，转化成string类型便于查看
	// 要有一个结构体才能有value：key这样子呈现出来
	fmt.Println("persistance: " + string(jsonStu))
	msg, err := json.Marshal("张三")
	fmt.Println("直接给: " + string(msg))

	// Unmarshal
	// 如果用map[string]interface{} 的话 number大于6位数就会自动传科学记数法
	//jsonStr := `{"number":1234567.888}`
	jsonStr := `{"number":1234567}`
	result := make(map[string]interface{})
	err = json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result["number"])
	// 1. 使用类型强制转换
	fmt.Println(int(result["number"].(float64)))
	// 2. 直接用结构体定义 ，然后var result Num再赋值就好了
	type Num struct {
		Number int `json:"number"`
	}
	var structresult Num
	err = json.Unmarshal([]byte(jsonStr), &structresult)
	fmt.Println(structresult.Number)
	// 3.使用UserNumber()方法
	d := json.NewDecoder(bytes.NewReader([]byte(jsonStr)))
	d.UseNumber()
	err = d.Decode(&result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result["number"])
	// 此时则要注意数据类型
	fmt.Println(fmt.Sprintf("type: %v", reflect.TypeOf(result["number"])))
	// 可转换其他类型：
	numInt, _ := result["number"].(json.Number).Int64()
	fmt.Println(fmt.Sprintf("value: %v, type: %v", numInt, reflect.TypeOf(numInt)))
	// 输出
	// map[number:1.234567e+06]

	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
	// fmt.Printf("%v: \n", vc) // {Jan Kersschot [0x126d2b80 0x126d2be0] none}:
	// JSON format:

	// 这里缺乏err处理
	js, _ := json.Marshal(vc)
	fmt.Printf("JSON format: %s", js)
	// using an encoder:
	// 创建一个文件
	file, _ := os.OpenFile("vcard.json", os.O_CREATE|os.O_WRONLY, 0666)
	// return后关闭文件
	defer file.Close()
	// 返回一个Encoder类型得指针，可调用方法Encode，将数据对象写入
	enc := json.NewEncoder(file)
	err = enc.Encode(vc)
	if err != nil {
		log.Println("Error in encoding json")
	}
}
