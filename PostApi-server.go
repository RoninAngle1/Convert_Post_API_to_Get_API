package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Config struct {
	Port         string `json:"port"`
	PostURL      string `json:"post_url"`
	PostPath     string `json:"post_path"`
	GetURL       string `json:"get_url"`
	GetUsername  string `json:"get_username"`
	GetPassword  string `json:"get_password"`
	GetFrom      string `json:"get_from"`
	PostUsername string `json:"post_username"`
	PostPassword string `json:"post_password"`
}

type PostRequestBody struct {
	PhoneNumber string `json:"PhoneNumber"`
	Message     string `json:"Message"`
}

type GetRequestBody struct {
	To   string `json:"to"`
	Text string `json:"text"`
}

func main() {
	var config Config
	configFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println("Error opening config file")
		return
	}
	defer configFile.Close()

	configBytes, err := io.ReadAll(configFile)
	if err != nil {
		fmt.Println("Error reading config file")
		return
	}

	err = json.Unmarshal(configBytes, &config)
	if err != nil {
		fmt.Println("Error parsing config file")
		return
	}

	http.HandleFunc(config.PostPath, func(w http.ResponseWriter, r *http.Request) {
		username := r.Header.Get("UserName")
		password := r.Header.Get("Password")

		if username != config.PostUsername || password != config.PostPassword {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		var requestBody PostRequestBody
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		getRequestBody := GetRequestBody{
			To:   requestBody.PhoneNumber,
			Text: requestBody.Message,
		}

		client := &http.Client{}
		req, err := http.NewRequest("GET", config.GetURL, nil)
		if err != nil {
			http.Error(w, "Error creating GET request", http.StatusInternalServerError)
			return
		}

		req.SetBasicAuth(config.GetUsername, config.GetPassword)
		q := req.URL.Query()
		q.Add("username", config.GetUsername)
		q.Add("password", config.GetPassword)
		q.Add("from", config.GetFrom)
		q.Add("to", getRequestBody.To)
		q.Add("text", getRequestBody.Text)
		req.URL.RawQuery = q.Encode()

		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, "Error sending GET request", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// Copy the response body from the GET request to the response writer
		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)
	})

	fmt.Printf("Server is listening on port %s\n", config.Port)
	http.ListenAndServe(":"+config.Port, nil)
}
