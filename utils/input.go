package utils

import "fmt"

func GetUserInputNumber() (int, error) {
	var i int
	_, err := fmt.Scanf("%d", &i)
	return i, err
}
