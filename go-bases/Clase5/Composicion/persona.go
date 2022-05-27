package main

import "fmt"

type Bootcamper struct {
	Name     string `json:"name"`
	Age      uint   `json:"age"`
	Position string `json:"position"`
	Bootcamp
}

type Bootcamp struct {
	Id   int    `json:"Id"`
	Tech string `json:"tech"`
}

func main() {
	student1 := Bootcamper{"Juan", 45, "Backend developer", Bootcamp{}}
	student1.Bootcamp = Bootcamp{0003, "Java"}
	student2 := Bootcamper{
		Name:     "Vincent",
		Age:      25,
		Position: "Backend Developer",
		Bootcamp: Bootcamp{
			Id:   0001,
			Tech: "Go",
		},
	}
	student3 := Bootcamper{}
	student3.Name = "Nicolas"
	student3.Age = 30
	student3.Position = "Frontend"
	student3.Bootcamp.Id = 0002
	student3.Bootcamp.Tech = "REact"

	fmt.Println("Student1:", student1)
	fmt.Println("Student2:", student2)
	fmt.Println("Student3:", student3)
}
