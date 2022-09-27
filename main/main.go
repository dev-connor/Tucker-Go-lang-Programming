package main

type Report interface {
	Report() string
}

type FinanceReport struct { // 회계 보고서
	report string
}

func (r *FinanceReport) Report() string {
	return r.report
}

type ReportSender struct {
	// ...
}

func (s *ReportSender) SendReport(report Report) { // 보고서 전송
	// Report 인터페이스 객체를 인수로 받음
	// ...
	print(report.Report())
}

func main() {
	f := FinanceReport{"재무 보고서 내용"}
	r := ReportSender{}

	r.SendReport(&f)

}
