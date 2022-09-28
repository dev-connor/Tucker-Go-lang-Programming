package main

type Mail struct {
	alarm Alarm
}

type Alarm struct {
}

func (m *Mail) OnRecv() { // OnRecv() 메서드는 메일 수신 시 호출됩니다.
	m.alarm.Alarm() // 알람을 울립니다.
}
