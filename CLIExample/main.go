package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	var optionalArgs = map[string]string{}
	if len(args) == 0 {
		fmt.Println("Please enter the correct arguments")
	} else {
		for index, options := range args {
			if options == "-full" {
				optionalArgs["full"] = "true"
			}
			if options == "-case" {
				if len(args) > index+1 {
					if args[index+1] == "upper" || args[index+1] == "lower" {
						optionalArgs["case"] = args[index+1]
					} else {
						fmt.Println("Please provide correct case value")
						return
					}
				} else {
					fmt.Println("Please provide correct case value")
					return
				}

			}
		}

		switch {
		case args[0] == "name":
			// printName(true, optionalArgs["case"])
			var name = "Pushpa"
			var result = handleOptionalArgs(name, optionalArgs)
			fmt.Println(result)
		case args[0] == "company":
			fmt.Println("Red Hat")
		case args[0] == "hometown":
			fmt.Println("Coastal")
		default:
			fmt.Println("invalid")
		}
	}
}
func handleOptionalArgs(name string, optionalArgs map[string]string) string {
	var result = name
	for option, values := range optionalArgs {
		switch {
		case option == "full":
			result = handleFullArg(result)
		case option == "case":
			result = handleCaseArg(values, result)
		default:
			result = name
		}
	}
	return (result)
}
func handleFullArg(name string) string {
	if strings.ToUpper("pushpa") == name {
		return "PUSHPA PADTI"
	} else if strings.ToLower("pushpa") == name {
		return "pushpa padti"
	} else {
		return "Pushpa Padti"
	}

}
func handleCaseArg(values string, name string) string {
	if values == "lower" {
		return strings.ToLower(name)
	} else if values == "upper" {
		return strings.ToUpper(name)
	} else {
		return "error"
	}

}

// func printName(isFull bool, kase string) error {
// 	fmt.Println(isFull, kase)
// 	finalName := "Pushpa"
// 	if isFull {
// 		finalName += " Padti"
// 	}

// 	switch kase {
// 	case "upper":
// 		finalName = strings.ToUpper(finalName)
// 	case "lower":
// 		finalName = strings.ToLower(finalName)
// 	default:
// 		return fmt.Errorf("unknown case value")
// 	}

// 	fmt.Println(finalName)
// 	return nil
// }
