package auth

import (
	"fmt"
	"testing"
)

func TestAuthenticate(t *testing.T) {
	fmt.Println("Testing Authenticate Function")
}

func TestBadToken(t *testing.T) {
	fmt.Println("Not Implemented")
	// tests := []struct {
	// 	request events.APIGatewayProxyRequest
	// 	expect  bool
	// 	err     error
	// }{
	// 	{
	// 		request: events.APIGatewayProxyRequest{},
	// 		expect:  false,
	// 	},
	// 	{
	// 		request: events.APIGatewayProxyRequest{
	// 			Headers: map[string]string{
	// 				"Authorization": "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IlJFTkZRMFU1UVRCQk9EbEZRME5GTURnMk5ETTFNVVkyT1VaRU56ZEVNVFV4T1RjNU1VWTJRdyJ9.eyJpc3MiOiJodHRwczovL3R1dG9yaWFsZWRnZS5ldS5hdXRoMC5jb20vIiwic3ViIjoiZ29vZ2xlLW9hdXRoMnwxMTc0ODUxNzAwMjg3MjkzMTk5ODgiLCJhdWQiOlsiaHR0cHM6Ly90dXRvcmlhbGVkZ2UuZXUuYXV0aDAuY29tL2FwaS92Mi8iLCJodHRwczovL3R1dG9yaWFsZWRnZS5ldS5hdXRoMC5jb20vdXNlcmluZm8iXSwiaWF0IjoxNTg5MTA4MTcyLCJleHAiOjE1ODkxMTUzNzIsImF6cCI6IkI2M0JNT3BKM1NneE9CMmpiNkxZN0VZRWI0ZFUzSjREIiwic2NvcGUiOiJvcGVuaWQgcHJvZmlsZSJ9.c2lo_iRymNq1AFB6Eq0SmXl45Q1W_UqF7PdJGz2JY-S-69zlR6bVQ0lw6j7jJNSpFJKjUscMy7eV3HfrVqwd71gYoIN4TmBedKvhRZiiSHJ3OYQ3oqdBGLJjljUo9NDCxWpci_YGnv6SddyQODYp3ZdkKqkfUQtlO_cwisQPF6XMqWDOoob2QuZa_5aDvuFE68nrIvmzUiJN037lgHWsSSQ3MKU7JWRizRtndO_oy5v6CTnF4Mjt2EL9wTJQnH06vxogbSTJoixrBSoyQcUEoS1vqTScHE9Ay8kY1uy9qXuDMlNa7SPHgtXlu09_6FVOKHcyN7uq4JDOdviACxV0JA",
	// 			},
	// 		},
	// 		expect: false,
	// 	},
	// }

	// for _, test := range tests {
	// 	response, tokenInfo := Authenticate(test.request)
	// 	t.Logf("%+v\n", tokenInfo)
	// 	assert.Equal(t, test.expect, response)
	// }

}
