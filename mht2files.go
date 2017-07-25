package main

import(
  "bufio"
  "fmt"
  "os"
  "strings"
)

func print(s string) {
  fmt.Println(s)
}

func createFile(s *bufio.Scanner) {
  line := s.Text()
  print(line)
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
      print("boundary:" + boundary)

    } else if boundary != "" && strings.Contains(line, boundary) {
      createFile(scanner)
    }
    //print(line)
  }
  if err := scanner.Err(); err != nil {
    panic(err)
  }

}
