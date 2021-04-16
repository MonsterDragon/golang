package main

import "fmt"

type MobileStorager interface {
	read()
	write()
}

type MobileDisks struct {

}

func (d *MobileDisks) read() {
	fmt.Println("移动硬盘在读取数据！")
}

func (d *MobileDisks) write() {
	fmt.Println("移动硬盘在写数据！")
}

type UDisk struct {

}

func (u *UDisk) read() {
	fmt.Println("U盘在读取数据！")
}

func (u *UDisk) write() {
	fmt.Println("U盘在写数据！")
}

func main() {
	m := MobileDisks{}
	u := UDisk{}
	x := []MobileStorager{&m, &u}
	for i:=0; i<2; i++ {
		x[i].read()
		x[i].write()
	}
}
