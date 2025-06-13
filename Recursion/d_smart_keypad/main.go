package smartkeypad

import "fmt"

func smartKeyPadHelper(number int, keypadMapping map[int]string, word string) {
	if number == 0 {
		fmt.Println(word)
		return
	}

	// lastdigit := number % 10
	// mapping := keypadMapping[number]

	// for _ v := range mapping {

	// }
}

func SmartKeypad(number int) {
	keypadMapping := map[int]string{
		1: "",
		2: "ABC",
		3: "DEF",
	}

	smartKeyPadHelper(number, keypadMapping, "")
}
