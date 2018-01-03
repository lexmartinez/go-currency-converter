package main

import (
	"os"
	"fmt"
	"net/http"
	"strconv"
	"io/ioutil"
	"strings"
)

func main() {
	args := os.Args[1:]
	qty, err := strconv.ParseFloat(args[0], 64)

	if err == nil {
		rate := rate(args[1], args[2])
		fmt.Printf("%.2f", (rate * qty))
	}

}

func rate(from, to string) float64 {

	url := "https://free.currencyconverterapi.com/api/v5/convert?q=" + from + "_" + to +"&compact=y"

	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
		return 0
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err2 := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)

		if (err2 == nil){
			strval := strings.Replace(bodyString, "{\""+ from + "_" + to +"\":{\"val\":", "", 1)
			strval = strings.Replace(strval, "}}", "", 1)
			rate, err3 := strconv.ParseFloat(strval, 64)
			if (err3 == nil){
				return rate
			}
		}
	}
	return 0
}
