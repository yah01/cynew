package cynew

import (
	"os"
	"sync"
)

const (
	Perm      = 0644
	Separator = string(os.PathSeparator)

	TemplateType_Empty      TemplateType = 0	// for creating a empty file
	TemplateType_SingleFile TemplateType = 1	// for creating single file with template
	TemplateType_Project    TemplateType = 2	// for creating a folder with template
)

var (
	FileDir     = os.Args[0][:len(os.Args[0])-len("cynew")]
	TemplateDir = FileDir + "/templates"
	WorkDir, _  = os.Getwd()

	WaitAllGoroutine = sync.WaitGroup{}
)
