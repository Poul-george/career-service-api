package create

type Input struct {
	ExternalUserID string
	Name           string
	Password       string
	MailAddress    string
	Comments       string
}

func InputData(ip Input) (*Input, error) {
	// var input Input
	// input = ip

	return &ip, nil
}
