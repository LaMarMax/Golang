package pair_go

import (
	"fmt"
)

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

// 具体的类型
type Book struct {
}

func (b *Book) ReadBook() {
	fmt.Println("Read a book")
}

func (b *Book) WriteBook() {
	fmt.Println("Write a book")
}

func Pair() {
	// b pair<type:Book, value:Book{}地址>
	b := &Book{}

	// r pair<type:, value:>
	var r Reader
	// r pair<type:Book, value:Book{}地址>
	r = b
	r.ReadBook()

	// r pair<type:, value:>
	var w Writer
	// r pair<type:Book, value:Book{}地址>
	w = r.(Writer) // 此处的断言为什么会成功，因为w r具体的type是一致的
	//因为Book重写了Reader和Writer，因此Reader和Writer都能指向Book对象，因此他们两个之间是能够互相断言的
	w.WriteBook()

	// //以读写状态打开当前Linux终端
	// //tty pair<tye:*os.File, value:"/dev/tty"文件描述符>
	// tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	// if err != nil {
	// 	fmt.Println("open file error", err)
	// 	return
	// }

	// // r pair<tye:  value:  >
	// var r io.Reader
	// // r pair<tye:*os.File, value:"/dev/tty"文件描述符>
	// r = tty

	// // w pair<tye:  value:  >
	// var w io.Writer
	// // w pair<tye:*os.File, value:"/dev/tty"文件描述符>
	// w = r.(io.Writer)

	// w.Write([]byte("hello, this is a test!!\n"))

	// var str string
	// // pair<type:string, value:"lisi">
	// str = "lisi"

	// var allType interface{}
	// // str的pair传递给allType了
	// // pair<type:string, value:"lisi">
	// allType = str

	// value, ok := allType.(string)
	// if ok {
	// 	fmt.Println(value)
	// }
}
