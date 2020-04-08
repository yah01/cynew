package cynew

import (
	"os"
	"sync"
)

const (
	Perm      = 0644
	Separator = string(os.PathSeparator)

	TemplateType_Empty      TemplateType = 0
	TemplateType_SingleFile TemplateType = 1
	TemplateType_Project    TemplateType = 2
)

var (
	FileDir     = os.Args[0][:len(os.Args[0])-len("cynew.exe")]
	TemplateDir = FileDir + "/templates"
	WorkDir, _  = os.Getwd()

	WaitAllGoroutine = sync.WaitGroup{}
)
