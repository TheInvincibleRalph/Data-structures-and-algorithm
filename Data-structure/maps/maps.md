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

## Concepts in Hashmaps

### Collision within Buckets

In a hash table, a **bucket** is a storage location where data is kept. Each bucket can hold one or more pieces of data. When multiple pieces of data end up in the same bucket (a situation called a "collision"), they are typically stored in a linked list within that bucket.

Here's a simplified explanation:

1. **Hashing**: When you add a key-value pair to a hash table, a hash function calculates an index based on the key. This index determines which bucket to store the data in.

2. **Buckets and Linked Lists**: If two different keys hash to the same bucket, the hash table uses a linked list to store these items. The bucket contains a pointer to the first node in the list. Each node contains a key, a value, and a pointer to the next node in the list.

3. **Handling Collisions**: When inserting a new item, if the target bucket already contains items (i.e., the linked list has nodes), the new item is added to the list. This helps manage collisions by chaining items together.


### Load Factor

- **Definition**: The load factor of a hash table is a measure of how full the hash table is. It is typically calculated as the ratio of the number of elements (`n`) to the number of buckets (`m`).
- **Formula**: Load Factor = `n / m`
  - `n`: Number of elements (or key-value pairs) stored in the hash table.
  - `m`: Number of buckets or slots in the hash table.

#### Why Resize?

- **Performance**: When a hash table has a high load factor, it means that many of the buckets are likely to have multiple elements (chains), leading to longer search times. This can degrade the performance of the hash table, particularly for operations like insert, delete, and search.
- **Threshold**: To maintain efficient operation, hash tables often resize (increase the number of buckets) when the load factor exceeds a certain threshold. A common threshold is 0.7, meaning that if more than 70% of the buckets are filled, the hash table is resized.

#### Resizing Process

1. **Increase Size**: The number of buckets is increased, usually doubled.
2. **Rehashing**: All the existing elements are rehashed and placed into the new, larger array of buckets. This is necessary because the hash function depends on the size of the array, and changing the array size alters the hash values.

### Open Addressing
Open addressing is a collision resolution strategy used in hash tables. Instead of using linked lists to store multiple elements that hash to the same slot (as in chaining), open addressing stores all elements directly in the hash table array. When a collision occurs, the algorithm seeks the next available slot in the array according to a specific sequence until an empty slot is found.

### Common Open Addressing Techniques
1. **Linear Probing**: Incrementally checks the next slot (with wrap-around) until an empty slot is found.
2. **Quadratic Probing**: Uses a quadratic function to determine the next slot to check.
3. **Double Hashing**: Applies a second hash function to resolve collisions, using the output to determine the step size.