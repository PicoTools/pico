package tls

import (
	"context"
	"crypto/sha1"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"

	"github.com/PicoTools/pico/internal/constants"
	"github.com/PicoTools/pico/internal/ent"
	"github.com/PicoTools/pico/internal/ent/pki"

	"github.com/go-faster/errors"
)

// NewTlsTransport returns TLS config for using with GRPC servers
func NewTlsTransport(ctx context.Context, db *ent.Client, t pki.Type) (*tls.Config, string, error) {
	// query CA
	ca, err := db.Pki.
		Query().
		Where(pki.TypeEQ(pki.TypeCa)).
		Only(ctx)
	if err != nil {
		return &tls.Config{}, "", errors.Wrap(err, "query CA key/cert from DB")
	}

	// query needing cert
	v, err := db.Pki.
		Query().
		Where(pki.TypeEQ(t)).
		Only(ctx)
	if err != nil {
		return &tls.Config{}, "", errors.Wrapf(err, "query %s key/cert from DB", t.String())
	}

	// create x509 keypair
	caX509, err := tls.X509KeyPair(ca.Cert, ca.Key)
	if err != nil {
		return &tls.Config{}, "", errors.Wrap(err, "create X509 keypair for CA")
	}
	vX509, err := tls.X509KeyPair(v.Cert, v.Key)
	if err != nil {
		return &tls.Config{}, "", errors.Wrapf(err, "create X509 keypair for %s", t.String())
	}

	// create pool with x509 certs
	caPool := x509.NewCertPool()
	caCrt, err := x509.ParseCertificate(caX509.Certificate[0])
	if err != nil {
		return &tls.Config{}, "", errors.Wrap(err, "parse X509 certificate for CA")
	}
	// add CA to pool
	caPool.AddCert(caCrt)

	// calculate fingerprint
	fingerprint := sha1.Sum(vX509.Certificate[0])

	return &tls.Config{
		CipherSuites:           constants.GrpcTlsCiphers,
		MinVersion:             constants.GrpcTlsMinVersion,
		SessionTicketsDisabled: true,
		Certificates:           []tls.Certificate{vX509},
		ClientAuth:             tls.NoClientCert,
		RootCAs:                caPool,
	}, hex.EncodeToString(fingerprint[:]), nil
}
