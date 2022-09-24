package main

import (
	"testing"
)

func TestImport(t *testing.T) {
	file, err := Import("./data/test-input.png")
	if err != nil {
		t.Error(err)
	}

	if file == nil {
		t.Error("file is nil")
	}
}

func TestEncode(t *testing.T) {
	file, err := Import("./data/test-input.png")
	if err != nil {
		t.Error(err)
	}

	data := Data{}

	data, err = Encode(file)
	if err != nil {
		t.Error(err)
	}

	if data.RGBA == nil {
		t.Error("data.RGBA is nil")
	}
}

func TestSave(t *testing.T) {
	file, err := Import("./data/test-input.png")
	if err != nil {
		t.Error(err)
	}

	data := Data{}

	data, err = Encode(file)
	if err != nil {
		t.Error(err)
	}

	err = data.Save("./data/test-out.png")
	if err != nil {
		t.Error(err)
	}

	file, err = Import("./data/test-out.png")
	if err != nil {
		t.Error(err)
	}

	if file == nil {
		t.Error("output file is nil")
	}
}

func TestJson(t *testing.T) {
	file, err := Import("./data/test-input.png")
	if err != nil {
		t.Error(err)
	}

	data := Data{}

	data, err = Encode(file)
	if err != nil {
		t.Error(err)
	}

	json, err := data.Json()
	if err != nil {
		t.Error(err)
	}

	if json == "" {
		t.Error("json is empty")
	}
}

func TestDecodeJson(t *testing.T) {
	file, err := Import("./data/test-input.png")
	if err != nil {
		t.Error(err)
	}

	data := Data{}

	data, err = Encode(file)
	if err != nil {
		t.Error(err)
	}

	json, err := data.Json()
	if err != nil {
		t.Error(err)
	}

	newData := Data{}

	err = newData.DecodeJson(json)
	if err != nil {
		t.Error(err)
	}

	if newData.RGBA == nil {
		t.Error("newData.RGBA is nil")
	}
}
