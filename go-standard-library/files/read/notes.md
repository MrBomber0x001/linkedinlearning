## File information

Lstat and stat return FileInfo structure

```go
type FileInfo interface {
    Name() string
    Size() int64
    Mode() FileMode
    ModTime() time.Time
    IsDir() bool
    Sys() interface{}
}

```

## Writing

There are many ways to write to a file, we're going to start with the easiest way!
`ioutil` package provides a high level functionalites for dealing with files!

```go

func main() {
 data1 := []byte("This is some sample text\n")
 fmt.Println(data1) // slice of bytes! [80 22 33]
 ioutil.WriteFile("sample.txt", data1, 0644) // writes the entire file!
}
```

There's no append function in Go as in Go 1.15, I don't know if there's any in 1.18
