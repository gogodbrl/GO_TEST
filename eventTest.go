package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"runtime"
	"sync"
)

func parser(ch_list []chan []string, wg *sync.WaitGroup){
	fi, err := os.Open("/dev/shm/TEST/sample_input.txt")
	if err != nil{
		panic(err)
	}
	defer fi.Close()

	data :=make([]string, 0)

	scanner :=bufio.NewScanner(fi)
	scanner.Split(bufio.ScanLines)

	var result_slice []string = make([]string,0)
	for scanner.Scan() {
		data = strings.Split(scanner.Text(),",")
		if data[0] == "END" {
			for ch_idx:= range ch_list{
				ch_list[ch_idx] <-result_slice
			}
			break
		}
		data2, _ := strconv.Atoi(data[2])
		data3, _ := strconv.Atoi(data[3])
		var num int = data2 + data3

		if num % 2 == 0 {
			tmp_str := scanner.Text()+",0"
			result_slice = append(result_slice, tmp_str)
		} else{
			tmp_str := scanner.Text()+",1"
			result_slice = append(result_slice, tmp_str)
		}
		if len(result_slice) == 100000 {
			for ch_idx := range ch_list {
				ch_list[ch_idx] <- result_slice
			}
			result_slice = nil
		}
	}
	wg.Done()
}

func summary(ch chan []string, key_list []int, val_list []int, wg *sync.WaitGroup) {
	result_map := make(map[string] []int)
	keys := ""
	vals := ""
	for k := range key_list {
		keys = keys + strconv.Itoa(key_list[k])
	}
	for v := range val_list {
		vals = vals + strconv.Itoa(val_list[v])
	}

	fmt.Println("go chan "+keys+"_"+vals+"START")
	for {
		result_slice := <- ch
		for data := range result_slice {
			split_data := strings.Split(result_slice[data], ",")

			key := ""
			for k := range key_list {
				key += split_data[k]
			}

			_, ok := result_map[key]
			if ok { //has key
				for idx :=0; idx < len(val_list); idx ++ {
					add_val, _ := strconv.Atoi(split_data[val_list[idx]])
					result_map[key][idx] += add_val
				}
			}else{
				result_map[key] = make([]int, len(val_list))
				for idx :=0; idx <len(val_list); idx ++{
					add_val, _ := strconv.Atoi(split_data[val_list[idx]])
					result_map[key][idx] += add_val
				}
			}
		}
		if len(result_slice) != 100000 {
			break
		}else {
			result_slice = nil
		}
	}
	file_path := "/dev/shm/TEST/sample_output_test_"+keys+"_"+vals+".txt"
	fo, err := os.Create(file_path)
	if err != nil {
		panic(err)
	}
	defer fo.Close()

	writer := bufio.NewWriter(fo)
	for key, value := range result_map {
		
		writer.WriteString(key)
		for i := range value {
			line := strconv.Itoa(value[i])
			writer.WriteString("," + line)
		}
		writer.WriteString("\n")
	}
	writer.Flush()
	wg.Done()
}
func main() {
	var wg sync.WaitGroup

	runtime.GOMAXPROCS(2)
	ch :=make(chan []string, 100000)
	ch1 :=make(chan []string, 100000)
	ch2 :=make(chan []string, 100000)

	var ch_list []chan []string
	
	ch_list = append(ch_list, ch)
	ch_list = append(ch_list, ch1)
	ch_list = append(ch_list, ch2)
	
	wg.Add(1)
	go parser(ch_list, &wg)
	wg.Add(1)
	go summary(ch, []int {1}, []int{3}, &wg)
	wg.Add(1)
	go summary(ch1,[]int {0,1}, []int {2,3},  &wg)
	wg.Add(1)
	go summary(ch2,[]int {0}, []int {2,3},  &wg)
	
	wg.Wait()
	fmt.Println("TEST DONE")
}
