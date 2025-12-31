package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"
	texttemplate "text/template"
)

// This program demonstrates template engines in Go

type User struct {
	Name  string
	Email string
	Age   int
}

type PageData struct {
	Title string
	Users []User
}

func main() {
	fmt.Println("=== Template Engine ===")
	fmt.Println()

	// 1. Text template
	fmt.Println("1. Text Template:")
	textTemplate()
	fmt.Println()

	// 2. HTML template
	fmt.Println("2. HTML Template:")
	htmlTemplate()
	fmt.Println()

	// 3. Template with variables
	fmt.Println("3. Template with Variables:")
	templateWithVariables()
	fmt.Println()

	// 4. Template with functions
	fmt.Println("4. Template with Functions:")
	templateWithFunctions()
	fmt.Println()

	// 5. Template with conditionals
	fmt.Println("5. Template with Conditionals:")
	templateWithConditionals()
	fmt.Println()

	// 6. Template with loops
	fmt.Println("6. Template with Loops:")
	templateWithLoops()
	fmt.Println()

	// 7. Template from file
	fmt.Println("7. Template from File:")
	templateFromFile()
	fmt.Println()

	// 8. Template nesting
	fmt.Println("8. Template Nesting:")
	templateNesting()
	fmt.Println()
}

func textTemplate() {
	tmpl := "Hello, {{.}}!\n"
	t := texttemplate.Must(texttemplate.New("greeting").Parse(tmpl))
	t.Execute(os.Stdout, "World")
}

func htmlTemplate() {
	tmpl := `<h1>{{.Title}}</h1>
<p>Welcome to {{.Title}}!</p>`
	
	t := template.Must(template.New("page").Parse(tmpl))
	data := struct {
		Title string
	}{
		Title: "Go Learning Lab",
	}
	t.Execute(os.Stdout, data)
}

func templateWithVariables() {
	tmpl := `Name: {{.Name}}
Email: {{.Email}}
Age: {{.Age}}`
	
	t := texttemplate.Must(texttemplate.New("user").Parse(tmpl))
	user := User{
		Name:  "Alice",
		Email: "alice@example.com",
		Age:   30,
	}
	t.Execute(os.Stdout, user)
}

func templateWithFunctions() {
	tmpl := `{{.Name | upper}} - {{.Email}}`
	
	funcMap := template.FuncMap{
		"upper": func(s string) string {
			return strings.ToUpper(s)
		},
	}
	
	t := texttemplate.Must(texttemplate.New("user").Funcs(funcMap).Parse(tmpl))
	user := User{Name: "Bob", Email: "bob@example.com"}
	t.Execute(os.Stdout, user)
}

func templateWithConditionals() {
	tmpl := `{{if .Age}}
{{.Name}} is {{.Age}} years old.
{{else}}
Age not specified.
{{end}}`
	
	t := texttemplate.Must(texttemplate.New("conditional").Parse(tmpl))
	
	user1 := User{Name: "Alice", Age: 30}
	fmt.Println("With age:")
	t.Execute(os.Stdout, user1)
	
	user2 := User{Name: "Bob", Age: 0}
	fmt.Println("Without age:")
	t.Execute(os.Stdout, user2)
}

func templateWithLoops() {
	tmpl := `Users:
{{range .Users}}
- {{.Name}} ({{.Email}})
{{end}}`
	
	t := texttemplate.Must(texttemplate.New("list").Parse(tmpl))
	data := PageData{
		Users: []User{
			{Name: "Alice", Email: "alice@example.com"},
			{Name: "Bob", Email: "bob@example.com"},
			{Name: "Charlie", Email: "charlie@example.com"},
		},
	}
	t.Execute(os.Stdout, data)
}

func templateFromFile() {
	// Create a temporary template file
	tmplContent := `Hello, {{.Name}}!`
	
	tmpFile, err := os.CreateTemp("", "template-*.txt")
	if err == nil {
		defer os.Remove(tmpFile.Name())
		tmpFile.WriteString(tmplContent)
		tmpFile.Close()
		
		t, err := texttemplate.ParseFiles(tmpFile.Name())
		if err == nil {
			user := User{Name: "David"}
			t.Execute(os.Stdout, user)
		}
	}
}

func templateNesting() {
	baseTmpl := `{{template "header" .}}
{{template "content" .}}
{{template "footer" .}}`
	
	headerTmpl := `{{define "header"}}<header>Header</header>{{end}}`
	contentTmpl := `{{define "content"}}<main>{{.}}</main>{{end}}`
	footerTmpl := `{{define "footer"}}<footer>Footer</footer>{{end}}`
	
	t := texttemplate.Must(texttemplate.New("base").Parse(baseTmpl))
	texttemplate.Must(t.Parse(headerTmpl))
	texttemplate.Must(t.Parse(contentTmpl))
	texttemplate.Must(t.Parse(footerTmpl))
	
	t.Execute(os.Stdout, "Content")
}

