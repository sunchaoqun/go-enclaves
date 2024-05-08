package main

import (
    "fmt"
    "io/ioutil"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/kms"
    "github.com/mdlayher/vsock"
)

func main() {
    // Create a new AWS session
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String("ap-southeast-1"),
    })
    if err != nil {
        fmt.Println("Session creation failed:", err)
        return
    }

    // Create a KMS client
    svc := kms.New(sess)

    // Initialize a vsock connection
    conn, err := vsock.Dial(3, 8000, nil) // Adjust the CID and port as needed
    if err != nil {
        fmt.Println("Failed to dial via vsock:", err)
        return
    }
    defer conn.Close()

    // Read data from vsock connection
    data, err := ioutil.ReadAll(conn)
    if err != nil {
        fmt.Println("Failed to read from vsock:", err)
        return
    }
    fmt.Println("Data received via vsock:", string(data))

    // Encrypt data
    result, err := svc.Encrypt(&kms.EncryptInput{
        KeyId:     aws.String("alias/v4"), // Specify your KMS key
        Plaintext: data,
    })
    if err != nil {
        fmt.Println("Encryption failed:", err)
        return
    }
    fmt.Println("Encrypted data:", result.CiphertextBlob)

    // Decrypt data
    decryptResult, err := svc.Decrypt(&kms.DecryptInput{
        CiphertextBlob: result.CiphertextBlob,
    })
    if err != nil {
        fmt.Println("Decryption failed:", err)
        return
    }
    fmt.Println("Decrypted data:", string(decryptResult.Plaintext))
}
