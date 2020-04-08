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
	fileContent, err := ioutil.ReadFile(TemplateDir + Separator + templateName+".json")
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

func CreateFile(path string, content TemplateContent) {
	defer WaitAllGoroutine.Done()
	newFile, _ := os.Create(path)
	if !content.IsEmpty() {
		file := content.GetFileList()[0]
		newFile.Write(file.Content)
	}
	newFile.Close()
}

func CreateFileTemplate(fileName string, content []byte) {
	var (
		templateName string
		info         string
		template     Template
	)

	fmt.Print("Template name: ")
	fmt.Scanln(&templateName)
	if templateName == "" {
		templateName = fileName
	}
	fmt.Print("Template information (no necessary): ")
	fmt.Scanln(&info)

	jsonFile, err := os.Create(TemplateDir + Separator + templateName + ".json")
	if err != nil {
		fmt.Println("Create template", templateName, "error:", err)
		return
	}

	template = Template{
		Name:    templateName,
		Info:    info,
		Type:    TemplateType_SingleFile,
		Folders: nil,
		Files: []File{
			{
				Name:    "", // Name is unused for template with type TemplateType_SingleFile
				Content: content,
			},
		},
	}

	templateJson, err := json.Marshal(template)
	jsonFile.Write(templateJson)
	jsonFile.Close()
}

func CreateFolderTemplate(folderName string, content []os.FileInfo) {

}
