package main

import (
    "fmt"
    "io"
    "log"
    "github.com/mdlayher/vsock"
)

func main() {
    // Define the CID and port. Adjust these values according to your enclave's configuration.
    const enclaveCID uint32 = 5 
    const port uint32 = 9090

    // Dial the enclave's VSOCK address
    conn, err := vsock.Dial(enclaveCID, port, nil)
    if err != nil {
        log.Fatalf("failed to dial: %v", err)
    }
    defer func() {
        // Close the connection and log if there is an error
        if err := conn.Close(); err != nil {
            log.Printf("failed to close connection: %v", err)
        }
    }()

    // Send a message to the enclave
    _, err = conn.Write([]byte("hello\n"))
    if err != nil {
        log.Fatalf("failed to write: %v", err)
    }

    // Optionally, wait for a response if expected
    buf := make([]byte, 1024) // Adjust buffer size according to expected response size
    n, err := conn.Read(buf)
    if err != nil {
        if err != io.EOF {
            log.Fatalf("failed to read: %v", err)
        }
    }
    fmt.Printf("Received response: %s", string(buf[:n]))

    fmt.Println("Message sent to enclave successfully.")
}
