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

var text = "enchanted forest, with a path between the trees, sun shining, close up, Cinematic Lighting, 3d, render, hyper detailed, 8k"
var text2 = "Summarize this : Microsoft Azure is a cloud computing platform developed by Microsoft. It provides a range of cloud services, including computing, storage, analytics, virtualization, networking and the web. Azure allows organizations to build, deploy and manage applications and services from any cloud environment. It also provides tools and services to help organizations develop, deploy and manage applications, platforms and systems. Some of the features of Azure include virtual machines, compute and storage services, networking, automation and identity management. Azure also provides tools to help developers create and deploy applications and services, as well as tools to create, manage and deploy applications and services in the cloud. In addition, Azure provides an integrated development environment (IDE) to help developers create and debug applications."

func TestCompletion(t *testing.T) {

	client := NewClient(apiKey, organization)
	request := make(CompletionRequest)
	request.SetModel(TEXT_DAVINCI_003)
	request.SetPrompt(text2 + " {}")
	request["temperature"] = 0.75
	request["max_tokens"] = 4096 - len(text2)
	request["stop"] = "{}"

	cr, err := client.Complete(request)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(cr)
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
