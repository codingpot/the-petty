# deepcopy

## import

`go get github.com/codingpod/the-petty/deepcopy`

`import github.com/codingpod/the-petty/deepcopy`

## how to use

```go
import (
    "fmt"

    "github.com/codingpod/the-petty/deepcopy"
)

type Point struct {
    X int
    Y int
}

func main() {
    a := Point{
        X: 4,
        Y: 4,
    }
    b, err := deepcopy.Copy(a)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(b)
}
```

## supported

* [X] pointer
* [X] primitive type
* [X] slice
* [X] array
* [X] channel
* [ ] interface
* [ ] function