package main

// #cgo LDFLAGS: -lasound
// #include <alsa/asoundlib.h>
import "C"
import "fmt"

func main() {
  var seq_handle *C.snd_seq_t

  if (C.snd_seq_open(&seq_handle, C.CString("hw"), C.SND_SEQ_OPEN_DUPLEX, 0) < 0) {
    fmt.Println("Error opening alsa lib")
  }
  fmt.Println("The ALSA libraries are installed.")
}
