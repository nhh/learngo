package main

import "fmt"

const englishPrefix = "Hello, "
const spanishPrefix = "Hola, "

func convertStringToBinary(str string) string {
	binaryString := ""
	for _, char := range str {
		binaryString = fmt.Sprintf("%s%.8b", binaryString, char)
	}
	return binaryString
}

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}

	switch lang {
	case "English":
		return englishPrefix + name
	case "Spanish":
		return spanishPrefix + name
	default:
		return convertStringToBinary(englishPrefix) + convertStringToBinary(name)
	}

}

func main() {
	fmt.Println(Hello("Niklas", "English"))
}
