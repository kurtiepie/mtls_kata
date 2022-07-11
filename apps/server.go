package main

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type EnvConfig struct {
	Cert   string
	Key    string
	Bundle string
}

func GetEnv() EnvConfig {
	cert := os.Getenv("CERT_PATH")
	key := os.Getenv("KEY_PATH")
	bundle := os.Getenv("BUNDLE_PATH")

	if cert == "" {
		log.Fatal("Error Environ CERT_PATH missing")
	}
	if key == "" {
		log.Fatal("Error Environ KEY_PATH missing")
	}
	if bundle == "" {
		log.Fatal("Error Environ BUNDLE_PATH missing")
	}

	return EnvConfig{
		Cert:   cert,
		Key:    key,
		Bundle: bundle,
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello, world\n")
}
func main() {
	http.HandleFunc("/hello", helloHandler)
	caCert, err := ioutil.ReadFile("../bundle.pem")
	if err != nil {
		log.Fatal(err)
	}
	// Create Cert Pool to pass to tls.Config as ClientCAs
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Create TLS Config with the ca pool and enable Client cert valdation
	tlsConfig := &tls.Config{
		ClientCAs:  caCertPool,
		ClientAuth: tls.RequireAndVerifyClientCert,
	}
	tlsConfig.BuildNameToCertificate()

	// Create a server to listen on 8443 with the TLS config
	server := &http.Server{
		Addr:      ":8443",
		TLSConfig: tlsConfig,
	}
	log.Fatal(server.ListenAndServeTLS("../certs/www.bestintheusa.us.crt", "../certs/www.bestintheusa.us.key"))

}
