package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {

	// ioutil.TempDir() will create directory and
	// it will return either directory name or error
	tempDir, err := ioutil.TempDir("", "username-")
	check(err)
	defer os.RemoveAll(tempDir)

	// Now we will create file inside our temDir directory
	// ioutil.TempFile() will create file and
	// it will return either file name or error
	tempFile, err := ioutil.TempFile(tempDir, "name-*.png")
	check(err)
	defer os.Remove(tempFile.Name())

	// This will print out full name and path of our image
	fmt.Println(tempFile.Name())

	// This will write data inside the file
	if _, err := tempFile.Write([]byte("Name: Ashutosh\n")); err != nil {
		fmt.Println(err)
	}
	// This will read data present inside our file
	// if it is unable to read printout the reason
	data, err := ioutil.ReadFile(tempFile.Name())
	check(err)

	// Convert stream data into string and print it in output
	fmt.Println(string(data))

}
