package main

import(
  "bufio"
  "encoding/base64"
  "fmt"
  "os"
  "path"
  "strings"
)

type File struct {
  BasePath string
  Name string
}

func (of *File) createFile(s *bufio.Scanner) {
  location := "Content-Location"
  var line string
  for s.Scan() {
    line = s.Text()
    if strings.Contains(line, location) {
      of.Name = line[strings.LastIndex(line, "/")+1: ]
      print(of.Name)
      break
    }
  }

  // Ignore next empty line.
  if !s.Scan() {
    // Return if scanner arrive at last boundary separator.
    return
  }
  line = s.Text()

  f, err := os.Create(path.Join(of.BasePath, of.Name))
  defer f.Close()
  if err != nil {
    panic(err)
  }

  enc := base64.StdEncoding
  for s.Scan() {
    line = s.Text()
    if len(line) == 0 {
      break
    }
    data, err := enc.DecodeString(line)
    _, err = f.Write(data)
    if err != nil {
      panic(err)
    }
  }
}

func print(s string) {
  fmt.Println(s)
}

func main() {
  var fp *os.File
  var err error

  if len(os.Args) < 2 {
    fp = os.Stdin
  } else {
    fp, err = os.Open(os.Args[1])
    if err != nil {
      panic(err)
    }
    defer fp.Close()
  }

  scanner := bufio.NewScanner(fp)
  var line string
  var boundary string
  for scanner.Scan() {
    line = scanner.Text()
    if strings.Contains(line, "boundary") {
      a := strings.Split(line, "=")
      boundary = strings.Trim(a[1], "\";")
    } else if boundary != "" && strings.Contains(line, boundary) {
      file := File{
        BasePath: "./output",
      }
      file.createFile(scanner)
    }
  }
  if err := scanner.Err(); err != nil {
    panic(err)
  }
}
