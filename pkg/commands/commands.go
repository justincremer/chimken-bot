package commands

import (
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func ExecuteCommand(s *discordgo.Session, m *discordgo.Message, t0 time.Time) {
	full := m.Content[1:]
	args := strings.Split(strings.TrimSpace(full), " ")
	cmd := args[0]
	// if len(msg) > 2 {
	// 	msg = strings.Split(strings.Split(m.Content, " ")[0], "!")[1]
	// }

	switch cmd {
	case "info":
		HandleInfoCommand(s, m, t0)
	case "help":
		HandleHelpCommand(s, m)
	case "whois":
		name := args[1]
		switch name {
		case "sophie":
			HandlePesonalMessage(s, m, name)
		case "justin":
			HandlePesonalMessage(s, m, name)
		case "liana":
			HandlePesonalMessage(s, m, name)
		case "sunny":
			HandlePesonalMessage(s, m, name)
		case "angela":
			HandlePesonalMessage(s, m, name)
		case "paul":
			HandlePesonalMessage(s, m, name)
		case "joseph":
			HandlePesonalMessage(s, m, name)
		case "siah":
			HandlePesonalMessage(s, m, name)
		case "fluzz":
			HandlePesonalMessage(s, m, name)
		case "kreiker":
			HandlePesonalMessage(s, m, name)
		default:
			HandleUnknownCommand(s, m, full)
		}
	default:
		HandleUnknownCommand(s, m, full)
	}
}

// case "strawpoll":
// 	titleBegin := strings.Index(m.Content, "{")
// 	titleEnd := strings.Index(m.Content, "}")
// 	if titleBegin == -1 || titleEnd == -1 {
// 		return
// 	}
// 	pollTitle := string(m.Content[titleBegin+1 : titleEnd])

// 	strpos := 0
// 	pollOptionsData := m.Content[titleEnd+1:]
// 	var pollOptions []string
// 	for strpos < len(pollOptionsData) {
// 		test := pollOptionsData[strpos:]
// 		beginPos := strings.Index(test, "[")
// 		endPos := strings.Index(test, "]")
// 		if beginPos == -1 || endPos == -1 || len(test[beginPos:endPos]) <= 0 {
// 			break
// 		}

// 		pollOptions = append(pollOptions, test[beginPos+1:endPos])
// 		strpos = strpos + (endPos + 1)
// 	}

// 	if len(pollOptions) < 2 {
// 		return
// 	}
// 	HandleStrawPollCommand(s, m, pollTitle, pollOptions)
