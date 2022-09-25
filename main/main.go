package main

import "fmt"

type passwordError struct {
	Len        int
	RequireLen int
}

func (err passwordError) Error() string {
	return "암호 길이가 짧습니다."
}

func RegisterAccount(name, password string) error {
	if len(password) < 8 {
		return passwordError{len(password), 8}
	}
	return nil
}

func main() {
	err := RegisterAccount("myID", "MyPw")
	if err != nil {
		if errInfo, ok := err.(passwordError); ok {
			fmt.Printf("%v Len:%d RequiredLen:%d\n",
				errInfo, errInfo.Len, errInfo.RequireLen)
		}
	} else {
		fmt.Println("회원 가입됐습니다.")
	}
}
