//go:build e2e
// +build e2e

package tests

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	BaseURL       = "http://localhost:8000"
	SimpleProgram = `package main

import "fmt"
	
func main() {
	fmt.Println("hello world")
}
`
)

func TestExecutionService(t *testing.T) {
	t.Run("the execution service can handle incoming execution requests", func(t *testing.T) {

		// client := resty.New()
		// resp, err := client.R().
		// 	SetBody(fmt.Sprintf(`{"code": %s, "tests": [
		// 		{
		// 			"name": "A test",
		// 			"code": %s,
		// 			"test": ""
		// 		}
		// 	]}`, SimpleProgram, SimpleProgram)).
		// 	Post(BaseURL + "/v1/execute")
		// assert.NoError(t, err)

		// assert.Equal(t, 200, resp.StatusCode())

	})
}
