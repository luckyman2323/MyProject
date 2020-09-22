package main

import "fmt"

type Man interface {
	getName() string
	getAge()	int
	getCountry() string
}

type Chinese struct {
	 Name 	 string
	 Age  	 int
	 Country string

}
type American struct {
	Name 	 string
	Age  	 int
	Country string

}
func (c Chinese) getName() string{
	return c.Name
}
func (c Chinese) getAge() int{
	return c.Age
}
func (c Chinese) getCountry() string{
	return c.Name
}

func (a American) getName() string{
	return a.Name
}
func (a American) getAge() int{
	return a.Age
}
func (a American) getCountry() string{
	return a.Name
}
//发送管道信息
func (c *Chinese)SendMessage(msg chan<- string,args []string){
	for _, v := range args {
		msg <- v
	}
	close(msg)
}
//获取管道信息
func (a *American)GetMessage(msg <-chan string){
	for {
		select {
		case m := <- msg:
			if m == "Hello World"{
				fmt.Println("收到信息：",m)
				continue
			}else if m == "Hello Kitty"{
				fmt.Println("收到信息：",m)
				continue
			}else if m == "break" {
				return
			}
			fmt.Println("非法消息！")
		}
	}
}
func main(){
	var man Man
	man = Chinese{"张三",23,"中国"}
	fmt.Println("age: ",man.getAge())
	fmt.Println("name: ",man.getName())
	fmt.Println("country: ",man.getCountry())
	man = American{"tom",33,"美国"}
	fmt.Println("age: ",man.getAge())
	fmt.Println("name: ",man.getName())
	fmt.Println("country: ",man.getCountry())

	c := Chinese{}
	a := American{}
	msg := make(chan string,10)
	var args []string =  []string{"Hello World","Hello Kitty","error test","break"}
	c.SendMessage(msg,args)
	a.GetMessage(msg,)
}