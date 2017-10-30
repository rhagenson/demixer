package ami

import (
	"testing"
)

func TestKCanBeUint16(t *testing.T) {
	k := *new(K)
  _ = (uint16)(k)
}

func TestKMinFollowsBauerEtAl2008(t *testing.T) {
	if minK != K(5) {
		t.Error("Minimum K no longer follows Bauer et al. (2008).")
	}
}

func TestKMaxFollowsBauerEtAl2008(t *testing.T) {
	if maxK != K(512) {
		t.Error("Maximum K no longer follows Bauer et al. (2008).")
	}
}
