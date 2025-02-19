package user

import (
	"fmt"
	"net/http"
	"time"

	"math/rand"

	"github.com/ecom-api/config"
	"github.com/ecom-api/models"
	"github.com/ecom-api/services/auth"
	"github.com/ecom-api/types"
	"github.com/ecom-api/utils"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)


type UserService struct {
	db *gorm.DB
}



func NewUserService(database *gorm.DB) *UserService {
	return &UserService{
		db: database,
	}
}


func (s *UserService) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/register",s.RegisterUser).Methods("POST")
	router.HandleFunc("/login",s.LoginUser).Methods("POST")
	router.HandleFunc("/verify-account",s.veryfy_account).Methods("POST")
	router.HandleFunc("/resend-verification-code",s.resend_verification_code).Methods("POST")
}




func (s *UserService) RegisterUser(w http.ResponseWriter, r *http.Request) {

	validate:=validator.New()
	
	var payload types.RegisterBodyType

	err:=utils.ParseJSON(r,&payload)

	if err!=nil{
		utils.WriteError(w,http.StatusBadRequest,err)
		return
	}

	err=validate.Struct(payload)

	if err!=nil{
		utils.WriteError(w,http.StatusBadRequest,err)
		return
	}

	// fmt.Printf("payload ")
	var u models.User
	results:=s.db.Where("email = ?", payload.Email).First(&u)
	

	if results.Error != nil && results.Error != gorm.ErrRecordNotFound {
    utils.WriteError(w, http.StatusInternalServerError, results.Error)
    return
	}

	if results.RowsAffected >0{
		// fmt.Println("user already exists")
		utils.WriteError(w,http.StatusBadRequest,fmt.Errorf("user already exists with email %s",payload.Email))
		return
	}
	var passwordhash string=""

	passwordhash,err=auth.HashPassword(payload.Password)

	if err!=nil{
		utils.WriteError(w,http.StatusInternalServerError,err)
		return
	}

	otp:=rand.Intn(900000) + 100000
	tx:=s.db.Create(&models.User{
		Email:payload.Email,
		Password:passwordhash,
		FirstName: payload.FirstName,
		LastName: payload.LastName,
		OTP: otp,
	})

	if tx.Error != nil {
		utils.WriteError(w, http.StatusInternalServerError, tx.Error)
		return
	}

	mailer:=utils.GetMailer()
	msg:=utils.GetMessage()

	

	msg.SetHeader("From",config.Envs.MailUser)
	msg.SetHeader("To",payload.Email)
	msg.SetHeader("Subject","Welcome to Ecom API")
	msg.SetBody("text/html",fmt.Sprintf("Hello %s, <br> Welcome to Ecom API <br/> Your OTP is %d",payload.FirstName,otp))

	go mailer.DialAndSend(msg)
	
	utils.WriteJSON(w,http.StatusCreated,map[string]string{
		"message": fmt.Sprintf("user created with email %s", payload.Email),
	})


}


func (s *UserService) LoginUser(w http.ResponseWriter, r *http.Request) {
	validate:=validator.New()
	
	var payload types.LoginBodyType

	

	err:=utils.ParseJSON(r,&payload)

	if err!=nil{
		utils.WriteError(w,http.StatusBadRequest,err)
		return
	}

	err=validate.Struct(payload)

	if err!=nil{
		utils.WriteError(w,http.StatusBadRequest,err)
		return
	}

	var u models.User
	results:=s.db.Where("email = ?", payload.Email).First(&u)
	

	if results.Error != nil && results.Error != gorm.ErrRecordNotFound {
		utils.WriteError(w, http.StatusInternalServerError, results.Error)
		return
	}

	if results.RowsAffected ==0{
		// fmt.Println("user already exists")
		utils.WriteError(w,http.StatusBadRequest,fmt.Errorf("user does not exists with email %s",payload.Email))
		return
	}

	if !auth.CheckPasswordHash(payload.Password,u.Password){
		utils.WriteError(w,http.StatusBadRequest,fmt.Errorf("password does not match"))
		return
	}

	token,err:=auth.GenerateJWT( jwt.MapClaims{
		"id": u.ID,
		"email": u.Email,
		"iss": "ecom-api",
		"sub": "auth",
		"time":time.Now().Unix(),

	})

	if err!=nil{
		utils.WriteError(w,http.StatusInternalServerError,err)
		return
	}

	utils.WriteJSON(w,http.StatusOK,map[string]string{
		"token": token,
	})
}

func (s *UserService) veryfy_account(w http.ResponseWriter,r *http.Request){
	validate:=validator.New()
	
	var payload types.VerifyBodyType

	

	err:=utils.ParseJSON(r,&payload)

	if err!=nil{
		utils.WriteError(w,http.StatusBadRequest,err)
		return
	}

	err=validate.Struct(payload)

	if err!=nil{
		utils.WriteError(w,http.StatusBadRequest,err)
		return
	}

	var u models.User
	results:=s.db.Where("email = ?", payload.Email).First(&u)
	

	if results.Error != nil && results.Error != gorm.ErrRecordNotFound {
		utils.WriteError(w, http.StatusInternalServerError, results.Error)
		return
	}

	if results.RowsAffected ==0{
		utils.WriteError(w,http.StatusBadRequest,fmt.Errorf("user does not exists with email %s",payload.Email))
		return
	}

	if u.OTP!=payload.OTP{
		utils.WriteError(w,http.StatusBadRequest,fmt.Errorf("OTP does not match"))
		return
	}

	u.Verified=true

	tx:=s.db.Save(&u)

	if tx.Error != nil {
		utils.WriteError(w, http.StatusInternalServerError, tx.Error)
		return
	}

	utils.WriteJSON(w,http.StatusOK,map[string]string{
		"message": "account verified",
	})
}

func (s *UserService) resend_verification_code(w http.ResponseWriter,r *http.Request){

	validate:=validator.New()
	
	var payload types.ResendBodyType


	err:=utils.ParseJSON(r,&payload)

	if err!=nil{
		utils.WriteError(w,http.StatusBadRequest,err)
		return
	}

	err=validate.Struct(payload)

	if err!=nil{
		utils.WriteError(w,http.StatusBadRequest,err)
		return
	}

	var u models.User
	results:=s.db.Where("email = ?", payload.Email).First(&u)
	

	if results.Error != nil && results.Error != gorm.ErrRecordNotFound {
		utils.WriteError(w, http.StatusInternalServerError, results.Error)
		return
	}

	if results.RowsAffected ==0{
		utils.WriteError(w,http.StatusBadRequest,fmt.Errorf("user does not exists with email %s",payload.Email))
		return
	}

	otp:=rand.Intn(900000) + 100000

	u.OTP=otp

	tx:=s.db.Save(&u)

	if tx.Error != nil {
		utils.WriteError(w, http.StatusInternalServerError, tx.Error)
		return
	}

	mailer:=utils.GetMailer()
	msg:=utils.GetMessage()

	

	msg.SetHeader("From",config.Envs.MailUser)
	msg.SetHeader("To",payload.Email)
	msg.SetHeader("Subject","Welcome to Ecom API")
	msg.SetBody("text/html",fmt.Sprintf("Hello %s, <br> Welcome to Ecom API <br/> Your OTP is %d",u.FirstName,otp))

	go mailer.DialAndSend(msg)
	
	utils.WriteJSON(w,http.StatusOK,map[string]string{
		"message": fmt.Sprintf("OTP sent to email %s", payload.Email),
	})
}
