package main

import (
	"bytes"
	"errors"
	"github.com/devypt/fly/core"
	"log"
	"strings"
	"testing"
)

var mockResponse interface{}

type mockFly core.Fly

func (m *mockFly) add() error {
	if mockResponse != nil {
		return mockResponse.(error)
	}
	m.Projects = append(m.Projects, core.Project{Name: "One"})
	return nil
}

func (m *mockFly) setup() error {
	if mockResponse != nil {
		return mockResponse.(error)
	}
	return nil
}

func (m *mockFly) start() error {
	if mockResponse != nil {
		return mockResponse.(error)
	}
	return nil
}

func (m *mockFly) clean() error {
	if mockResponse != nil {
		return mockResponse.(error)
	}
	return nil
}

func (m *mockFly) remove() error {
	if mockResponse != nil {
		return mockResponse.(error)
	}
	m.Projects = []core.Project{}
	return nil
}

func TestFly_add(t *testing.T) {
	m := mockFly{}
	mockResponse = nil
	if err := m.add(); err != nil {
		t.Error("Unexpected error")
	}
	if len(m.Projects) == 0 {
		t.Error("Unexpected error")
	}

	m = mockFly{}
	m.Projects = []core.Project{{Name: "Default"}}
	mockResponse = nil
	if err := m.add(); err != nil {
		t.Error("Unexpected error")
	}
	if len(m.Projects) != 2 {
		t.Error("Unexpected error")
	}

	m = mockFly{}
	mockResponse = errors.New("error")
	if err := m.clean(); err == nil {
		t.Error("Expected error")
	}
	if len(m.Projects) != 0 {
		t.Error("Unexpected error")
	}
}

func TestFly_start(t *testing.T) {
	m := mockFly{}
	mockResponse = nil
	if err := m.add(); err != nil {
		t.Error("Unexpected error")
	}
}

func TestFly_setup(t *testing.T) {
	m := mockFly{}
	mockResponse = nil
	if err := m.setup(); err != nil {
		t.Error("Unexpected error")
	}
}

func TestFly_clean(t *testing.T) {
	m := mockFly{}
	mockResponse = nil
	if err := m.clean(); err != nil {
		t.Error("Unexpected error")
	}
	mockResponse = errors.New("error")
	if err := m.clean(); err == nil {
		t.Error("Expected error")
	}
}

func TestFly_remove(t *testing.T) {
	m := mockFly{}
	mockResponse = nil
	if err := m.remove(); err != nil {
		t.Error("Unexpected error")
	}

	m = mockFly{}
	mockResponse = nil
	m.Projects = []core.Project{{Name: "Default"}, {Name: "Default"}}
	if err := m.remove(); err != nil {
		t.Error("Unexpected error")
	}
	if len(m.Projects) != 0 {
		t.Error("Unexpected error")
	}

	mockResponse = errors.New("error")
	if err := m.clean(); err == nil {
		t.Error("Expected error")
	}
}

func TestFly_version(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	version()
	if !strings.Contains(buf.String(), core.RVersion) {
		t.Error("Version expted", core.RVersion)
	}
}
