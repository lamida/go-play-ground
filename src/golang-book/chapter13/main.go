package main

import(
	"fmt"
	"os"
)

func main(){
	file, err := os.Open("test.txt")
	if err != nil {
		// handle error
		fmt.Println("Erros is happening", err)
		return
	}

	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		fmt.Println("error is happening", err)
		return
	}

	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		fmt.Println("erros is happening again", err)	
		return
	}

	str := string(bs)
	fmt.Println(str)
}

