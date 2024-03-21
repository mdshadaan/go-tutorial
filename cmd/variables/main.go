package main
import "fmt"
import "unicode/utf8"

func main(){
	var intNumber int = 1234
	//int will default to either 32 or 64 bits depending on your machine
	var anotherInt int64 = 67890122
	var unsignedInt uint32 = 345
	//unsigned int stores only positive integers
	fmt.Println(intNumber , anotherInt, unsignedInt)

	var myName string = "shadaan";
	var multiLine string = `Hi
	I'm a sentence
	on another line
	`
	fmt.Println(myName)
	fmt.Println(multiLine)
	//rune are a data type in go for stroing characters
	fmt.Println(utf8.RuneCountInString(myName))

	var myChar rune = 'A';
	fmt.Println(myChar)

	var myBool bool = true
	fmt.Println(myBool)

	//default values for all ints,floats , unsigned ints and rune is 0
	// for string it is "" and for boolean its false
	var defaultVal int
	fmt.Println(defaultVal)

	shortHand := "Strings can be defined like this"  //we can omit var keyword  and datatype if we initialise the value like this.
	fmt.Println(shortHand)

	const MyConst string = "This is a constant , I can't change it's value afterwards"
	fmt.Println(MyConst)

}