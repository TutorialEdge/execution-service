package execute

import (
	"io/ioutil"
	"os"
	"os/exec"
	"time"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

// ExecuteCode does the job of taking the Go code that has
// been sent to API from a snippet and executing it before
// returning the response
func Execute(c *fiber.Ctx) error {

	req := new(ExecuteRequest)
	if err := c.BodyParser(req); err != nil {
		log.Fatal(err)
	}

	tmpfile, err := ioutil.TempFile("/tmp", "main.*.go")
	if err != nil {
		log.Error(err)
	}
	log.Info("Created File: " + tmpfile.Name())
	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write([]byte(req.Code)); err != nil {
		log.Error(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Error(err)
	}

	out, err := exec.Command("ls").CombinedOutput()
	if err != nil {
		log.Info(err)
	}

	log.Println(string(out))
	log.Println(string(tmpfile.Name()))

	t1 := time.Now()

	out, err = exec.Command("go", "run", tmpfile.Name()).CombinedOutput()
	if err != nil {
		log.Error(err)
	}

	t2 := time.Now()

	log.Infof("go run output: %s\n", string(out))

	return c.JSON(ExecuteResponse{
		Output:   string(out),
		ExitCode: "",
		Time:     t2.Sub(t1).String(),
	})
}
