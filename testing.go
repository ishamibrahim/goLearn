package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"regexp"

	"github.com/olekukonko/tablewriter"
)

//TableFgHiBlueColor table foreground Blue color
var TableFgHiBlueColor = tablewriter.Colors{tablewriter.FgHiBlueColor}

//TableFgHiCyanColor table foreground Blue color
var TableFgHiCyanColor = tablewriter.Colors{tablewriter.FgHiCyanColor}

func readUserInput(promptMessage, optionRegex, inputKey, subExpName string) (string, bool) {
	var content string
	var isMatched bool
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(promptMessage)
	text, _ := reader.ReadString('\n')
	fmt.Println("text : ", text)
	if optionRegex != "" {
		re := regexp.MustCompile(optionRegex)
		if subExpName == "" {
			content = re.FindString(text)
			isMatched = re.MatchString(text)
		} else {
			matches := reSubMatchMap(re, text)
			content = matches[subExpName]
		}
		if !isMatched {
			errMsg := fmt.Sprintf("Not a valid %s. Please try again.", inputKey)
			fmt.Println(errMsg)
			return readUserInput(promptMessage, optionRegex, inputKey, subExpName)
		}
	} else {
		content = text
		isMatched = true
	}
	fmt.Println("Matchded text : ", content)
	return content, isMatched

}

func reSubMatchMap(r *regexp.Regexp, str string) map[string]string {
	match := r.FindStringSubmatch(str)
	subMatchMap := make(map[string]string)
	if len(match) > 0 {
		for i, name := range r.SubexpNames() {
			if i != 0 {
				subMatchMap[name] = match[i]
			}
		}
	}
	return subMatchMap
}

func printCSVTableData(inputData [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	var headerColorSlice []tablewriter.Colors
	var columnColorSlice []tablewriter.Colors

	lenData := len(inputData[0])
	for i := 0; i < lenData; i++ {
		headerColorSlice = append(headerColorSlice, TableFgHiBlueColor)
	}
	for i := 0; i < lenData; i++ {
		columnColorSlice = append(columnColorSlice, TableFgHiCyanColor)
	}

	table.SetHeader(inputData[0])
	table.SetHeaderColor(headerColorSlice...)
	table.SetColumnColor(columnColorSlice...)
	table.SetRowLine(true)
	for ind, inputRow := range inputData {
		fmt.Println("Inputs : ", inputRow)
		if ind != 0 {
			table.Append(inputRow)
		}
	}

	table.Render()
}

//readInputFile reads data from input CSV file
func readInputFile(filePath string) ([][]string, error) {
	var inputData [][]string

	if _, err := os.Stat(filePath); err != nil {
		return inputData, err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return inputData, err
	}

	r := csv.NewReader(file)
	records, err := r.ReadAll()

	if err != nil {
		return inputData, err
	}

	return records, err
}

func main() {
	// filepathString := readUserInput("Please upload CSV filepath using command  {-file ~[path to file]} : ", `\-file\s+(?P<fPath>[\w\d\\\/\_\-~\.]+)`, "Use '-file' before filepath", "fPath")
	// if filepathString != "" {
	// 	fmt.Println(" filepathString : ", filepathString)
	// 	filepathString = "/Users/isibrahi/Downloads/integration1.csv"
	// 	inputData, err := readInputFile(filepathString)
	// 	if err != nil {
	// 		fmt.Println("Error found", err.Error())
	// 		os.Exit(1)
	// 	}
	// 	printCSVTableData(inputData)
	// }

	templateOption, isMatched := readUserInput("Do you want to download CSV template to fill in property details? (Y/N) ?", `^((N|n)o|(Y|y)es|(Y|y)|(N|n))\n`, "Option", "")
	if isMatched {
		if matched, _ := regexp.Match(`(Y|y)es|(Y|y)`, []byte(templateOption)); matched {
			fmt.Println("MAtchd : ", matched)
		} else if matched, _ := regexp.Match(`N|n)o|(N|n)`, []byte(templateOption)); matched {
			fmt.Println("Not matched : ", matched)
		}
	}
}
