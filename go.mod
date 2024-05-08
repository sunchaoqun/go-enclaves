module go-enclave

go 1.18

require (
	github.com/aws/aws-sdk-go v1.52.0
	github.com/mdlayher/vsock v1.2.1
)

require (
	github.com/fxamacker/cbor/v2 v2.2.0 // indirect
	github.com/hf/nsm v0.0.0-20220930140112-cd181bd646b9 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/mdlayher/socket v0.4.1 // indirect
	github.com/x448/float16 v0.8.4 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
)

replace github.com/aws/aws-sdk-go-v2/service/kms => github.com/edgebitio/nitro-enclaves-sdk-go/kms v0.0.0-20221110205443-8a5476ff3cc2
