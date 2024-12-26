package pki

import (
	"context"
	"crypto/tls"
	"net"

	"github.com/PicoTools/pico/internal/ent"
	"github.com/PicoTools/pico/internal/ent/pki"

	"github.com/go-faster/errors"
)

// Init initialize PKI for GRPC
func (c *Config) Init(ctx context.Context, db *ent.Client, lIP, oIP, mIP net.IP) error {
	var ca *ent.Pki
	var err error

	tx, err := db.Tx(ctx)
	if err != nil {
		return errors.Wrap(err, "unable create tx")
	}

	// ca
	ca, err = tx.Pki.
		Query().
		Where(pki.TypeEQ(pki.TypeCa)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			// generate new CA
			ca, err = c.CA.newCa(ctx, tx)
			if err != nil {
				return errors.Wrap(err, "unable create CA")
			}
			_, err = tx.Pki.Delete().Where(pki.TypeIn(pki.TypeListener, pki.TypeOperator)).Exec(ctx)
			if err != nil {
				return errors.Wrap(err, "unable remove listener and operator key/cert from DB")
			}
		} else {
			return errors.Wrap(err, "query CA's pki")
		}
	}

	caTLS, err := tls.X509KeyPair(ca.Cert, ca.Key)
	if err != nil {
		return errors.Wrap(err, "create CA x509 keypair")
	}

	// listener
	_, err = tx.Pki.
		Query().
		Where(pki.TypeEQ(pki.TypeListener)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			_, err = c.Listener.newCert(ctx, caTLS, tx, lIP)
			if err != nil {
				return errors.Wrap(err, "create listener's key-cert")
			}
		} else {
			return errors.Wrap(err, "query listener's pki")
		}
	}

	// operator
	_, err = tx.Pki.
		Query().
		Where(pki.TypeEQ(pki.TypeOperator)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			_, err = c.Operator.newCert(ctx, caTLS, tx, oIP)
			if err != nil {
				return errors.Wrap(err, "create operator's key-cert")
			}
		} else {
			return errors.Wrap(err, "query operator's pki")
		}
	}

	// management
	_, err = tx.Pki.
		Query().
		Where(pki.TypeEQ(pki.TypeManagement)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			_, err = c.Management.newCert(ctx, caTLS, tx, mIP)
			if err != nil {
				return errors.Wrap(err, "create management key/cert")
			}
		} else {
			return errors.Wrap(err, "query management's key/cert")
		}
	}

	if err = tx.Commit(); err != nil {
		return errors.Wrap(err, "unable commit tx")
	}
	return nil
}
