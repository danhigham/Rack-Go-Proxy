package main

import (
  "io"
  "net"
  "log"
  "strconv"
  "os"
  "regexp"
)

func reader(c net.Conn) {
  r := io.Reader(c)
  w := io.Writer(c)

  uri_regex, _ := regexp.Compile(`\/(?P<func>[^?\/]+)\?(?P<param1>[^&\/]+)&?(?P<param2>[^&\/]+)?&?(?P<param3>[^&\/]+)?`)

  buf := make([]byte, 1024)

  for {

    n, err := r.Read(buf[:])
    if err != nil {
      return
    }
    
    uri := string(buf[0:n])
    
    if uri_regex.Find([]byte(uri)) != nil {

      log.Print("URI is ", uri)

      // decipher uri using following pattern /[function-name]?[parm-name]=[value]&[parm-name]=[value]
      matches := uri_regex.FindStringSubmatch(uri)

      func_name := matches[1]
      param1 := matches[2]

      log.Print(func_name)
      log.Print(param1)

      switch func_name { 
        case "sayHello": 
          sayHello(w)
        case "sayGoodbye": 
          sayGoodbye(w)
      } 
    }
  }
}

func sayHello(w io.Writer) {
  w.Write([]byte("Hello"))
  w.Write([]byte("\n"))
}

func sayGoodbye(w io.Writer) {
  w.Write([]byte("Goodbye"))
  w.Write([]byte("\n"))
}

func main() {
  fd, _ := strconv.ParseUint(os.Args[1], 8, 64)
  c, err := net.FileConn(os.NewFile(uintptr(fd), "server"))

  if err != nil {
    log.Panic(err)
  }

  defer c.Close()

  reader(c)
}
