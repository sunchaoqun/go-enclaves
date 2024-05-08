docker build -t myapp-image .
nitro-cli terminate-enclave --all
nitro-cli build-enclave --docker-uri myapp-image:latest --output-file myapp-image.eif
nitro-cli run-enclave --eif-path myapp-image.eif --memory 1024 --cpu-count 2 --enclave-cid 5 --debug-mode
nitro-cli console --enclave-id $(nitro-cli describe-enclaves | jq -r '.[0].EnclaveID')
