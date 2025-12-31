# Template Engine

Learn about template engines in Go: `text/template` and `html/template` for dynamic content generation.

## Learning Objectives

- Understand text templates
- Learn about HTML templates
- Master template syntax
- Understand template variables and functions
- Learn about conditionals and loops
- Understand template nesting

## Concepts Covered

### Text Template

For plain text generation:

```go
import "text/template"

tmpl := "Hello, {{.}}!"
t := template.Must(template.New("greeting").Parse(tmpl))
t.Execute(os.Stdout, "World")
```

### HTML Template

For HTML generation (auto-escapes):

```go
import "html/template"

tmpl := `<h1>{{.Title}}</h1>`
t := template.Must(template.New("page").Parse(tmpl))
t.Execute(os.Stdout, data)
```

### Template Variables

Access struct fields:

```go
tmpl := `Name: {{.Name}}, Email: {{.Email}}`
t.Execute(os.Stdout, user)
```

### Template Functions

Use functions in templates:

```go
funcMap := template.FuncMap{
    "upper": strings.ToUpper,
}
t := template.Must(template.New("test").Funcs(funcMap).Parse(tmpl))
```

### Conditionals

```go
tmpl := `{{if .Age}}
{{.Name}} is {{.Age}} years old.
{{else}}
Age not specified.
{{end}}`
```

### Loops

```go
tmpl := `{{range .Users}}
- {{.Name}}
{{end}}`
```

### Template Nesting

```go
baseTmpl := `{{template "header" .}}
{{template "content" .}}`

headerTmpl := `{{define "header"}}<header>...</header>{{end}}`
```

## Running the Example

```bash
# Navigate to this directory
cd 06-Template-Engine

# Initialize the module (if not already done)
go mod init github.com/codinsec/go-learning-lab/06-standard-library-web/template-engine

# Run the program
go run main.go
```

## Running Tests

```bash
# Run all tests
go test

# Run tests with verbose output
go test -v

# Run tests with coverage
go test -cover
```

## Key Takeaways

1. **text/template** - For plain text
2. **html/template** - For HTML (auto-escapes)
3. **{{.}}** - Current context
4. **{{.Field}}** - Access struct fields
5. **{{if}}** - Conditionals
6. **{{range}}** - Loops

## Common Patterns

**Basic template:**
```go
tmpl := "Hello, {{.Name}}!"
t := template.Must(template.New("test").Parse(tmpl))
t.Execute(writer, data)
```

**Template from file:**
```go
t, err := template.ParseFiles("template.html")
t.Execute(writer, data)
```

**Template with functions:**
```go
t := template.Must(template.New("test").Funcs(funcMap).Parse(tmpl))
```

## Best Practices

- **Use html/template** - For HTML (prevents XSS)
- **Validate input** - Before templating
- **Keep templates simple** - Move logic to Go code
- **Use template functions** - For formatting
- **Cache templates** - Parse once, execute many

## Important Notes

- **html/template auto-escapes** - Prevents XSS attacks
- **text/template** - No escaping (for plain text)
- **Template parsing** - Can be expensive (cache templates)
- **Context (.)** - Current data object

## Next Steps

After understanding template engines, proceed to:
- [07-Logging](../07-Logging/README.md) - Learn about logging

---
**Created by:** [Codinsec](https://codinsec.com) | [info@codinsec.com](mailto:info@codinsec.com)  
**Author:** Barbaros Kaymak

