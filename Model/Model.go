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
	NameLa          string       `json:"name_la"`
	NameEn          string       `json:"name_en"`
	Address         string       `json:"address"`
	TelephoneNo     string       `json:"telephone_no"`
	WhatsAppNo      string       `json:"whatsapp_no"`
	IDCard          string       `json:"id_card"`
	ExpiryDate      string       `json:"expiry_date"`
	IssuedAt        string       `json:"issued_at"`
	Age             string       `json:"age"`
	DateOfBirth     string       `json:"date_of_birth"`
	Height          string       `json:"height"`
	Weight          string       `json:"weight"`
	Nationality     string       `json:"nationality"`
	Occupation      string       `json:"occupation"`
	Position        string       `json:"position"`
	JobDescription  string       `json:"job_description"`
	WorkPlace       string       `json:"work_place"`
	OfficePhoneNo   string       `json:"office_phone_no"`
	OtherOccupation string       `json:"other_occupation"`
	AnswerInfo      []AnswerInfo `json:"answer_info"`
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
