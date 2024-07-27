### Maps and Dictionaries in Go: Crash Course

In Go, maps (also known as dictionaries in other languages) are unordered collections of key-value pairs. Each key in a map is unique, and you can quickly look up the value associated with a given key. Here’s an in-depth look at how maps work in Go, along with examples of common operations.

#### Creating and Initializing Maps

To declare a map, you use the `map` keyword followed by the key type and value type in square brackets. 

```go
var myMap map[string]int // Declares a map with string keys and int values
```

To initialize a map, you can use the `make` function or a map literal.

```go
// Using make
myMap = make(map[string]int)

// Using map literal
myMap = map[string]int{
    "apple":  5,
    "banana": 10,
}
```

#### Basic Operations

1. **Adding and Updating Elements:**
   You can add or update elements in a map by assigning a value to a key.

```go
myMap["orange"] = 15 // Adds a new key "orange" with value 15
myMap["apple"] = 7   // Updates the value of "apple" to 7
```

2. **Retrieving Elements:**
   You can retrieve a value by specifying its key.

```go
value := myMap["apple"]
fmt.Println(value) // Output: 7
```

3. **Deleting Elements:**
   Use the `delete` function to remove a key-value pair.

```go
delete(myMap, "banana")
```

4. **Checking for Key Existence:**
   To check if a key exists, you can use the two-value assignment form.

```go
value, exists := myMap["banana"]
if exists {
    fmt.Println("banana exists with value", value)
} else {
    fmt.Println("banana does not exist")
}
```

#### Iterating Over Maps

You can iterate over all key-value pairs in a map using a `for` loop.

```go
for key, value := range myMap {
    fmt.Println(key, value)
}
```

#### Example Program: Word Frequency Counter

Here's a complete example that reads a string and counts the frequency of each word using a map.

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    text := "This is a simple test. This test is only a test."

    // Initialize a map to hold word counts
    wordCount := make(map[string]int)

    // Split the text into words
    words := strings.Fields(text)

    // Count the frequency of each word
    for _, word := range words {
        word = strings.ToLower(word) //converts the current "word" to lowercase to ensure case insensitivity in word counting
        word = strings.Trim(word, ".") //removes all leading or trailing period from the current "word"
        wordCount[word]++ //increments a particular word by 1 if it never exists prior, also keeps count of the words as the loop furthers
    }

    // Print the word counts
    for word, count := range wordCount {
        fmt.Printf("%s: %d\n", word, count)
    }
}
```

#### Nested Maps

Maps in Go can also be nested, which means you can have a map whose values are other maps.

```go
// A map of maps
nestedMap := make(map[string]map[string]int)

nestedMap["fruits"] = map[string]int{
    "apple":  5,
    "banana": 10,
}

nestedMap["vegetables"] = map[string]int{
    "carrot": 3,
    "pepper": 7,
}

// Accessing nested map values
fmt.Println(nestedMap["fruits"]["apple"])  // Output: 5
fmt.Println(nestedMap["vegetables"]["carrot"]) // Output: 3
```