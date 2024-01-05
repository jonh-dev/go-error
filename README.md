<h1 align="center"> Go-Error </h1>

<p align="center">A error handling library for Go applications.</p>

<p align="center">
 <a href="#technologies">Technologies</a> •
 <a href="#running">Running</a> •
 <a href="#author">Author</a>
</p>

![image](https://github.com/jonh-dev/go-error/assets/101439670/3260ccc0-95da-444f-9fca-751f7e0240cd)

##

### Technologies

The following tools were used in building the project:

- Go
- IDE Visual Studio Code

##

### Running

**1.** First, you need to install Go on your system. You can do this by following the instructions at the following link: https://golang.org/dl/

**2.** Choose an IDE of your choice, in this case we will use Visual Studio Code. To download it follow the link: https://code.visualstudio.com/download

**3.** Open your terminal and use `go get` to download and install the `go-error` library. Replace `github.com/jonh-dev/go-error/errors` with the path to your `go-error` library:

```bash
$ go get github.com/jonh-dev/go-error/errors
```

**4.** Now you can import the go-error library into your Go project. Here’s an example of how you can do this:

```go
import (
    "github.com/jonh-dev/go-error/errors"
)
```

**5.** Now you can use the go-error library in your code. Here’s an example of how you can do this:

```go
func main() {
    err := errors.New(codes.Internal, "This is an error message")

    if err != nil {
        status := err.(*errors.Error).GRPCStatus()
        fmt.Println(status.Code(), status.Message())
    }
}
```

**6.** Finally, run the project using the command go run main.go to run the application.

##

### Author

![avatar](https://user-images.githubusercontent.com/101439670/181940218-4f68ffb9-0d35-40df-b8e9-86629333d244.png)

Made by Jonh Dev 🙏

[![LinkedIn Badge](https://img.shields.io/badge/-LINKEDIN-blue?style=flat-square&logo=Linkedin&logoColor=white&link="https://www.linkedin.com/in/jo%C3%A3o-carlos-schwab-zanardi-752591213/)](https://www.linkedin.com/in/jo%C3%A3o-carlos-schwab-zanardi-752591213/)
