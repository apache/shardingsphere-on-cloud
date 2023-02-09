.PHONY:openssl-local
openssl-local:
	mkdir -p certs && \
  	cd certs  && \
	openssl req -new -SHA256 -newkey rsa:2048 -nodes -keyout tls.key -out tls.csr -subj "/C=CN/ST=beijing/L=beijing/O=/OU=/" && \
	openssl x509 -req -sha256 -days 365 -in tls.csr -signkey tls.key -out tls.crt

