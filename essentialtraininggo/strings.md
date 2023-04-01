```go
book := "The color of my book is green!"
 fmt.Println(len(book))
 fmt.Printf("book[0] = %v (type %T)\n", book[0], book[0])
 //strings are immutable
 // book[0] = "b" // error

 fmt.Println(book[3:11])

 newBook := "t" + book[:]
 fmt.Println(newBook)

 poem := `
 this is raw string
 ...book`
 fmt.Println(poem)
 n := 42
 s := fmt.Sprintf("%d", n) // "42"
 fmt.Printf("s = %q (type %T)\n", s, s)
```

print even-ended number

```go
 count := 0

 start := time.Now()
 for a := 1000; a <= 9999; a++ {
  for b := a; b <= 9999; b++ {
   n := a * b
   s := fmt.Sprintf("%d", n)
   if s[0] == s[len(s)-1] {
    count++
   }
  }
 }
 fmt.Println(count)
 elapsed := time.Since(start)

 fmt.Println(elapsed)
```
