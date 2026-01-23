package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Course struct {
	Courseid    string `json:"courseid"`
	Coursename  string `json:"coursename"`
	Courseprice string `json:"courseprice"`
	Author      string `json:"author"`
}

func main() {
	connectDB()
	fmt.Println("Welcome to API")
	r := mux.NewRouter()

	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/courses", createOneCourse).Methods("POST")
	r.HandleFunc("/courses/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/courses/{id}", deleteOneCourse).Methods("DELETE")

	//listen to port
	log.Fatal(http.ListenAndServe(":1000", enableCORS(r)))

}
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (c Course) isempty() bool {
	// return c.Courseid == "" && c.Coursename == ""
	return c.Courseid == ""

}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by DEV</h1"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses")
	w.Header().Set("Content-type", "application/json")

	rows, err := DB.Query("SELECT * FROM courses ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	defer rows.Close()
	var courses []Course

	for rows.Next() {
		var c Course
		var id int

		rows.Scan(
			&id,
			&c.Coursename,
			&c.Courseprice,
			&c.Author,
		)
		c.Courseid = strconv.Itoa(id)
		courses = append(courses, c)
	}
	json.NewEncoder(w).Encode(courses)
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one Course")
	w.Header().Set("Content-type", "application/json")

	var c Course

	// Decode JSON body
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validation
	if c.Coursename == "" || c.Courseprice == "" || c.Author == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}

	query := `INSERT INTO COURSES
	(coursename,courseprice,author)
	VALUES(?,?,?)`

	result, err := DB.Exec(query,
		c.Coursename,
		c.Courseprice,
		c.Author,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, _ := result.LastInsertId()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"id": strconv.FormatInt(id, 10), "message": "Course Created"})

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update the course")
	w.Header().Set("Content-type", "application/json")

	//grab request
	params := mux.Vars(r)

	var c Course
	json.NewDecoder(r.Body).Decode(&c)

	_, err := DB.Exec(`
	UPDATE courses SET 
	coursename=?,courseprice=?,author=?
	WHERE id=?`,
		c.Coursename,
		c.Courseprice,
		c.Author,
		params["id"],
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode("Course Updated")

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete One Course")
	w.Header().Set("Content-type", "application/json")

	//grab request
	params := mux.Vars(r)

	_, err := DB.Exec("DELETE FROM courses WHERE id =?", params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode("Course Deleted")
}
