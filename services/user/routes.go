package user

import (
	"fmt"
	"math/rand"
	"net/http"

	"os"
	"time"

	"github.com/ecom-api/config"
	"github.com/ecom-api/middlewares"
	"github.com/ecom-api/models"
	"github.com/ecom-api/services/auth"
	"github.com/ecom-api/types"
	"github.com/ecom-api/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
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
func (s *UserService) RegisterRoutes(router fiber.Router) {

	router.Post("/register",s.RegisterUser)
	router.Post("/login",s.LoginUser)
	router.Post("/verify-account",s.veryfy_account)

	//only called 2 times in 5 minutes
	router.Post("/resend-verification-code",limiter.New(limiter.Config{
		Max: 2,
		Expiration: 5* time.Minute,
	}),middlewares.Is_authenticated,s.resend_verification_code)

	oauthService:=auth.NewOAuthService(s.db)
	router.Get("/google", oauthService.GoogleLogin)
	router.Get("/google/callback", oauthService.GoogleCallback)
	
	router.Use("/protected",middlewares.Is_authenticated,func (c *fiber.Ctx) error {
		return utils.WriteJSON(c,http.StatusCreated,map[string]string{
			"message": fmt.Sprintf("protected route %s",c.Locals("user_id")),
		})
	})
}
func (s *UserService) RegisterUser(c *fiber.Ctx) error {


	validate:=validator.New()
	
	var payload types.RegisterBodyType

	err:=c.BodyParser(&payload)

	if err!=nil{
		return utils.WriteError(c,http.StatusBadRequest,err)
		
	}

	err=validate.Struct(payload)

	if err!=nil{
		return utils.WriteError(c,http.StatusBadRequest,err)
		
	}

	// fmt.Printf("payload ")
	var u models.User
	results:=s.db.Where("email = ?", payload.Email).First(&u)
	

	if results.Error != nil && results.Error != gorm.ErrRecordNotFound {
    return utils.WriteError(c, http.StatusInternalServerError, results.Error)
    
	}

	if results.RowsAffected >0{
		// fmt.Println("user already exists")
		return utils.WriteError(c,http.StatusBadRequest,fmt.Errorf("user already exists with email %s",payload.Email))
		
	}
	var passwordhash string=""

	passwordhash,err=auth.HashPassword(payload.Password)

	if err!=nil{
		return utils.WriteError(c,http.StatusInternalServerError,err)
		
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
		return utils.WriteError(c, http.StatusInternalServerError, tx.Error)
		
	}

	mailer:=utils.GetMailer()
	msg:=utils.GetMessage()


	var path string ="./uploads/"+payload.Email;

	info,error:=os.Stat(path);

	fmt.Println(info);
	fmt.Println(error.Error());

	if os.IsNotExist(error){
			err=os.MkdirAll(path, os.ModePerm)
			if err!=nil{
				return utils.WriteError(c,http.StatusInternalServerError,err)
				
			}
		
	}else {
		if info.IsDir() {
			path="./uploads/"+payload.Email+"/"+time.Now().Format("2006-01-02")
			err=os.MkdirAll(path, os.ModePerm)
			if err!=nil{
				return utils.WriteError(c,http.StatusInternalServerError,err)
			}
		}
	}

	msg.SetHeader("From",config.Envs.MailUser)
	msg.SetHeader("To",payload.Email)
	msg.SetHeader("Subject","Welcome to Ecom API")
	msg.SetBody("text/html",fmt.Sprintf("Hello %s, <br> Welcome to Ecom API <br/> Your OTP is %d",payload.FirstName,otp))

	go mailer.DialAndSend(msg)
	
	return utils.WriteJSON(c,http.StatusCreated,map[string]string{
		"message": fmt.Sprintf("user created with email %s", payload.Email),
	})


}


func (s *UserService) LoginUser(c *fiber.Ctx) error {
	validate:=validator.New()
	
	var payload types.LoginBodyType

	

	err:=c.BodyParser(&payload)

	if err!=nil{
		return utils.WriteError(c,http.StatusBadRequest,err)
		
	}

	err=validate.Struct(payload)

	if err!=nil{
		return utils.WriteError(c,http.StatusBadRequest,err)
		
	}

	var u models.User
	results:=s.db.Where("email = ?", payload.Email).First(&u)
	

	if results.Error != nil && results.Error != gorm.ErrRecordNotFound {
		return utils.WriteError(c, http.StatusInternalServerError, results.Error)
		
	}

	if results.RowsAffected ==0{
		// fmt.Println("user already exists")
		return utils.WriteError(c,http.StatusBadRequest,fmt.Errorf("user does not exists with email %s",payload.Email))
		
	}

	if !auth.CheckPasswordHash(payload.Password,u.Password){
		return utils.WriteError(c,http.StatusBadRequest,fmt.Errorf("password does not match"))
		
	}

	token,err:=auth.GenerateJWT(u)

	if err!=nil{
		return utils.WriteError(c,http.StatusInternalServerError,err)
		
	}

	return utils.WriteJSON(c,http.StatusOK,map[string]string{
		"token": token,
	})
}

func (s *UserService) veryfy_account(c *fiber.Ctx) error{
	validate:=validator.New()
	
	var payload types.VerifyBodyType

	

	err:=c.BodyParser(&payload)

	if err!=nil{
		return utils.WriteError(c,http.StatusBadRequest,err)
		
	}

	err=validate.Struct(payload)

	if err!=nil{
		return utils.WriteError(c,http.StatusBadRequest,err)
		
	}

	var u models.User
	results:=s.db.Where("email = ?", payload.Email).First(&u)
	

	if results.Error != nil && results.Error != gorm.ErrRecordNotFound {
		return utils.WriteError(c, http.StatusInternalServerError, results.Error)
		
	}

	if results.RowsAffected ==0{
		utils.WriteError(c,http.StatusBadRequest,fmt.Errorf("user does not exists with email %s",payload.Email))
		
	}

	if u.OTP!=payload.OTP{
		return utils.WriteError(c,http.StatusBadRequest,fmt.Errorf("OTP does not match"))
		
	}

	u.Verified=true

	tx:=s.db.Save(&u)

	if tx.Error != nil {
		return utils.WriteError(c, http.StatusInternalServerError, tx.Error)
		
	}

	return utils.WriteJSON(c,http.StatusOK,map[string]string{
		"message": "account verified",
	})
}

func (s *UserService) resend_verification_code(c *fiber.Ctx) error{

	validate:=validator.New()
	
	var payload types.ResendBodyType


	err:=c.BodyParser(&payload)

	if err!=nil{
		return utils.WriteError(c,http.StatusBadRequest,err)
		
	}

	err=validate.Struct(payload)

	if err!=nil{
		return utils.WriteError(c,http.StatusBadRequest,err)
		
	}

	var u models.User
	results:=s.db.Where("email = ?", payload.Email).First(&u)
	

	if results.Error != nil && results.Error != gorm.ErrRecordNotFound {
		return utils.WriteError(c, http.StatusInternalServerError, results.Error)
		
	}

	if results.RowsAffected ==0{
		return utils.WriteError(c,http.StatusBadRequest,fmt.Errorf("user does not exists with email %s",payload.Email))
		
	}

	otp:=rand.Intn(900000) + 100000

	u.OTP=otp

	tx:=s.db.Save(&u)

	if tx.Error != nil {
		return utils.WriteError(c, http.StatusInternalServerError, tx.Error)
		
	}

	mailer:=utils.GetMailer()
	msg:=utils.GetMessage()

	

	msg.SetHeader("From",config.Envs.MailUser)
	msg.SetHeader("To",payload.Email)
	msg.SetHeader("Subject","Welcome to Ecom API")
	msg.SetBody("text/html",fmt.Sprintf("Hello %s, <br> Welcome to Ecom API <br/> Your OTP is %d",u.FirstName,otp))

	go mailer.DialAndSend(msg)
	
	return utils.WriteJSON(c,http.StatusOK,map[string]string{
		"message": fmt.Sprintf("OTP sent to email %s", payload.Email),
	})
}
