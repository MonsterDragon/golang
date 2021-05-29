package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var time_original int = 1598306400
var usage_string string = "Usage:\n\ttimecheck [cluster_nums] (single or mutiple)\nFor example:\n\ttimecheck f054420 || timecheck f054418 f0396720 f0227684"

func getTime() (last_time int, next_time int) {

	f, err := os.Open("/home/shuzhan/proving-info")
	if err != nil {
		fmt.Println("err= ", err)
		return
	}
	defer f.Close()

	r := bufio.NewReader(f)

	for {
		buf, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("err= ", err)
		}
		string_line := string(buf)
		if ok := strings.Contains(string_line, "Proving Period Start"); ok {
			str_slice := strings.Split(string_line, ":")
			str_slice_2 := strings.Trim(str_slice[1], " ")
			str_slice_3 := strings.Split(str_slice_2, " ")
			time_space, err := strconv.Atoi(str_slice_3[0])
			if err !=nil {
				fmt.Println("string to int, err: ", err)
				break
			}
			last_time = time_space * 30 + time_original
		} else if next := strings.Contains(string_line, "Next Period Start"); next {
			next_slice := strings.Split(string_line, ":")
			next_slice_2 := strings.Split(next_slice[1], " ")
			time_space_2, err := strconv.Atoi(next_slice_2[7])
			if err != nil {
				fmt.Println("string to int, err: ", err)
				break
			}
			next_time = time_space_2 * 30 + time_original
		}
	}
	return
}

func getStep(start_time int) (finish_time int) {
	f, err := os.Open("/home/shuzhan/proving-deadlines")
	if err != nil {
		fmt.Println("err= ", err)
		return
	}
	defer f.Close()

	r := bufio.NewReader(f)
	i := 0
	sum := 0

	for {
		buf, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("err= ", err)
		}
		if i != 0 && i != 1 {
			if i>=2 && i <=11 {
				string_line := string(buf)
				string_slice := strings.Split(string_line, " ")
				if string_slice[9] == "1" {
					sum++
				}
			} else {
				string_line := string(buf)
				string_slice := strings.Split(string_line, " ")
				if string_slice[8] == "1" {
					sum++
				}
			}
		}
		i++
	}

	finish_time = start_time + sum * 1800

	return finish_time
}

func writeFile(path string, info string)  {
	f, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	buf := []byte(info)
	n, err := f.Write(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(buf) != n {
		os.Exit(2)
	}
	defer f.Close()
}

func exec_comm(name string, args ...string) (result string) {
	cmd := exec.Command(name, args...)
	// stdout为输出对象，类似于文件
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	// 函数结束关闭文件句柄
	defer stdout.Close()
	// 执行命令
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	// 读取输出结果
	opBytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		log.Fatal(err)
	}

	result = string(opBytes)
	return
}

func getCluster() (clusterName []string) {
	s := make([]string, 10)
	if len(os.Args) == 1 {
		fmt.Printf("%s\n", usage_string)
		os.Exit(2)
	} else {
		s = os.Args
	}
	return s
}

func main() {
	cluster_names := getCluster()
	for m, n := range cluster_names {
		if m != 0 {
			info_result := exec_comm("lotus-poster", "--actor="+n, "proving", "info")
			writeFile("/home/shuzhan/proving-info", info_result)
			deadline_result := exec_comm("lotus-poster", "--actor="+n, "proving", "deadlines")
			writeFile("/home/shuzhan/proving-deadlines", deadline_result)
			last_time, next_time := getTime()
			finish_time := getStep(last_time)

			format_lasttime := exec_comm("date", "-d", "@"+strconv.Itoa(last_time-60*10), "+"+"\"%Y-%m-%d %H:%M:%S\"")
			format_nexttime := exec_comm("date", "-d", "@"+strconv.Itoa(next_time+60*20), "+"+"\"%Y-%m-%d %H:%M:%S\"")
			format_finishtime := exec_comm("date", "-d", "@"+strconv.Itoa(finish_time), "+"+"\"%Y-%m-%d %H:%M:%S\"")
			fmt.Printf("集群%s:\n", n)
			fmt.Printf("\t上一次时空证明启动时间是：%s\t上一次时空证明结束时间是：%s\n", format_lasttime, format_finishtime)
			fmt.Printf("\t下一次时空证明启动时间是：%s\n", format_nexttime)
		}
	}
}


