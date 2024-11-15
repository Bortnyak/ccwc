package ccwc

import (
	"flag"
	"fmt"
	"os"
)

func Run() {
	cFlag := flag.Bool("c", false, "Enables bytes counting")

	flag.Parse()
	if *cFlag {
		// fmt.Println("Bytes counting is turned on")

		args := flag.Args()
		if len(args) < 1 {
			fmt.Println("Please provide a file name")
			return
		}

		fileName := args[0]

		fileInfo, err := os.Stat(fileName)
		if err != nil {
			fmt.Println("Error getting file info:", err)
			return
		}

		// Get the size in bytes
		fileSize := fileInfo.Size()

		fmt.Println(fileSize, fileName)
	}

}
