package types

import "context"

type Signer interface {
	Sign(ctx context.Context, payload []byte) ([]byte, error)
	PublicKey(ctx context.Context) ([]byte, error)
}
