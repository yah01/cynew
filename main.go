package main

import (
	"encoding/json"
	"fmt"
	"github.com/yah01/cyflag"
	"github.com/yah01/cynew/store"
	. "github.com/yah01/cynew/type"
	"io/ioutil"
	"log"
	"os"
)

var (
	helpFlag        bool
	listFlag        bool
	templateFlag    string
	addTemplateFlag string
	deleteFlag      string
	infoFlag        string
)

func init() {
	cyflag.BoolVar(&helpFlag, "-h", false, "show help information")
	cyflag.BoolVar(&listFlag, "-ls", false, "list all template(s)")
	cyflag.StringVar(&templateFlag, "-t", "", "create file/folder with template")
	cyflag.StringVar(&addTemplateFlag, "-a", "", "make a template with the file/folder")
	cyflag.StringVar(&deleteFlag, "-d", "", "delete template")
	cyflag.StringVar(&infoFlag, "-i", "", "show information of template")

	_, err := ioutil.ReadDir(TemplateDir)
	if err != nil {
		os.Mkdir(TemplateDir, Perm)
	}
}

func main() {
	err := flagProcess()
	if err != nil {
		log.Fatalln(err)
	}
	template := store.ReadTemplateFile(templateFlag)
	WaitAllGoroutine.Add(len(cyflag.Args))
	if template.Type == TemplateType_Empty || template.Type == TemplateType_SingleFile {
		for _, fileName := range cyflag.Args {
			go store.CreateFile(WorkDir+Separator+fileName, template)
		}
	} else {
		for _, dirName := range cyflag.Args {
			go store.CreateDir(WorkDir+Separator+dirName, template)
		}
	}

	WaitAllGoroutine.Wait()
}

func trimSuffixName(suf string) string {
	name := []byte(suf)

	for i := len(name) - 1; i >= 0; i-- {
		if name[i] == '.' {
			name = name[0:i]
			break
		}
	}

	return string(name)
}

func hasSuffixName(suf string) bool {
	return suf != trimSuffixName(suf)
}

// Parse flags and execute what the flags mean
func flagProcess() error {
	err := cyflag.Parse()

	if err != nil {
		cyflag.PrintUsage()
	}

	if helpFlag {
		cyflag.PrintUsage()
	}

	if listFlag == true {
		if dir, err := ioutil.ReadDir(TemplateDir); err != nil {
			fmt.Println("reading dir error:", err)
		} else {
			for _, fileInfo := range dir {
				if !fileInfo.IsDir() {
					if fileContent, err := ioutil.ReadFile(TemplateDir + Separator + fileInfo.Name()); err == nil {
						var template Template
						if err = json.Unmarshal(fileContent, &template); err == nil {
							fmt.Printf("%v\t%v\n", template.Name, template.Info)
						}
					}
				}
			}
		}
	}

	if templateFlag != "" {
		if len(cyflag.Args) == 0 {
			return fmt.Errorf("no file/folder name")
		}
	}

	if addTemplateFlag != "" {
		var (
			file   []byte
			folder []os.FileInfo
		)

		file, err := ioutil.ReadFile(addTemplateFlag)
		if err != nil {
			file = nil
			folder, err = ioutil.ReadDir(addTemplateFlag)
		}

		if err != nil {
			fmt.Println("Read file/folder error:", addTemplateFlag)
		} else {
			if file != nil {
				store.CreateFileTemplate(addTemplateFlag)
			} else if folder != nil {
				store.CreateFolderTemplate(addTemplateFlag)
			}
		}
	}

	if deleteFlag != "" {
		if err := os.Remove(TemplateDir + Separator + deleteFlag); err != nil {
			fmt.Println("Delete templateFlag", deleteFlag, "error:", err)
		}
	}

	if infoFlag != "" {
		if fileContent, err := ioutil.ReadFile(TemplateDir + Separator + infoFlag); err == nil {
			var template Template
			if err = json.Unmarshal(fileContent, &template); err == nil {
				fmt.Printf("%v\t%v\n", template.Name, template.Info)
			}
		}
	}

	return err
}
