package lever

type listCandidatesResp struct {
	Data    []*Candidate `json:"data"`
	Next    string       `json:"next"`
	HasNext bool         `json:"hasNext"`
}

type Candidate struct {
	ID           string
	Name         string
	Stage        string
	StageChanges []*StageChange
}

type StageChange struct {
	ToStageID    string
	ToStageIndex int
	UpdatedAt    int
	UserID       string
}
