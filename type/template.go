package cynew

func (template Template) GetFolderList() []Folder {
	return template.Folders
}

func (template Template) GetFileList() []File {
	return template.Files
}

func (template Template) IsEmpty() bool {
	return template.Type == TemplateType_Empty
}