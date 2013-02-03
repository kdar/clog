package clog

import (
  "fmt"
  "github.com/wsxiaoys/terminal/color"
  "io"
  "log"
  "os"
  //"sync"
)

var std = log.New(os.Stderr, "", log.LstdFlags)

//var std_mu synx.Mutex

// SetOutput sets the output destination for the standard logger.
func SetOutput(w io.Writer) {
  //std_mu.Lock()
  //defer std_mu.Unlock()
  //std.out = w
  std = log.New(w, "", log.LstdFlags)
}

// Flags returns the output flags for the standard logger.
func Flags() int {
  return std.Flags()
}

// SetFlags sets the output flags for the standard logger.
func SetFlags(flag int) {
  std.SetFlags(flag)
}

// Prefix returns the output prefix for the standard logger.
func Prefix() string {
  return std.Prefix()
}

// SetPrefix sets the output prefix for the standard logger.
func SetPrefix(prefix string) {
  std.SetPrefix(prefix)
}

// These functions write to the standard logger.

// Print calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Print.
func Print(v ...interface{}) {
  std.Output(2, color.Sprint(v...))
}

// Printf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func Printf(format string, v ...interface{}) {
  std.Output(2, color.Sprintf(format, v...))
}

// Println calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Println.
func Println(v ...interface{}) {
  std.Output(2, color.Sprint(fmt.Sprintln(v...)))
}

// Fatal is equivalent to Print() followed by a call to os.Exit(1).
func Fatal(v ...interface{}) {
  std.Output(2, color.Sprint(v...))
  os.Exit(1)
}

// Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
func Fatalf(format string, v ...interface{}) {
  std.Output(2, color.Sprintf(format, v...))
  os.Exit(1)
}

// Fatalln is equivalent to Println() followed by a call to os.Exit(1).
func Fatalln(v ...interface{}) {
  std.Output(2, color.Sprint(fmt.Sprintln(v...)))
  os.Exit(1)
}

// Panic is equivalent to Print() followed by a call to panic().
func Panic(v ...interface{}) {
  s := color.Sprint(v...)
  std.Output(2, s)
  panic(s)
}

// Panicf is equivalent to Printf() followed by a call to panic().
func Panicf(format string, v ...interface{}) {
  s := color.Sprintf(format, v...)
  std.Output(2, s)
  panic(s)
}

// Panicln is equivalent to Println() followed by a call to panic().
func Panicln(v ...interface{}) {
  s := color.Sprint(fmt.Sprintln(v...))
  std.Output(2, s)
  panic(s)
}
