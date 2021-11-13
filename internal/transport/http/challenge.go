package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/TutorialEdge/execution-service/internal/challenge"
	log "github.com/sirupsen/logrus"
)

// ChallengeCode does the job of taking the Go code that has
// been sent to API from a snippet and executing it before
// returning the response
func (h Handler) ExecuteChallenge(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	log.Info("Execute Challenge Endpoint Hit")

	req := new(challenge.ChallengeRequest)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendErrorResponse(w, "Failed to decode JSON Body", err)
		return
	}

	dir, err := ioutil.TempDir("/tmp", "challenge*")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir)

	tmpfn := filepath.Join(dir, "main.go")
	if err := ioutil.WriteFile(tmpfn, []byte(req.Code), 0666); err != nil {
		log.Fatal(err)
	}
	var response challenge.ChallengeResponse
	mod := exec.Command("go", "mod", "init", "test")
	mod.Dir = dir
	output, err := mod.CombinedOutput()
	if err != nil {
		log.Error(err)
	}

	log.Info(string(output))

	t1 := time.Now()
	out, err := exec.Command("go", "run", tmpfn).CombinedOutput()
	if err != nil {
		log.Error(err)
	}
	t2 := time.Now()

	response.Output = string(out)
	response.Built = true
	response.Time = t2.Sub(t1).String()

	for _, test := range req.Tests {
		tmpfn := filepath.Join(dir, test.Name+".go")
		if err := ioutil.WriteFile(tmpfn, []byte(test.Code), 0666); err != nil {
			log.Fatal(err)
		}

		cmd := exec.Command("go", "test", "-run", test.Test)
		cmd.Dir = dir
		out, err := cmd.CombinedOutput()
		if err != nil {
			log.Error(err)
			test.Output = err.Error()
			test.Passed = false
		} else {
			test.Output = string(out)
			test.Passed = true
		}

		response.Tests = append(response.Tests, test)
		log.Infof("go test %s\n", tmpfn)
		log.Infof("%+v\n", string(out))
	}

	log.Infof("go run output: %s\n", string(out))

	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}
