package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var arr [30000]byte
	var curr byte = 0
	var args = os.Args[1:]

	if len(args) != 1 {
		fmt.Print("Missing file argument")
		return
	}

	file, err := ioutil.ReadFile(args[0])
	if err != nil {
		fmt.Print(err)
		return
	}

	for i := 0; i < len(file); i++ {
		switch file[i] {
		// +
		case 43:
			arr[curr] = arr[curr] + 1
		// -
		case 45:
			arr[curr] = arr[curr] - 1
		// ,
		case 44:
			reader := bufio.NewReader(os.Stdin)
			char, err := reader.ReadByte()
			if err != nil {
				fmt.Print(err)
				return
			}
			arr[curr] = char

		// .
		case 46:
			fmt.Print(string(arr[curr]))

		// [
		case 91:
			if arr[curr] == 0 {
				var skips int = 0
				for true {
					i = i + 1
					if file[i] == 91 {
						skips = skips + 1
					}
					if file[i] == 93 {
						if skips == 0 {
							break
						} else {
							skips = skips - 1
						}

					}
				}
			}

		// ]
		case 93:
			if arr[curr] != 0 {
				var skips int = 0
				for true {
					i = i - 1
					if file[i] == 93 {
						skips = skips + 1
					}
					if file[i] == 91 {
						if skips == 0 {
							break
						} else {
							skips = skips - 1
						}

					}
				}
			}

		// <
		case 60:
			curr = curr - 1

		// >
		case 62:
			curr = curr + 1

		}
	}
}
