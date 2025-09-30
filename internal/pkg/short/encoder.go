package short

import "crypto/rand"

const (
	Alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890_"
	Length   = 10
)

func NewCode() (string, error) {
	b := make([]byte, Length)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	out := make([]byte, Length)
	for i := 0; i < Length; i++ {
		out[i] = Alphabet[int(b[i])%len(Alphabet)]
	}
	return string(out), nil
}
