package models

type ServicesConfig func(*Services) error

func WithTezTools() ServicesConfig {
	return func(s *Services) error {
		s.TezTools = newTezTools()
		return nil
	}
}

// NewServices creates a new services provider.
func NewServices(cfgs ...ServicesConfig) (*Services, error) {
	var s Services
	for _, cfg := range cfgs {
		if err := cfg(&s); err != nil {
			return nil, err
		}
	}
	return &s, nil
}

// Services stores all provided services.
type Services struct {
	TezTools TezTools
}
