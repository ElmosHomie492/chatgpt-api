# chatgpt-api

## About
You already know what [ChatGPT](https://platform.openai.com/docs/introduction) is. This API is a simple implementation of the OpenAI API for GPT. In short...you can ask it questions!

## Running
Before you begin, create a new API key [here](). Then, just set an environment variable called `CHATGPT_API_KEY` to the API key that was just generated. Then run `go run main.go`.

To test the API, run the following cURL command (replace APIKey with your base64 encoded API key):
```bash
curl --request POST \
  --url http://localhost:8080/askGPT \
  --header 'APIKey: YOUR_BASE64_ENCODED_API_KEY_HERE' \
  --header 'Content-Type: multipart/form-data' \
  --form 'question=what is 2+2?'
```

## Base64 Encoding Help
If you don't know how to base64 encode a string, don't fear! It's really easy. Just run this: 
```bash
echo $CHATGPT_API_KEY | base64
```