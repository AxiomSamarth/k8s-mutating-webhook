package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"

	endpoints "github.com/AxiomSamarth/k8s-mutating-webhook/pkg/endpoints"
)

type ServerParameters struct {
	port     int    // webhook server port
	certFile string // path to the x509 certificate for https
	keyFile  string // path to the x509 private key matching `CertFile`
}

var (
	parameters ServerParameters
)

func main() {

	flag.IntVar(&parameters.port, "port", 8443, "Webhook server port.")
	flag.StringVar(&parameters.certFile, "tlsCertFile", "/etc/webhook/certs/tls.crt", "File containing the x509 Certificate for HTTPS.")
	flag.StringVar(&parameters.keyFile, "tlsKeyFile", "/etc/webhook/certs/tls.key", "File containing the x509 private key to --tlsCertFile.")
	flag.Parse()

	http.HandleFunc("/", endpoints.HandleRoot)
	http.HandleFunc("/mutate", endpoints.HandleMutate)
	log.Fatal(http.ListenAndServeTLS(":"+strconv.Itoa(parameters.port), parameters.certFile, parameters.keyFile, nil))

}
