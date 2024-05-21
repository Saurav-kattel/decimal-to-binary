package converter

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Converter struct {
	userInput string
}

func (c *Converter) getUserInput() {
	fmt.Println("Enter The Decimal Number: ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Split(text, "\n")[0]
	c.userInput = text
}

func (c *Converter) parseData(text string) (int, error) {
	paresdData, err := strconv.Atoi(text)
	if err != nil {
		return 0, err
	}
	return paresdData, nil
}
func hasDecimalPoint(number float64) bool {
	str := strconv.FormatFloat(number, 'f', -1, 64)
	return strings.Contains(str, ".")
}

func (c *Converter) generateIntBinary(digit int) string {
	if digit == 0 {

		return ""
	}
	remainder := digit % 2
	quotient := int(math.Floor(float64(digit) / 2))

	binary := c.generateIntBinary(quotient)
	return binary + fmt.Sprint(remainder)
}

func (c *Converter) generateFloatBinary(digit int, depth *int) string {

	if *depth > 20 {
		return ""
	}

	*depth++
	if digit == 0 {
		return ""
	}

	validDecimal := "0." + fmt.Sprint(digit)
	floatNumber, _ := strconv.ParseFloat(validDecimal, 64)
	quotient := floatNumber * 2
	intValue := math.Floor(quotient)
	if !hasDecimalPoint(quotient) {
		return fmt.Sprint(intValue)
	}
	str, _ := strconv.ParseInt(strings.Split(fmt.Sprint(quotient), ".")[1], 10, 64)
	res := c.generateFloatBinary(int(str), depth)
	return res + fmt.Sprint(intValue)

}

func (c *Converter) reverse(str string) string {
	runes := []rune(str)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}

func (c *Converter) Convert() string {
	depth := 0
	c.getUserInput()
	slicedData := strings.Split(c.userInput, ".")
	integerPart, _ := c.parseData(slicedData[0])
	intVal := c.generateIntBinary(integerPart)
	if len(slicedData) > 1 {
		floatingPart, _ := c.parseData(slicedData[1])
		floatVal := c.generateFloatBinary(floatingPart, &depth)
		str := c.reverse(floatVal)
		return intVal + "." + str
	}
	return intVal
}
