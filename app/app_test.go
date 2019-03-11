package app

import (
	"testing"
)

func TestConfig(t *testing.T) {
	_, err := getConfig()
	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestDBConnect(t *testing.T) {
	db, err := dbConnect()
	if err != nil {
		t.Errorf("%v", err)
	}
	defer db.Close()

	if _, err := db.Exec("SELECT version()"); err != nil {
		t.Errorf("%v", err)
	}
}

func TestWriteEntity(t *testing.T) {
	wc := webComic{
		Month:      "12",
		Num:        10,
		Year:       "2019",
		Transcript: "Hello, World!",
		Day:        "01",
	}

	if err := wc.write(); err != nil {
		t.Errorf("%v", err)
	}
}

func TestReadEntity(t *testing.T) {
	var wc webComic

	if err := wc.read(10); err != nil {
		t.Errorf("%v", err)
	}
}
