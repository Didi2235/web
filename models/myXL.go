package models

import (
	"fmt"
	"net/http"
	"strconv"

	"../plugin"
)

type Account struct {
	Status int    `json:"Status"`
	ReqOTP string `json:"test"`
	Resu   map[string]interface{}
	Token  string
	RToken string
	Msisdn string
}

func Otp(Number int) Account {
	Numb := strconv.Itoa(Number)

	url := "https://otp-service.apps.dp.xl.co.id/v1/generate/" + Numb + "/MYXLAPP_LOGIN_ID"
	var Pilot Account
	payload := map[string]string{
		"test": "",
	}

	body := plugin.Body(payload)
	header := map[string]string{
		"authorization":     "Basic ZGVtb2NsaWVudDpkZW1vY2xpZW50c2VjcmV0",
		"x-apicache-bypass": "true",
		"Content-type":      "application/json"}
	result, err := plugin.Init("POST", url, header, body)
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
		Pilot.Status = result.Status
		Pilot.Resu = map[string]interface{}{"Status": "error"}
		return Pilot
	}
	if result.Status == 200 {
		Pilot.Resu = map[string]interface{}{"Status": "OK"}
	} else {
		Pilot.Resu = map[string]interface{}{"Status": "Failed"}
	}
	Pilot.Status = result.Status
	v := result.Body["statusCode"]
	fmt.Println(v)

	return Pilot
}

func Loginxl(Number int, Otp string) Account {
	Numb := strconv.Itoa(Number)
	url := "https://login-controller-service.apps.dp.xl.co.id/v1/login/otp/auth?msisdn=" + Numb + "&imei=dcb39e12-2472-070c-0779-007925599283&otp=" + Otp + "&channel=MYXLAPP_LOGIN_ID"
	var Pilot Account
	header := map[string]string{
		"authorization":     "Basic ZGVtb2NsaWVudDpkZW1vY2xpZW50c2VjcmV0",
		"x-apicache-bypass": "true",
		"Content-type":      "application/json"}

	result, err := plugin.Init("GET", url, header, nil)
	if err != nil {
		Pilot.Status = http.StatusInternalServerError
		Pilot.Resu = map[string]interface{}{"Message": err}
		return Pilot
	}
	Pilot.Status = http.StatusOK
	v := result.Body["result"]
	b := v.(map[string]interface{})
	c := b["user"].(map[string]interface{})
	Pilot.Msisdn = c["msisdn_enc"].(string)
	Pilot.Token = b["accessToken"].(string)
	Pilot.RToken = b["refreshToken"].(string)
	Pilot.Resu = map[string]interface{}{"Message": "success"}
	return Pilot
}
