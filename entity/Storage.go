package entity

import (
	"fmt"  
    "regexp"
    "os"
) 

func IsEmail(str string) bool {  
    var b bool  
    b, _ = regexp.MatchString("^([a-z0-9_\\.-]+)@([\\da-z\\.-]+)\\.([a-z\\.]{2,6})$", str)  
    if false == b {  
        return b  
    }
    return b  
}

func IsCellphone(str string) bool {  
    var b bool  
    b, _ = regexp.MatchString("^1[0-9]{10}$", str)  
    if false == b {  
        return b  
    }  
    return b  
}  

func RegisterUser(name string, password string, email string, phone string) bool {
	var user User
	err := false
	if (IsEmail(email) == false) {
		fmt.Println("The email is error!")
		err = true
	}
	if (IsCellphone(phone) == false) {
        	fmt.Println("The phone is error!")
		err = true
	}
	if (len(password) < 6) {
		fmt.Println("The length of password can't be less than 6!")
		err = true
	}
	_, isExit, _:= Query_user(name)
	if (isExit) {
		fmt.Println("The username has already exited, please use another username!")
		err = true
	}
	if (err) {
		return false
	}
	user.Name = name
	user.Password = password
	user.Email = email
	user.Phone = phone
	uData = append(uData,user)
	User_WriteToFile(uData)
	fmt.Println("Regist successfully!")
	return true
}

