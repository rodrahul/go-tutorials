package main

import (
	"encoding/json"
	"fmt"
)

type HTTPCodes struct {
	Code    int    `json:"Code"`
	Descrip string `json:"Descrip"`
}

func main() {
	var h1 []HTTPCodes
	rcvd := `[{"Code":200,"Descrip":"StatusOK"},{"Code":301,"Descrip":"StatusMovedPermanently"},{"Code":302,"Descrip":"StatusFound"},{"Code":303,"Descrip":"StatusSeeOther"},{"Code":307,"Descrip":"StatusTemporaryRedirect"},{"Code":400,"Descrip":"StatusBadRequest"},{"Code":401,"Descrip":"StatusUnauthorized"},{"Code":402,"Descrip":"StatusPaymentRequired"},{"Code":403,"Descrip":"StatusForbidden"},{"Code":404,"Descrip":"StatusNotFound"},{"Code":405,"Descrip":"StatusMethodNotAllowed"},{"Code":418,"Descrip":"StatusTeapot"},{"Code":500,"Descrip":"StatusInternalServerError"}]`

	err := json.Unmarshal([]byte(rcvd), &h1)
	if err != nil {
		fmt.Println("Unmarshal err: ", err)
	}
	fmt.Println(h1)

	for _, v := range h1 {
		fmt.Println(v.Code, "-", v.Descrip)
	}

}