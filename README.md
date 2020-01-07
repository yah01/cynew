# cynew

## Build

1. Install [Go](https://golang.org/)
2. Clone the repo
> git clone https://github.com/yah01/cynew.git
3. Change directory to where cynew exists
> cd cynew
4. Run command
> go build
5. There will be a executable file

## Usage

I highly recommend that adding the directory of cynew into environment variable **Path** such that you could run it anywhere.

### simply create a new file named *hello.cpp*
> cynew *hello.cpp*

### create a new file *hello.cpp* with a template *temp*
> cynew *hello.cpp* *temp*

### create more files *file1,file2,...* with the same template *temp*
> cynew *file1* *file2* ... *temp*

### create more files *file1,file2,...* without any template
> cynew -c *file1* *file2* ...

### set the default template to *temp*
> cynew -t *temp*

### set the default suffix to *.cpp*
> cynew -s *cpp*

### list all template(s)
> cynew -ls

## flag

**All flags should be before another parameters**

- [-ls]: list all template(s)
- [-c]: create file(s) without template
- [-t *temp*]: set default template to *temp*
- [-s *suf*]: set default suffix to *.suf*
- [-a *filename*]: add *filename* into templates
- [-i *temp*]: show information of *temp*
- [-o]: open with OS default program