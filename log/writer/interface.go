package writer

type MessageWriter interface {
	Write(level, msg string) error
}
