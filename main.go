package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strconv"
	"bytes"
)

func testErr(err error){
	if err != nil{
		log.Fatal(err);
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		//Get Resource usage
		free, err := exec.Command("free", "-h").Output()
		testErr(err);
		//Get Uptime
		up, err := exec.Command("uptime").Output()
		testErr(err);
		//Get and convert CPU Temperature
		tempRaw, err := ioutil.ReadFile("/sys/class/thermal/thermal_zone0/temp")
		tempRaw = bytes.Trim(tempRaw, "\n");
		testErr(err);
		tempFloat, err := strconv.ParseFloat(string(tempRaw), 32);
		testErr(err);
		tempStr := fmt.Sprintf("%f", tempFloat/1000);

		//Format Response
		result := fmt.Sprintf("{\"free\":%q, \"uptime\":%q, \"cputemp\":%q}", string(free), string(up), tempStr);

		// Server Headers
		w.Header().Set("Content-Type", "application/json");
		w.Header().Set("Access-Control-Allow-Origin", "*");
		fmt.Fprintf(w, result);
	});

	log.Fatal(http.ListenAndServe(":8081", nil));

}
