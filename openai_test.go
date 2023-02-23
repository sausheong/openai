package openai

import (
	"encoding/base64"
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

var apiKey, organization string

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Failed to load the env vars: %v", err)
	}
	apiKey = os.Getenv("OPENAI_API_KEY")
	organization = os.Getenv("OPENAI_ORGANIZATION")
}

var header = "Suggest an inspirational quote based on "
var text = "enchanted forest, with a path between the trees, sun shining, close up, Cinematic Lighting, 3d, render, hyper detailed, 8k"

func TestCompletion(t *testing.T) {

	client := NewClient(apiKey, organization)
	request := make(CompletionRequest)
	request.SetUser("test-user")
	request.SetModel(TEXT_DAVINCI_002)
	request.SetPrompt(fmt.Sprintf("%s:%s", header, text))
	request["temperature"] = 0.75
	request["max_tokens"] = 50

	cr, err := client.Complete(request)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(cr.Text())
	fmt.Println(cr.Id())
	fmt.Println(cr.Model())
	fmt.Println(cr.Object())
	fmt.Println(cr.Created())

}

func TestGenerateImage(t *testing.T) {

	client := NewClient(apiKey, organization)
	request := make(ImageRequest)
	request.SetUser("test-user")
	request.SetPrompt(text)

	cr, err := client.GenerateImage(request)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(cr.URL())
	fmt.Println(cr.Created())
}

func TestGenerateImageB64(t *testing.T) {

	client := NewClient(apiKey, organization)
	request := make(ImageRequest)
	request.SetUser("test-user")
	request.SetPrompt(text)
	request.SetFormat("b64_json")

	cr, err := client.GenerateImage(request)
	if err != nil {
		t.Error(err)
	}
	toFile(cr.ImageBase64())
	fmt.Println(cr.Created())
}

func toFile(b64 string) {
	dec, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		panic(err)
	}

	f, err := os.Create("image.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err := f.Write(dec); err != nil {
		panic(err)
	}
	if err := f.Sync(); err != nil {
		panic(err)
	}
}
