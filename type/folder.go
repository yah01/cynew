package cynew

func (folder Folder) GetFolderList() []Folder {
	return folder.Folders
}

func (folder Folder) GetFileList() []File {
	return folder.Files
}

func (folder Folder) IsEmpty() bool {
	return len(folder.Folders) == 0 && len(folder.Files) == 0
}