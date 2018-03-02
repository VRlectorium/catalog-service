package data

import (
	"fmt"
	"os"

	"github.com/go-pg/pg"
)

type PsqlStore struct {
	db *pg.DB
}

type Subject struct {
	Id       int
	Courseid int
	Name     string
}

type Course struct {
	Id       int
	Name     string
	Subjects []Subject
}

func NewPsqlStore(dbName string) (*PsqlStore, error) {
	var db *pg.DB
	if os.Getenv("TEST") == "TEST" {
		db = pg.Connect(&pg.Options{
			User:     "root",
			Password: "root",
			Addr:     "catalog-db:5432",
			Database: dbName,
		})
	} else {
		db = pg.Connect(&pg.Options{
			User:     "root",
			Password: "root",
			Database: dbName,
		})
	}

	return &PsqlStore{db: db}, nil
}

func (p *PsqlStore) Test() {
	var n int
	_, err := p.db.Query(pg.Scan(&n), "SELECT ?", 42)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	fmt.Println(n)
}

func (p *PsqlStore) CreateCourse(course *Course) error {
	_, err := p.db.QueryOne(course, `
		INSERT INTO courses (name) VALUES (?name)
		RETURNING name`, course)
	return err
}

type ORMres struct {
	Id        int
	Name      string
	Subject   string
	Subjectid int
}

func (p *PsqlStore) GetCourses() ([]Course, error) {
	var courses []Course
	var res []ORMres
	_, err := p.db.Query(&res,
		`SELECT courses.id, courses.name, subjects.name as subject, subjects.id as subjectid
		FROM courses INNER JOIN subjects ON courses.id=subjects.courseid`)
	for _, r := range res {
		if len(courses) > r.Id-1 && courses[r.Id-1].Id == r.Id {
			courses[r.Id-1].Subjects = append(courses[r.Id-1].Subjects, Subject{
				Id:       r.Subjectid,
				Courseid: r.Id,
				Name:     r.Subject})
		} else {
			course := Course{
				Id:   r.Id,
				Name: r.Name,
				Subjects: []Subject{Subject{
					Id:       r.Id,
					Courseid: r.Subjectid,
					Name:     r.Subject}}}
			courses = append(courses, course)
		}
	}
	return courses, err
}

func (p *PsqlStore) GetCourse(id string) (Course, error) {
	var course Course
	_, err := p.db.QueryOne(&course, `SELECT * FROM courses WHERE id=?`, id)
	return course, err
}

func (p *PsqlStore) CreateSubject(subject *Subject) error {
	_, err := p.db.QueryOne(subject, `
		INSERT INTO subjects (courseid, name) VALUES (?courseid,?name)
		RETURNING name`, subject)
	return err
}

func (p *PsqlStore) GetSubjects() ([]Subject, error) {
	var subjects []Subject
	_, err := p.db.Query(&subjects, `SELECT * FROM subjects`)
	return subjects, err
}

func (p *PsqlStore) GetSubject(id string) (Subject, error) {
	var subject Subject
	_, err := p.db.QueryOne(&subject, `SELECT * FROM subjects WHERE id=?`, id)
	return subject, err
}
