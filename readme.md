go install github.com/aws/aws-sdk-go@latest
go install github.com/aws/aws-sdk-go/aws/session@latest
go install github.com/aws/aws-sdk-go/service/kms@latest

----

go mod init go-enclave

go get github.com/aws/aws-sdk-go/aws
go get github.com/aws/aws-sdk-go/aws/session
# go get github.com/aws/aws-sdk-go/service/kms



go run main.go

docker build -t myapp-image .

nitro-cli build-enclave --docker-uri myapp-image:latest --output-file myapp-image.eif

nitro-cli run-enclave --eif-path myapp-image.eif --memory 1024 --cpu-count 2 --enclave-cid 5 --debug-mode

nitro-cli run-enclave --eif-path hello.eif --memory 128 --cpu-count 2 --enclave-cid 5 --debug-mode

                "StringEqualsIgnoreCase": {
                    "kms:RecipientAttestation:PCR3": "a0f2cb3ec6a8c113ee9e82e1c993534979647e41bed8a65863cc1e2f68e67ac0ed2c02a6fb80ec42f77bd192bbb20edc",
                    "kms:RecipientAttestation:PCR2": "9a1ada2f4287f6bc192c372082eee5181ca94fa1dad51d9d06b3475493fbb99170badd2c58f932d0d57d7e88e25199e2",
                    "kms:RecipientAttestation:PCR1": "bcdf05fefccaa8e55bf2c8d6dee9e79bbff31e34bf28a99aa19e6b29c37ee80b214a414b7607236edf26fcb78654e63f",
                    "kms:RecipientAttestation:PCR0": "16987b39366f8a3841cbca4dad24316f6d68cdbb40c63f9c9f1395b684b5c63db1fc196799902bcf9173c743b39f0c6e",
                    "kms:RecipientAttestation:PCR4": "17cde91d1a64b502c756b229c40cb8cbb985bc43dd3a1c3ee4afc742bc5435ad85137e912bbb95af1f65e4374a4ade21"
                }

,
            "Condition": {
                "StringEqualsIgnoreCase": {
                    "kms:RecipientAttestation:PCR4": ""
                }
            }

KMS_KEY_ARN="arn:aws:kms:ap-southeast-1:932423224465:key/4257a12c-9129-4466-9801-96c7914a264c"

MESSAGE="Hello, KMS\!"
CIPHERTEXT=$(aws kms encrypt --key-id --region ap-southeast-1 "$KMS_KEY_ARN" --plaintext "$MESSAGE" --query CiphertextBlob --output text)
echo $CIPHERTEXT

CMK_REGION=ap-southeast-1 # Must match above
ENCLAVE_CID=$(nitro-cli describe-enclaves | jq -r .[0].EnclaveCID)
# Run docker with network host to allow it to fetch IAM credentials with IMDSv2
docker run --network host --security-opt seccomp=unconfined -it kmstool-instance \
    /kmstool_instance --cid "$ENCLAVE_CID" --region "$CMK_REGION" "$CIPHERTEXT"


    INSTANCE_ID="i-038607602f0dcdcf3"
    python -c "import hashlib, sys; \
    h=hashlib.sha384(); h.update('\0'*48); \
    h.update(\"$INSTANCE_ID\".encode('utf-8')); \
    print(h.hexdigest())"

    ROLE_ARN="arn:aws:iam::932423224465:role/PVRE-SSMOnboardingRole-N3BHvdwo3rWc"
    python -c "import hashlib, sys; \
    h=hashlib.sha384(); h.update('\0'*48); \
    h.update(\"$ROLE_ARN\".encode('utf-8')); \
    print(h.hexdigest())"


    aws sts assume-role --role-arn "arn:aws:iam::932423224465:role/PVRE-SSMOnboardingRole-N3BHvdwo3rWc" --role-session-name AWSCLI-Session


    gcc -o main main.c -I/usr/include -L/path/to/libs -lnitro_enclaves_sdk -lnitro_enclaves_kms


gcc -o main encryption.c -I/usr/include -L/usr/lib -lnitro_enclaves_sdk -lnitro_enclaves_kms



aws sts assume-role --role-arn "arn:aws:iam::932423224465:role/PVRE-SSMOnboardingRole-N3BHvdwo3rWc" --role-session-name AWSCLI-Session

CMK_REGION="ap-southeast-1"

MESSAGE="Hello, KMS\!"
CIPHERTEXT=$(aws kms encrypt --region ap-southeast-1 --key-id "$KMS_KEY_ARN" --plaintext "$MESSAGE" --query CiphertextBlob --output text)
echo $CIPHERTEXT

docker run --network host --security-opt seccomp=unconfined -it kmstool-instance \
    /kmstool_instance --cid "$ENCLAVE_CID" --region "$CMK_REGION" "$CIPHERTEXT"




curl -o /dev/null -s -w "Total: %{time_total}s\n" http://localhost

curl -o /dev/null -s -w "Total: %{time_total}s\n" http://10.0.3.200


curl -o /dev/null -s -w "Total: %{time_total}s\n" http://10.0.3.109

curl -o /dev/null -s -w "Total: %{time_total}s\n" http://10.0.3.81


