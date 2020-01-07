package main

import (
	"cyFlag"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

type Config struct {
	DefaultTemp   string
	DefaultSuffix string
}

func getSelfPath() string {
	return os.Args[0][0 : len(os.Args[0])-len("cynew.exe")]
}

const (
	Perm = 0644
)

var (
	help   bool
	list   bool
	create bool
	open   bool
	temp   string
	suffix string
	add    string
	info   string

	config Config

	fileDir    = getSelfPath()
	tempDir    = fileDir + "/templates"
	workDir, _ = os.Getwd()
	configPath = fileDir + "/config.json"
)

func init() {
	configFile, _ := ioutil.ReadFile(configPath)
	json.Unmarshal(configFile, &config)

	cyFlag.BoolVar(&help, "-h", false, "show help information")
	cyFlag.BoolVar(&list, "-ls", false, "list all template(s)")
	cyFlag.BoolVar(&create, "-c", false, "create file(s) without template")
	cyFlag.BoolVar(&open, "-o", false, "open with OS default program")
	cyFlag.StringVar(&temp, "-t", config.DefaultTemp, "set default template")
	cyFlag.StringVar(&suffix, "-s", config.DefaultSuffix, "set default suffix")
	cyFlag.StringVar(&add, "-a", "", "add *filename* into templates")
	cyFlag.StringVar(&info, "-i", "", "show information of *temp*")
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
	err := cyFlag.Parse()

	if err != nil {
		cyFlag.Usage()
	}

	if help == true {
		cyFlag.Usage()
	}

	if list == true {
		dir, _ := ioutil.ReadDir(tempDir)

		for i := 0; i < len(dir); i++ {

			fmt.Println(dir[i].Name())
		}
	}

	config.DefaultTemp = temp
	if suffix == "none" {
		config.DefaultSuffix = ""
	} else if config.DefaultSuffix != suffix {
		config.DefaultSuffix = "." + suffix
	}

	if add != "" {
		file, err := ioutil.ReadFile(add)
		if err != nil {
			fmt.Println("No such file:", add)
		} else {
			ioutil.WriteFile(tempDir+"/"+trimSuffixName(add), file, Perm)
		}
	}

	if info != "" {

	}

	JSON, _ := json.Marshal(&config)
	ioutil.WriteFile(configPath, JSON, Perm)
}

func main() {
	flagProcess()

	// fmt.Println(create)
	// fmt.Println(cyFlag.Args)

	if len(cyFlag.Args) == 1 {
		create = true
	}

	if create == true {
		for _, name := range cyFlag.Args {
			if !hasSuffixName(name) {
				name += config.DefaultSuffix
			}

			ioutil.WriteFile(workDir+"/"+name, nil, Perm)

			if open {
				cmd := exec.Command("cmd", "/k", "start", workDir+"/"+name)
				cmd.Start()
			}
		}
	} else if len(cyFlag.Args) > 0 {
		tempName := cyFlag.Args[len(cyFlag.Args)-1]
		file, err := ioutil.ReadFile(tempDir + "/" + tempName)

		if err != nil {
			fmt.Println(cyFlag.Args)
			fmt.Println("No such templates:", tempName)
		} else {
			for i := 0; i < len(cyFlag.Args)-1; i++ {
				name := cyFlag.Args[i]
				if !hasSuffixName(name) {
					name += config.DefaultSuffix
				}

				ioutil.WriteFile(workDir+"/"+name, file, Perm)

				if open {
					cmd := exec.Command("cmd", "/k", "start", workDir+"/"+name)
					cmd.Start()
				}
			}
		}
	}
}
