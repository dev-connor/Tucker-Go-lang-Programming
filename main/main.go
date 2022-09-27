package main

func SendReport(r *Report, method SendType, receiver string) {
	switch method {
	case Email:
		// 이메일 전송
	case Fax:
		// 팩스 전송
	case PDF:
		// pdf 파일 생성
	case Printer:
		// 프린팅
		// ...
	}
}

func main() {
	f := FinanceReport{"재무 보고서 내용"}
	r := ReportSender{}

	r.SendReport(&f)

}
