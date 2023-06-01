package integration_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMathSolverAPI(t *testing.T) {
	assert := assert.New(t)

	// Test the evaluate endpoint
	expression := map[string]string{"expression": "What is 5 plus 5?"}
	expressionBytes, _ := json.Marshal(expression)
	response, err := http.Post("http://localhost:8080/evaluate", "application/json", bytes.NewBuffer(expressionBytes))
	assert.NoError(err)
	assert.Equal(http.StatusOK, response.StatusCode)
	decoder := json.NewDecoder(response.Body)
	var result map[string]int
	err = decoder.Decode(&result)
	assert.NoError(err)
	assert.Equal(10, result["result"])

	// Test the validate endpoint
	response, err = http.Post("http://localhost:8080/validate", "application/json", bytes.NewBuffer(expressionBytes))
	assert.NoError(err)
	assert.Equal(http.StatusOK, response.StatusCode)
	decoder = json.NewDecoder(response.Body)
	var validation map[string]bool
	err = decoder.Decode(&validation)
	assert.NoError(err)
	assert.True(validation["valid"])

	// Test the errors endpoint
	response, err = http.Get("http://localhost:8080/errors")
	assert.NoError(err)
	assert.Equal(http.StatusOK, response.StatusCode)
	decoder = json.NewDecoder(response.Body)
	var errors []map[string]interface{}
	err = decoder.Decode(&errors)
	assert.NoError(err)
	// No errors should be present yet
	assert.Len(errors, 0)

	// Send invalid expression and check errors
	expression = map[string]string{"expression": "What is 5 pluss 5?"}
	expressionBytes, _ = json.Marshal(expression)
	response, err = http.Post("http://localhost:8080/evaluate", "application/json", bytes.NewBuffer(expressionBytes))
	assert.NoError(err)
	assert.Equal(http.StatusBadRequest, response.StatusCode)

	// Check errors endpoint again
	response, err = http.Get("http://localhost:8080/errors")
	assert.NoError(err)
	assert.Equal(http.StatusOK, response.StatusCode)
	decoder = json.NewDecoder(response.Body)
	err = decoder.Decode(&errors)
	assert.NoError(err)
	// Now one error should be present
	assert.Len(errors, 1)
}
