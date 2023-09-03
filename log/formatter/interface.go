package formatter

import "github.com/asstroneer/go-log/log/message"

type MessageFormatter interface {
	Format(message *message.LogMessage) (string, error)
}
