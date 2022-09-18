package auth

import (
	_ "embed"
)

//go:embed cert/secret.pem
var rowPrivKey []byte

//go:embed cert/public.pem
var rowPubKey []byte
