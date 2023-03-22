// Code generated by github.com/spacemeshos/go-scale/scalegen. DO NOT EDIT.

// nolint
package shared

import (
	"github.com/spacemeshos/go-scale"
)

func (t *MerkleProof) EncodeScale(enc *scale.Encoder) (total int, err error) {
	{
		n, err := scale.EncodeByteSlice(enc, t.Root)
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		n, err := scale.EncodeSliceOfByteSlice(enc, t.ProvenLeaves)
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		n, err := scale.EncodeSliceOfByteSlice(enc, t.ProofNodes)
		if err != nil {
			return total, err
		}
		total += n
	}
	return total, nil
}

func (t *MerkleProof) DecodeScale(dec *scale.Decoder) (total int, err error) {
	{
		field, n, err := scale.DecodeByteSlice(dec)
		if err != nil {
			return total, err
		}
		total += n
		t.Root = field
	}
	{
		field, n, err := scale.DecodeSliceOfByteSlice(dec)
		if err != nil {
			return total, err
		}
		total += n
		t.ProvenLeaves = field
	}
	{
		field, n, err := scale.DecodeSliceOfByteSlice(dec)
		if err != nil {
			return total, err
		}
		total += n
		t.ProofNodes = field
	}
	return total, nil
}
