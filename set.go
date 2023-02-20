package main

//UniqueLetters - O(1) Set of Unique ASCII lowercase letters
type UniqueLetters map[byte]struct{}

// Add a letter to the set
func (u UniqueLetters) Add(letter byte) {
	u[letter] = struct{}{}
}

// Remove a letter from the set
func (u UniqueLetters) Remove(letter byte) {
	delete(u, letter)
}

// Length return how many keys are in the set
func (u UniqueLetters) Length() int {
	counter := 0
	for range u {
		counter++
	}
	return counter
}

// Clear all keys from the set
func (u UniqueLetters) Clear() {
	for key := range u {
		delete(u, key)
	}
}

//Has - Check if a letter exists in the set
func (u UniqueLetters) Has(letter byte) bool {
	_, ok := u[letter]
	return ok
}
