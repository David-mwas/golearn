// go install -v golang.org/x/tools/gopls@latest
package json

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"name"`
	Price    string   `json:"price"`
	Platform string   `json:"platform"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	// EncodeJson()
	DecodeJson()
}
func EncodeJson() {
	lcoCourses := []course{
		{"React", "299", "LearnCodeOnline", "abc123", []string{"web-dev", "js"}},
		{"MERN", "199", "LearnCodeOnline", "abc123", []string{"web-dev", "js"}},
		{"Angular", "299", "LearnCodeOnline", "abc123", nil},
	}

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(finalJson))
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
                "name": "React",
                "price": "299",
                "platform": "LearnCodeOnline",
                "tags": [
                        "web-dev",
                        "js"
                ]
        }`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)

		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON was not valid")
	}

	var lcoCourse2 map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &lcoCourse2)
	fmt.Printf("%#v\n", lcoCourse2)

	for k, v := range lcoCourse2 {
		fmt.Printf("Key is: %v and value is: %v and type is %T\n", k, v, v)
	}

}
