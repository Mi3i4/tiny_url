package shortener

type Service struct {
	cfg Config
}

func New(opts ...Option) *Service {
	cfg := defaultConfig()

	for _, opt := range opts {
		opt(&cfg)
	}

	return &Service{cfg: cfg}
}

func (s *Service) CodeLength() int {
	return s.cfg.codeLength
}

func (s *Service) Validate(url string) error {
	return s.cfg.validator(url)
}
