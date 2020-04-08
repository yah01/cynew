package store

import (
	"encoding/json"
	"fmt"
	. "github.com/yah01/cynew/type"
	"io/ioutil"
	"os"
)

func ReadTemplateFile(templateName string) Template {
	var template Template
	fileContent, err := ioutil.ReadFile(TemplateDir + Separator + templateName)
	if err != nil {
		return Template{}
	}

	if err = json.Unmarshal(fileContent, &template); err != nil {
		return Template{}
	}

	return template
}

func CreateDir(path string, content TemplateContent) {
	defer WaitAllGoroutine.Done()

	err := os.Mkdir(path, Perm)
	if err != nil {
		fmt.Println("Create dir", path, "error:", err)
		return
	}

	folders := content.GetFolderList()
	files := content.GetFileList()
	WaitAllGoroutine.Add(len(folders) + len(files))
	for _, folder := range folders {
		go CreateDir(path+Separator+folder.Name, folder)
	}
	for _, file := range files {
		go file.Create(path + Separator + file.Name)
	}
}

func CreateFile(path string,content TemplateContent) {
	defer WaitAllGoroutine.Done()
	newFile,_ := os.Create(path)
	if !content.IsEmpty() {
		file := content.GetFileList()[0]
		newFile.Write(file.Content)
	}
	newFile.Close()
}
