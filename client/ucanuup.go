/*设计要点：
1.按行读取并处理，因为目标文本大小未知，可能是10G，避免一次占用大量内存
2.按行处理时采用多协程并发处理以提升效率
*/
package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
)

var Matchstr = "UCanUup"
var count int

func handle(msg string, ch chan int) {
	ct := strings.Count(msg, Matchstr)
	ch <- ct
}
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fd, err := os.Open(`E:\QMDownload\UCloud.txt`)
	if err != nil {
		log.Fatal("open file err:", err)
	}
	var chs []chan int
	rder := bufio.NewReader(fd)
	for i := 0; ; i++ {
		ln, err := rder.ReadString('\n')
		if err != nil && err != io.EOF {
			log.Fatal("rder.ReadString:", err)
		}
		if err == io.EOF {
			break
		}
		chi := make(chan int)
		chs = append(chs, chi)
		go handle(ln, chs[i])
	}
	for _, ch := range chs {
		ct := <-ch
		count += ct
	}
	log.Println("match count:", count)
}
