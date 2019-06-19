package twofer

func ExampleShareWith() {
	for _, test := range tests {
		ShareWith(test.name)
		// Output: One for you, one for me.
		// One for Alice, one for me.
		// One for Bob, one for me.
	}

}
