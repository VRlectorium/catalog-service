package data

import (
	"fmt"

	"github.com/go-pg/pg"
)

type PsqlStore struct {
	db *pg.DB
}

type Subject struct {
	Id       int
	CourseId int
	Name     string
}

type Course struct {
	Id   int
	Name string
}

func NewPsqlStore(dbName string) (*PsqlStore, error) {
	db := pg.Connect(&pg.Options{
		User:     "root",
		Password: "root",
		Addr:     "catalog-db:5432",
		Database: dbName,
	})
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

func (p *PsqlStore) GetCourses() ([]Course, error) {
	var courses []Course
	_, err := p.db.Query(&courses, `SELECT * FROM courses`)
	return courses, err
}

func (p *PsqlStore) GetCourse(id string) (Course, error) {
	var course Course
	_, err := p.db.QueryOne(&course, `SELECT * FROM courses WHERE id=?`, id)
	return course, err
}
