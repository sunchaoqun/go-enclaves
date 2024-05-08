#include "encryption.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <aws/nitro_enclaves/kms.h>

void encrypt_data(const char *key_id, const char *plaintext) {
    NeKmsContext ctx;
    ne_kms_initialize_context(&ctx);

    NeKmsEncryptResponse *response;
    int status = ne_kms_encrypt(&ctx, key_id, plaintext, strlen(plaintext), &response);

    if (status == NE_KMS_SUCCESS) {
        printf("Encrypted data: %s\n", response->ciphertextBlob);
        ne_kms_free_encrypt_response(response);
    } else {
        printf("Encryption failed\n");
    }

    ne_kms_free_context(&ctx);
}

