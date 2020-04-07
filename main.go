package main

import (
	"encoding/json"
	"fmt"
	"github.com/yah01/cyflag"
	. "github.com/yah01/cynew/type"
	"io/ioutil"
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
	addTemplateFlag string
	deleteFlag      string
	infoFlag        string

	fileDir     = getSelfPath()
	templateDir = fileDir + "/templates"
	workDir, _  = os.Getwd()
)

func init() {
	cyflag.BoolVar(&helpFlag, "-h", false, "show help information")
	cyflag.BoolVar(&listFlag, "-ls", false, "list all template(s)")
	cyflag.StringVar(&addTemplateFlag, "-t", "", "make a template with the file/folder")
	cyflag.StringVar(&deleteFlag, "-d", "", "delete template")
	cyflag.StringVar(&infoFlag, "-i", "", "show information of template")
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
func flagProcess() {
	err := cyflag.Parse()

	if err != nil || helpFlag == true {
		cyflag.PrintUsage()
	}

	if listFlag == true {
		if dir, err := ioutil.ReadDir(templateDir); err != nil {
			fmt.Println("reading dir error:", err)
		} else {
			for _, fileInfo := range dir {
				if !fileInfo.IsDir() {
					if fileContent, err := ioutil.ReadFile(templateDir + "/" + fileInfo.Name()); err == nil {
						var template Template
						if err = json.Unmarshal(fileContent, &template); err == nil {
							fmt.Printf("%v\t%v\n", template.Name, template.Info)
						}
					}
				}
			}
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
		// todo
	}

	if infoFlag != "" {
		// todo
	}
}

func main() {
	flagProcess()

	if len(cyflag.Args) == 1 {
		fileName := cyflag.Args[0]

		_, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, Perm)
		if err != nil {
			fmt.Println("Can't create file", fileName)
		}
	} else {

	}
}
