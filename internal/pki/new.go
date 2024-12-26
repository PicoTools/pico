package pki

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"time"

	"github.com/PicoTools/pico/internal/constants"
	"github.com/PicoTools/pico/internal/ent"
	"github.com/PicoTools/pico/internal/ent/pki"
	"github.com/PicoTools/pico/internal/utils"

	"github.com/go-faster/errors"
)

// newCa generate CA keypair and saves to database
func (cfg *CAConfig) newCa(ctx context.Context, tx *ent.Tx) (*ent.Pki, error) {
	// generate public/private
	privateKey, err := ecdsa.GenerateKey(utils.ChooseEC(constants.ECLength), rand.Reader)
	if err != nil {
		return nil, errors.Wrap(err, "generate CA ECDSA private key")
	}
	publicKey := privateKey.Public()
	// create certificate
	certTpl := &x509.Certificate{
		SerialNumber: big.NewInt(cfg.Serial),
		Subject: pkix.Name{
			OrganizationalUnit: []string{cfg.Subject.OrganizationalUnit},
			Organization:       []string{cfg.Subject.Organization},
			Country:            []string{cfg.Subject.Country},
			CommonName:         cfg.Subject.CommonName,
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().AddDate(constants.CertValidTime, 0, 0),
		KeyUsage:  x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage: []x509.ExtKeyUsage{
			x509.ExtKeyUsageClientAuth,
			x509.ExtKeyUsageServerAuth,
		},
		BasicConstraintsValid: true,
	}
	// self-signed CA certificate
	cert, err := x509.CreateCertificate(rand.Reader, certTpl, certTpl, publicKey, privateKey)
	if err != nil {
		return nil, errors.Wrap(err, "generate CA certificate")
	}
	// format certificate in PEM
	certPem := new(bytes.Buffer)
	if err = pem.Encode(certPem, &pem.Block{Type: "CERTIFICATE", Bytes: cert}); err != nil {
		return nil, errors.Wrap(err, "create CA PEM certificate")
	}
	// format key in PEM
	key, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		return nil, errors.Wrap(err, "marshal CA ECDSA private key")
	}
	keyPem := new(bytes.Buffer)
	if err = pem.Encode(keyPem, &pem.Block{Type: "EC PRIVATE KEY", Bytes: key}); err != nil {
		return nil, errors.Wrap(err, "create CA PEM key")
	}
	return tx.Pki.
		Create().
		SetType(pki.TypeCa).
		SetCert(certPem.Bytes()).
		SetKey(keyPem.Bytes()).
		Save(ctx)
}

// newCert generate listener keypair and saves to database
func (cfg *ListenerConfig) newCert(ctx context.Context, catls tls.Certificate, tx *ent.Tx, ip net.IP) (*ent.Pki, error) {
	// generate public/private
	privateKey, err := ecdsa.GenerateKey(utils.ChooseEC(constants.ECLength), rand.Reader)
	if err != nil {
		return nil, errors.Wrap(err, "generate listener ECDSA private key")
	}
	publicKey := privateKey.Public()
	// create certificate
	certTpl := &x509.Certificate{
		SerialNumber: big.NewInt(cfg.Serial),
		Subject: pkix.Name{
			Organization: []string{cfg.Subject.Organization},
			Country:      []string{cfg.Subject.Country},
			Province:     []string{cfg.Subject.Province},
			Locality:     []string{cfg.Subject.Locality},
			CommonName:   cfg.Subject.CommonName,
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().AddDate(constants.CertValidTime, 0, 0),
		KeyUsage:  x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{
			x509.ExtKeyUsageClientAuth,
			x509.ExtKeyUsageServerAuth,
		},
		IPAddresses: []net.IP{ip},
	}
	ca, err := x509.ParseCertificate(catls.Certificate[0])
	if err != nil {
		return nil, errors.Wrap(err, "parse CA certificate")
	}
	// sign certificate
	cert, err := x509.CreateCertificate(rand.Reader, certTpl, ca, publicKey, catls.PrivateKey)
	if err != nil {
		return nil, errors.Wrap(err, "create listener certificate")
	}
	// format cert in PEM
	certPem := new(bytes.Buffer)
	if err = pem.Encode(certPem, &pem.Block{Type: "CERTIFICATE", Bytes: cert}); err != nil {
		return nil, err
	}
	// format key in PEM
	key, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		return nil, errors.Wrap(err, "marshal listener ECDSA private key")
	}
	keyPem := new(bytes.Buffer)
	if err = pem.Encode(keyPem, &pem.Block{Type: "EC PRIVATE KEY", Bytes: key}); err != nil {
		return nil, errors.Wrap(err, "create listener PEM key")
	}
	return tx.Pki.
		Create().
		SetType(pki.TypeListener).
		SetCert(certPem.Bytes()).
		SetKey(keyPem.Bytes()).
		Save(ctx)
}

// newCert generate operator keypair and saves to database
func (cfg *OperatorConfig) newCert(ctx context.Context, catls tls.Certificate, tx *ent.Tx, ip net.IP) (*ent.Pki, error) {
	// generate public/private
	privateKey, err := ecdsa.GenerateKey(utils.ChooseEC(constants.ECLength), rand.Reader)
	if err != nil {
		return nil, errors.Wrap(err, "generate listener ECDSA private key")
	}
	publicKey := privateKey.Public()
	// create certificate
	certTpl := &x509.Certificate{
		SerialNumber: big.NewInt(cfg.Serial),
		Subject: pkix.Name{
			Organization: []string{cfg.Subject.Organization},
			Country:      []string{cfg.Subject.Country},
			Province:     []string{cfg.Subject.Province},
			Locality:     []string{cfg.Subject.Locality},
			CommonName:   cfg.Subject.CommonName,
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().AddDate(constants.CertValidTime, 0, 0),
		KeyUsage:  x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{
			x509.ExtKeyUsageClientAuth,
			x509.ExtKeyUsageServerAuth,
		},
		IPAddresses: []net.IP{ip},
	}
	ca, err := x509.ParseCertificate(catls.Certificate[0])
	if err != nil {
		return nil, errors.Wrap(err, "parse CA certificate")
	}
	// sign certificate
	cert, err := x509.CreateCertificate(rand.Reader, certTpl, ca, publicKey, catls.PrivateKey)
	if err != nil {
		return nil, errors.Wrap(err, "create listener certificate")
	}
	// format cert in PEM
	certPem := new(bytes.Buffer)
	if err = pem.Encode(certPem, &pem.Block{Type: "CERTIFICATE", Bytes: cert}); err != nil {
		return nil, err
	}
	// format key in PEM
	key, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		return nil, errors.Wrap(err, "marshal listener ECDSA private key")
	}
	keyPem := new(bytes.Buffer)
	if err = pem.Encode(keyPem, &pem.Block{Type: "EC PRIVATE KEY", Bytes: key}); err != nil {
		return nil, errors.Wrap(err, "create listener PEM key")
	}
	return tx.Pki.
		Create().
		SetType(pki.TypeOperator).
		SetCert(certPem.Bytes()).
		SetKey(keyPem.Bytes()).
		Save(ctx)
}

// newCert generate management keypair and saves to database
func (cfg *ManagementConfig) newCert(ctx context.Context, catls tls.Certificate, tx *ent.Tx, ip net.IP) (*ent.Pki, error) {
	// generate public/private
	privateKey, err := ecdsa.GenerateKey(utils.ChooseEC(constants.ECLength), rand.Reader)
	if err != nil {
		return nil, errors.Wrap(err, "generate management ECDSA private key")
	}
	publicKey := privateKey.Public()
	// create certificate
	certTpl := &x509.Certificate{
		SerialNumber: big.NewInt(cfg.Serial),
		Subject: pkix.Name{
			Organization: []string{cfg.Subject.Organization},
			Country:      []string{cfg.Subject.Country},
			Province:     []string{cfg.Subject.Province},
			Locality:     []string{cfg.Subject.Locality},
			CommonName:   cfg.Subject.CommonName,
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().AddDate(constants.CertValidTime, 0, 0),
		KeyUsage:  x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{
			x509.ExtKeyUsageClientAuth,
			x509.ExtKeyUsageServerAuth,
		},
		IPAddresses: []net.IP{ip},
	}
	ca, err := x509.ParseCertificate(catls.Certificate[0])
	if err != nil {
		return nil, errors.Wrap(err, "parse CA certificate")
	}
	// sign certificate
	cert, err := x509.CreateCertificate(rand.Reader, certTpl, ca, publicKey, catls.PrivateKey)
	if err != nil {
		return nil, errors.Wrap(err, "create management certificate")
	}
	// format cert in PEM
	certPem := new(bytes.Buffer)
	if err = pem.Encode(certPem, &pem.Block{Type: "CERTIFICATE", Bytes: cert}); err != nil {
		return nil, err
	}
	// format key in PEM
	key, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		return nil, errors.Wrap(err, "marshal management ECDSA private key")
	}
	keyPem := new(bytes.Buffer)
	if err = pem.Encode(keyPem, &pem.Block{Type: "EC PRIVATE KEY", Bytes: key}); err != nil {
		return nil, errors.Wrap(err, "create management PEM key")
	}
	return tx.Pki.
		Create().
		SetType(pki.TypeManagement).
		SetCert(certPem.Bytes()).
		SetKey(keyPem.Bytes()).
		Save(ctx)
}
