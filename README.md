# Go-gpt-functions: AI-Powered Magic Functions🪄
🧙‍♂️ Make a spell to get all you want！🔧 Using Openai's gpt models to implement magic function for golang.



## Usage:

```go
func main() {
	ctx := context.TODO()
	magicObj := magic.NewMagicObject("YOUR_API_KEY")
    	// magicObj.SetDefaultEngineModel("gpt-4-0314") //if you have access to gpt-4 model
	// magicObj.SetDefaultTemperature(0.5)
	
	response, err := magicObj.DoMagic(ctx, "Give me a random number between 0 and 100", nil)
	if err != nil {
		log.Print("doMagic err: ", err)
	}
	log.Print("magicResponse random: ", response)
    	// Output: magicResponse random: 62

	response, err = magicObj.DoMagic(ctx, "Give me a simple user tip for server errors", nil)
	if err != nil {
		log.Print("doMagic err: ", err)
	}
	log.Print("magicResponse server err info: ", response)
    	// Output: magicResponse server err info: Try refreshing the page or clearing your browser's cache and cookies. If the error persists, contact the website's support team for assistance

	response, err = magicObj.DoMagic(ctx, "Output information about the word 'default', in json format", nil)
	if err != nil {
		log.Print("doMagic err: ", err)
	}
	log.Print("magicResponse json word info: ", response)
    	// Output: magicResponse json word info: {"word": "default", "type": "noun", "definition": "the preselected option or parameter in a software application or other system", "synonyms": ["setting", "preference", "choice"], "antonyms": ["nonpayment", "delinquency"]}

	response, err = magicObj.DoMagic(ctx, "Generate me a random user information with five random properties and values, in json format", nil)
	if err != nil {
		log.Print("doMagic err: ", err)
	}
	log.Print("magicResponse user info: ", response)
    	// Output: magicResponse user info: 
	// {"name": "Emily", "age": 27, "email": "emily@example.com", "city": "New York", "hobby": "photography"}

    response, err = magicObj.DoMagic(ctx, "Generate me a random user information with six random properties and values, in json format.", map[string]interface{}{
		"name": "ycace",
		"age":  "18",
	})
	if err != nil {
		log.Print("doMagic err: ", err)
	}
	log.Print("magicResponse user info with input: ", response)
    	// Output: magicResponse user info with input: 
	// {"name": "ycace", "age": 18, "gender": "male", "email": "ycace@example.com", "address": "123 Main St", "phone": "555-555-5555"}

}
```
