package handlers

import (
	"encoding/json"
	"net/http"
	"net/smtp"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/k8shop/go-rest-api/pkg/informer"
	"github.com/k8shop/go-rest-api/pkg/models"
)

//Registration handler
type Registration struct {
	db       *gorm.DB
	informer *informer.Informer
}

//NewRegistrationHandler returns a new users handles
func NewRegistrationHandler() Interface {
	return &Registration{}
}

//Slug for this handler
func (r *Registration) Slug() string {
	return "registration"
}

//Register routes
func (r *Registration) Register(db *gorm.DB, informer *informer.Informer, router *mux.Router) {
	r.db = db
	r.informer = informer
	router.HandleFunc("/new", r.handleNew).Methods("POST")
	router.HandleFunc("/account/{id:[0-9]+}/verify", r.handleVerify).Methods("POST")
	router.HandleFunc("/account/{id:[0-9]+}/send_verification", r.handleSendVerification).Methods("GET")
}

func (r *Registration) handleNew(res http.ResponseWriter, req *http.Request) {
	user := &models.User{
		Email:         req.FormValue("Email"),
		FirstName:     req.FormValue("FirstName"),
		LastName:      req.FormValue("LastName"),
		EmailVerified: false,
	}

	user.SetPassword(req.FormValue("Password"))

	errs := r.db.Create(user).GetErrors()
	if len(errs) > 0 {
		res.WriteHeader(http.StatusBadRequest)
		resBytes, err := json.Marshal(errs)
		if err != nil {
			res.Write([]byte("{\"error\": \"" + err.Error() + "\"}"))
			return
		}
		res.Write(resBytes)
		return
	}

	resBytes, err := json.Marshal(user)
	if err != nil {
		res.Write([]byte("{\"error\": \"" + err.Error() + "\"}"))
		return
	}
	res.Write(resBytes)
}

func (r *Registration) handleVerify(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		panic(err)
	}
	codes := []*models.VerificationCode{}
	errs := r.db.Debug().Find(
		&codes,
		"AccountID = ? AND Code = ? AND Used = false AND Expiry > NOW()",
		id,
		req.FormValue("Code"),
	).GetErrors()
	if len(errs) > 0 {
		res.WriteHeader(http.StatusBadRequest)
		resBytes, err := json.Marshal(errs)
		if err != nil {
			res.Write([]byte("{\"error\": \"" + err.Error() + "\"}"))
			return
		}
		res.Write(resBytes)
		return
	}

	user := &models.User{}
	if len(codes) > 0 {
		r.db.Find(user, id)
		user.EmailVerified = true
		r.db.Save(user)
		for _, code := range codes {
			code.Used = true
			r.db.Save(code)
		}
	}
	resBytes, err := json.Marshal(user)
	if err != nil {
		res.Write([]byte("{\"error\": \"" + err.Error() + "\"}"))
		return
	}
	res.Write(resBytes)
}

func (r *Registration) handleSendVerification(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		panic(err)
	}
	user := &models.User{}
	errs := r.db.Find(user, id).GetErrors()
	if len(errs) > 0 {
		res.WriteHeader(http.StatusBadRequest)
		resBytes, err := json.Marshal(errs)
		if err != nil {
			res.Write([]byte("{\"error\": \"" + err.Error() + "\"}"))
			return
		}
		res.Write(resBytes)
		return
	}
	verification, err := models.NewVerificationCode(id)
	if err != nil {
		res.Write([]byte("{\"error\": \"" + err.Error() + "\"}"))
		return
	}

	r.db.Create(verification)

	//// TODO: send email with code to user.Email
	auth := smtp.PlainAuth(
		"",
		os.Getenv("SMTP_USER"),
		os.Getenv("SMTP_PASS"),
		os.Getenv("SMTP_HOST"),
	)
	msg := []byte("To: " + user.Email + "\r\n" +
		"From: " + os.Getenv("SMTP_FROM") + " \r\n" +
		"Subject: k8shop verification code\r\n" +
		"\r\n" +
		"Code: " + verification.Code + ".\r\n")
	err = smtp.SendMail(
		os.Getenv("SMTP_HOST")+":"+os.Getenv("SMTP_PORT"),
		auth,
		os.Getenv("SMTP_FROM"),
		[]string{user.Email},
		msg,
	)
	if err != nil {
		res.Write([]byte("{\"error\": \"" + err.Error() + "\"}"))
		return
	}

	res.Write([]byte("{\"result\": \"success\"}"))
}
