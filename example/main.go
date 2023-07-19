/**
* @program: utils
*
* @description:
*
* @author: lemo
*
* @create: 2020-07-11 13:08
**/

package main

import (
	"github.com/lemonyxk/utils/array"
	"log"
)

type A struct {
	Name string `json:"name"`
}

type B struct {
	Name int `json:"name"`
}

func main() {

	// log.Println(utils.Sort(1, 2, -1).Asc(1))
	// log.Println(utils.Sort(1, 2, -1).Desc())
	//
	// var a = utils.Ternary(true, 2, 1)
	//
	// log.Println(a)
	//
	// var b = 2
	//
	// b = a
	//
	// log.Println(b)

	// var a = A{Name: "hello"}
	// var b = A{Name: "world"}
	// var c = []A{a, b}
	// var d = []any{a, b, c}
	// var e = map[string]any{"Name": 111}
	// var res = utils.Extract.Src(d).Field("Name").String()
	// log.Println(res)
	//
	// log.Println(utils.Structure.GetTags(A{}))
	//
	// utils.Assign.Dest(&a).Src(&e).AllowWeak().Do()
	//
	// log.Println(a)
	//
	// var aa = A{Name: "50"}
	// var bb = B{Name: 11111}
	// utils.Assign.Dest(&bb).Src(&aa).AllowWeak().Do()
	// log.Println(bb)

	// var res = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//
	// var arr = utils.Array(&res)
	//
	// var add = []int{100, 200, 300, 400, 500, 600, 700, 800, 900, 1000, 1100, 1200, 1300, 1400, 1500, 1600, 1700, 1800, 1900, 2000, 2100, 2200, 2300, 2400, 2500, 2600, 2700, 2800, 2900, 3000, 3100, 3200, 3300, 3400, 3500, 3600, 3700, 3800, 3900, 4000, 4100, 4200, 4300, 4400, 4500, 4600, 4700, 4800, 4900, 5000}
	//
	// arr.Splice(3, 2, add...)
	//
	// log.Println(res)

	var res = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var arr = array.Ordered(res)

	log.Println(arr.Slice(3, 5))

	var a = []int{1, 2, 3, -9}

	array.Any(a).Sort(func(a int, b int) bool {
		return b > a
	})

	log.Println(a)

	// log.Println(utils.Compress.From("./").TarGz("./z.tar.gz"))
	// log.Println(utils.Compress.From("../asasasa/z.tar.gz").UnTarGz("../asasasa"))

	// log.Println(utils.Compress.From("./").Zip("./z.zip"))
	// log.Println(utils.Compress.From("../asasasa/z.zip").UnZip("../asasasa/"))

	// log.Println(arr.Slice(0, 1))
}
