package api

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"strings"
	"encoding/json"
)

type Response struct {
	Content string `json:"content"`
}

func Respond(apiUrl, sentence, authorId string) string {
	sentence = strings.Replace(sentence, " ", "%20", -1)
	apiUrl = apiUrl + "/api/response"

	payload := strings.NewReader("sentence=" + sentence + "&authorId=" + authorId)
	req, err := http.NewRequest("POST", apiUrl, payload)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	var response Response
	json.Unmarshal(body, &response)

	return response.Content
}