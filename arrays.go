package main

import (
	"sort"
	"strings"
	"unicode/utf8"
)

func IsUnique(input string) bool {
    input_array := strings.Split(input, ""); 

    sort.Slice(input_array, func(i, j int) bool {
        a, _ := utf8.DecodeRuneInString(input_array[i]);
        b, _ := utf8.DecodeRuneInString(input_array[j]);
        return a < b;     
    });

    var prev string = "a" 
    for index, element := range input_array {
        if index != 0 && prev == element {
            return false
        }
        prev = element;
    }
    
    return true;
}

func CheckPermutation(inputOne string, inputTwo string) bool {
    if (len(inputOne) != len(inputTwo)) {
        return false;
    }

    inputOneArray := strings.Split(inputOne, "");
    inputTwoArray := strings.Split(inputTwo, "");

    sort.Slice(inputOneArray, func(i, j int) bool {
        return inputOneArray[i] < inputOneArray[j];
    });

    sort.Slice(inputTwoArray, func(i, j int) bool {
        return inputTwoArray[i] < inputTwoArray[j];
    });

    for index, el:= range inputOneArray {
        if el != inputTwoArray[index] {
            return false;
        }
    }

    return true;
}

func Urlify(input string) string {
    inputArray := strings.Split(input, "");
    for index, letter := range inputArray {
        if letter == " " {
            inputArray[index] = "%20";
        }
    }

    return strings.Join(inputArray, "");
}

func IsPermOfPalnidrome(input string) bool {
    if len(input) == 0 || len(input) == 1 {
        return true;
    }
    inputArray := strings.Split(input, "");

    letterCount := make(map[string]int);
    for _, letter := range inputArray {
        letterCount[letter] += 1;
    }

    foundOdd := false;
    for _, val := range letterCount {
        if val % 2 != 0 && foundOdd {
            return false;
        } else if val % 2 != 0 {
            foundOdd = true;
        }
    }

    return true;
}
