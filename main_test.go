package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloWithName(t *testing.T) {
	gin.SetMode("test")

	server := httptest.NewServer(GetMainEngine("Привет"))
	defer server.Close()

	response, err := http.Get(server.URL + "/hello?name=Петр")
	defer response.Body.Close()

	assert.Nil(t, err)
	assert.Equal(t, response.StatusCode, http.StatusOK)

	decoder := json.NewDecoder(response.Body)
	var response_data map[string]interface{}
	err = decoder.Decode(&response_data)

	assert.Nil(t, err)
	assert.Equal(t, "Привет, Петр", response_data["message"])

}
