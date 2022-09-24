# imgx

JSON encoder and decoder for PNG images. Example can be found [here](data/example.json).

## Usage

### Encoding

```go
package main

import (
	"fmt"

	imgx "github.com/eminmuhammadi/imgx"
)

data := imgx.Data{}

file, err := imgx.Import("input.png")
if err != nil {
	panic(err)
}

data, err = imgx.Encode(file)
if err != nil {
	panic(err)
}

json, err := data.Json()
if err != nil {
	panic(err)
}

fmt.Println(json)
```

### Decoding

```go
package main

import (
    "fmt"

    imgx "github.com/eminmuhammadi/imgx"
)

data := imgx.Data{}
json := "{...}"

err := data.DecodeJson(json)
if err != nil {
	t.Error(err)
}

err = data.Save("output.png")
if err != nil {
	panic(err)
}
```
