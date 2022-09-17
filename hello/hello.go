package main

func UploadFile() (string, bool) {
	return "file.txt", true
}

func main() {
	day := 3

	switch day {
	case 1:
		println("첫째 날입니다.")
		println("오늘은 팀미팅이 있습니다.")
	case 2:
		println("둘째 날입니다.")
		println("오늘은 면접이 있습니다.")
	case 3:
		println("셋째 날입니다.")
		println("설계안을 확정하는 날입니다.")
	case 4:
		println("넷째 날입니다.")
		println("예산을 확정하는 날입니다.")
	case 5:
		println("다섯째 날입니다.")
		println("최종 계약하는 날입니다.")
	default:
		println("프로젝트를 진행하세요.")
	}
}
