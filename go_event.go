package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"io/ioutil"
)

func parser(ch chan []string){
	fi, err := os.Open("sample_raw.txt")
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
			result_slice = append(result_slice,data[0])
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
	}
	ch <- result_slice
}

func summary(ch chan []string){
	result_map := make(map[string][]int)
	result_slice := <-ch

	for data:= range result_slice{
		if result_slice[data] == "END" {
			break		
		}
		split_data := strings.Split(result_slice[data],",")
		key := split_data[0]+split_data[1]

		if _, ok := result_map[key]; ok == false{
			result_map[key] = []int{0,0}
		}
		//ab의 값중에 짝수는 x번 , 홀수는 y번 출현했다.	
		odd_count := result_map[key][0]
		eval_count := result_map[key][1]
		show_data, _ := strconv.Atoi(split_data[4])

		if show_data == 0 { //eval
			eval_count = eval_count+1
			result_map[key] = []int{odd_count,eval_count}  
		}else {
			odd_count = odd_count+1
			result_map[key] = []int{odd_count,eval_count}
		}
	}
	fo, err := os.Create("go_sample_result.txt")
	if err != nil {
		panic(err)
	}
	defer fo.Close()
	for key, value := range result_map {
		fo.Write([]byte(value))
	}
}
func main() {
	ch :=make(chan []string,10)
	go parser(ch)
	go summary(ch)

	fmt.Scanln()
}
