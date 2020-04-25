package cynew

type TemplateContent interface {
	GetFolderList() []Folder
	GetFileList() []File
	IsEmpty() bool
}

type TemplateType uint8
type Template struct {
	Name    string
	Info    string
	Type    TemplateType
	Folders []Folder `json:"folders"`
	Files   []File   `json:"files"`
}

type Folder struct {
	Name    string   `json:"name"`
	Folders []Folder `json:"folders"`
	Files   []File   `json:"files"`
}

type File struct {
	Name    string `json:"name"`
	Content []byte `json:"content"`
}
