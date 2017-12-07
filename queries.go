package main

import (
	"fmt"
	"reflect"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Queries ...
func Queries() {
	// sanity check
	fmt.Println("queries.go has run")

	// declarations
	var student Students
	fmt.Println("student =", reflect.TypeOf(student))

	// connection
	connection, err := gorm.Open("postgres", "host=localhost dbname=go_getters_g_portal sslmode=disable")
	fmt.Println("connection =", reflect.TypeOf(connection))
	// on connection err
	if err != nil {
		panic("failed to connect database")
	}
	defer connection.Close()

	// Drops the table to rest the database
	connection.DropTable(&Students{})

	// Migrate taht ol schema
	connection.AutoMigrate(&Students{})

	// I'm pretty sure these are seeds ;) // Create
	connection.Create(&Students{ // Clark Walker
		// ID:              1,                                  // uint
		Name:            "Clark Walker",                     // string
		Email:           "303walker@gmail.com",              // string
		Cohort:          64,                                 // int
		GithubURL:       "https://github.com/ClarkWalker",   // string
		LinkedinURL:     "www.linkedin.com/in/clark-walker", // string
		SlackHandle:     "cMonster",                         // string
		Password:        "password123",                      // string
		Role:            "snarky student",                   // string
		SlackPassword:   "123password",                      // string
		AbsentExcused:   2,                                  // int
		AbsentUnexcused: 1})                                 // int

	connection.Create(&Students{ // Dakota Pfeifer
		// ID:              2,                                  // uint
		Name:            "Dakota Pfeifer",             // string
		Email:           "dpfeif@outlook.com",         // string
		Cohort:          64,                           // int
		GithubURL:       "https://github.com/dpfeif",  // string
		LinkedinURL:     "www.linkedin.com/in/dpfeif", // string
		SlackHandle:     "dpfeif",                     // string
		Password:        "password123",                // string
		Role:            "fancy lad",                  // string
		SlackPassword:   "123password",                // string
		AbsentExcused:   2,                            // int
		AbsentUnexcused: 1})                           // int

	// Read
	connection.Find(&student)
	var getAll = connection.Find(&student)
	fmt.Println(getAll)

	// // Update
	// connection.Model(&student).Update("Role", "vary snarky student")

	// // Delete
	// connection.Delete(&student)
} // end main

// GetAllStudents ...
// func GetAllStudents(connection *gorm.DB) {
// 	var student Students
// 	connection.Find(&student)
// }