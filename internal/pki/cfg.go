package pki

type Config struct {
	CA         CAConfig         `json:"ca"`
	Listener   ListenerConfig   `json:"listener"`
	Operator   OperatorConfig   `json:"operator"`
	Management ManagementConfig `json:"management"`
}

type CAConfig struct {
	Serial  int64         `json:"serial" default:"1"`
	Subject SubjectConfig `json:"subject"`
}

type ListenerConfig struct {
	Serial  int64         `json:"serial" default:"1"`
	Subject SubjectConfig `json:"subject"`
}

type OperatorConfig struct {
	Serial  int64         `json:"serial" default:"1"`
	Subject SubjectConfig `json:"subject"`
}

type ManagementConfig struct {
	Serial  int64         `json:"serial" default:"1"`
	Subject SubjectConfig `json:"subject"`
}

type SubjectConfig struct {
	OrganizationalUnit string `json:"ou" default:"Google Inc. CA"`
	Organization       string `json:"o" default:"google.com"`
	CommonName         string `json:"cn" default:"google.com"`
	Country            string `json:"c" default:"USA"`
	Province           string `json:"p" default:"California"`
	Locality           string `json:"l" default:"California"`
}
