package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// TODO => model for COURSE entity [file]
type Course struct {
	CourseId     string  `json:"id"`
	CourseName   string  `json:"name"`
	CoursePrice  int     `json:"price"`
	CourseAuthor *Author `json:"author"`
}

// the course is said to be empty name is empty
func (course Course) IsEmpty() bool {
	return course.CourseName == ""
}

// TODO => model for AUTHOR entity [file]
type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// TODO => Fake DB
var Courses = []Course{}

// TODO => controller action method to get all courses
func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses action method is called")
	// STUB => can set header keys-values pairs
	w.Header().Set("Content-Type", "application/json")
	// STUB => send the response in json format
	json.NewEncoder(w).Encode(Courses)
}

// TODO => controller action method to get specific course
func GetCourseById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Course By Id action method is called")
	// STUB => get the id from the request
	params := mux.Vars(r)
	fmt.Printf("The Params from the request are -> %v \n", params)
	// STUB => now loop through the courses from DB and get this course
	for _, course := range Courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	// if we find nothing
	json.NewEncoder(w).Encode("No Course found with given id")
}

// TODO => controller action method to create new course
func CreateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create Course action method is called")
	// STUB => set the header props
	r.Header.Set("Content-Type", "application/json")

	// STUB => handle if the request is completely empty [nothing is sent in the body]
	if r.Body == nil {
		json.NewEncoder(w).Encode("Request body can't be empty for POST requests")
	}

	// STUB => handle if the request body is an empty json "{}"
	// variable for the new received entity
	var course Course
	// we are putting data into variable so we have to send it by ref
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Course entity is expected, instead you have sent empty json object")
		return
	}
	// if its not empty .. so we can add it to the databse and return the response
	// => generate a random id for the new course [db work later]
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(200)
	course.CourseId = strconv.Itoa(id)
	Courses = append(Courses, course)
	// send the response back
	json.NewEncoder(w).Encode(course)
	return
}

// TODO => controller action method to update an existing course
func UpdateCourseById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update Course action method is called")

	// set the response header
	w.Header().Set("Content-Type", "application/json")

	// handle the request body
	if r.Body == nil {
		json.NewEncoder(w).Encode("Request body can't be empty")
	}

	// grap the id from the request
	params := mux.Vars(r)

	for idx, course := range Courses {
		if course.CourseId == params["id"] {
			// Courses[idx] = UpdatedCourse
			Courses = append(Courses[:idx], Courses[idx+1:]...)
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			Courses = append(Courses, course)
			// return the response
			json.NewEncoder(w).Encode(course)
			return
		}
	}
}

// TODO => controller action method to delete an existing course
func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete Course action method is called")
	// grap the id
	params := mux.Vars(r)
	for idx, course := range Courses {
		if course.CourseId == params["id"] {
			Courses = append(Courses[:idx], Courses[idx+1:]...)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
}

func main() {
	// TODO => data seeding
	Courses = append(Courses, Course{
		CourseId:    "1",
		CourseName:  "Nuxtjs master class",
		CoursePrice: 299,
		CourseAuthor: &Author{
			FullName: "fady gamil",
			Website:  "frontend_masters.com",
		},
	})
	Courses = append(Courses, Course{
		CourseId:    "2",
		CourseName:  "MERN Stack Step By Step",
		CoursePrice: 560,
		CourseAuthor: &Author{
			FullName: "magy magdy",
			Website:  "Udemy.com",
		},
	})

	// TODO => define the router
	r := mux.NewRouter()

	// TODO => handle routes
	r.HandleFunc("/courses", GetAllCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", GetCourseById).Methods("GET")
	r.HandleFunc("/courses", CreateCourse).Methods("POST")
	r.HandleFunc("/courses/{id}", UpdateCourseById).Methods("PUT")
	r.HandleFunc("/courses/{id}", DeleteCourse).Methods("DELETE")

	// TODO => listen to requests on specific port
	log.Fatal(http.ListenAndServe(":4000", r))
}
