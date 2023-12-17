package identifier

type UserID uint64

func (u UserID) Uint64() uint64 {
	return uint64(u)
}
