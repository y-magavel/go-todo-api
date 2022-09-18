package auth

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/y-magavel/go-todo-api/clock"
	"github.com/y-magavel/go-todo-api/entity"
)

//go:embed cert/secret.pem
var rowPrivKey []byte

//go:embed cert/public.pem
var rowPubKey []byte

type JWTer struct {
	PrivateKey, PublicKey jwk.Key
	Store                 Store
	Clocker               clock.Clocker
}

//go:generate go run github.com/matryer/moq -out jwt_moq_test.go . Store
type Store interface {
	Save(ctx context.Context, key string, userID entity.UserID) error
	Load(ctx context.Context, key string) (entity.UserID, error)
}

func NewJWTer(s Store) (*JWTer, error) {
	j := &JWTer{Store: s}
	privkey, err := parse(rowPrivKey)
	if err != nil {
		return nil, fmt.Errorf("failed in NewJWTer: private key: %w", err)
	}

	pubkey, err := parse(rowPubKey)
	if err != nil {
		return nil, fmt.Errorf("failed in NewJWTer: public key: %w", err)
	}

	j.PrivateKey = privkey
	j.PublicKey = pubkey
	j.Clocker = clock.RealClocker{}

	return j, nil
}

func parse(rawKey []byte) (jwk.Key, error) {
	key, err := jwk.ParseKey(rawKey, jwk.WithPEM(true))
	if err != nil {
		return nil, err
	}
	return key, nil
}
