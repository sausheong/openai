# OpenAI

OpenAI is a package that wraps around the OpenAI HTTP APIs. To use this package properly you will need to have an OpenAI account (currently in beta), which will give you the API key that you need. You will also be assigned to an organization, which you can optionally add to the API request.

Important to note that each application developed using the OpenAI APIs need to be approved by OpenAI. 

## Completion

OpenAI models can be used for multiple text related tasks. Given a prompt and some parameters, the model can return one or more predicted completions and other information.

To start using the API, create a client and set the API key and organization.

````go
client := NewClient(apiKey, organization)
````

Next, we need to create a `CompletionRequest`, which is a `map[string]any` with some attached methods. The 2 mandatory fields in the request are model and prompt, which sets the model to use, as well as the prompt to send for completion. You should also send the ID for the end-user even though it is optional, because OpenAI will review your application and this is a stated criteria in the review.

````go
request := make(CompletionRequest)
request.SetUser("test-user")
request.SetModel(TEXT_DAVINCI_002)
request.SetPrompt("Suggest three names for a horse that is a superhero.")
````
You can also set other parameters, which you can get from the API documentation, by treating the request as a map.

````go
request["temperature"] = 0.75
request["max_tokens"] = 50
````

To send the completion to OpenAI, use the `Complete` method in the client, passing it the request. 

````go
cr, err := client.Complete(request)
if err != nil {
    // resolve the error
}
fmt.Println(cr.Text())
````

Once OpenAI responds, it will provide a `CompletionResponse` which is also a `map[string]any` that has a few methods. You will most likely just use the `Text` method to extract the text response, but you can also get the other response data by using the methods or treating it like a map. 