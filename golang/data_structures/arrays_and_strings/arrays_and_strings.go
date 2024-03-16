/*
Author: Nennii Cally-Ntete
Date updated: March 15th 2024
Description: Implementation of ArrayLists and Hash Tables
*/

package arraysandstrings

import "fmt"

type ArrayList struct {
}

/*
Hash tables are data structures that map keys to values for highly efficient lookup
*/

// Golangs native implementation of a hash table
var myMap = make(map[string]int)

// Base implementation
type HashTable struct {
	size  int
	table [][]KeyValuePair // Slice containing a slice of key value pairs
}

type KeyValuePair struct {
	key   string
	value int
}

// Hash function to generate an index from a key
func (ht *HashTable) hash(key string) int {
	hash := 0
	for _, char := range key {
		hash += int(char)
	}
	return hash % ht.size
}

// Insert a key value pair into the hash table
func (ht *HashTable) Set(key string, value int) {
	// Create a hash from the key
	index := ht.hash(key)

	// Set the key value pair in the hash table
	ht.table[index] = append(ht.table[index], KeyValuePair{key: key, value: value})
}

// Get the value associated with a key
func (ht *HashTable) Get(key string) (int, error) {
	// Get the index of the key
	index := ht.hash(key)

	// Get the key value pair from the hash table
	keyValuePairs := ht.table[index]
	for _, entry := range keyValuePairs {
		if entry.key == key {
			return entry.value, nil
		}
	}
	return 0, fmt.Errorf("Key not found.")
}

// Delete a key from the hash table
func (ht *HashTable) Delete(key string) error {
	// Get the hash
	index := ht.hash(key)

	// Delete from hash table
	keyValuePairs := ht.table[index]
	for i, entry := range keyValuePairs {
		if entry.key == key {
			// Remove the entry from the slice
			ht.table[index] = append(keyValuePairs[:i], keyValuePairs[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Key not found.")
}
