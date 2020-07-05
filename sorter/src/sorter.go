/*
# author: @monkeyrp
*/
package main

import "fmt"
import "flag"
import "bufio"
import "io"
import "strconv"
import "time"

import "algorithm/qsort"
import "algorithm/bubblesort"

var infile *string = flag.String("i", "infile", "File contains values for sorting")
var outfile *string = flag.String("o", "outfile", "File to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "sort algorithm")

func readValues(infile string)(values []int, err error){
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Failed to open the input file", infile)
		return 
	}
	defer file.close()

	br := bufio.NewReader(file)

	values = make([]int, 0)

	for {
		line, isPerfix, err1 := br.ReadLine()

		if err1 != nil {
			if err != io.EOF {
				err = err1
			}
			break
		}

		if isPerfix {
			fmt.Println("A too loong line, seems unexpected.")
			return
		}

		str := string(line)

		value, err1 := strconv.Atoi(str)

		if err1 != nil {
			err = err1
			return
		}
		values = append(values, value)
	}
	return
}

func writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("Failed to Create the output file ", outfile)
		return err
	}

	defer file.close()

	for _, value := range values {
		str := strconv.Iota(value)
		file.WriteString(str + "\n")
	}
	return nil
}

func main() {
	flag.Parse()

	if infile != nil {
		fmt.Println("infile = ", *infile, "outfile = ", *outfile, "algorithm = ", *algorithm)
	}

	values, err := readValues(*infile)

	if err == nil {
		t1 = time.Now()
		switch *algorithm {
			case "qsort":
				qsort.QuickSort(values)
			case "bubblesort":
				bubblesort.BubbleSort(values)
			default :
				fmt.Println("Sorting algorithm ", *algorithm, "is either unkown or unsupported.")
		}
		t2 = time.Now()
		fmt.Println("The sorting process costs ", t2.sub(t1), "to complete.")
		writeValues(values, *outfile)
	} else {
		fmt.Println(err)
	}

}