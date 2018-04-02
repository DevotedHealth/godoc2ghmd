package main

// Must will panic if err isn't nil
func Must(err error) {
	if err != nil {
		panic(err)
	}
}
