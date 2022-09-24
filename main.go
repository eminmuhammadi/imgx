package main

import (
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
)

type Data struct {
	Type   string `json:"type"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	RGBA   []RGBA `json:"rgba"`
}

type RGBA struct {
	R uint32 `json:"r"`
	G uint32 `json:"g"`
	B uint32 `json:"b"`
	A uint32 `json:"a"`
}

// Import PNG file
func Import(path string) (*os.File, error) {
	// path
	path, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	// if file is not png
	if filepath.Ext(path) != ".png" {
		return nil, fmt.Errorf("file is not png")
	}

	// import image
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return file, nil
}

// Encode PNG file to Data
func Encode(file *os.File) (Data, error) {
	file.Seek(0, 0)

	// Decode image
	img, err := png.Decode(file)
	if err != nil {
		return Data{}, err
	}

	defer file.Close()

	// data (image.Image) to rgb pixels
	width, height := img.Bounds().Max.X, img.Bounds().Max.Y
	data := &Data{
		Type:   "image/png",
		Width:  width,
		Height: height,
		RGBA:   nil,
	}

	// Get pixels
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			data.RGBA = append(data.RGBA, RGBA{
				R: r,
				G: g,
				B: b,
				A: a,
			})
		}
	}

	return *data, nil
}

// Save PNG file
func (data *Data) Save(path string) error {
	// Create new image from rgb pixels
	newImage := image.NewRGBA(image.Rect(0, 0, data.Width, data.Height))

	for y := 0; y < data.Height; y++ {
		for x := 0; x < data.Width; x++ {
			c := color.RGBA{
				R: uint8(data.RGBA[y*data.Width+x].R),
				G: uint8(data.RGBA[y*data.Width+x].G),
				B: uint8(data.RGBA[y*data.Width+x].B),
				A: uint8(data.RGBA[y*data.Width+x].A),
			}
			newImage.Set(x, y, c)
		}
	}

	// absoulte path
	path, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	// *image.RGBA to png file
	out, err := os.Create(path)
	if err != nil {
		return err
	}

	defer out.Close()

	err = png.Encode(out, newImage)
	if err != nil {
		return err
	}

	return nil
}

// Json encode Data to json
func (data *Data) Json() (string, error) {
	// return Data as json
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

// Json data to PNG file
func (data *Data) DecodeJson(jsonData string) error {
	// json to Data
	err := json.Unmarshal([]byte(jsonData), data)
	if err != nil {
		return err
	}

	return nil
}
