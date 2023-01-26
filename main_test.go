package gokzgbenchmarks

import (
	"fmt"
	"math/rand"
	"testing"

	gokzg "github.com/protolambda/go-kzg/eth"
	"github.com/stretchr/testify/require"
)

const (
	BytesPerFieldElement = 32
	BytesPerBlob         = 4096 * BytesPerFieldElement
)

///////////////////////////////////////////////////////////////////////////////
// Types
///////////////////////////////////////////////////////////////////////////////

type GoKzgBlobImpl []byte

func (b GoKzgBlobImpl) Len() int {
	return 4096
}

func (b GoKzgBlobImpl) At(index int) [32]byte {
	var blob [32]byte
	copy(blob[:], b[index*32:(index+1)*32])
	return blob
}

type GoKzgBlobSequenceImpl []GoKzgBlobImpl

func (b GoKzgBlobSequenceImpl) Len() int {
	return len(b)
}

func (b GoKzgBlobSequenceImpl) At(index int) gokzg.Blob {
	return b[index]
}

///////////////////////////////////////////////////////////////////////////////
// Benchmarks
///////////////////////////////////////////////////////////////////////////////

func GetRandFieldElement(seed int64) [32]byte {
	rand.Seed(seed)
	bytes := make([]byte, 31)
	_, err := rand.Read(bytes)
	if err != nil {
		panic("failed to get random field element")
	}

	var fieldElementBytes [32]byte
	copy(fieldElementBytes[:], bytes)
	return fieldElementBytes
}

func GetRandBlob(seed int64) GoKzgBlobImpl {
	var blob [BytesPerBlob]byte
	for i := 0; i < BytesPerBlob; i += BytesPerFieldElement {
		fieldElementBytes := GetRandFieldElement(seed + int64(i))
		copy(blob[i:i+BytesPerFieldElement], fieldElementBytes[:])
	}
	return blob[:]
}

func Benchmark(b *testing.B) {
	const length = 64
	blobs := GoKzgBlobSequenceImpl{}
	commitments := gokzg.KZGCommitmentSequenceImpl{}
	for i := 0; i < length; i++ {
		blobs = append(blobs, GetRandBlob(int64(i)))
		commitment, _ := gokzg.BlobToKZGCommitment(blobs.At(i))
		commitments = append(commitments, commitment)
	}
	z := [32]byte{1, 2, 3}
	y := [32]byte{4, 5, 6}
	proof, _ := gokzg.ComputeAggregateKZGProof(blobs[:1])

	b.Run("BlobToKZGCommitment", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_, ret := gokzg.BlobToKZGCommitment(blobs[0])
			require.True(b, ret)
		}
	})

	b.Run("VerifyKZGProof", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_, err := gokzg.VerifyKZGProof(commitments[0], z, y, proof)
			require.NoError(b, err)
		}
	})

	for i := 1; i <= len(blobs); i *= 2 {
		b.Run(fmt.Sprintf("ComputeAggregateKZGProof(blobs=%v)", i), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				_, err := gokzg.ComputeAggregateKZGProof(blobs[:i])
				require.NoError(b, err)
			}
		})
	}

	for i := 1; i <= len(blobs); i *= 2 {
		b.Run(fmt.Sprintf("VerifyAggregateKZGProof(blobs=%v)", i), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				_, err := gokzg.VerifyAggregateKZGProof(blobs[:i], commitments[:i], proof)
				require.NoError(b, err)
			}
		})
	}
}
