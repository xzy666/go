package main

import (
	"flag"
	"fmt"
	"os"
	"bufio"
	"io"
	"strconv"
	"algorithm"
)

var infile *string = flag.String("i", "in.txt", "file contains values for sorting")
var outfile *string = flag.String("o", "out.txt", "file contains values for sorting")
var algo *string = flag.String("a", "sort", "sort algorithm")

func main() {
	flag.Parse()
	if infile != nil {
		fmt.Println("infile = ", *infile, "outfile = ", *outfile, "algorithm =", *algo)
	}
	values, err := readValues(*infile)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("read values : ", values)
		switch *algo {
		case "Qsort":
			algorithm.QSort(values,0,len(values)-1)
			break
		case "Bsort":
			algorithm.BSort(values)
			break
		default:
			break
		}
		fmt.Println(values)
		err := WriterFile(values, *outfile)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("OK")
		}
	}
}

/**
	数据读取
 */
func readValues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Failed to open the input file ", infile)
		return
	}
	defer file.Close()
	br := bufio.NewReader(file)
	values = make([]int, 0)
	for {
		line, isPrefix, errTmp := br.ReadLine()
		if errTmp != nil {
			if errTmp != io.EOF {
				err = errTmp
			}
			break
		}
		if isPrefix {
			fmt.Println("a too long line,seems unexpected.")
			return
		}
		str := string(line)
		value, errTmp := strconv.Atoi(str)
		if errTmp != nil {
			err = errTmp
			return
		}
		values = append(values, value)
	}
	return
}

/**
	数据写入
 */
func WriterFile(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("Failed to create file", err)
	}
	defer file.Close()
	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}
