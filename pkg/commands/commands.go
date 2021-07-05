package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/justincremer/chimkin-bot/pkg/currency"
	"strconv"
	"strings"
	"time"
)

func ExecuteCommand(s *discordgo.Session, m *discordgo.Message, t0 time.Time, bank *currency.Bank) {
	full := m.Content[1:]
	args := strings.Split(strings.TrimSpace(full), " ")
	cmd := args[0]

	switch cmd {
	case "info":
		HandleInfo(s, m, t0)
	case "help":
		HandleHelp(s, m)
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
			HandleUnknown(s, m, full)
		}
	case "monies":
		cmd = args[1]
		switch cmd {
		case "new":
			currency.HandleCreateAccount(s, m, bank)
		case "balance":
			currency.HandleBalance(s, m, bank)
		case "pay":
			user := m.Author.Username
			n, err := strconv.Atoi(args[2])
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s is not a valid number", args[2]))
				return
			}
			currency.HandlePayment(s, m, bank, user, n)
		case "gamble":
			n, err := strconv.Atoi(args[2])
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s is not a valid number", args[2]))
				return
			}
			currency.HandleGamble(s, m, bank, n)
		default:
			HandleUnknown(s, m, full)
		}
	default:
		HandleUnknown(s, m, full)
	}
}
