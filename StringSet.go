package main

/*
  Represents a distinct set of strings.
*/
type StringSet struct {
	values []string
}

/*
  Adds a string to this set, if it does not already exist in the set.
*/
func (this *StringSet) Add(items... string) {

	for _, item := range items {
		if !this.Contains(item) {
			this.values = append(this.values, item)
		}
	}
}

/*
  Returns true if this set contains the given [item], false otherwise.
*/
func (this StringSet) Contains(item string) bool {

	for _, extant := range this.values {

		if extant == item {
			return true
		}
	}

	return false
}

/*
  Returns a slice representing all items contained by this set.
*/
func (this StringSet) GetSlice() []string {
	return this.values
}

func (this StringSet) Length() int {
	return len(this.values)
}
