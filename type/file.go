package cynew

import "os"

func (file File) Create(path string) {
	newFile,_ := os.Create(path)

	newFile.Close()
}