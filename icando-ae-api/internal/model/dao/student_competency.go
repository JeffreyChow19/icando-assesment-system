package dao

type StudentCompetencyDao struct {
	CompetencyID   string `json:"competencyId"`
	CompetencyName string `json:"competencyName"`
	CorrectCount   int    `json:"correctCount"`
	TotalCount     int    `json:"totalCount"`
}
