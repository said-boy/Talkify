package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/google/generative-ai-go/genai"
	"github.com/spf13/viper"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"github.com/said-boy/talkify/entity"
)

// mengambil api key gemini dari file env/config.yaml
func getApiKeyGemini() string {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}
	return viper.GetString("apiKey.gemini")
}

// Func untuk menyimpan riwayat percakapan ke dalam file
// func SaveConversationHistory(history []*genai.Content, filename string) error {
func SaveConversationHistory(oldHistory []*genai.Content, history []*genai.Content, filename string) error {

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var conversations entitiy.Conversations
	var conf []*entitiy.Conversation
	
	counter := len(oldHistory)+2
	for i, v := range history {
		if i == counter {
			break
		}

		for _, j := range v.Parts {
			text := fmt.Sprintf("%v", j)
			conversation := &entitiy.Conversation{
				Parts: []entitiy.Part{
					{
						Text: genai.Text(text),
					},
				},
				Role: v.Role,
			}

			conf = append(conf, conversation)

		}
	}
	conversations.History = conf

	// Marshal Conversations ke dalam bentuk JSON
	jsonData, err := json.Marshal(conversations)
	if err != nil {
		return err
	}

	// Simpan data JSON ke dalam file
	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

// Func untuk memuat kembali riwayat percakapan dari file
func LoadConversationHistory(filename string) ([]*genai.Content, error) {
	// Baca konten file JSON
	jsonData, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Deklarasikan variabel untuk menyimpan data hasil unmarshal
	var conversation = entitiy.Conversations{}

	// Unmarshal JSON ke dalam variabel conversation
	if err := json.Unmarshal(jsonData, &conversation); err != nil {
		return nil, err
	}

	var result []*genai.Content
	for _, v := range conversation.History {
		for _, j := range v.Parts {
			content := &genai.Content{
				Parts: []genai.Part{
					genai.Text(j.Text),
				},
				Role: v.Role,
			}
			result = append(result, content)
		}
	}

	// Kembalikan riwayat percakapan
	return result, nil
}

func SendGeminiMessage(message string) []*genai.Content {
	ctx := context.Background()
	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(ctx, option.WithAPIKey(getApiKeyGemini()))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// For text-only input, use the gemini-pro model
	model := client.GenerativeModel("gemini-pro")
	// Initialize the chat
	cs := model.StartChat()

	// mengambil percakapan lama.
	history, err := LoadConversationHistory("log/conversation_history.json")


	if err != nil {
		log.Fatal(err)
	}

	cs.History = history

	iter := cs.SendMessageStream(ctx, genai.Text(message))
	if err != nil {
		log.Fatal(err)
	}

	for {
		resp, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		for _, v := range resp.Candidates {
			cs.History = append(cs.History, v.Content)
		}

	}

	// menyimpan percakapan baru.
	if err := SaveConversationHistory(history, cs.History, "log/conversation_history.json"); err != nil {
		log.Fatal(err)
	}

	return cs.History

}

func GeminiToSayAnswer() {
	// Ganti "audio.mp3" dengan nama file MP3 yang ingin Anda putar
	file := "data/speech.mp3"

	// Membuat perintah untuk menjalankan "mpg123" dengan nama file sebagai argumen
	cmd := exec.Command("mpg123", file)

	// Menjalankan perintah dan menangani kesalahan jika terjadi
	err := cmd.Run()
	if err != nil {
		fmt.Println("Gagal memutar file MP3:", err)
		return
	}
}
