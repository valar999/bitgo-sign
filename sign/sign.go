package sign

import (
	"bytes"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
	"math/big"
	"strings"
	"time"
)

// SingTx calculates a hash of some data from transaction
// and signature of this hash
func SignTx(key, address, amountStr string, expire time.Time, seqId uint) (hash, sig []byte, err error) {
	buf := new(bytes.Buffer)
	buf.WriteString("ETHER")

	addr, err := hex.DecodeString(strings.TrimPrefix(address, "0x"))
	if err != nil {
		return
	}
	buf.Write(addr)

	amount := new(big.Int)
	_, ok := amount.SetString(amountStr, 10)
	if !ok {
		return
	}
	buf.Write(math.U256Bytes(amount))

	buf.Write(math.U256Bytes(big.NewInt(expire.Unix())))
	buf.Write(math.U256Bytes(big.NewInt(int64(seqId))))

	h := sha3.NewLegacyKeccak256()
	h.Write(buf.Bytes())
	hash = h.Sum(nil)

	k, err := crypto.HexToECDSA(key)
	if err != nil {
		return
	}
	sig, err = crypto.Sign(hash, k)
	if err != nil {
		return
	}

	return
}
