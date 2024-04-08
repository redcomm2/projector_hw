package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const serverPort = 8000

type User struct {
	ID       int
	Login    string
	Password string
	Role     string
}

type Grade struct {
	Subject string
	Score   float64
}

type StudentDetails struct {
	User   User
	Grades []Grade
}

type SchoolClass struct {
	Name     string
	Teacher  User
	Students map[int]StudentDetails
}

var classes = []SchoolClass{
	{
		Name: "Math",
		Teacher: User{
			ID:       1,
			Login:    "teacher",
			Password: "x",
			Role:     "teacher",
		},
		Students: map[int]StudentDetails{
			2: {
				User: User{ID: 2, Login: "student", Password: "x", Role: "student"},
				Grades: []Grade{
					{Subject: "Math", Score: 92},
					{Subject: "Science", Score: 88},
				},
			},
			3: {
				User: User{ID: 3, Login: "student2", Password: "x", Role: "student"},
				Grades: []Grade{
					{Subject: "Math", Score: 85},
					{Subject: "Science", Score: 90},
				},
			},
		},
	},
	{
		Name: "Phys",
		Teacher: User{
			ID:       4,
			Login:    "teacher2",
			Password: "x",
			Role:     "teacher",
		},
		Students: map[int]StudentDetails{
			5: {
				User: User{ID: 5, Login: "student3", Password: "x", Role: "student"},
				Grades: []Grade{
					{Subject: "Physics", Score: 95},
					{Subject: "Math", Score: 89},
				},
			},
			6: {
				User: User{ID: 6, Login: "student4", Password: "x", Role: "student"},
				Grades: []Grade{
					{Subject: "Physics", Score: 93},
					{Subject: "Math", Score: 88},
				},
			},
		},
	},
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/login", loginHandler)
	mux.HandleFunc("/students", authorize("teacher", studentsListHandler))
	mux.HandleFunc("/student/", authorize("teacher", studentHandler))

	s := &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}
	log.Printf("Server starting on port %d\n", serverPort)
	if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Printf("server serve failed: %s", err)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		login := r.FormValue("login")
		password := r.FormValue("password")

		for classIndex, class := range classes {
			if class.Teacher.Login == login && class.Teacher.Password == password {
				setSession(w, class.Teacher.ID, classIndex)
				http.Redirect(w, r, "/students", http.StatusSeeOther)
				return
			}

			for _, studentDetails := range class.Students {
				if studentDetails.User.Login == login && studentDetails.User.Password == password {
					setSession(w, studentDetails.User.ID, classIndex)
					fmt.Fprintf(w, "Login successful. Welcome, %s!", studentDetails.User.Login)
					return
				}
			}
		}

		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	} else {
		fmt.Fprint(w, `<html><body><form action="/login" method="post">
            Username: <input type="text" name="login"><br>
            Password: <input type="password" name="password"><br>
            <input type="submit" value="Login">
        </form></body></html>`)
	}
}

func studentsListHandler(w http.ResponseWriter, r *http.Request) {
	sessionToken, _ := r.Cookie("session_token")
	tokenParts := strings.Split(sessionToken.Value, ":")
	classIndex, _ := strconv.Atoi(tokenParts[2])

	class := classes[classIndex]
	json.NewEncoder(w).Encode(class.Students)
}

func studentHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/student/")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid student ID format", http.StatusBadRequest)
		return
	}

	sessionToken, err := r.Cookie("session_token")
	if err != nil {
		http.Error(w, "Error reading session token", http.StatusBadRequest)
	}

	tokenParts := strings.Split(sessionToken.Value, ":")
	classIndex, _ := strconv.Atoi(tokenParts[2])

	class := classes[classIndex]
	studentDetails, exists := class.Students[id]
	if !exists {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(studentDetails)
}
