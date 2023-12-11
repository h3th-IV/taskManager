package tasker

// import "fmt"

// type Student struct {
// 	Name   string
// 	Age    uint
// 	Gender string
// }

// type Class struct {
// 	Students []Student
// }

// // Method to add a student to the class
// func (c *Class) AddStudent(student Student) {
// 	c.Students = append(c.Students, student)
// }

// // Method to remove a student from the class based on name
// func (c *Class) RemoveStudentByName(name string) {
// 	for i, student := range c.Students {
// 		if student.Name == name {
// 			// Remove the student from the slice using append and slicing
// 			c.Students = append(c.Students[:i], c.Students[i+1:]...)
// 			return
// 		}
// 	}
// 	fmt.Printf("Student with name %s not found in the class.\n", name)
// }

// // Method to display information about all students in the class
// func (c *Class) DisplayStudents() {
// 	for _, student := range c.Students {
// 		fmt.Printf("Name: %s, Age: %d, Gender: %s\n", student.Name, student.Age, student.Gender)
// 	}
// }

// func main() {
// 	// Create a class
// 	class := Class{}

// 	// Add students to the class
// 	class.AddStudent(Student{Name: "Alice", Age: 20, Gender: "Female"})
// 	class.AddStudent(Student{Name: "Bob", Age: 22, Gender: "Male"})

// 	// Display info about 'em students in the class
// 	fmt.Println("Students in the class:")
// 	class.DisplayStudents()

// 	// Remove a student by name
// 	class.RemoveStudentByName("Alice")

// 	// Display information about the updated list of students
// 	fmt.Println("\nAfter removing a student:")
// 	class.DisplayStudents()
// }
