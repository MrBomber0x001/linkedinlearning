```go
 stocks := map[string]float64{
  "AMZN": 2034.5,
  "GOOG": 3434.5,
  "MSTF": 365.5,
 }

 fmt.Println(len(stocks))
 fmt.Println(stocks["AMZN"])
 fmt.Println(stocks["NOT_HERE"]) // get the zero value of the map which is 0 because of float64

 value, ok := stocks["TESLA"]
 if !ok {
  fmt.Println("cant find tesla")
 } else {
  fmt.Println(value)
 }

 // Set
 stocks["new_one"] = 545.6
 fmt.Println(stocks)

 delete(stocks, "new_one")
 fmt.Println(stocks)

 for key := range stocks {
  fmt.Println(key)
 }

 for k, v := range stocks {
  fmt.Println(k, v)
 }
```

count words in text

```go
func main() {
 text := `
 Yousef Meska
 Omar Meska
 Yousef Mahmoud
 Mahmoud Meska`
 // countWords(text)
 count := countWords(text)
 fmt.Println(count)
}

func timeTracker(start time.Time, name string) {
 elapsed := time.Since(start)
 log.Printf("%s taken %s", name, elapsed)
}

func countWords(text string) map[string]int {

 // lowercase the string and split the string into words
 // check if each word appear on the map, if not increment the count
 defer timeTracker(time.Now(), "countWords")
 words := strings.Fields(text)
 count := map[string]int{}
 for _, word := range words {
  count[strings.ToLower(word)]++
 }
 return count

}

```
