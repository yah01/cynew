package main

import (
	"encoding/json"
	"fmt"
	"github.com/yah01/cyflag"
	. "github.com/yah01/cynew/type"
	"io/ioutil"
	"os"
	"os/exec"
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

func flagProcess() {
	err := cyflag.Parse()

	if err != nil || helpFlag == true {
		cyflag.PrintUsage()
	}

	if listFlag == true {
		if dir, err := ioutil.ReadDir(templateDir); err != nil {
			fmt.Println("error when reading dir: %+v", err)
		} else {
			for _, fileInfo := range dir {
				if !fileInfo.IsDir() {
					if fileContent, err := ioutil.ReadFile(templateDir + "/" + fileInfo.Name()); err == nil {
						var template Template
						if err = json.Unmarshal(fileContent, &template); err == nil {
							fmt.Println("%v\t%v", template.Name, template.Info)
						}
					}
				}
			}
		}
	}

	if addTemplateFlag != "" {
		var(
			file []byte
			folder []os.FileInfo
		)
		file, err := ioutil.ReadFile(addTemplateFlag)
		if err != nil {
			folder,err = ioutil.ReadDir(addTemplateFlag)
		}
		if err != nil {
			fmt.Println("Read file/folder error: %v", addTemplateFlag)
		} else {
			if file != nil {

			}
			ioutil.WriteFile(templateDir+"/"+trimSuffixName(addFlag), file, Perm)
		}
	}

	if deleteFlag != "" {

	}

	if infoFlag != "" {

	}
}

func main() {
	flagProcess()

	if len(cyflag.Args) == 1 {
		createFlag = true
	}

	if createFlag == true {
		for _, name := range cyflag.Args {
			if !hasSuffixName(name) {
				name += config.DefaultSuffix
			}

			ioutil.WriteFile(workDir+"/"+name, nil, Perm)

			if openFlag {
				cmd := exec.Command("cmd", "/k", "start", workDir+"/"+name)
				cmd.Start()
			}
		}
	} else if len(cyflag.Args) > 0 {
		tempName := cyflag.Args[len(cyflag.Args)-1]
		file, err := ioutil.ReadFile(templateDir + "/" + tempName)

		if err != nil {
			fmt.Println(cyflag.Args)
			fmt.Println("No such templates:", tempName)
		} else {
			for i := 0; i < len(cyflag.Args)-1; i++ {
				name := cyflag.Args[i]
				if !hasSuffixName(name) {
					name += config.DefaultSuffix
				}

				ioutil.WriteFile(workDir+"/"+name, file, Perm)

				if openFlag {
					cmd := exec.Command("cmd", "/k", "start", workDir+"/"+name)
					cmd.Start()
				}
			}
		}
	}
}
