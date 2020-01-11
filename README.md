# cynew

## Build

1. Install [Go](https://golang.org/)
2. Clone the repo
```shell
go get github.com/yah01/cynew
```
3. Change directory to where cynew exists
```shell
cd $GOPATH/src/github.com/yah01/cynew
```
4. Run command
```shell
go build
```
5. There will be a executable file

## Usage

I highly recommend that adding the directory of cynew into environment variable **Path** such that you could run it anywhere.

### simply create a new file named *hello.cpp*
```shell
cynew hello.cpp
```

### create a new file *hello.cpp* with a template *temp*
```shell
cynew hello.cpp temp
```

### create more files *file1,file2,...* with the same template *temp*
```shell
cynew file1 file2 ... temp
```

### create more files *file1,file2,...* without any template
```shell
cynew -c file1 file2 ...
```

### set the default template to *temp*
```shell
cynew -t temp
```

### set the default suffix to *.cpp*
```shell
cynew -s cpp
```

### list all template(s)
```shell
cynew -ls
```

## flag

- [-ls]: list all template(s)
- [-c]: create file(s) without template
- [-t *temp*]: set default template to *temp*
- [-s *suf*]: set default suffix to *.suf*
- [-a *filename*]: add *filename* into templates
- [-i *temp*]: show information of *temp*
- [-o]: open with OS default program