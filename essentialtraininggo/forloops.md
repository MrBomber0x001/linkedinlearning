```go
n := 10

 switch n {
 case 3:
  fmt.Println("welcome")
 }
 switch {
 case n > 10:
  fmt.Println("welcome")
 case n < 10:
  fmt.Println("welcome welcome")
 default:
  fmt.Println("default")
 }

 for i := 0; i < 3; i++ {
  if i == 1 {
   break
  }
  if i == 2 {
   continue
  }
 }

 a := 4
 //like while loop
 for a < 3 {
  a++
 }
 b := 2
 // while true
 for {
  if b < 2 {
   break
  }
  b++
 }
```

```go
for i := 0; i <= 20; i++ {
  if i%3 == 0 && i%5 == 0 {
   fmt.Println("FizzBuzz")
  } else if i%3 == 0 {
   fmt.Println("Fizz")
  } else if i%5 == 0 {
   fmt.Println("Buzz")
  } else {

   fmt.Println(i)
  }
 }
```
