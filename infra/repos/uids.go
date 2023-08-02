package repos

import (
	crand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math/rand"

	"github.com/Soreing/growler/domain/general/uids"
)

type UidsRepository struct {
	rng *rand.Rand
}

func NewUidsRepository() (*UidsRepository, error) {
	bytseed := make([]byte, 8)

	n, err := crand.Read(bytseed)
	if err != nil {
		return nil, err
	} else if n != len(bytseed) {
		return nil, fmt.Errorf("failed to generate seed")
	}

	seed := binary.BigEndian.Uint64(bytseed)
	src := rand.NewSource(int64(seed)) // This is scuffed

	return &UidsRepository{
		rng: rand.New(src),
	}, nil
}

var _ uids.IRepository = (*UidsRepository)(nil)

func (r *UidsRepository) GetHexString(
	digits int,
) string {
	charset := "0123456789abcdef"
	bytes := make([]byte, digits)
	data := int64(1)

	for i := 0; i < digits; i++ {
		if data < 16 {
			data = r.rng.Int63()
		}
		bytes[i] = charset[data&0xF]
		data >>= 4
	}

	return string(bytes)
}
