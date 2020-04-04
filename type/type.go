package cynew

type rootType string

type Template struct {
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