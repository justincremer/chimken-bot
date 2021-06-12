package commands

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/justincremer/chimken-bot/pkg/logger"
	"strings"
)

func HandleInfoCommand(s *discordgo.Session, m *discordgo.Message, t0 time.Time) {
	t1 := time.Now()
	channel, err := s.Channel(m.ChannelID)
	logger.Must("Unknown channel error: ", err)

	title := "ChimkenBot Info Panel"
	channelName := channel.Name
	message := "```txt\n%s\n%s\n%-16s%-20s\n%-16s%-20s\n%-16s%-20s```"
	message = fmt.Sprintf(message, title, strings.Repeat("-", len(title)), "ChannelID", m.ChannelID, "Channel Name", channelName, "Uptime", (t1.Sub(t0).String()))
	s.ChannelMessageSend(m.ChannelID, message)
}

func HandleHelpCommand(s *discordgo.Session, m *discordgo.Message) {
	message := fmt.Sprintf("```txt%s```", strings.Repeat("\n%s : %s", 5))
	message = fmt.Sprintf(message, "info", "listen to chimken bot talk about her life", "help", "where you are now", "brib", "birbie:)", "paul", "who is sweet potato", "poll", "work the polls")
	s.ChannelMessageSend(m.ChannelID, message)
}

func HandleSunnyCommand(s *discordgo.Session, m *discordgo.Message) {
	s.ChannelMessageSend(m.ChannelID, "Sunnu nation must rise\n┻━┻ ︵ ＼(’0’)/／ ︵ ┻━┻")
}

func HandleLianaCommand(s *discordgo.Session, m *discordgo.Message) {
	s.ChannelMessageSend(m.ChannelID, "WHEN\nWENH\nWHEN YOU\nWHEN OU\nWHEN\nwHEN YOU")
}

func HandleAngelaCommand(s *discordgo.Session, m *discordgo.Message) {
	s.ChannelMessageSend(m.ChannelID, "S tier troglodite")
}

func HandleJosephCommand(s *discordgo.Session, m *discordgo.Message) {
	s.ChannelMessageSend(m.ChannelID, "Da bling")
}

func HandleUnknownCommand(s *discordgo.Session, m *discordgo.Message, msg string) {
	c, err := s.UserChannelCreate(m.Author.ID)
	logger.Must("Unknown command error: ", err)

	s.ChannelMessageSend(c.ID, "The command \""+msg+"\" in not recognized")
}
