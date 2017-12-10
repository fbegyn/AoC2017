package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	list := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, 125, 126, 127, 128, 129, 130, 131, 132, 133, 134, 135, 136, 137, 138, 139, 140, 141, 142, 143, 144, 145, 146, 147, 148, 149, 150, 151, 152, 153, 154, 155, 156, 157, 158, 159, 160, 161, 162, 163, 164, 165, 166, 167, 168, 169, 170, 171, 172, 173, 174, 175, 176, 177, 178, 179, 180, 181, 182, 183, 184, 185, 186, 187, 188, 189, 190, 191, 192, 193, 194, 195, 196, 197, 198, 199, 200, 201, 202, 203, 204, 205, 206, 207, 208, 209, 210, 211, 212, 213, 214, 215, 216, 217, 218, 219, 220, 221, 222, 223, 224, 225, 226, 227, 228, 229, 230, 231, 232, 233, 234, 235, 236, 237, 238, 239, 240, 241, 242, 243, 244, 245, 246, 247, 248, 249, 250, 251, 252, 253, 254, 255}

	prob1(list)
	prob2(list)
}

func prob1(li []int) {
	list := make([]int, len(li))
	copy(list, li)
	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	scanner := bufio.NewScanner(f)
	var line string
	if scanner.Scan() {
		line = scanner.Text()
	}
	strNumbers := strings.Split(line, ",")
	sequence := make([]int, len(strNumbers))
	for i, el := range strNumbers {
		number, err := strconv.Atoi(el)
		if err != nil {
			log.Fatalln(err)
		}
		sequence[i] = number
	}
	hash(list, sequence, 1)
	fmt.Printf("Multiplication of the first 2 numbers = %d\n", list[0]*list[1])
}

func prob2(list []int) {
	chars, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	sequence := make([]int, len(chars))
	for i, ch := range chars {
		sequence[i] = int(ch)
	}
	sequence = append(sequence, []int{17, 31, 73, 47, 23}...)
	hash(list, sequence, 64)
	fmt.Print("The hash is: \n")
	for i := 16; i <= 256; i += 16 {
		hashPart := 0
		for _, el := range list[i-16 : i] {
			hashPart ^= el
		}
		fmt.Printf("%.02x", hashPart)
	}
	fmt.Println()
}

func mod(number int, length int) int {
	if number < 0 {
		return mod(length+number, length)
	} else {
		return number % length
	}
}

func reverse(list []int, start int, stop int, length int) {
	copyList := make([]int, len(list))
	copy(copyList, list)
	for i := 0; i < length; i++ {
		index1 := mod(start+i, len(list))
		index2 := mod(stop-i, len(list))
		list[index1] = copyList[index2]
	}
}

func hash(list []int, inputLengths []int, numberRounds int) {
	start := 0
	skipSize := 0
	for i := 0; i < numberRounds; i++ {
		for _, length := range inputLengths {
			startInvert := (start) % len(list)
			endInvert := (start + length - 1) % len(list)
			reverse(list, startInvert, endInvert, length)
			start += length + skipSize
			skipSize++
		}
	}

}
