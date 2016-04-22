package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
)

// ROW struct
type ROW struct {
	YMD    int32
	Open   int32
	High   int32
	Low    int32
	Close  int32
	Turnover float32
	Volumn int32
	B      int32
}

func main() {
	bts, err := ioutil.ReadFile("sh600000.day")
	if err != nil {
		log.Fatalln(err)
	}
	// var ymd uint32
	// buf := bytes.NewReader(bts[4:8])
	// err = binary.Read(buf, binary.LittleEndian, &ymd)
	// if err != nil {
	// 	fmt.Println("binary.Read failed:", err)
	// }
	// fmt.Printf("%d\n", ymd)
    // a := make(map[uint32]uint32)
	rn := len(bts) / 32
	// for i := 0; i < rn; i++ {
	// 	buf := bytes.NewReader(bts[32*i : 32*(i+1)])
	// 	var row ROW
	// 	err = binary.Read(buf, binary.LittleEndian, &row)
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// 	fmt.Println(row.YMD, row.Open, row.High, row.Low, row.Close, row.Turnover, row.Volumn, row.B)
    //     a[row.B]++
	// }
    // for k, v := range a {
    //     fmt.Println(k,v )
    // }
    rows := make([]ROW, rn)
    buf := bytes.NewReader(bts)
    err = binary.Read(buf, binary.LittleEndian, rows)
    fmt.Println(rows)
}
