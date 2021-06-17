# webUtils
## tls CA self signature
Generate a private key.</br>
openssl genrsa -out client.key 2048 

Generate a certificate signing request (CSR)
openssl req -new -key client.key -out client.csr

Send the CSR to a certificate authority (CA) to request an X.509 certificate
openssl x509 -req -in client.csr -CA clientRootCA.crt -CAkey clientRootCA.key -CAcreateserial -out client.crt

also could see the link below for more info about SAN extensions, which used in golang1.15+:
https://blog.csdn.net/weixin_40280629/article/details/113563351
