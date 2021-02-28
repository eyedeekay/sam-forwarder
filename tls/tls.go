package i2ptls

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"log"
	"os"
)

func CheckFile(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func Certificate(Cert, Pem string, names []string, priv *ed25519.PrivateKey) (*tls.Certificate, error) {
	template := x509.Certificate{
		//		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"Acme Co"},
		},
		//		NotBefore:             notBefore,
		//		NotAfter:              notAfter,
		//		KeyUsage:              keyUsage,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}
	template.DNSNames = append(template.DNSNames, names...)

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, priv.Public().(ed25519.PublicKey), priv)
	if err != nil {
		log.Fatalf("Failed to create certificate: %v", err)
	}
	certOut, err := os.Create(Cert)
	if err != nil {
		log.Fatalf("Failed to open "+Cert+" for writing: %v", err)
	}
	if err := pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes}); err != nil {
		log.Fatalf("Failed to write data to "+Cert+": %v", err)
	}
	if err := certOut.Close(); err != nil {
		log.Fatalf("Error closing "+Cert+": %v", err)
	}
	log.Print("wrote " + Cert + "\n")

	keyOut, err := os.OpenFile(Pem, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Failed to open "+Pem+" for writing: %v", err)
	}
	privBytes, err := x509.MarshalPKCS8PrivateKey(priv)
	if err != nil {
		log.Fatalf("Unable to marshal private key: %v", err)
	}
	if err := pem.Encode(keyOut, &pem.Block{Type: "PRIVATE KEY", Bytes: privBytes}); err != nil {
		log.Fatalf("Failed to write data to key.pem: %v", err)
	}
	if err := keyOut.Close(); err != nil {
		log.Fatalf("Error closing key.pem: %v", err)
	}
	return tls.LoadX509KeyPair(Cert, Pem)
}

func TLSConfig(Cert, Pem string, names []string) *tls.Config {
	if CheckFile(Cert) && CheckFile(Pem) {
		cert, err := tls.LoadX509KeyPair(Cert, Pem)
		if err != nil {
			log.Fatal(err)
		}
		return &tls.Config{Certificates: []tls.Certificate{cert}}
	} else {
		_, priv, err := ed25519.GenerateKey(rand.Reader)
		if err != nil {
			log.Fatal(err)
		}
		cert, err := Certificate(Cert, Pem, names, priv)
		if err != nil {
			log.Fatal(err)
		}
		return &tls.Config{Certificates: []tls.Certificate{cert}}
	}
}
