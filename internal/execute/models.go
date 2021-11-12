package execute


type ExecuteResponse struct {
	ExitCode string `json:"exit_code"`
	Output   string `json:"output"`
	Time     string `json:"time"`
}

type ExecuteRequest struct {
	Code          string `json:"code"`
	ChallengePath string `json:"challenge_path"`
}