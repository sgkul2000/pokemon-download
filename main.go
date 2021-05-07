package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/apoorvam/goterminal"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Please provide valid epidode numbers.")
		os.Exit(1)
	}
	i, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Please provide valid epidode numbers.")
		os.Exit(1)
	}
	n, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Please provide valid epidode numbers.")
		os.Exit(1)
	}

	if _, err := os.Stat("Pokemon"); os.IsNotExist(err) {
		os.Mkdir("Pokemon", 0777)
	}

	start := time.Now()

	for ; i <= n; i++ {
		url := fmt.Sprintf("https://streamcdnuservi4api.in/xiWvvKV/pokmn/pkmn1_IL/Dub/%v_Pokemon_Dubbed_Pokemon360.com.mp4", GetNumber(i))
		DownloadFile(url, fmt.Sprintf("Pokemon-%v.mp4", GetNumber(i)))
	}

	elapsed := time.Since(start)
	fmt.Printf("Download completed in %s\n", elapsed)

}

func GetNumber(i int) string {
	if 1/10 < 1 {
		return "0" + fmt.Sprint(i)
	} else {
		return fmt.Sprint(i)
	}
}

func PrintDownloadPercent(done chan int64, path string, total int64, file string) {

	var stop bool = false
	writer := goterminal.New(os.Stdout)

	for {
		select {
		case <-done:
			writer.Clear()
			writer.Reset()
			fmt.Printf("Downloaded: %v\n", file)
			stop = true
		default:

			file, err := os.Open(path)
			if err != nil {
				log.Fatal(err)
			}

			fi, err := file.Stat()
			if err != nil {
				log.Fatal(err)
			}

			size := fi.Size()

			if size == 0 {
				size = 1
			}

			var percent float64 = float64(size) / float64(total) * 100
			// fmt.Println(percent)
			writer.Clear()
			// add your text to writer's buffer
			fmt.Fprintf(writer, "Downloading %v [%v>%v]\n", strings.Repeat("=", int(percent/4)), file, strings.Repeat(" ", int((100-percent)/4)))

			// write to terminal
			writer.Print()

		}

		if stop {
			break
		}

		time.Sleep(time.Second)
	}
}

func DownloadFile(url string, dest string) {

	file := dest

	var path bytes.Buffer
	path.WriteString("Pokemon")
	path.WriteString("/")
	path.WriteString(file)

	out, err := os.Create(path.String())

	if err != nil {
		fmt.Println(path.String())
		panic(err)
	}

	defer out.Close()

	headResp, err := http.Head(url)

	if err != nil {
		panic(err)
	}

	defer headResp.Body.Close()

	size, err := strconv.Atoi(headResp.Header.Get("Content-Length"))

	if err != nil {
		panic(err)
	}

	done := make(chan int64)

	go PrintDownloadPercent(done, path.String(), int64(size), file)

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	n, err := io.Copy(out, resp.Body)

	if err != nil {
		panic(err)
	}

	done <- n

}
