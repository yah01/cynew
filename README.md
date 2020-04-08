# cynew

## Installation
You can install cynew by the way you prefer.

### Release
1. Download [cynew](https://github.com/yah01/cynew/releases)
2. Add cynew into environment variable PATH
```shell
export PATH=$PATH:[folder where you install cynew]
```

### Build
You need install [Golang](https://golang.org/) first
1. Clone the repo
```shell
go get github.com/yah01/cynew
```
2. Change directory to where cynew exists
```shell
cd $GOPATH/src/github.com/yah01/cynew
```
3. Run build command
```shell
go build
```
4. There will be a executable file
5. Add cynew into environment variable PATH
```shell
export PATH=$PATH:[folder where you install cynew]
```

## Usage

### simply create a empty file with name *hello.cpp*
```shell
cynew hello.cpp
```

### create more empty files with name *first.go*, *second.json*
```shell
cynew first.go second.json
```

Examples above are simple and nothing awesome, they can't show power of cynew. You may want to create a file/folder with some contents (what I call *template*), you can indicate a template when you create file/folder, cynew will write the contents what the template saving into the file/folder.

### create file with template
Maybe you are building a awesome project, and many source code file have some common contents (template my-template):
```shell
cynew awesome.go -t my-template
```
You can create more than one file by just adding more file names

### create folder with template
Maybe you used scaffold what creates a folder with some contents (files,folders,... I would call also it template), cynew is a scaffold creator and builder!
```
cynew my-app -t my-awesome-project
```
command above will create a folder with name my-app, and fill it with contents what template my-awesome-project saving.

### create template with existing file/folder
Now you know how to create file/folder with template, but where is the template? You can create template with cynew easily
```shell
cynew -a hello.go
```
This will create a template with content what hello.go has, then cynew will ask you to input a template name (if you input nothing, the template name is the file name). You can create a folder template by the same way.

## flag

- [-h]: show help information
- [-ls]: list all template(s)
- [-t *template*]: create file/folder with template
- [-a *filename*]: add *filename* into templates (folder is OK)
- [-d *template*]: delete template
- [-i *template*]: show information of template