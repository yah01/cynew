package store

import (
	"fmt"
	"github.com/yah01/cybuf-go"
	. "github.com/yah01/cynew/type"
	"io/ioutil"
	"os"
	"path/filepath"
)

func ReadTemplateFile(templateName string) Template {
	var template Template
	fileContent, err := ioutil.ReadFile(TemplateDir + Separator + templateName + ".cybuf")
	if err != nil {
		return Template{}
	}
	if err = cybuf.Unmarshal(fileContent, &template); err != nil {
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
		go CreateDir(path+Separator+folder.Name, &folder)
	}
	for _, file := range files {
		go CreateFile(path+Separator+file.Name, Folder{
			Folders: nil,
			Files: []File{
				file,
			},
		})
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

func CreateFileTemplate(fileName string) {
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

	cybufFile, err := os.Create(TemplateDir + Separator + templateName + ".cybuf")
	if err != nil {
		fmt.Println("Create template", templateName, "error:", err)
		return
	}

	content, err := ioutil.ReadFile(WorkDir + Separator + fileName)

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

	templateCybuf, err := cybuf.Marshal(template)
	if err != nil {
		fmt.Println("Error: %+v", err)
		return
	}
	cybufFile.Write(templateCybuf)
	cybufFile.Close()
}

func CreateFolderTemplate(folderName string) {
	var (
		templateName string
		info         string
		template     Template
		folder       Folder
	)

	fmt.Print("Template name: ")
	fmt.Scanln(&templateName)
	if templateName == "" {
		templateName = folderName
	}
	fmt.Print("Template information (no necessary): ")
	fmt.Scanln(&info)

	cybufFile, err := os.Create(TemplateDir + Separator + templateName + ".cybuf")
	if err != nil {
		fmt.Println("Create template", templateName, "error:", err)
		return
	}

	template = Template{
		Name:    templateName,
		Info:    info,
		Type:    TemplateType_Project,
		Folders: nil,
		Files:   nil,
	}

	path := WorkDir + Separator + folderName
	TransformFolderToTemplate(path, &folder)
	template.Folders, template.Files = folder.Folders, folder.Files
	templateCybuf, err := cybuf.Marshal(template)
	cybufFile.Write(templateCybuf)
	cybufFile.Close()
}

func TransformFolderToTemplate(path string, folderContent *Folder) {

	dir, err := ioutil.ReadDir(path)

	if err != nil {
		fmt.Println("Read dir", filepath.Base(path), "error:", err)
		return
	}

	for _, file := range dir {
		if file.IsDir() {
			folder := Folder{
				Name:    file.Name(),
				Folders: nil,
				Files:   nil,
			}
			TransformFolderToTemplate(path+Separator+file.Name(), &folder)
			folderContent.Folders = append(folderContent.Folders, folder)
		} else {
			content, err := ioutil.ReadFile(path + Separator + file.Name())
			if err != nil {
				fmt.Println("Read file", file.Name(), "error:", err)
			}

			folderContent.Files = append(folderContent.Files, File{
				Name:    file.Name(),
				Content: content,
			})
		}
	}
}
