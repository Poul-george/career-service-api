package identifier

type ExternalUserID string

func (e ExternalUserID) String() string {
	return string(e)
}
