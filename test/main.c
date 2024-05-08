#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "aws/nitro_enclaves/kms.h"

int main() {
    // Initialize the AWS Nitro Enclaves SDK
    ne_initialize();

    // Initialize the KMS context
    NeKmsContext ctx;
    ne_kms_initialize_context(&ctx);

    // Specify the KMS key ID
    const char *key_id = "arn:aws:kms:us-west-2:123456789012:key/abcd1234-a123-456a-a12b-a123b4cd56ef";

    // Specify the plaintext to be encrypted
    const char *plaintext = "Hello, Enclave!";

    // Encrypt the plaintext using KMS
    NeKmsEncryptResponse *encrypt_response;
    int encrypt_status = ne_kms_encrypt(&ctx, key_id, plaintext, strlen(plaintext), &encrypt_response);

    if (encrypt_status == NE_KMS_SUCCESS) {
        printf("Encrypted data: %s\n", encrypt_response->ciphertextBlob);

        // Decrypt the ciphertext using KMS
        NeKmsDecryptResponse *decrypt_response;
        int decrypt_status = ne_kms_decrypt(&ctx, key_id, encrypt_response->ciphertextBlob, encrypt_response->ciphertextBlobSize, &decrypt_response);

        if (decrypt_status == NE_KMS_SUCCESS) {
            printf("Decrypted data: %s\n", decrypt_response->plaintext);
            ne_kms_free_decrypt_response(decrypt_response);
        } else {
            printf("Decryption failed\n");
        }

        ne_kms_free_encrypt_response(encrypt_response);
    } else {
        printf("Encryption failed\n");
    }

    // Cleanup resources
    ne_kms_free_context(&ctx);
    ne_terminate();

    return 0;
}

