package logger

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"time"
)

// Constants for the different colors. See more on https://en.wikipedia.org/wiki/ANSI_escape_code
const (
	Black = (iota + 30)
	Red
	Green
	Yellow
	Blue
	Pink
	Cyan
	White
	_
	_
	ReverseBlack
	ReverseRed
	ReverseGreen
	ReverseYellow
	ReverseBlue
	ReversePink
	ReverseCyan
	ReverseWhite
	_
	ResetBack
)

// Constants for bright variations of the colors
const (
	Grey = (iota + 90)
	BrightRed
	BrightGreen
	BrightYellow
	BrightBlue
	BrightPink
	BrightCyan
	BrightWhite
	_
	_
	ReverseGrey
	ReverseBrightRed
	ReverseBrightGreen
	ReverseBrightYellow
	ReverseBrightBlue
	ReverseBrightPink
	ReverseBrightCyan
	ReverseBrightWhite
)

// Constants for special ANSI characters. They may not be widely accepted in all terminals.
const (
	Reset = iota
	Bold
	Faint
	Italic
	Underline
	_
	_
	ReverseVideo
	Concealed
	CrossedOut
)

// Constants for format strings used all over the package
const (
	ansiColorCode = "\033[%dm"
	defaultFormat = "> #%d %s ⇒  %s > %s\n"
	stackFormat   = "> #%d %s ⇒  ON %s:%s ⇒  %s > %s\n"
)

// Logger represents a logger instance. It can be used to log messages. It receives some options to customize the logger.
// shouldColor: if true, the logger will print messages with colors.
// shouldDate: if true, the logger will print messages with date.
// shouldStack: if true, the logger will print messages with a stack trace (file and line).
// timeFormat: the format of the date.
// out: the output stream.
type Logger struct {
	shouldColor bool
	shouldDate  bool
	shouldStack bool
	out         io.Writer
	timeFormat  string
	id          uint64
	richOutput  bool
}

// Message represents a log message. It contains the message type, the message text, the message color, id
// and the file and line where the message was created.
type Message struct {
	Text  string
	Color int
	File  string
	Line  string
	Id    uint64
	Type  string
}

// Parse returns an string that is ready to print from the received Message.
func (m *Message) Parse(shouldStack, shouldColor, richOutput bool, timeFormat string) string {
	date := time.Now().Format(timeFormat)
	message := m.Text + "\n"
	if richOutput {
		message = fmt.Sprintf(defaultFormat, m.Id, date, m.Type, m.Text)
		if shouldStack {
			message = fmt.Sprintf(stackFormat, m.Id, date, m.File, m.Line, m.Type, m.Text)
		}
	}
	if shouldColor {
		color := fmt.Sprintf(ansiColorCode, m.Color)
		reset := fmt.Sprintf(ansiColorCode, Reset)
		resetBack := fmt.Sprintf(ansiColorCode, ResetBack)
		message = fmt.Sprintf("%s%s%s%s", color, message, resetBack, reset)
	}
	return message
}

// New returns a new Logger instance with received options.
func New(shouldColor, shouldDate, shouldStack, allowData bool, out io.Writer, timeFormat string) (*Logger, error) {
	return &Logger{
		shouldColor: shouldColor && !isWindows(),
		shouldDate:  shouldDate,
		shouldStack: shouldStack,
		out:         out,
		timeFormat:  timeFormat,
		id:          1,
		richOutput:  allowData,
	}, nil
}

// isWindows returns true if the current OS is Windows.
func isWindows() bool {
	return runtime.GOOS == "windows"
}

// getLineAndFile returns the line and file where the function was called.
func getLineAndFile(shouldStack bool) (string, string) {
	if shouldStack {
		_, file, line, _ := runtime.Caller(2)
		return fmt.Sprint(line), file
	}
	return "", ""
}

// incrementId increments the message counter id of the logger.
func (l *Logger) incrementId() {
	l.id++
}

// SetShouldColor sets the shouldColor field of the logger.
func (l *Logger) SetShouldColor(shouldColor bool) {
	l.shouldColor = shouldColor && !isWindows()
}

// SetShouldDate sets the shouldDate field of the logger.
func (l *Logger) SetShouldDate(shouldDate bool) {
	l.shouldDate = shouldDate
}

// SetStack sets the shouldStack field of the logger.
func (l *Logger) SetShouldStack(shouldStack bool) {
	l.shouldStack = shouldStack
}

// SetTimeFormat sets the time format of the logger.
func (l *Logger) SetTimeFormat(timeFormat string) {
	l.timeFormat = timeFormat
}

// SetOut sets the output stream of the logger.
func (l *Logger) SetOut(out io.Writer) {
	l.out = out
}

// SetRichOutput sets the richOutput field of the logger.
func (l *Logger) SetRichOutput(allowData bool) {
	l.richOutput = allowData
}

// print is the internal function that prints the message. It receives the message and if the message is a Debug message.
// If the message is a Debug message, it will always print the stack trace.
func (l *Logger) print(message Message, isDebug bool) {
	l.out.Write([]byte(message.Parse(l.shouldStack || isDebug, l.shouldColor, l.richOutput, l.timeFormat)))
	l.incrementId()
}

// Print prints a message in INFO mode. It receives the message to print. It prints the message in WHITE color.
func (l *Logger) Print(input string) {
	line, file := getLineAndFile(l.shouldStack)
	l.print(Message{
		Text:  input,
		Color: White,
		Type:  "INF",
		Id:    l.id,
		Line:  line,
		File:  file,
	}, false)
}

// Printf prints a message in INFO mode. It receives the message to print. It prints the message in WHITE color. It will
// format the message in the same way as fmt.Printf
func (l *Logger) Printf(format string, args ...interface{}) {
	l.Print(fmt.Sprintf(format, args...))
}

// Info prints a message in INFO mode. It receives the message to print. It prints the message in WHITE color.
func (l *Logger) Info(input string) {
	l.Print(input)
}

// Infof prints a message in INFO mode. It receives the message to print. It prints the message in WHITE color. It will
// format the message in the same way as fmt.Printf
func (l *Logger) Infof(format string, args ...interface{}) {
	l.Printf(format, args...)
}

// Success prints a message in SUCCESS mode. It receives the message to print. It prints the message in GREEN color.
func (l *Logger) Success(input string) {
	line, file := getLineAndFile(l.shouldStack)
	l.print(Message{
		Text:  input,
		Color: Green,
		Type:  "SUC",
		Id:    l.id,
		Line:  line,
		File:  file,
	}, false)
}

// Successf prints a message in SUCCESS mode. It receives the message to print. It prints the message in GREEN color.
// It will format the message in the same way as fmt.Printf
func (l *Logger) Successf(format string, args ...interface{}) {
	l.Success(fmt.Sprintf(format, args...))
}

// Warning prints a message in WARNING mode. It receives the message to print. It prints the message in YELLOW color.
func (l *Logger) Warning(input string) {
	line, file := getLineAndFile(l.shouldStack)
	l.print(Message{
		Text:  input,
		Color: Yellow,
		Type:  "WAR",
		Id:    l.id,
		Line:  line,
		File:  file,
	}, false)
}

// WarningF prints a message in WARNING mode. It receives the message to print. It prints the message in YELLOW color.
// It will format the message in the same way as fmt.Printf
func (l *Logger) Warningf(format string, args ...interface{}) {
	l.Warning(fmt.Sprintf(format, args...))
}

// Error prints a message in ERROR mode. It receives the message to print. It prints the message in RED color.
func (l *Logger) Error(input string) {
	line, file := getLineAndFile(l.shouldStack)
	l.print(Message{
		Text:  input,
		Color: Red,
		Type:  "ERR",
		Id:    l.id,
		Line:  line,
		File:  file,
	}, false)
}

// Errorf prints a message in ERROR mode. It receives the message to print. It prints the message in RED color. It will
// format the message in the same way as fmt.Printf
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.Error(fmt.Sprintf(format, args...))
}

// Fatal works in the same way as Error, but it will exit the program with an error code.
func (l *Logger) Fatal(input string) {
	l.Error(input)
	os.Exit(1)
}

// Fatalf works in the same way as Error, but it will exit the program with an error code. It will format the message
// in the same way as fmt.Printf
func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.Fatal(fmt.Sprintf(format, args...))
}

// Critical prints a message in CRITICAL mode. It receives the message to print. It prints the message in PINK color.
func (l *Logger) Critical(input string) {
	line, file := getLineAndFile(l.shouldStack)
	l.print(Message{
		Text:  input,
		Color: Pink,
		Type:  "CRT",
		Id:    l.id,
		Line:  line,
		File:  file,
	}, false)
}

// CriticalF prints a message in CRITICAL mode. It receives the message to print. It prints the message in PINK color.
// It will format the message in the same way as fmt.Printf
func (l *Logger) Criticalf(format string, args ...interface{}) {
	l.Critical(fmt.Sprintf(format, args...))
}

// Debug prints a message in DEBUG mode. It receives the message to print. It prints the message in BLUE color.
func (l *Logger) Debug(input string) {
	line, file := getLineAndFile(true)
	l.print(Message{
		Text:  input,
		Color: Blue,
		Type:  "DBG",
		Id:    l.id,
		Line:  line,
		File:  file,
	}, true)
}

// Debugf prints a message in DEBUG mode. It receives the message to print. It prints the message in BLUE color. It will
// format the message in the same way as fmt.Printf
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.Debug(fmt.Sprintf(format, args...))
}

// CustomAnsiPrint prints a message in a custom color. It receives the message to print and the color to use.
func (l *Logger) CustomAnsiPrint(input string, color int) {
	line, file := getLineAndFile(false)
	l.print(Message{
		Text:  input,
		Color: color,
		Type:  "CUS",
		Id:    l.id,
		Line:  line,
		File:  file,
	}, false)
}

// CustomAnsiPrintf prints a message in a custom color. It receives the message to print and the color to use. It will
// format the message in the same way as fmt.Printf
func (l *Logger) CustomAnsiPrintf(format string, color int, args ...interface{}) {
	l.CustomAnsiPrint(fmt.Sprintf(format, args...), color)
}
