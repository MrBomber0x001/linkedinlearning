```go
// spread operator
func main(){
    values := []int{1,2,3};
    sum := add(values...);
    fmt.Println(sum);
}

func add(values ...int) int {
    sum := 0;
    for _, v := range values {
        sum += v;
    }
    return sum;
}
```

```go
loans := []string{"yousef", "ahmed", "hussien"}
 fmt.Println(loans)
 fmt.Printf("loans = %v (type %T)", loans, loans)

 fmt.Println(len(loans))

 fmt.Println(loans[1])

 fmt.Println(loans[1:])

 for i := 0; i < len(loans); i++ {
  fmt.Println(loans[i])
 }

 for _, name := range loans {
  fmt.Println(name)
 }

 loans = append(loans, "meska")
 fmt.Println("new", loans)
```

print max value

```go
func main() {
 nums := []int{13, 26, 66, 44, 35}
 max := printMax(nums)
 fmt.Println(max)
}

func timeTracker(start time.Time, name string) {
 elapsed := time.Since(start)
 log.Printf("%s taken %s", name, elapsed)
}

func printMax(nums []int) int {
 defer timeTracker(time.Now(), "printMax")
 max := nums[0]
 for i := range nums {
  if max < nums[i] {
   max = nums[i]
  }
 }
 return max
}
```
