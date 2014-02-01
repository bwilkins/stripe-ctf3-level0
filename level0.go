package main

import (
  "os"
  "fmt"
  "log"
  "io"
  "bufio"
  "io/ioutil"
  "strings"
  "bytes"
  "sort"
  "math"
)

type FastSearch []string

func NewFastSearch(filename string) FastSearch {

  file, err := os.Open(filename)
  defer file.Close()
  if err != nil {
    log.Fatal("Unable to open dictionary file for reading")
  }
  contents, err := ioutil.ReadAll(file)
  if err != nil {
    log.Fatal("Unable to read the contents of the dictionary file")
  }

  contentsStr := string(contents)

  words := sort.StringSlice(strings.Split(contentsStr, "\n"))
  words.Sort()

  return FastSearch(words)
}

func (fs *FastSearch) Contains(needle []byte) bool {
    var start_point, mid_point, end_point int
    start_point = 0
    end_point = len(*fs)-1
    mid_point = start_point + int(math.Trunc(float64((end_point - start_point)/2)))

    for start_point != mid_point && end_point != mid_point {
      switch bytes.Compare(needle, []byte((*fs)[mid_point])) {
        case -1:
          end_point = mid_point
        case 0:
          return true
        case 1:
          start_point = mid_point
      }

      mid_point = start_point + int(math.Trunc(float64((end_point - start_point)/2)))

      if bytes.Equal([]byte((*fs)[start_point]), needle) ||
         bytes.Equal([]byte((*fs)[mid_point]), needle) ||
         bytes.Equal([]byte((*fs)[end_point]), needle) {
        return true
      }
    }
    return false

}

type Sipper struct {
  reader bufio.Reader
  eof bool
}

func NewSipper(reader io.Reader) (Sipper) {
  return Sipper{*bufio.NewReader(reader), false}
}

func (s *Sipper) SipWord() (word []byte, err error) {
  var char byte

  for {
    char, err = (*s).reader.ReadByte()
    if err != nil {
      return
    }

    if char != '\n' && char != ' ' {
      word = append(word, char)
    } else {
      (*s).reader.UnreadByte()
      break
    }
  }

  return
}

func (s *Sipper) SipWhitespace() (space []byte, err error) {
  var char byte

  for {
    char, err = (*s).reader.ReadByte()
    if err != nil {
      return
    }
    if char == '\n' || char == ' ' {
      space = append(space, char)
    } else {
      (*s).reader.UnreadByte()
      break
    }
  }

  return
}

func main() {
  var path string
  var input Sipper
  var dictionary FastSearch
  var word, space []byte
  var err error
  if len(os.Args) > 1 {
    path = os.Args[1]
  } else {
    path = "/usr/share/dict/words"
  }

  dictionary = NewFastSearch(path)
  input = NewSipper(os.Stdin)

  for {
    word, err = input.SipWord()
    if err != nil {
      return
    }

    if dictionary.Contains(bytes.ToLower(word)) {
      fmt.Printf("%s", word)
    } else {
      fmt.Printf("<%s>", word)
    }

    space, err = input.SipWhitespace()

    fmt.Printf("%s", space)
  }

  return
}