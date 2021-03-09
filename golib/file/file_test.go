package file

import (
	"fmt"
	"os"
	"testing"
)

func TestFile(t *testing.T) {
	file, _ := os.Open("")
	defer file.Close()

	file.Seek(0, 0)

	data := make([]byte, 100)
	readed, err := file.Read(data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("read: %d\n", readed)
}
