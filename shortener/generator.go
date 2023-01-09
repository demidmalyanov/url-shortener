package shortener

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"

	"github.com/itchyny/base58-go"
)

func sha256Input(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))

	return algorithm.Sum(nil)
}

func base58Encoded(bytes []byte) string {
	encoding := base58.BitcoinEncoding

	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return string(encoded)
}

func GenerateTokenForUrl(url string) string {
	urlHashBytes := sha256Input(url)
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	token := base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))

	return token[:8] // first 8 letters
}
