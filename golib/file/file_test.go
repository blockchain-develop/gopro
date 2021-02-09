package file

import (
	"os"
	"testing"
)

func TestFile(t *testing.T) {
	file, _ := os.Open("")
	defer file.Close()

	file.Seek(0, 0)

	data := make([]byte, 100)
	readed, _ := file.Read(data)
}
