package main

import (
	"encoding/json"
	"log"
	"net/http"
	"st-go-api/database"
	"st-go-api/model"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetMarital(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := database.ConnectDb()
	selDB, err := db.Query("SELECT marital_id, marital_name FROM  tb_marital_status ORDER BY marital_id ASC")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	marital := model.Marital{}
	res := []model.Marital{}

	for selDB.Next() {
		var MaritalId int
		var MaritalName string
		err = selDB.Scan(&MaritalId, &MaritalName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		marital.MaritalId = MaritalId
		marital.MaritalName = MaritalName
		res = append(res, marital)
	}

	jsonConverter, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(jsonConverter)
	defer db.Close()
}

func GetPlan(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := database.ConnectDb()
	selDB, err := db.Query("SELECT plan_id, plan_name FROM tb_plan ORDER BY plan_id ASC")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	plan := model.Plan{}
	res := []model.Plan{}

	for selDB.Next() {
		var PlanId int
		var PlanName string
		err = selDB.Scan(&PlanId, &PlanName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		plan.PlanId = PlanId
		plan.PlanName = PlanName
		res = append(res, plan)
	}

	jsonConverter, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(jsonConverter)
	defer db.Close()
}

func GetOccupationClass(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := database.ConnectDb()
	selDB, err := db.Query("SELECT occupation_class_id, occupation_class_name FROM tb_occupation_class ORDER BY occupation_class_id ASC")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ocp := model.OccupationClass{}
	res := []model.OccupationClass{}

	for selDB.Next() {
		var OccupationClassId int
		var OccupationClassName string
		err = selDB.Scan(&OccupationClassId, &OccupationClassName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		ocp.OccupationClassId = OccupationClassId
		ocp.OccupationClassName = OccupationClassName
		res = append(res, ocp)
	}

	jsonConverter, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(jsonConverter)
	defer db.Close()
}

func GetInsuranceType(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := database.ConnectDb()
	selDB, err := db.Query("SELECT insurance_type_id, insurance_type_name FROM tb_insurance_type ORDER BY insurance_type_id ASC")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ins := model.InsuranceType{}
	res := []model.InsuranceType{}

	for selDB.Next() {
		var InsuranceTypeId int
		var InsuranceTypeName string
		err = selDB.Scan(&InsuranceTypeId, &InsuranceTypeName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		ins.InsuranceTypeId = InsuranceTypeId
		ins.InsuranceTypeName = InsuranceTypeName
		res = append(res, ins)
	}

	jsonConverter, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(jsonConverter)
	defer db.Close()
}

func RegisterInfo(w http.ResponseWriter, r *http.Request) {

	//db := dbConn()

	var RegisterID string
	var NameLA string
	var NameEN string
	var Address string
	var TelephoneNo string
	var WhatsAppNo string
	var IDCard string
	var ExpiryDate string
	var IssuedAt string
	var Age string
	var DateOfBirth string
	var Height string
	var Weight string
	var Nationality string
	var Occupation string
	var Position string
	var JobDescription string
	var WorkPlace string
	var OfficePhoneNo string
	var OtherOccupation string
	var OccupationID string
	var UserID string

	if r.Method == "POST" {

		var req model.RegisterRequest

		err := json.NewDecoder(r.Body).Decode(&req)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		RegisterID = req.RegisterID
		NameLA = req.NameLa
		NameEN = req.NameEn
		Address = req.Address
		TelephoneNo = req.TelephoneNo
		WhatsAppNo = req.WhatsAppNo
		IDCard = req.IDCard
		ExpiryDate = req.ExpiryDate
		IssuedAt = req.IssuedAt
		Age = req.Age
		DateOfBirth = req.DateOfBirth
		Height = req.Height
		Weight = req.Weight
		Nationality = req.Nationality
		Occupation = req.Occupation
		Position = req.Position
		JobDescription = req.JobDescription
		WorkPlace = req.WorkPlace
		OfficePhoneNo = req.OfficePhoneNo
		OtherOccupation = req.OtherOccupation
		OccupationID = req.OccupationID
		UserID = req.UserID

		InsertRegisterInfo(req)

		log.Println("INFO: RegisterID: " + RegisterID)
		log.Println("INFO: NameLA: " + NameLA)
		log.Println("INFO: NameEN: " + NameEN)
		log.Println("INFO: Address: " + Address)
		log.Println("INFO: TelephoneNo: " + TelephoneNo)
		log.Println("INFO: WhatsAppNo: " + WhatsAppNo)
		log.Println("INFO: IDCard: " + IDCard)
		log.Println("INFO: ExpiryDate: " + ExpiryDate)
		log.Println("INFO: IssuedAt: " + IssuedAt)
		log.Println("INFO: Age: " + Age)
		log.Println("INFO: DateOfBirth: " + DateOfBirth)
		log.Println("INFO: Height: " + Height)
		log.Println("INFO: Weight: " + Weight)
		log.Println("INFO: Nationality: " + Nationality)
		log.Println("INFO: Occupation: " + Occupation)
		log.Println("INFO: Position: " + Position)
		log.Println("INFO: JobDescription: " + JobDescription)
		log.Println("INFO: WorkPlace: " + WorkPlace)
		log.Println("INFO: OfficePhoneNo: " + OfficePhoneNo)
		log.Println("INFO: OtherOccupation: " + OtherOccupation)
		log.Println("INFO: OccupationClassID: " + OccupationID)
		log.Println("INFO: UserID: " + UserID)

		for i := 0; i < len(req.AnswerInfo); i++ {
			questionID := req.AnswerInfo[i].QuestionID
			log.Println("INFO: QuestionID: " + questionID)
			choiceSelected := req.AnswerInfo[i].AnswerChoices
			for j := 0; j < len(choiceSelected); j++ {
				choiceID := choiceSelected[j].ChoiceID
				choiceTitle := choiceSelected[j].ChoiceTitle
				subQuestion := choiceSelected[j].SubQuestion
				log.Println("INFO: ChoiceID: " + choiceID)
				log.Println("INFO: ChoiceTitle: " + choiceTitle)
				InsertInfoAnswer(questionID, choiceID, RegisterID)
				for k := 0; k < len(subQuestion); k++ {
					subQuestionID := subQuestion[k].SubQuestionID
					subQuestionAnsw := subQuestion[k].SubQuestionAnwser
					log.Println("INFO: SubQuestionID: " + subQuestionID)
					log.Println("INFO: SubQuestionAnswer: " + subQuestionAnsw)
				}
			}
		}

		//name_x := req.Name
		//city_x := req.City

		//insForm, err := db.Prepare("INSERT INTO Employee(name, city) VALUES(?,?)")
		//insForm, err := db.Prepare("CALL InsertEmployee(?,?)")
		//if err != nil {
		//	panic(err.Error())
		//}
		//insForm.Exec(name_x, city_x)
		//log.Println("INSERT: Name: " + name + " | City: " + city)
		//log.Println("INSERT: Name: " + name_x + " | City: " + city_x)
	}
	//defer db.Close()
	//http.Redirect(w, r, "/", 301)
}

func InsertRegisterInfo(m model.RegisterRequest) bool {

	var result bool = false

	db := database.ConnectDb()

	res, err := db.Prepare("insert into tb_register(register_id, name_la, name_en, address, telephone_no, whatsapp_no, id_card, expiry_date, issued_at, age, date_of_birth, height, weight, nationality, occupation, position, job_description, work_place, office_phone_no, other_occupation, occupation_class_id, user_id) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")

	if err != nil {
		panic(err.Error())
	}

	expiryDate, err := time.Parse("2006-01-02", m.ExpiryDate)

	if err != nil {
		panic(err.Error())
	}

	dateOfBirth, err := time.Parse("2006-01-02", m.DateOfBirth)

	if err != nil {
		panic(err.Error())
	}

	res.Exec(m.RegisterID, m.NameLa, m.NameEn, m.Address, m.TelephoneNo, m.WhatsAppNo, m.IDCard, expiryDate, m.IssuedAt, m.Age, dateOfBirth, m.Height, m.Weight, m.Nationality, m.Occupation, m.Position, m.JobDescription, m.WorkPlace, m.OfficePhoneNo, m.OtherOccupation, m.OccupationID, m.UserID)

	if res != nil {
		result = true
	} else {
		panic(err.Error())
	}

	defer db.Close()

	return result
}

func InsertInfoAnswer(QuestionID string, ChoiceID string, RegisterID string) bool {

	var result bool = false

	db := database.ConnectDb()

	res, err := db.Prepare("insert into tb_answer(question_id, choice_id, register_id) values (?,?,?)")

	if err != nil {
		panic(err.Error())
	}

	res.Exec(QuestionID, ChoiceID, RegisterID)

	if res != nil {
		result = true
	}

	return result

}

func InsertInfoAnswerDetail(AnswerID int, ChoiceID string, QuestionID string, SubQuestionID string, Text string) bool {

	var result bool = false

	db := database.ConnectDb()

	res, err := db.Prepare("insert into tb_sub_answer(answer_id, choice_id, question_id, sub_question_id, text) values (?,?,?,?,?)")

	if err != nil {
		panic(err.Error())
	}

	res.Exec(AnswerID, ChoiceID, QuestionID, SubQuestionID, Text)

	if res != nil {
		result = true
	}

	return result

}

func main() {
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/plan", GetPlan)
	http.HandleFunc("/marital", GetMarital)
	http.HandleFunc("/occupationClass", GetOccupationClass)
	http.HandleFunc("/insuranceType", GetInsuranceType)
	http.HandleFunc("/RegisterInfo", RegisterInfo)
	http.ListenAndServe(":8080", nil)
}
