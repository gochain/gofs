package gofs

import (
	"bytes"
	"math/big"
	"testing"
)

func TestUnpackPinInputs(t *testing.T) {
	cid, gbh := []byte("test"), big.NewInt(10)
	data, err := pinMethod.Inputs.Pack(cid, gbh)
	if err != nil {
		t.Fatal(err)
	}

	pi, err := UnpackPinInputs(data)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(cid, pi.CID) {
		t.Error("cid wrong")
	}
	if pi.GBH.Cmp(gbh) != 0 {
		t.Error("gbh wrong")
	}

	pi, err = UnpackPinInputs(append(pinMethod.Id(), data...))
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(cid, pi.CID) {
		t.Error("cid wrong")
	}
	if pi.GBH.Cmp(gbh) != 0 {
		t.Error("gbh wrong")
	}
}
