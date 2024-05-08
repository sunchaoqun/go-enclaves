package main

import (
    "bufio"
    "fmt"
    "net"
    "strconv"
    "strings"
    // "os/exec"
    // "encoding/hex"
    // "encoding/base64"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/kms"
)

type Credential struct {
    Region          string
    AccessKeyId     string
    SecretAccessKey string
    SessionToken    string
}

func main() {
    credential := Credential{
        Region:          "ap-southeast-1",
        AccessKeyId:     "",
        SecretAccessKey: "",
        SessionToken:    "",
    }

    // Set up AWS session
    // sess, err := session.NewSession(&aws.Config{
    //     Region: aws.String(credential.Region),
    // })

	// 创建一个含有临时凭证的 AWS 会话
	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(credential.AccessKeyId, credential.SecretAccessKey, credential.SessionToken),
		Region:      aws.String(credential.Region),
	})

    if err != nil {
        fmt.Println("Error creating AWS session:", err)
        return
    }

    // Create KMS client
    svc := kms.New(sess)

    // Listen on TCP port 9000
    ln, err := net.Listen("tcp", ":9000")
    if err != nil {
        fmt.Println("Error setting up TCP listener:", err)
        return
    }
    defer ln.Close()

    fmt.Println("Listening on port 9000...")
    for {
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println("Error accepting connection:", err)
            continue
        }
        go handleConnection(conn, credential, svc)
    }
}

func handleConnection(conn net.Conn, credential Credential, svc *kms.KMS) {
    defer conn.Close()
    reader := bufio.NewReader(conn)

    for{
        // Read data from the connection
        data, err := reader.ReadBytes('\n')
        if err != nil {
            fmt.Println("Error reading data:", err)
            return
        }

        // Encrypt data
        result, err := svc.Encrypt(&kms.EncryptInput{
            KeyId:     aws.String("alias/v4"),
            Plaintext: data,
        })
    
        // 构建命令行参数
        // cmd := exec.Command("kmstool_enclave_cli",
        //     "encrypt",
        //     "--region", credential.Region,
        //     "--proxy-port", "8000",
        //     "--aws-access-key-id", credential.AccessKeyId,
        //     "--aws-secret-access-key", credential.SecretAccessKey,
        //     "--key-id", "alias/v4",
        //     "--aws-session-token", credential.SessionToken,
        //     "--plaintext", string(data),
        // )
    
        // 执行命令
        // encryptedOutput, err := cmd.CombinedOutput()
        encryptedOutput := result.CiphertextBlob

        if err != nil {
            fmt.Println("Encryption error:", err)
            return
        }

        fmt.Println("Encrypted data:", encryptedOutput)

        // if err != nil {
        //     fmt.Println("Encryption failed:", err)
        //     return
        // }
        // fmt.Println("Encrypted data:", result.CiphertextBlob)

        // Send encrypted data back to client
        // _, err = conn.Write(encryptedOutput)

        if err != nil {
            fmt.Println("Error sending encrypted data:", err)
        }

        // Convert encrypted data to hexadecimal
        // encryptedHex := hex.EncodeToString(result.CiphertextBlob)

        // Send hexadecimal string back to client
        // _, err = conn.Write([]byte(encryptedHex + "\n"))

        // Encode CiphertextBlob to Base64
        // encryptedBase64 := base64.StdEncoding.EncodeToString(result.CiphertextBlob)

        // // Send Base64-encoded string back to client
        // _, err = conn.Write([]byte(encryptedBase64 + "\n"))
        
        // if err != nil {
        //     fmt.Println("Error sending encrypted data:", err)
        // }

        // formattedOutput := formatByteArray(result.CiphertextBlob)
        formattedOutput := formatByteArray(encryptedOutput) 

        // // Send formatted byte array string back to client
        _, err = conn.Write([]byte(formattedOutput + "\n"))
        if err != nil {
            fmt.Println("Error sending formatted data:", err)
        }

        // Decrypt data
        decryptResult, err := svc.Decrypt(&kms.DecryptInput{
            CiphertextBlob: encryptedOutput,
        })
        if err != nil {
            fmt.Println("Decryption failed:", err)
            return
        }
        fmt.Println("Decrypted data:", string(decryptResult.Plaintext))
    }
}

// Helper function to format byte array
func formatByteArray(b []byte) string {
    var sb strings.Builder
    sb.WriteString("[")
    for i, byteVal := range b {
        sb.WriteString(strconv.Itoa(int(byteVal)))
        if i < len(b)-1 {
            sb.WriteString(" ")
        }
    }
    sb.WriteString("]")
    return sb.String()
}