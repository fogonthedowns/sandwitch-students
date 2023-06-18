package main

func countStudents(students []int, sandwiches []int) int {
	counter := 0
	for len(students) > 0 {
		if students[0] == sandwiches[0] {
			counter = 0
			// match, student and sandwitch leave line
			students = students[1:]
			sandwiches = sandwiches[1:]
			continue
		}
		counter++
		if counter > len(students) {
			// done
			break
		}
		// pop student
		student := students[0]
		// send to back of the line
		students = append(students[1:], student)

	}
	// students without a match :(
	return len(students)
}
