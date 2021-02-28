package i2ptls

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"log"
	"math/big"
	"os"
)

func CheckFile(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func CertPemificate(CertPem, KeyPem string, names []string, priv ed25519.PrivateKey) (tls.Certificate, error) {

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		log.Fatalf("Failed to generate serial number: %v", err)
	}
	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"Acme Co"},
		},
		//		NotBefore:             notBefore,
		//		NotAfter:              notAfter,
		//		KeyUsage:              keyUsage,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}
	if len(names) > 0 {
		template.DNSNames = append(template.DNSNames, names...)
	}
	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, priv.Public().(ed25519.PublicKey), priv)
	if err != nil {
		log.Fatalf("Failed to create certificate: %v", err)
	}
	certOut, err := os.Create(CertPem)
	if err != nil {
		log.Fatalf("Failed to open "+CertPem+" for writing: %v", err)
	}
	if err := pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes}); err != nil {
		log.Fatalf("Failed to write data to "+CertPem+": %v", err)
	}
	if err := certOut.Close(); err != nil {
		log.Fatalf("Error closing "+CertPem+": %v", err)
	}
	log.Print("wrote " + CertPem + "\n")

	keyOut, err := os.OpenFile(KeyPem, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Failed to open "+KeyPem+" for writing: %v", err)
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
	return tls.LoadX509KeyPair(CertPem, KeyPem)
}

func TLSConfig(CertPem, KeyPem string, names []string) (*tls.Config, error) {
	var ServerName string
	if len(names) > 0 {
		ServerName = names[0]
	}

	if CheckFile(CertPem) && CheckFile(KeyPem) {
		cert, err := tls.LoadX509KeyPair(CertPem, KeyPem)
		if err != nil {
			return nil, err
		}

		return &tls.Config{Certificates: []tls.Certificate{cert}, ServerName: ServerName}, nil
	} else {
		_, priv, err := ed25519.GenerateKey(rand.Reader)
		if err != nil {
			return nil, err
		}
		cert, err := CertPemificate(CertPem, KeyPem, names, priv)
		if err != nil {
			return nil, err
		}
		return &tls.Config{Certificates: []tls.Certificate{cert}, ServerName: ServerName}, nil
	}
}
