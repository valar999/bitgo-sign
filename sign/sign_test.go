package sign

import (
	"encoding/hex"
	"testing"
	"time"
)

const key = "e07ba3e9785d2b70b3096c47f83e13327ad58b5afc93694e0ae63b8aabd71b98"

func TestSignTx(t *testing.T) {
	hash, sig, err := SignTx(
		key,
		"0x3633abC37C15142F438BC508202E70ad36AA5fF0",
		"100000000000000",
		time.Unix(1673969264, 0),
		50,
	)
	if err != nil {
		t.Error(err)
	}

	if hex.EncodeToString(hash) != "0831b36b5c6aea6f48aaef0445cdcba2774340f874c34779d566ec871ddc3066" {
		t.Error("wrong hash")
	}
	if hex.EncodeToString(sig) != "bb25efe3c9e63da2720f7a6d428efcf3414170581e227c55188a0eb7b90819c77135c7167f0a903fb7d106afc0219357d717b6023f18f75fbfabaa85302f22ee00" {
		t.Error("wrong signature")
	}
}
