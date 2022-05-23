package pk

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"strings"
)

func GenPK(prefix string) (string, error) {
	if len(prefix) > 5 {
		return "", errors.New("max prefix 3 chrac")
	}

	ctx, cancel := context.WithCancel(context.Background())
	resC := make(chan string, 1)

	for i := 0; i < 16; i++ {
		i := i
		go func(ctx context.Context) {
			cnt := 0
			for {
				select {
				case <-ctx.Done():
					return
				default:
					privateKey, err := crypto.GenerateKey()
					if err != nil {
						log.Fatal(err)
					}
					publicKey := privateKey.Public()
					publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
					address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

					cnt++
					if cnt%10 == 0 && i == 0 {
						fmt.Println("calculate", cnt*16, "times")
					}

					if strings.HasPrefix(address, prefix) {
						privateKeyBytes := crypto.FromECDSA(privateKey)
						pkstr := hexutil.Encode(privateKeyBytes)[2:]
						fmt.Println("found", address, pkstr)
						resC <- pkstr
						return
					}
				}
			}
		}(ctx)
	}

	for {
		select {
		case res := <-resC:
			cancel()
			return res, nil
		}
	}
}
