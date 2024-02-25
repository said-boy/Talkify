package api

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var apiKeyWit string

func getApiKeyWit() string {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}
	apiKeyWit = viper.GetString("apiKey.wit")
	return apiKeyWit
}

func SendWitVoice(fileRef string) string {

	audio, err := ioutil.ReadFile(fileRef)
	if err != nil {
		log.Fatal("Error reading file:\n%v\n", err)

	}

	reader := bytes.NewReader(audio)

	url := "https://api.wit.ai/speech?v=20141022"
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, reader)

	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+getApiKeyWit())
	req.Header.Set("Content-Type", "audio/wav")
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(body)
}

func SendWitMessage(message string) string {
	url := "https://api.wit.ai/message?v=20160225&q=" + convert(message)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+getApiKeyWit())
	client := &http.Client{}
	resp, _ := client.Do(req)
	contents, _ := ioutil.ReadAll(resp.Body)
	return string(contents)
}

func convert(message string) string {
	arrString := strings.Split(message, " ")
	var ret string
	for x := 0; x < len(arrString); x++ {
		ret += arrString[x] + "%20"
	}
	return ret
}

func GetWitVoiceFromMessage(message string) {
	// url := "https://api.wit.ai/synthesize?v=20230215&q=" + convert(message)
	// req, _ := http.NewRequest("GET", url, nil)
	// req.Header.Set("Authorization", "Bearer "+getApiKeyWit())
	// req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Accept", "audio/pcm16")

	// URL endpoint dan token Anda
	url := "https://api.wit.ai/synthesize?v=20230215"
	token := getApiKeyWit()

	// Data yang akan dikirim dalam body permintaan
	data := `{"q": "` + message + `", "voice": "Rebecca", "speed": 80, "style": "soft", "pitch": 110}`

	// Membuat permintaan HTTP POST
	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Menambahkan header ke permintaan
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "audio/mpeg")

	// Mengirimkan permintaan
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Mengecek kode status respons
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Response status:", resp.Status)
		return
	}

	// Membuka file untuk menyimpan suara
	file, err := os.Create("data/speech.mp3")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Menyalin body respons ke file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("Error copying response body to file:", err)
		return
	}

}
