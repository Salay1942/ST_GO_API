package model

type Plan struct {
	PlanId   int
	PlanName string
}

type Marital struct {
	MaritalId   int
	MaritalName string
}

type OccupationClass struct {
	OccupationClassId   int
	OccupationClassName string
}

type InsuranceType struct {
	InsuranceTypeId   int
	InsuranceTypeName string
}

type RegisterRequest struct {
	NameLa      string       `json:"name_la"`
	NameEn      string       `json:"name_en"`
	Address     string       `json:"address"`
	TelephoneNo string       `json:"telephone_no"`
	AnswerInfo  []AnswerInfo `json:"answer_info"`
}

type AnswerInfo struct {
	QuestionID    string         `json:"question_id"`
	QuestionName  string         `json:"question_name"`
	AnswerChoices []AnswerChoice `json:"answer_choices"`
}

type AnswerChoice struct {
	ChoiceID    string        `json:"choice_id"`
	ChoiceTitle string        `json:"choice_name"`
	SubQuestion []SubQuestion `json:"sub_question,omitempty"`
}

type SubQuestion struct {
	SubQuestionID     string `json:"sub_question_id"`
	SubQuestionAnwser string `json:"sub_question_anwser"`
}
