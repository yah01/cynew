package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	DefaultTemp string
	DefaultSuffix string
}

func getSelfPath() string {
	return os.Args[0][0:len(os.Args[0])-len("cynew.exe")]
}

const (
	Perm = 0644
)

var (
	help bool
	list bool
	create bool
	temp string
	suffix string
	add string
	info string

	config Config

	fileDir = getSelfPath()
	tempDir = fileDir+"/templates"
	workDir,_ = os.Getwd()
	configPath = fileDir+"/config.json"
)

func init() {
	configFile, _ := ioutil.ReadFile(configPath)
	json.Unmarshal(configFile, &config)

	flag.BoolVar(&help,"h",false,"show help information")
	flag.BoolVar(&list,"ls",false,"list all template(s)")
	flag.BoolVar(&create,"c",false,"create file(s) without template")
	flag.StringVar(&temp,"t",config.DefaultTemp,"set default template")
	flag.StringVar(&suffix,"s",config.DefaultSuffix,"set default suffix")
	flag.StringVar(&add,"a","","add *filename* into templates")
	flag.StringVar(&info,"i","","show information of *temp*")
}

func trimSuffixName(suf string) string {
	name := []byte(suf)

	for i:=len(name)-1;i>=0;i-- {
		if name[i] == '.' {
			name = name[0:i]
			break
		}
	}

	return string(name)
}

func flagProcess() {
	flag.Parse()

	if help==true {
		flag.Usage()
	}

	if list==true {
		dir, _ := ioutil.ReadDir(tempDir)

		for i := 0; i < len(dir); i++ {

			fmt.Println(dir[i].Name())
		}
	}

	config.DefaultTemp = temp
	config.DefaultSuffix = suffix

	if add != "" {
		file,err := ioutil.ReadFile(add)
		if err != nil {
			fmt.Println("No such file:",add)
		} else {
			ioutil.WriteFile(tempDir + "/" + trimSuffixName(add), file,Perm)
		}
	}

	if info != "" {

	}
}

func main() {
	flagProcess()

	if flag.NArg()==1 {
		create = true
	}

	if create == true {
		for _,name := range flag.Args() {
			ioutil.WriteFile(workDir+"/"+name,nil,Perm)
		}
	} else {
		tempName := flag.Arg(flag.NArg()-1)
		file,err := ioutil.ReadFile(tempDir+"/"+tempName)

		if err != nil {
			fmt.Println("No such templates:",tempName)
		} else {
			for i:=0;i<flag.NArg()-1;i++ {
				name := flag.Arg(i)
				ioutil.WriteFile(workDir+"/"+name, file, Perm)
			}
		}
	}
}
