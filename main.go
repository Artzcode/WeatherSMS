package main

import (
	"net/http"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"github.com/tidwall/gjson" // a way to parse json easily
)

type Config struct {
	User string `json:"user"`
	Pass string `json:"pass"`
	ApiKey string `json:"apikey"`
}

func LoadConfig(filename string) (Config, error) {
	var config Config
	configFile, err := os.Open(filename)
	defer configFile.Close()
	if(err != nil) {
		return config, err
	}
	jsonParse := json.NewDecoder(configFile)
	err = jsonParse.Decode(&config)
	return config, err
}

func main() {

	config, _ := LoadConfig("config.json")

	user := config.User
	pass := config.Pass
	apiKey := config.ApiKey
	
	weather := string(getWeather(apiKey))
	
	location := gjson.Get(weather, "location.name").String()
	maxTemp := strings.Replace(strings.Replace(gjson.Get(weather, "forecast.forecastday.#.day.maxtemp_c").String(), "[", "", -1), "]", "", -1)
	minTemp := strings.Replace(strings.Replace(gjson.Get(weather, "forecast.forecastday.#.day.mintemp_c").String(), "[", "", -1), "]", "", -1)
	
	println(location)
	println(minTemp)
	println(maxTemp)
	
	msg := fmt.Sprintf("Salut Tyoma ! Il fera %s degre a %s ce matin et %s degre cette apres-midi", minTemp, location, maxTemp)
	response := sendMsg(user, pass, msg)
	fmt.Println(response)
}

func sendMsg(user string, pass string, msg string) string {

	apiUrl := "https://smsapi.free-mobile.fr/sendmsg"

	prepareData := map[string]string{"user": user, "pass": pass, "msg": msg}
	jsonValues, _ := json.Marshal(prepareData)

	client := &http.Client{}
	request, _ := http.NewRequest("POST", apiUrl, bytes.NewBuffer(jsonValues))
	request.Header.Add("Content-Type", "application/json")
	response, _ := client.Do(request)
	
	return response.Status
}


func getWeather(apiKey string) (data []byte){
	
	apiUrl := fmt.Sprintf("https://api.apixu.com/v1/forecast.json?key=%s&q=Paris", apiKey)

	json, err := http.Get(apiUrl)

	if err != nil {
		fmt.Print(err.Error())
	}

	response, err := ioutil.ReadAll(json.Body)
	if err != nil {
		log.Fatal(err)
	}

	return response
}