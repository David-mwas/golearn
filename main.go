// go install -v golang.org/x/tools/gopls@latest
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/exp/rand"
)

type Course struct {
	CourseID    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice float64 `json:"courseprice"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	WebSite  string `json:"website"`
}

// fake db
var courses []Course

// middleware,helper functions
func (c *Course) isEmpty() bool {
	// return c.CourseID == "" && c.CourseName == ""

	return c.CourseName == ""
}

func main() {
	fmt.Println("Starting the server...")

	r := mux.NewRouter()

	// dummy data
	courses = append(courses, Course{CourseID: "1", CourseName: "Go Programming", CoursePrice: 100.0, Author: &Author{FullName: "John Doe", WebSite: "http://johndoe.com"}})

	// routes
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/courses", createCourse).Methods("POST")
	r.HandleFunc("/courses/{courseid}", getCourse).Methods("GET")
	r.HandleFunc("/courses/{courseid}", updateCourse).Methods("PUT")
	r.HandleFunc("/courses/{courseid}", deleteCourse).Methods("DELETE")

	http.ListenAndServe(":8080", r)

}

// controllers

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the home page</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get a course")
	w.Header().Set("Content-Type", "application/json")

	// get the course id from the url
	courseid := mux.Vars(r)["courseid"]

	for _, course := range courses {
		if course.CourseID == courseid {
			json.NewEncoder(w).Encode(course)
			return
		}

	}
	json.NewEncoder(w).Encode("Course with id: " + courseid + " not found")
	// return
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create a course")
	w.Header().Set("Content-Type", "application/json")
	fmt.Println(r.Body)
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please provide course details")
		return
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.isEmpty() {
		json.NewEncoder(w).Encode("Please provide course name")
		return
	}

	// generate a course ,string
	rand.Seed(uint64(time.Now().UnixNano()))
	course.CourseID = strconv.Itoa(rand.Intn(100))

	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)

}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update a course")
	w.Header().Set("Content-Type", "application/json")

	// get the course id from the url
	courseid := mux.Vars(r)["courseid"]

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please provide course details")
		return
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.isEmpty() {
		json.NewEncoder(w).Encode("Please provide course details")
		return
	}

	for index, course := range courses {
		if course.CourseID == courseid {
			courses = append(courses[:index], courses[index+1:]...)

			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseID = courseid
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
		}

	}
	json.NewEncoder(w).Encode("Course with id: " + courseid + " not found")
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete a course")
	w.Header().Set("Content-Type", "application/json")

	// get the course id from the url
	courseid := mux.Vars(r)["courseid"]
	for index, course := range courses {
		if course.CourseID == courseid {
			courses = append(courses[:index], courses[index+1:]...)

			json.NewEncoder(w).Encode("Course with id: " + courseid + " deleted")
			return
		}
	}

	json.NewEncoder(w).Encode("Course with id: " + courseid + " not found")
}
