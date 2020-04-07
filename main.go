package main

import (
	"encoding/json"
	"fmt"
	"github.com/yah01/cyflag"
	. "github.com/yah01/cynew/type"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	DefaultTemp   string `json:"default_temp"`
	DefaultSuffix string `json:"default_suffix"`
}

func getSelfPath() string {
	return os.Args[0][:len(os.Args[0])-len("cynew.exe")]
}

var (
	helpFlag        bool
	listFlag        bool
	templateFlag    string
	addTemplateFlag string
	deleteFlag      string
	infoFlag        string

	fileDir     = getSelfPath()
	templateDir = fileDir + "/templates"
	workDir, _  = os.Getwd()
)

func init() {
	cyflag.BoolVar(&helpFlag, "-h", false, "show help information")
	cyflag.BoolVar(&listFlag, "-ls", false, "list all templateFlag(s)")
	cyflag.StringVar(&templateFlag, "-t", "", "create file/folder with templateFlag")
	cyflag.StringVar(&addTemplateFlag, "-a", "", "make a templateFlag with the file/folder")
	cyflag.StringVar(&deleteFlag, "-d", "", "delete templateFlag")
	cyflag.StringVar(&infoFlag, "-i", "", "show information of templateFlag")
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
		return err
	}

	if helpFlag {
		cyflag.PrintUsage()
	}

	if listFlag == true {
		if dir, err := ioutil.ReadDir(templateDir); err != nil {
			fmt.Println("reading dir error:", err)
		} else {
			for _, fileInfo := range dir {
				if !fileInfo.IsDir() {
					if fileContent, err := ioutil.ReadFile(templateDir + Separator + fileInfo.Name()); err == nil {
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
			folder, err = ioutil.ReadDir(addTemplateFlag)
		}
		if err != nil {
			fmt.Println("Read file/folder error:", addTemplateFlag)
		} else {
			if file != nil {
				// todo
			} else if folder != nil {
				// todo
			}
		}
	}

	if deleteFlag != "" {
		if err := os.Remove(templateDir + Separator + deleteFlag); err != nil {
			fmt.Println("Delete templateFlag", deleteFlag, "error:", err)
		}
	}

	if infoFlag != "" {
		if fileContent, err := ioutil.ReadFile(templateDir + Separator + infoFlag); err == nil {
			var template Template
			if err = json.Unmarshal(fileContent, &template); err == nil {
				fmt.Printf("%v\t%v\n", template.Name, template.Info)
			}
		}
	}

	return nil
}

func main() {
	err := flagProcess()
	if err != nil {
		log.Fatalln(err)
	}

	for _, projectName := range cyflag.Args {
		// Without template
		if templateFlag == "" {
			_, err = os.Create(projectName)
			if err != nil {
				fmt.Println("Create file", projectName, "error:", err)
			}
		} else {
			// todo
		}
	}
}
