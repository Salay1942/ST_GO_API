package main

import (
	"encoding/json"
	"log"
	"net/http"
	"st-go-api/database"
	"st-go-api/model"

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

	if r.Method == "POST" {

		var req model.RegisterRequest

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		nameLa := req.NameLa
		nameEn := req.NameEn
		Address := req.Address
		TelephoneNo := req.TelephoneNo

		log.Println("INFO: NameEN: " + nameEn)
		log.Println("INFO: NameLA: " + nameLa)
		log.Println("INFO: Address: " + Address)
		log.Println("INFO: TelephoneNo: " + TelephoneNo)

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

func main() {
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/plan", GetPlan)
	http.HandleFunc("/marital", GetMarital)
	http.HandleFunc("/occupationClass", GetOccupationClass)
	http.HandleFunc("/insuranceType", GetInsuranceType)
	http.HandleFunc("/RegisterInfo", RegisterInfo)
	http.ListenAndServe(":8080", nil)
}
