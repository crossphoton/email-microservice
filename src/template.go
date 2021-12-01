package src

import (
	"bytes"
	"fmt"
	"html/template"
)

func getTemplate(templateName string) (*template.Template, error) {
	t, err := template.New(fmt.Sprintf("%s.html", templateName)).ParseFiles(fmt.Sprintf("./templates/%s.html", templateName))
	if err != nil {
		return nil, fmt.Errorf("invalid template: %v", err)
	}

	return t, nil
}

// getMessageFromTemplate fills in the template parameters and
// returns the final message formed
func getMessageFromTemplate(templateName string, params map[string]string) ([]byte, error) {
	t, err := getTemplate(templateName)
	if err != nil {
		return nil, err
	}

	var message bytes.Buffer
	err = t.Execute(&message, params)
	if err != nil {
		return nil, err
	}

	return message.Bytes(), nil
}
