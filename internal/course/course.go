package course

import (
	"github.com/TutorialEdge/execution-service/internal/database"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

// Service -
type Service struct {
	Store   database.Store
	Courses []Course
}

// LessonStatus -
type LessonStatus struct {
	gorm.Model
	Sub      string `json:"sub"`
	AuthorID string `json:"author_id"`
	Slug     string `json:"slug"`
	Course   string `json:"course"`
}

// Course - the struct that contains the response for the courses
type CourseProgress struct {
	Title    string
	Slug     string
	Length   int
	Complete int
	URL      string
}

// Course - the struct that contains the course definition
type Course struct {
	Title    string `json:"title"`
	Slug     string
	Sections []Section `json:"sections"`
}

// Section - represents each section of a course
type Section struct {
	Name     string    `json:"name"`
	Weight   int       `json:"weight"`
	Articles []Article `json:"articles"`
}

// Article - represents an individual article in a course
type Article struct {
	Title   string `json:"title"`
	Premium bool   `json:"premium"`
	Time    string `json:"time"`
	State   string `json:"status"`
	URL     string `json:"url"`
}

// New - returns a new comment services
func New(db database.Store) Service {
	courses, err := PopulateCourses()
	if err != nil {
		log.Error(err)
	}

	return Service{
		Store:   db,
		Courses: courses,
	}
}

// GetCoursesStarted -
func (s Service) GetCoursesStarted(authorID string) []CourseProgress {
	var courses []CourseProgress
	rows, err := s.Store.DB.Raw("select distinct(course) from lesson_statuses where sub = ?", authorID).Rows()
	if err != nil {
		log.Error(err.Error())
		return courses
	}
	defer rows.Close()

	for rows.Next() {
		var course string
		err := rows.Scan(&course)
		if err != nil {
			log.Error(err.Error())
		}

		type CountResult struct {
			Count int
		}
		var count CountResult
		s.Store.DB.Raw("select count(*) from lesson_statuses where sub = ? and course = ?", authorID, course).Scan(&count)

		var title string
		var lessons int
		for _, c := range s.Courses {
			if c.Slug == course {
				for _, section := range c.Sections {
					lessons += len(section.Articles)
				}
				title = c.Title
			}
		}

		courses = append(courses, CourseProgress{
			Title:    title,
			Slug:     course,
			Length:   lessons,
			Complete: count.Count,
			URL:      "/courses/" + course + "/",
		})
	}

	return courses
}

// GetCourseStatus -
func (s Service) GetCourseStatus(course, authorID string) []LessonStatus {
	log.Printf("Fetching course status with course id %s and sub %s\n", course, authorID)
	var lessons []LessonStatus
	s.Store.DB.Where("course = ? AND sub = ?", course, authorID).Find(&lessons)
	return lessons
}

// PostCourseUpdate -
// Adds a new comment to the site
func (s Service) PostCourseUpdate(lesson LessonStatus) error {
	log.Info("Posting Course Status Update")
	if s.Store.DB.Where("sub = ? AND slug = ?", lesson.Sub, lesson.Slug).Find(&lesson).RecordNotFound() {
		if err := s.Store.DB.Create(&lesson).Error; err != nil {
			return err
		}
	}
	return nil
}
