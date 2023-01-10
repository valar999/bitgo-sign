package sign

import (
	"bytes"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common/math"
	"golang.org/x/crypto/sha3"
	"math/big"
	"strings"
	"time"
)

func SignTx(address string, amountStr string, expire time.Time, seqId uint) ([]byte, error) {
	buf := new(bytes.Buffer)
	buf.WriteString("ETHER")

	addr, err := hex.DecodeString(strings.TrimPrefix(address, "0x"))
	if err != nil {
		return nil, err
	}
	buf.Write(addr)

	amount := new(big.Int)
	_, ok := amount.SetString(amountStr, 10)
	if !ok {
		return nil, err
	}
	buf.Write(math.U256Bytes(amount))

	buf.Write(math.U256Bytes(big.NewInt(expire.Unix())))
	buf.Write(math.U256Bytes(big.NewInt(int64(seqId))))

	h := sha3.NewLegacyKeccak256()
	h.Write(buf.Bytes())
	return h.Sum(nil), nil
}
