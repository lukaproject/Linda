# /bin/bash

cert_output="$root_dir/output/certs/"
cert_path="$test_script_dir/certs"

mkdir -p $cert_output

openssl req -newkey rsa:4096 -nodes -keyout $cert_output/rsa_private.key -x509 -days 3650 -out $cert_output/cert.crt
