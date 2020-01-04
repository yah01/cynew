package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	DefaultTemp string
}

func getSelfPath() string {
	return os.Args[0][0:len(os.Args[0])-len("cynew.exe")]
}

func main() {
	var config Config
	configFile, _ := ioutil.ReadFile(getSelfPath() + "/config.json")
	json.Unmarshal(configFile, &config)
	path, _ := os.Getwd()

	switch len(os.Args) {
	case 2:
		switch os.Args[1] {
		case "-ls":
			dir, _ := ioutil.ReadDir(getSelfPath())

			for i := 0; i < len(dir); i++ {
				switch dir[i].Name() {
				case "config.json", "cynew.exe":
					continue

				default:
					fmt.Println(dir[i].Name())
				}
			}

		default:
			fmt.Println("create file", os.Args[1])
			ioutil.WriteFile(path+"/"+os.Args[1], nil, 0644)
		}
	case 3:
		switch os.Args[1] {
		case "-d":
			config.DefaultTemp = os.Args[2]
			configFile, _ = json.Marshal(config)
			ioutil.WriteFile(getSelfPath()+"/config.json", configFile, 0644)
		default:
			template, err := ioutil.ReadFile(getSelfPath() + "/" + os.Args[2])

			if err != nil {
				fmt.Println("not find template", os.Args[2])
				ioutil.WriteFile(path+"/"+os.Args[1], nil, 0644)
			} else {
				fmt.Println("create template", os.Args[2])
				ioutil.WriteFile(path+"/"+os.Args[1], template, 0644)
			}
		}
	}
}
