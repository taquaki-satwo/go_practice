package main

import "fmt"

// MyError コメント
type MyError struct {
	message string
}

func (err MyError) Error() string {
	return err.message
}

func main() {
	val1, err1 := hex2int("1")
	fmt.Println(val1, err1)

	val2, err2 := hex2int("abcd")
	fmt.Println(val2, err2)

	val3, err3 := hex2int("z")
	fmt.Println(val3, err3)
}

func hex2int(hex string) (val int, err error) {
	for _, r := range hex {
		val *= 16
		switch {
		case '0' <= r && r <= '9':
			val += int(r - '0')
		case 'a' <= r && r <= 'f':
			val += int(r-'a') + 10
		default:
			return 0, MyError{"不正な文字です:" + string(r)}
		}
	}

	return
}
