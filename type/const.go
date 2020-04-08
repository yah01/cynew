package cynew

import (
	"os"
	"sync"
)

const (
	Perm      = 0777
	Separator = string(os.PathSeparator)

	TemplateType_Empty      TemplateType = 0 // for creating a empty file
	TemplateType_SingleFile TemplateType = 1 // for creating single file with template
	TemplateType_Project    TemplateType = 2 // for creating a folder with template
)

var (
	tempPath, _ = os.Executable()
	FileDir     = tempPath[:len(tempPath)-len("/cynew")]
	TemplateDir = FileDir + Separator + "templates"
	WorkDir, _  = os.Getwd()

	WaitAllGoroutine = sync.WaitGroup{}
)
