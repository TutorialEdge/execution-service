package course

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	log "github.com/sirupsen/logrus"
)

var (
	ClonePath = "/tmp/tutorialedge"
)

func cloneRepo() error {
	githubToken := os.Getenv("GH_ACCESS_TOKEN")
	_, err := git.PlainClone(ClonePath, false, &git.CloneOptions{
		URL:      "https://github.com/TutorialEdge/tutorialedge.net",
		Progress: os.Stdout,
		Auth: &http.BasicAuth{
			Username: "elliotforbes",
			Password: githubToken,
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func PopulateCourses() ([]Course, error) {
	err := os.RemoveAll(ClonePath)
	if err != nil {
		return nil, err
	}

	err = cloneRepo()
	if err != nil {
		return nil, err
	}

	files, err := ioutil.ReadDir(ClonePath + "/data/courses")
	if err != nil {
		log.Fatal(err)
	}

	var courses []Course
	for _, f := range files {
		var course Course
		dat, err := ioutil.ReadFile(ClonePath + "/data/courses/" + f.Name())
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(dat, &course)
		if err != nil {
			return nil, err
		}

		slug := strings.Split(f.Name(), ".")
		course.Slug = slug[0]

		courses = append(courses, course)
	}
	return courses, nil
}
