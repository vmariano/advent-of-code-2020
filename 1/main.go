package main

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
)

func loadInputFile() []int {
	fileContent, err := ioutil.ReadFile("./1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	const SIZEOF_INT32 = 4 // bytes
	data := make([]int, len(fileContent)/SIZEOF_INT32)
	for i := range data {
		data[i] = int(binary.LittleEndian.Uint32(fileContent[i*SIZEOF_INT32 : (i+1)*SIZEOF_INT32]))
	}
	return data
}

func getFullList() []int {
	return loadInputFile()
}

func getPartialList() []int {
	return []int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}
}

func twoEntries(numberList []int) {
	for i := 0; i < len(numberList); i++ {
		for j := i + 1; j < len(numberList); j++ {
			a := numberList[i]
			b := numberList[j]
			if (a + b) == 2020 {
				fmt.Print("Number ", a)
				fmt.Println(" and number ", b)
				fmt.Println("multiply ", a*b)
			}
		}
	}
}

func threeEntries(numberList []int) {
	for i := 0; i < len(numberList); i++ {
		for j := i + 1; j < len(numberList); j++ {
			for k := j + 1; k < len(numberList); k++ {
				a := numberList[i]
				b := numberList[j]
				c := numberList[k]
				if (a + b + c) == 2020 {
					fmt.Print("Number ", a)
					fmt.Print(" , ", b)
					fmt.Println("  and ", c)
					fmt.Println("multiply ", a*b*c)
				}
			}
		}
	}
}

func main() {
	twoEntries(getPartialList())
	twoEntries(getFullList())
	threeEntries(getPartialList())
	threeEntries(getFullList())
}
