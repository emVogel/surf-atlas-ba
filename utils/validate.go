package utils

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

/**
* validate a upperCased FilterKey against the Spot struct
* if FilterKEy is part of the struct than the its valdiated positively
 */
	func ValidateFilterKey(key string, spotStructValue reflect.Value) bool {
		upperCaseKey := strings.ToUpper(key[:1]) + key[1:]
		fmt.Println("uppaercase key",upperCaseKey)
		 structFiledType:= spotStructValue.FieldByName(upperCaseKey)
		 fmt.Print(structFiledType)
		 return structFiledType.CanSet()
	}
	
	func ValidateFilterValue(value string) bool{
		forbiddenChars := []string{"/", "-", "\\", ";", "\"", "'", "*", "--", "!", "$"}
		forbiddenWords := []string{"Select", "Or", "And", "Drop", "Table", "Where", "Insert", "Into"}
		for index:= range forbiddenChars {
			if strings.ContainsAny(value, forbiddenChars[index]) {
				return false
			}
		}
		for ind:= range forbiddenWords{
			str := fmt.Sprintf("%s(%s|%s)", (forbiddenWords[ind][:1]), strings.ToUpper(forbiddenWords[ind][1:]), strings.ToLower(forbiddenWords[ind][1:])) // regexp would be S(elect|ELECT)
			fmt.Print(str)
			regExp, _ := regexp.Compile(str)
			matched := regExp.Match([]byte(value))
			if matched {
				return false
			}
	
		}
		return true
	}