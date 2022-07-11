# mtls_kata
Openssl PKI Examples with golang app as functional test

# Create CA
openssl req -new -config etc/root-ca.conf -out ca/root-ca.csr -keyout ca/root-ca/private/root-ca.key
openssl ca -selfsign -config etc/root-ca.conf -out ca/root-ca.crt -extensions root_ca_ext

# Create signing CA
openssl req -new -config etc/signing-ca.conf -out ca/signing-ca.csr -keyout ca/signing-ca/private/signing-ca.key

# Sign signingca CSR
openssl ca -config etc/root-ca.conf -in ca/signing-ca.csr -out ca/signing-ca.crt -extensions signing_ca_ext

# Create new Server cert
openssl req -new -config etc/www.bestintheusa.us.conf -out certs/www.bestintheusa.us.csr -keyout certs/www.bestintheusa.us.key
openssl ca -config etc/signing-ca.conf -in certs/www.bestintheusa.us.csr -out certs/www.bestintheusa.us.crt -extensions server_ext

# Create New Signing Cert
openssl req -new -config etc/client.bestintheusa.us -out certs/client.bestintheusa.us.csr -keyout certs/client.bestintheusa.us.key
openssl ca -config etc/signing-ca.conf -in certs/client.bestintheusa.us.csr -out certs/client.bestintheusa.us.crt -extensions server_ext

# Create cert chain
cat root-ca.pem signing-ca.pem > bundle.pem
# Edit /etc/hosts and test
curl --cacert bundle.pem -H'Host: client.bestintheusa.us' https://client.bestintheusa.us:8443/hello
