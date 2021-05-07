package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
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
		fmt.Println("hello")
		os.Mkdir("Pokemon", 0777)
	}
	for ; i <= n; i++ {
		url := fmt.Sprintf("https://streamcdnuservi4api.in/xiWvvKV/pokmn/pkmn1_IL/Dub/%v_Pokemon_Dubbed_Pokemon360.com.mp4", GetNumber(i))
		err := DownloadFile(url, fmt.Sprintf("Pokemon/Pokemon-%v", GetNumber(i)))
		if err != nil {
			panic(err)
		}
	}

}

// DownloadFile will download a url and store it in local filepath.
// It writes to the destination file as it downloads it, without
// loading the entire file into memory.
func DownloadFile(url string, filepath string) error {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func GetNumber(i int) string {
	if 1/10 < 1 {
		return "0" + fmt.Sprint(i)
	} else {
		return fmt.Sprint(i)
	}
}
