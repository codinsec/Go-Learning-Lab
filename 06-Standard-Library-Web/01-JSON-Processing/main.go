package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// This program demonstrates JSON processing in Go

// 1. Basic struct with JSON tags
type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Email   string `json:"email,omitempty"` // Omit if empty
	Address string `json:"-"`                // Ignore this field
}

// 2. Nested structs
type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Country string `json:"country"`
}

type Employee struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Address Address `json:"address"`
}

// 3. Custom marshaling
type CustomDate struct {
	time.Time
}

func (d CustomDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Time.Format("2006-01-02"))
}

func (d *CustomDate) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	t, err := time.Parse("2006-01-02", str)
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}

type Event struct {
	Title string     `json:"title"`
	Date  CustomDate `json:"date"`
}

// 4. Anonymous structs
type Response struct {
	Status  string          `json:"status"`
	Data    json.RawMessage `json:"data"` // Raw JSON
	Message string          `json:"message,omitempty"`
}

func main() {
	fmt.Println("=== JSON Processing ===")
	fmt.Println()

	// 1. Marshal (Go struct to JSON)
	fmt.Println("1. Marshal (Go to JSON):")
	person := Person{
		Name:    "Alice",
		Age:     30,
		Email:   "alice@example.com",
		Address: "123 Main St", // This will be ignored
	}
	
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
	} else {
		fmt.Printf("   JSON: %s\n", string(jsonData))
	}
	fmt.Println()

	// 2. Marshal with indentation
	fmt.Println("2. Marshal with Indentation:")
	jsonIndented, err := json.MarshalIndent(person, "", "  ")
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
	} else {
		fmt.Printf("   JSON:\n%s\n", string(jsonIndented))
	}
	fmt.Println()

	// 3. Unmarshal (JSON to Go struct)
	fmt.Println("3. Unmarshal (JSON to Go):")
	jsonStr := `{"name":"Bob","age":25,"email":"bob@example.com"}`
	var person2 Person
	err = json.Unmarshal([]byte(jsonStr), &person2)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
	} else {
		fmt.Printf("   Person: %+v\n", person2)
	}
	fmt.Println()

	// 4. Nested structs
	fmt.Println("4. Nested Structs:")
	employee := Employee{
		ID:   1001,
		Name: "Charlie",
		Address: Address{
			Street:  "456 Oak Ave",
			City:    "New York",
			Country: "USA",
		},
	}
	empJSON, _ := json.MarshalIndent(employee, "", "  ")
	fmt.Printf("   Employee JSON:\n%s\n", string(empJSON))
	fmt.Println()

	// 5. Custom marshaling
	fmt.Println("5. Custom Marshaling:")
	event := Event{
		Title: "Conference",
		Date:  CustomDate{Time: time.Now()},
	}
	eventJSON, _ := json.MarshalIndent(event, "", "  ")
	fmt.Printf("   Event JSON:\n%s\n", string(eventJSON))
	fmt.Println()

	// 6. Unmarshal to map
	fmt.Println("6. Unmarshal to Map:")
	jsonStr2 := `{"name":"David","age":35,"city":"Boston"}`
	var data map[string]interface{}
	json.Unmarshal([]byte(jsonStr2), &data)
	fmt.Printf("   Map: %+v\n", data)
	fmt.Printf("   Name: %v\n", data["name"])
	fmt.Println()

	// 7. Raw JSON
	fmt.Println("7. Raw JSON (json.RawMessage):")
	rawJSON := json.RawMessage(`{"type":"user","id":123}`)
	response := Response{
		Status: "success",
		Data:   rawJSON,
	}
	respJSON, _ := json.MarshalIndent(response, "", "  ")
	fmt.Printf("   Response JSON:\n%s\n", string(respJSON))
	fmt.Println()

	// 8. Omitempty tag
	fmt.Println("8. Omitempty Tag:")
	person3 := Person{
		Name:  "Eve",
		Age:   28,
		Email: "", // Will be omitted
	}
	json3, _ := json.MarshalIndent(person3, "", "  ")
	fmt.Printf("   Person (empty email omitted):\n%s\n", string(json3))
	fmt.Println()

	// 9. Arrays and slices
	fmt.Println("9. Arrays and Slices:")
	people := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 35},
	}
	peopleJSON, _ := json.MarshalIndent(people, "", "  ")
	fmt.Printf("   People array:\n%s\n", string(peopleJSON))
	fmt.Println()

	// 10. Error handling
	fmt.Println("10. Error Handling:")
	invalidJSON := `{"name":"Frank","age":"not a number"}`
	var person4 Person
	err = json.Unmarshal([]byte(invalidJSON), &person4)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
	} else {
		fmt.Printf("   Person: %+v\n", person4)
	}
}

