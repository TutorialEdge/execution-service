package challenge

import (
	"github.com/jinzhu/gorm"
)

// Challenge - holds the users challenges
type Challenge struct {
	gorm.Model
	Slug          string `json:"slug"`
	AuthorID      string `json:"sub"`
	Code          string `json:"code"`
	Score         int    `json:"score"`
	Passed        bool   `json:"passed"`
	ExecutionTime string `json:"execution_time"`
}

// ChallengeResponse contains the response from
// executing the Go code
type ChallengeResponse struct {
	ExitCode string          `json:"exit_code"`
	Output   string          `json:"output"`
	Built    bool            `json:"built"`
	Time     string          `json:"time"`
	Tests    []ChallengeTest `json:"tests"`
}

// ChallengeRequest takes in the source code
// from the editor as well as a number of tests
// which are written into a file and ran
type ChallengeRequest struct {
	Code  string          `json:"code"`
	Tests []ChallengeTest `json:"tests"`
}

// ChallengeTest is a struct which contains
// the source code for a test file as well as
// the metadata such as the test name
type ChallengeTest struct {
	Name   string `json:"name"`
	Code   string `json:"code"`
	Test   string `json:"test"`
	Output string `json:"output"`
	Passed bool   `json:"passed"`
}
