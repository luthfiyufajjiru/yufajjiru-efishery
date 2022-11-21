package auth

import (
	"bytes"
	"encoding/json"
	"fetchapp/modules/auth/models"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func TokenClaim(inp string) (models.UserDTO, error) {
	host := os.Getenv("authapp_url")
	url := fmt.Sprintf("%s/api/v1/account/claim", host)

	var result models.UserDTO

	client := &http.Client{
		Timeout: time.Second * 30,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		x := err.Error()
		log.Println(x)
		return result, fmt.Errorf("got error %s", err.Error())
	}

	req.Header.Set("Token", inp)

	response, err := client.Do(req)

	if err != nil {
		x := err.Error()
		log.Println(x)
		return result, fmt.Errorf("got error %s", err.Error())
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return result, nil
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return result, err
	}

	err = json.Unmarshal(body, &result)

	if err != nil {
		return result, err
	}

	return result, nil
}

func Register(phone, name, role string) (string, error) {
	host := os.Getenv("authapp_url")
	url := fmt.Sprintf("%s/api/v1/account/sign-up", host)

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	data := map[string]interface{}{
		"phone_number": phone,
		"name":         name,
		"role":         role,
	}

	jsonValue, _ := json.Marshal(data)
	u := bytes.NewReader(jsonValue)

	req, err := http.NewRequest("POST", url, u)
	if err != nil {
		return "", fmt.Errorf("got error %s", err.Error())
	}

	req.Header.Set("Content-Type", "application/json")

	response, err := client.Do(req)

	if err != nil {
		return "", fmt.Errorf("got error %s", err.Error())
	}

	body, err := io.ReadAll(response.Body)
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", nil
	}

	if err != nil {
		return "", fmt.Errorf("got error %s", err.Error())
	}

	return string(body), nil
}

func SignIn(phone, pass string) (string, error) {
	host := os.Getenv("authapp_url")
	url := fmt.Sprintf("%s/api/v1/account/sign-in", host)

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	data := map[string]interface{}{
		"phone_number": phone,
		"password":     pass,
	}

	jsonValue, _ := json.Marshal(data)
	u := bytes.NewReader(jsonValue)

	req, err := http.NewRequest("POST", url, u)
	if err != nil {
		return "", fmt.Errorf("got error %s", err.Error())
	}

	req.Header.Set("Content-Type", "application/json")

	response, err := client.Do(req)

	if err != nil {
		return "", fmt.Errorf("got error %s", err.Error())
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", nil
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return "", fmt.Errorf("got error %s", err.Error())
	}

	return string(body), nil
}
