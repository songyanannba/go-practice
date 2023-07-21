package main

import (
	"fmt"
	"strings"
)

func main() {

	//wild_1&null&bar_high&null&bar_low&null&wild_1&null&bar_7&
	//null&bar_high&null&wild_1&null&7_high&null&bar_med&null&
	//bar_7&null&7_high&null&wild_1&null&bar_med&null&bar_low
	//&null&wild_1&null&bar_med&null&bar_7&null
	//wild_100&null&bar_high&null&bar_low&
	//null&wild_100&null&bar_7&null&bar_high
	//&null&wild_100&null&7_high&null&bar_med&
	//null&bar_7&null&7_high&null&wild_100&
	//null&bar_med&null&bar_low&null&wild_100
	//&null&bar_med&null&bar_7&null&wild_10
	//&null&bar_high&null&bar_low&null&wild_10
	//&null&bar_7&null&bar_high&null&wild_10&
	//null&7_high&null&bar_med&null&bar_7
	//&null&7_high&null&wild_10&null&
	//bar_med&null&bar_low&null&wild_10&
	//null&bar_med&null&bar_7&null&wild_5&null&
	//bar_high&null&bar_low&null&wild_5&
	//null&bar_7&null&bar_high&null&wild_5
	//&null&7_high&null&bar_med&null&bar_7
	//&null&7_high&null&wild_5&null&bar_med
	//&null&bar_low&null&wild_5&null&
	//bar_med&null&bar_7&null&wild_2&
	//null&bar_high&null&bar_low&null&
	//wild_2&null&bar_7&null&bar_high&null&w
	//ild_2&null&7_high&null&bar_med&null&
	//bar_7&null&7_high&null&wild_2&null&
	//bar_med&null&bar_low&null&wild_2&
	//null&bar_med&null&bar_7&null

	//b_slot_event
	//[0,0]&[0,1]&[1,2]&[2,5]&[5,10]&[10,20Strings方法练习]&[20Strings方法练习,30]&[30,40]
	//&[40,50]&[50,100]&[100,200]&[200,500]&[500,1000]&[1000,10000]&[10000,99999]
	//@1&6308251&9308251&9708251&9808251&9848251&9938251&9963251&9975751&9981951
	//&9991951&9993951&9995951&9997951&9999951&10000001

	//160@4@ 16&20Strings方法练习&24@ 1&4&6&7 @0&1&2&3@ 1&2&2&2&2
	//20Strings方法练习@
	//2@
	//8&10&12@
	//1&4&6&7@
	//0&1&2&3@
	//1&1&2&2&2

	//1&2&3&4&5&6&7&8&9&10&11&12&13&14&15&16&17&18&19&20Strings方法练习
	//@
	//1&1&1&1&1&6&11&16&21结构体的复制练习传指针&23&24&24&24&24&24&24&24&24&24&24&24

	//wild_1&null&bar_high&null&bar_low&null&wild_1&null&bar_7&null&
	//bar_high&null&wild_1&null&7_high&null&bar_med&null&bar_7&null&7_high&null&
	//wild_1&null&bar_med&null&bar_low&null&wild_1&null&bar_med&null&bar_7&null
	//wild_100& wild_10& wild_5& wild_2& wild_1& 7_high& bar_7&
	//bar_high& bar_med& bar_low& null  //11
	//@
	//1& 1& 1& 1& 1& 501& 1301& 1751& 2301& 2951& 3701& 7401 //12

	str := "3_45"

	before, after, found := strings.Cut(str, "_")

	fmt.Println("before, after, found ", before, after, found)
}
