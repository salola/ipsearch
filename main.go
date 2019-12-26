package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type ipHashTable map[byte]ipHashTable

var rootHashTable ipHashTable

var hashTablesCount int = 0

var adr1 = [4]byte{127, 0, 0, 1}
var adr2 = [4]byte{127, 0, 10, 2}
var adr3 = [4]byte{192, 168, 1, 2}

func addIPToHashTable(root *ipHashTable, addr [4]byte, level byte) {

	newroot, found := (*root)[addr[level]]

	if !found {
		if level < 3 {
			(*root)[addr[level]] = make(ipHashTable, 255)
			hashTablesCount++
			addIPToHashTable(root, addr, level)
		} else {
			(*root)[addr[level]] = nil
			return
		}

	} else {
		if level == 3 {
			return
		}
		addIPToHashTable(&newroot, addr, level+1)
	}

}

func loadIPfromFile(fileName string) {

	file, err := os.Open(fileName)

	var normalip int = 0
	var badip int = 0

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	r, _ := regexp.Compile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)
	fmt.Print("Loading.")
	for scanner.Scan() {
		str := scanner.Text()
		if r.MatchString(str) {
			addIPToHashTable(&rootHashTable, parseIPtoArray(str), 0)
			normalip++
			//fmt.Print(".")
		} else {
			fmt.Println("\nBad IP:", str)
			badip++
		}
	}

	fmt.Println("\nNormal IP:", normalip, "Bad IP:", badip)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func parseIPtoArray(str string) [4]byte {

	var adr [4]byte

	strs := strings.Split(str, ".")

	for i := 0; i < 4; i++ {
		x, _ := strconv.Atoi(strs[i])
		adr[i] = byte(x)
	}

	return adr
}

func search(adr [4]byte) bool {
	_, found := rootHashTable[adr[0]][adr[1]][adr[2]][adr[3]]
	return found
}

func main() {
	rootHashTable = make(ipHashTable, 255)
	loadIPfromFile("ip_list.txt")
	fmt.Printf("HashTables added %d\n", hashTablesCount)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter IP")
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		start := time.Now()
		res := search(parseIPtoArray(text))
		elapsed := time.Since(start)
		fmt.Println(res)
		fmt.Printf("Search took %s\n", elapsed)

	}

}