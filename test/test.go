package main

// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -L. -lencryption
// #include "encryption.h"
import "C"
import "unsafe"

func main() {
    keyID := C.CString("arn:aws:kms:region:account-id:key/key-id")
    plaintext := C.CString("Data to encrypt")
    defer C.free(unsafe.Pointer(keyID))
    defer C.free(unsafe.Pointer(plaintext))

    C.encrypt_data(keyID, plaintext)
}

