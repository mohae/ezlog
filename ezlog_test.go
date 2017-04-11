package ezlog

import (
	"bufio"
	"bytes"
	"fmt"
	"testing"
)

func TestLogFilenameFormat(t *testing.T) {
	var buf bytes.Buffer
	tst := New(LogDebug, Full, &buf, "", Lshortfile)
	s := "oh no Mr. Bill!"
	g := "Gumby!!!"
	tst.Error(s)
	if buf.String() != fmt.Sprintf("ezlog_test.go:15: ERROR: %s\n", s) {
		t.Errorf("error: got %q want \"ezlog_test.go:15: ERROR: %s\n\"", buf.String(), s)
	}
	buf.Reset()
	tst.Errorf("%s %s", s, g)
	if buf.String() != fmt.Sprintf("ezlog_test.go:20: ERROR: %s %s\n", s, g) {
		t.Errorf("errorf: got %q want \"ezlog_test.go:20: ERROR: %s %s\n\"", buf.String(), s, g)
	}
	buf.Reset()
	tst.Errorln(s)
	if buf.String() != fmt.Sprintf("ezlog_test.go:25: ERROR: %s\n", s) {
		t.Errorf("errorln: got %q want \"ezlog_test.go:25: ERROR: %s\n\"", buf.String(), s)
	}
	buf.Reset()
	tst.Info(s)
	if buf.String() != fmt.Sprintf("ezlog_test.go:30: INFO: %s\n", s) {
		t.Errorf("info: got %q want \"ezlog_test.go:30: INFO: %s\n\"", buf.String(), s)
	}
	buf.Reset()
	tst.Infof("%s %s", s, g)
	if buf.String() != fmt.Sprintf("ezlog_test.go:35: INFO: %s %s\n", s, g) {
		t.Errorf("infof: got %q want \"ezlog_test.go:35: INFO: %s %s\n\"", buf.String(), s, g)
	}
	buf.Reset()
	tst.Infoln(s)
	if buf.String() != fmt.Sprintf("ezlog_test.go:40: INFO: %s\n", s) {
		t.Errorf("infoln: got %q want \"ezlog_test.go:40: INFO: %s\n\"", buf.String(), s)
	}
	buf.Reset()
	tst.Debug(s)
	if buf.String() != fmt.Sprintf("ezlog_test.go:45: DEBUG: %s\n", s) {
		t.Errorf("debug: got %q want \"ezlog_test.go:45: DEBUG: %s\n\"", buf.String(), s)
	}
	buf.Reset()
	tst.Debugf("%s %s", s, g)
	if buf.String() != fmt.Sprintf("ezlog_test.go:50: DEBUG: %s %s\n", s, g) {
		t.Errorf("debugf: got %q want \"ezlog_test.go:50: DEBUG: %s %s\n\"", buf.String(), s, g)
	}
	buf.Reset()
	tst.Debugln(s)
	if buf.String() != fmt.Sprintf("ezlog_test.go:55: DEBUG: %s\n", s) {
		t.Errorf("debugln: got %q want \"ezlog_test.go:55: DEBUG: %s\n\"", buf.String(), s)
	}
}

func TestSeverityByName(t *testing.T) {
	tests := []struct {
		name  string
		level Level
		ok    bool
	}{
		{"", 0, false},
		{"warn", 0, false},
		{"none", LogNone, true},
		{"error", LogError, true},
		{"info", LogInfo, true},
		{"debug", LogDebug, true},
		{"DEBUG", LogDebug, true},
		{"DeBuG", LogDebug, true},
	}

	for _, test := range tests {
		v, ok := LevelByName(test.name)
		if ok != test.ok {
			t.Errorf("%s: got %v; want %v", test.name, ok, test.ok)
		}
		if v != test.level {
			t.Errorf("%s: got %v; want %v", test.name, v, test.level)
		}
	}
}

func TestParseLogFlag(t *testing.T) {
	tests := []struct {
		v        string
		expected int
		err      error
	}{
		{"", 0, UnknownFlagError{}},
		{"zdate", 0, UnknownFlagError{"zdate"}},
		{"Ldate", Ldate, nil},
		{"Date", Ldate, nil},
		{"LTIME", Ltime, nil},
		{"TIME", Ltime, nil},
		{"lmicroseconds", Lmicroseconds, nil},
		{"MicroSeconds", Lmicroseconds, nil},
		{"llongfile", Llongfile, nil},
		{"longfile", Llongfile, nil},
		{"LShortFile", Lshortfile, nil},
		{"SHORTFILE", Lshortfile, nil},
		{"lUTC", LUTC, nil},
		{"UTC", LUTC, nil},
		{"lstdflags", LstdFlags, nil},
		{"stdflags", LstdFlags, nil},
		{"none", 0, nil},
	}

	for _, test := range tests {
		v, err := ParseLogFlag(test.v)
		if err != nil {
			if err != test.err {
				t.Errorf("%q: got %s; want %s", test.v, err, test.err)
			}
			continue
		}
		if test.err != nil {
			t.Errorf("%q: got no error; wanted %s", test.v, test.err)
			continue
		}
		if v != test.expected {
			t.Errorf("%q: got %d; want %d", test.v, v, test.expected)
		}
	}
}

func TestNoneLogger(t *testing.T) {
	var buf bytes.Buffer
	SetLevel(LogNone)
	Error("error")
	if buf.Len() > 0 {
		t.Errorf("write error line: expected no bytes to be written, %d were", buf.Len())
	}
	Errorf("errorf: %d", 42)
	if buf.Len() > 0 {
		t.Errorf("write errorf line: expected no bytes to be written, %d were", buf.Len())
	}
	Errorln("error")
	if buf.Len() > 0 {
		t.Errorf("write error line: expected no bytes to be written, %d were", buf.Len())
	}
	Info("info")
	if buf.Len() > 0 {
		t.Errorf("write info line: expected no bytes to be written, %d were", buf.Len())
	}
	Infof("infof: %d", 42)
	if buf.Len() > 0 {
		t.Errorf("write infof line: expected no bytes to be written, %d were", buf.Len())
	}
	Infoln("info")
	if buf.Len() > 0 {
		t.Errorf("write info line: expected no bytes to be written, %d were", buf.Len())
	}
	Debug("debug")
	if buf.Len() > 0 {
		t.Errorf("write debug line: expected no bytes to be written, %d were", buf.Len())
	}
	Debugf("debugf: %d", 42)
	if buf.Len() > 0 {
		t.Errorf("write debug line: expected no bytes to be written, %d were", buf.Len())
	}
	Debugln("debugln")
	if buf.Len() > 0 {
		t.Errorf("write debug line: expected no bytes to be written, %d were", buf.Len())
	}
}

func TestErrorLogger(t *testing.T) {
	var buf bytes.Buffer
	l := New(LogError, Short, &buf, "", 0)
	l.Error("error", 42)
	if buf.String() != "ERR: error42\n" {
		t.Errorf("write error line: got %q; want \"ERROR: error42\n\"", buf.String())
	}
	buf.Reset()
	l.Errorf("errorf: %d %s", 42, "zaphod")
	if buf.String() != "ERR: errorf: 42 zaphod\n" {
		t.Errorf("write error line: got %q; want \"ERROR: errorf: 42 zaphod\n\"", buf.String())
	}
	buf.Reset()
	l.Errorln("errorln", 42)
	if buf.String() != "ERR: errorln 42\n" {
		t.Errorf("write errorln line: got %q; want \"ERROR: errorln 42\n\"", buf.String())
	}
	buf.Reset()
	l.Info("info")
	if buf.Len() > 0 {
		t.Errorf("write info line: expected no bytes to be written, %d were", buf.Len())
	}
	l.Infof("infof: %d", 42)
	if buf.Len() > 0 {
		t.Errorf("write infof line: expected no bytes to be written, %d were", buf.Len())
	}
	l.Infoln("infoln")
	if buf.Len() > 0 {
		t.Errorf("write infoln line: expected no bytes to be written, %d were", buf.Len())
	}
	l.Debug("debug")
	if buf.Len() > 0 {
		t.Errorf("write debug line: expected no bytes to be written, %d were", buf.Len())
	}
	l.Debugf("debugf: %d", 42)
	if buf.Len() > 0 {
		t.Errorf("write debugf line: expected no bytes to be written, %d were", buf.Len())
	}
	l.Debugln("debugln")
	if buf.Len() > 0 {
		t.Errorf("write debugln line: expected no bytes to be written, %d were", buf.Len())
	}
}

func TestInfoLogger(t *testing.T) {
	var buf bytes.Buffer
	l := New(LogInfo, Full, &buf, "", 0)
	l.Error("error")
	if buf.String() != "ERROR: error\n" {
		t.Errorf("write error line: got %q; want \"ERROR: error\n\"", buf.String())
	}
	buf.Reset()
	l.Errorf("errorf: %d", 42)
	if buf.String() != "ERROR: errorf: 42\n" {
		t.Errorf("write error line: got %q; want \"ERROR: errorf: 42\n\"", buf.String())
	}
	buf.Reset()
	l.Errorln("errorln:", 42)
	if buf.String() != "ERROR: errorln: 42\n" {
		t.Errorf("write errorln line: got %q; want \"ERROR: errorln: 42\n\"", buf.String())
	}
	buf.Reset()
	l.Info("info", "trillian", "arthur")
	if buf.String() != "INFO: infotrillianarthur\n" {
		t.Errorf("write info line: got %q; want \"INFO: infotrillianarthur\n\"", buf.String())
	}
	buf.Reset()
	l.Infof("infof: %d %d", 42, 11)
	if buf.String() != "INFO: infof: 42 11\n" {
		t.Errorf("write info line: got %q; want \"INFO: infof: 42 11\n\"", buf.String())
	}
	buf.Reset()
	l.Infoln("infoln:", 42, 11)
	if buf.String() != "INFO: infoln: 42 11\n" {
		t.Errorf("write infoln line: got %q; want \"INFO: infoln: 42 11\n\"", buf.String())
	}
	buf.Reset()
	l.Debug("debug")
	if buf.Len() > 0 {
		t.Errorf("write debug line: expected no bytes to be written, %d were", buf.Len())
	}
	l.Debugf("debugf: %d", 42)
	if buf.Len() > 0 {
		t.Errorf("write debugf line: expected no bytes to be written, %d were", buf.Len())
	}
	l.Debug("debugln")
	if buf.Len() > 0 {
		t.Errorf("write debugln line: expected no bytes to be written, %d were", buf.Len())
	}

}

func TestDebugLogger(t *testing.T) {
	var buf bytes.Buffer
	l := New(LogDebug, Full, &buf, "", 0)
	l.Error("error")
	if buf.String() != "ERROR: error\n" {
		t.Errorf("write error line: got %q; want \"ERROR: error\n\"", buf.String())
	}
	buf.Reset()
	l.Errorf("errorf: %d", 42)
	if buf.String() != "ERROR: errorf: 42\n" {
		t.Errorf("write error line: got %q; want \"ERROR: errorf: 42\n\"", buf.String())
	}
	buf.Reset()
	l.Errorln("errorln:", 42)
	if buf.String() != "ERROR: errorln: 42\n" {
		t.Errorf("write errorln line: got %q; want \"ERROR: errorln: 42\n\"", buf.String())
	}
	buf.Reset()
	l.Info("info")
	if buf.String() != "INFO: info\n" {
		t.Errorf("write info line: got %q; want \"INFO: info\n\"", buf.String())
	}
	buf.Reset()
	l.Infof("infof: %d", 42)
	if buf.String() != "INFO: infof: 42\n" {
		t.Errorf("write info line: got %q; want \"INFO: infof: 42\n\"", buf.String())
	}
	buf.Reset()
	l.Infoln("infoln:", 42)
	if buf.String() != "INFO: infoln: 42\n" {
		t.Errorf("write infoln line: got %q; want \"INFO: infoln: 42\n\"", buf.String())
	}
	buf.Reset()
	l.Debug("debug", "hoopy", "frood")
	if buf.String() != "DEBUG: debughoopyfrood\n" {
		t.Errorf("write debug line: %q; want \"DEBUG: debughoopyfrood\n\"", buf.String())
	}
	buf.Reset()
	l.Debugf("debugf: %d %d", 42, 1999)
	if buf.String() != "DEBUG: debugf: 42 1999\n" {
		t.Errorf("write debug line: %q; want \"DEBUG: debugf: 42 1999\n\"", buf.String())
	}
	buf.Reset()
	l.Debugln("debugln:", 42, 1999)
	if buf.String() != "DEBUG: debugln: 42 1999\n" {
		t.Errorf("write debugln line: %q; want \"DEBUG: debugln: 42 1999\n\"", buf.String())
	}
}

func TestUseCharFlagsPrefix(t *testing.T) {
	var buf bytes.Buffer
	l := New(LogDebug, Full, &buf, "", 0)
	l.SetFlags(LstdFlags)
	f := l.Flags()
	if f != LstdFlags {
		t.Errorf("flags: got %d; want %d", f, LstdFlags)
	}
	l.SetFlags(0)
	l.Error("error")
	if buf.String() != "ERROR: error\n" {
		t.Errorf("write error line: got %q; want \"ERROR: error\n\"", buf.String())
	}
	buf.Reset()
	l.Errorf("errorf: %d", 42)
	if buf.String() != "ERROR: errorf: 42\n" {
		t.Errorf("write error line: got %q; want \"ERROR: errorf: 42\n\"", buf.String())
	}
	buf.Reset()
	l.Errorln("errorln:", 42)
	if buf.String() != "ERROR: errorln: 42\n" {
		t.Errorf("write errorln line: got %q; want \"ERROR: errorln: 42\n\"", buf.String())
	}
	buf.Reset()
	// change to use char
	l.SetLevelStringType(Char)
	l.Error("error")
	if buf.String() != "E: error\n" {
		t.Errorf("write error line: got %q; want \"E: error\n\"", buf.String())
	}
	buf.Reset()
	l.Errorf("errorf: %d", 42)
	if buf.String() != "E: errorf: 42\n" {
		t.Errorf("write error line: got %q; want \"E: errorf: 42\n\"", buf.String())
	}
	buf.Reset()
	l.Errorln("errorln:", 42)
	if buf.String() != "E: errorln: 42\n" {
		t.Errorf("write errorln line: got %q; want \"E: errorln: 42\n\"", buf.String())
	}
	buf.Reset()
	l.SetLevelStringType(Short)
	l.Info("info")
	if buf.String() != "INF: info\n" {
		t.Errorf("write info line: got %q; want \"INF: info\n\"", buf.String())
	}
	buf.Reset()
	l.Infof("infof: %d", 42)
	if buf.String() != "INF: infof: 42\n" {
		t.Errorf("write info line: got %q; want \"INF: infof: 42\n\"", buf.String())
	}
	buf.Reset()
	l.Infoln("infoln:", 42)
	if buf.String() != "INF: infoln: 42\n" {
		t.Errorf("write infoln line: got %q; want \"INF: infoln: 42\n\"", buf.String())
	}
	l.SetPrefix("xyz")
	p := l.Prefix()
	if p != "xyz" {
		t.Errorf("prefix: got %q; want \"xyz\"", p)
	}
	buf.Reset()
	l.Debug("debug")
	if buf.String() != "xyzDBG: debug\n" {
		t.Errorf("write debug line: %q; want \"xyzDEBUG: debugf: 42\n\"", buf.String())
	}
	buf.Reset()
	l.Debugf("debugf: %d", 42)
	if buf.String() != "xyzDBG: debugf: 42\n" {
		t.Errorf("write debug line: %q; want \"xyzDEBUG: debugf: 42\n\"", buf.String())
	}
	buf.Reset()
	l.Debugln("debugln:", 42)
	if buf.String() != "xyzDBG: debugln: 42\n" {
		t.Errorf("write debugln line: %q; want \"xyzDEBUG: debugln: 42\n\"", buf.String())
	}
}

// This also tests the package global logger
func TestSetLogLevelPrefixFlags(t *testing.T) {
	var buf bytes.Buffer
	tst := New(LogDebug, Full, &buf, "", Lshortfile)
	tst.SetFlags(0)
	f := tst.Flags()
	if f != 0 {
		t.Errorf("flags: got %d; want 0", f)
	}
	tst.SetLevel(LogDebug)
	tst.Error("error")
	if buf.String() != "ERROR: error\n" {
		t.Errorf("write error line: got %q; want \"ERROR: error\n\"", buf.String())
	}
	buf.Reset()
	tst.Errorf("errorf: %d %s", 42, "eleven")
	if buf.String() != "ERROR: errorf: 42 eleven\n" {
		t.Errorf("write errorf line: got %q; want \"ERROR: errorf: 42 eleven\n\"", buf.String())
	}
	buf.Reset()
	tst.Errorln("errorln:", 42, "eleven")
	if buf.String() != "ERROR: errorln: 42 eleven\n" {
		t.Errorf("write errorln line: got %q; want \"ERROR: errorln: 42 eleven\n\"", buf.String())
	}
	buf.Reset()
	tst.Info("info")
	if buf.String() != "INFO: info\n" {
		t.Errorf("write info line: got %q; want \"INFO: info\n\"", buf.String())
	}
	buf.Reset()
	tst.Infof("infof: %d", 42)
	if buf.String() != "INFO: infof: 42\n" {
		t.Errorf("write infof line: got %q; want \"INFO: infof: 42\n\"", buf.String())
	}
	buf.Reset()
	tst.Infoln("infoln:", 42)
	if buf.String() != "INFO: infoln: 42\n" {
		t.Errorf("write infoln line: got %q; want \"INFO: infoln: 42\n\"", buf.String())
	}
	buf.Reset()
	tst.SetPrefix("abc")
	p := tst.Prefix()
	if p != "abc" {
		t.Errorf("prefix: got %q; want \"abc\"", p)
	}
	tst.SetLevelStringType(Char)
	tst.Debug("debug")
	if buf.String() != "abcD: debug\n" {
		t.Errorf("write debug line: %q; want \"abcD: debug\n\"", buf.String())
	}
	buf.Reset()
	tst.Debugf("debugf: %d", 42)
	if buf.String() != "abcD: debugf: 42\n" {
		t.Errorf("write debugf line: %q; want \"abcD: debugf: 42\n\"", buf.String())
	}
	buf.Reset()
	tst.Debugln("debugln:", 42)
	if buf.String() != "abcD: debugln: 42\n" {
		t.Errorf("write debugln line: %q; want \"abcD: debugln: 42\n\"", buf.String())
	}
	buf.Reset()
	tst.SetLevel(LogNone)
	if tst.GetLevel() != LogNone {
		t.Errorf("logger severity level: got %s, want %s", tst.GetLevel(), LogNone)
	}
	tst.Error("error")
	if buf.Len() > 0 {
		t.Errorf("write error line: expected no bytes to be written, %d were", buf.Len())
	}
	tst.Info("info")
	if buf.Len() > 0 {
		t.Errorf("write info line: expected no bytes to be written, %d were", buf.Len())
	}
	tst.Debug("debug")
	if buf.Len() > 0 {
		t.Errorf("write debug line: expected no bytes to be written, %d were", buf.Len())
	}
}

func TestPrint(t *testing.T) {
	var buf bytes.Buffer
	s := "Time is an illusion. Lunchtime double so."
	SetOutput(&buf)
	SetLevel(LogDebug)
	l := GetLevel()
	if l != LogDebug {
		t.Errorf("level: got %s want %s", l, LogDebug)
	}
	SetLevel(LogNone)
	SetLevelStringType(Short)
	v := GetLevelStringType()
	if v != Short {
		t.Errorf("GetSetLevelStringType: got %v want %v", v, Short)
	}
	SetLevelStringType(Full)
	v = GetLevelStringType()
	if v != Full {
		t.Errorf("GetSetLevelStringType: got %v want %v", v, Full)
	}

	Print(s)
	if buf.Len() > 0 {
		t.Errorf("print: expected len to be 0; got %d", buf.Len())
	}
	Printf("%s", s)
	if buf.Len() > 0 {
		t.Errorf("printf: expected len to be 0; got %d", buf.Len())
	}
	Println(s)
	if buf.Len() > 0 {
		t.Errorf("println: expected len to be 0; got %d", buf.Len())
	}
	SetPrefix("---")
	p := Prefix()
	if p != "---" {
		t.Errorf("got %q; want \"---\"", p)
	}
	SetFlags(Lshortfile)
	f := Flags()
	if f != Lshortfile {
		t.Errorf("got %d; want %d", f, Lshortfile)
	}
	tst := New(LogDebug, Short, &buf, "", 0)
	tst.Print(s)
	if buf.String() != fmt.Sprintf("%s\n", s) {
		t.Errorf("print: got %q; want \"%s\n\"", buf.String(), s)
	}
	buf.Reset()
	tst.Printf("%s", s)
	if buf.String() != fmt.Sprintf("%s\n", s) {
		t.Errorf("printf: got %q; want \"%s\n\"", buf.String(), s)
	}
	buf.Reset()
	tst.Println(s)
	if buf.String() != fmt.Sprintf("%s\n", s) {
		t.Errorf("println: got %q; want \"%s\n\"", buf.String(), s)
	}
	buf.Reset()
}

func TestFuncs(t *testing.T) {
	var buf bytes.Buffer
	b := bufio.NewWriter(&buf)
	l := New(LogError, Full, b, "", 0)
	l.AddFunc(b.Flush)
	val := "oh no!"
	l.Error(val)
	l.Close()
	val = "ERROR: " + val + "\n"
	s := buf.String()
	if s != val {
		t.Errorf("got %s; want %s", s, val)
	}
}
