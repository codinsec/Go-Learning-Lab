package main

import (
	"bytes"
	"strings"
	"testing"
	texttemplate "text/template"
)

func TestTextTemplate(t *testing.T) {
	tmpl := "Hello, {{.}}!"
	t := texttemplate.Must(texttemplate.New("test").Parse(tmpl))
	
	var buf bytes.Buffer
	t.Execute(&buf, "World")
	
	result := buf.String()
	if !strings.Contains(result, "Hello") || !strings.Contains(result, "World") {
		t.Errorf("Expected 'Hello, World!', got %s", result)
	}
}

func TestTemplateWithVariables(t *testing.T) {
	tmpl := `Name: {{.Name}}`
	t := texttemplate.Must(texttemplate.New("test").Parse(tmpl))
	
	user := User{Name: "Alice"}
	var buf bytes.Buffer
	t.Execute(&buf, user)
	
	result := buf.String()
	if !strings.Contains(result, "Alice") {
		t.Errorf("Expected 'Alice', got %s", result)
	}
}

func TestTemplateWithConditionals(t *testing.T) {
	tmpl := `{{if .Age}}{{.Name}}{{else}}No age{{end}}`
	t := texttemplate.Must(texttemplate.New("test").Parse(tmpl))
	
	user := User{Name: "Bob", Age: 30}
	var buf bytes.Buffer
	t.Execute(&buf, user)
	
	result := buf.String()
	if !strings.Contains(result, "Bob") {
		t.Errorf("Expected 'Bob', got %s", result)
	}
}

func TestTemplateWithLoops(t *testing.T) {
	tmpl := `{{range .}}{{.Name}}{{end}}`
	t := texttemplate.Must(texttemplate.New("test").Parse(tmpl))
	
	users := []User{
		{Name: "Alice"},
		{Name: "Bob"},
	}
	
	var buf bytes.Buffer
	t.Execute(&buf, users)
	
	result := buf.String()
	if !strings.Contains(result, "Alice") || !strings.Contains(result, "Bob") {
		t.Errorf("Expected 'Alice' and 'Bob', got %s", result)
	}
}

