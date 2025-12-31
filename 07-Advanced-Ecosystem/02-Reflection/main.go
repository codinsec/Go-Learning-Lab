package main

import (
	"fmt"
	"reflect"
)

// This program demonstrates reflection in Go

// Example struct for reflection
type Person struct {
	Name    string `json:"name" validate:"required"`
	Age     int    `json:"age" validate:"min=0,max=150"`
	Email   string `json:"email" validate:"email"`
	private string // unexported field
}

func (p Person) GetName() string {
	return p.Name
}

func (p *Person) SetName(name string) {
	p.Name = name
}

func (p Person) String() string {
	return fmt.Sprintf("Person{Name: %s, Age: %d, Email: %s}", p.Name, p.Age, p.Email)
}

func main() {
	fmt.Println("=== Reflection ===")
	fmt.Println()

	// 1. Type reflection
	fmt.Println("1. Type Reflection:")
	var x int = 42
	t := reflect.TypeOf(x)
	fmt.Printf("   Type of x: %s\n", t)
	fmt.Printf("   Kind: %s\n", t.Kind())
	fmt.Printf("   Size: %d bytes\n", t.Size())
	fmt.Println()

	// 2. Value reflection
	fmt.Println("2. Value Reflection:")
	v := reflect.ValueOf(x)
	fmt.Printf("   Value: %d\n", v.Int())
	fmt.Printf("   Can set: %t\n", v.CanSet())
	fmt.Println()

	// 3. Setting values (pointer required)
	fmt.Println("3. Setting Values:")
	var y int = 10
	vp := reflect.ValueOf(&y)
	fmt.Printf("   Original y: %d\n", y)
	
	// Get the element that vp points to
	elem := vp.Elem()
	fmt.Printf("   Can set: %t\n", elem.CanSet())
	
	// Set the value
	elem.SetInt(20)
	fmt.Printf("   New y: %d\n", y)
	fmt.Println()

	// 4. Struct reflection
	fmt.Println("4. Struct Reflection:")
	p := Person{
		Name:  "Alice",
		Age:   30,
		Email: "alice@example.com",
	}
	
	structType := reflect.TypeOf(p)
	fmt.Printf("   Struct type: %s\n", structType)
	fmt.Printf("   Number of fields: %d\n", structType.NumField())
	fmt.Println()

	// 5. Iterate over struct fields
	fmt.Println("5. Iterate Over Struct Fields:")
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		fmt.Printf("   Field %d: %s (type: %s)\n", i, field.Name, field.Type)
		fmt.Printf("      Tag: %s\n", field.Tag)
		
		// Get specific tag values
		if jsonTag := field.Tag.Get("json"); jsonTag != "" {
			fmt.Printf("      JSON tag: %s\n", jsonTag)
		}
		if validateTag := field.Tag.Get("validate"); validateTag != "" {
			fmt.Printf("      Validate tag: %s\n", validateTag)
		}
	}
	fmt.Println()

	// 6. Get field values
	fmt.Println("6. Get Field Values:")
	structValue := reflect.ValueOf(p)
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		fieldValue := structValue.Field(i)
		
		// Only access exported fields
		if fieldValue.CanInterface() {
			fmt.Printf("   %s = %v\n", field.Name, fieldValue.Interface())
		} else {
			fmt.Printf("   %s = <unexported>\n", field.Name)
		}
	}
	fmt.Println()

	// 7. Set struct field values
	fmt.Println("7. Set Struct Field Values:")
	p2 := Person{Name: "Bob", Age: 25}
	p2Value := reflect.ValueOf(&p2).Elem()
	
	nameField := p2Value.FieldByName("Name")
	if nameField.IsValid() && nameField.CanSet() {
		nameField.SetString("Charlie")
		fmt.Printf("   Updated name: %s\n", p2.Name)
	}
	
	ageField := p2Value.FieldByName("Age")
	if ageField.IsValid() && ageField.CanSet() {
		ageField.SetInt(35)
		fmt.Printf("   Updated age: %d\n", p2.Age)
	}
	fmt.Println()

	// 8. Method reflection
	fmt.Println("8. Method Reflection:")
	personType := reflect.TypeOf(p)
	fmt.Printf("   Number of methods: %d\n", personType.NumMethod())
	
	for i := 0; i < personType.NumMethod(); i++ {
		method := personType.Method(i)
		fmt.Printf("   Method %d: %s\n", i, method.Name)
		fmt.Printf("      Type: %s\n", method.Type)
	}
	fmt.Println()

	// 9. Call methods via reflection
	fmt.Println("9. Call Methods via Reflection:")
	personValue := reflect.ValueOf(p)
	
	// Call GetName method
	getNameMethod := personValue.MethodByName("GetName")
	if getNameMethod.IsValid() {
		results := getNameMethod.Call(nil)
		if len(results) > 0 {
			fmt.Printf("   GetName() result: %s\n", results[0].String())
		}
	}
	
	// Call String method
	stringMethod := personValue.MethodByName("String")
	if stringMethod.IsValid() {
		results := stringMethod.Call(nil)
		if len(results) > 0 {
			fmt.Printf("   String() result: %s\n", results[0].String())
		}
	}
	fmt.Println()

	// 10. Call methods on pointer
	fmt.Println("10. Call Methods on Pointer:")
	p3 := &Person{Name: "David", Age: 40}
	p3Value := reflect.ValueOf(p3)
	
	setNameMethod := p3Value.MethodByName("SetName")
	if setNameMethod.IsValid() {
		args := []reflect.Value{reflect.ValueOf("Daniel")}
		setNameMethod.Call(args)
		fmt.Printf("   After SetName: %s\n", p3.Name)
	}
	fmt.Println()

	// 11. Type checking
	fmt.Println("11. Type Checking:")
	var values []interface{} = []interface{}{42, "hello", 3.14, true, p}
	
	for _, val := range values {
		t := reflect.TypeOf(val)
		v := reflect.ValueOf(val)
		fmt.Printf("   Value: %v, Type: %s, Kind: %s\n", val, t, v.Kind())
	}
	fmt.Println()

	// 12. Slice and map reflection
	fmt.Println("12. Slice and Map Reflection:")
	slice := []int{1, 2, 3, 4, 5}
	sliceValue := reflect.ValueOf(slice)
	fmt.Printf("   Slice length: %d\n", sliceValue.Len())
	fmt.Printf("   Slice capacity: %d\n", sliceValue.Cap())
	fmt.Printf("   Element type: %s\n", sliceValue.Type().Elem())
	
	// Get element at index
	if sliceValue.Len() > 0 {
		firstElem := sliceValue.Index(0)
		fmt.Printf("   First element: %d\n", firstElem.Int())
	}
	
	// Map reflection
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	mapValue := reflect.ValueOf(m)
	fmt.Printf("   Map type: %s\n", mapValue.Type())
	fmt.Printf("   Map length: %d\n", mapValue.Len())
	
	// Iterate over map
	iter := mapValue.MapRange()
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		fmt.Printf("   %s -> %d\n", key.String(), value.Int())
	}
	fmt.Println()

	// 13. Create new instances
	fmt.Println("13. Create New Instances:")
	personType2 := reflect.TypeOf(Person{})
	newPerson := reflect.New(personType2).Elem()
	
	// Set fields
	nameField2 := newPerson.FieldByName("Name")
	if nameField2.IsValid() && nameField2.CanSet() {
		nameField2.SetString("Eve")
	}
	
	ageField2 := newPerson.FieldByName("Age")
	if ageField2.IsValid() && ageField2.CanSet() {
		ageField2.SetInt(28)
	}
	
	personPtr := newPerson.Addr().Interface().(*Person)
	fmt.Printf("   Created person: %s\n", personPtr)
	fmt.Println()

	// 14. Key concepts
	fmt.Println("14. Key Concepts:")
	fmt.Println("   - Reflection allows runtime type inspection")
	fmt.Println("   - Use reflect.TypeOf() for type information")
	fmt.Println("   - Use reflect.ValueOf() for value information")
	fmt.Println("   - Need pointer to set values")
	fmt.Println("   - Reflection has performance overhead")
	fmt.Println("   - Use sparingly, prefer type assertions when possible")
}

