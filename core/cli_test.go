package core

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
	"time"
)

func TestFly_Stop(t *testing.T) {
	r := Fly{}
	r.Projects = append(r.Schema.Projects, Project{exit: make(chan os.Signal, 1)})
	r.Stop()
	_, ok := <-r.Projects[0].exit
	if ok != false {
		t.Error("Unexpected error", "channel should be closed")
	}
}

func TestFly_Start(t *testing.T) {
	r := Fly{}
	err := r.Start()
	if err == nil {
		t.Error("Error expected")
	}
	r.Projects = append(r.Projects, Project{Name: "test", exit: make(chan os.Signal, 1)})
	go func() {
		time.Sleep(100)
		close(r.Projects[0].exit)
		_, ok := <-r.Projects[0].exit
		if ok != false {
			t.Error("Unexpected error", "channel should be closed")
		}
	}()
	err = r.Start()
	if err != nil {
		t.Error("Unexpected error", err)
	}
}

func TestFly_Prefix(t *testing.T) {
	r := Fly{}
	input := "test"
	result := r.Prefix(input)
	if len(result) == 0 && !strings.Contains(result, input) {
		t.Error("Unexpected error")
	}
}

func TestLogWriter_Write(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	w := LogWriter{}
	input := ""
	val, err := w.Write([]byte(input))
	if err != nil || val > 0 {
		t.Error("Unexpected error", err, "string length should be 0 instead", val)
	}
}
