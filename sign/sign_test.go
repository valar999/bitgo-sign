package sign

import (
	"encoding/hex"
	"testing"
	"time"
)

func TestSignTx(t *testing.T) {
	h, err := SignTx(
		"0x3633abC37C15142F438BC508202E70ad36AA5fF0",
		"100000000000000",
		time.Unix(1673969264, 0),
		50,
	)
	if err != nil {
		t.Error(err)
	}

	if hex.EncodeToString(h) != "0831b36b5c6aea6f48aaef0445cdcba2774340f874c34779d566ec871ddc3066" {
		t.Error("keccak256 hash does not match")
	}
}
