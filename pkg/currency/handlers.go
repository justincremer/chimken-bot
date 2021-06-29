package currency

import (
	"fmt"
	"math/rand"

	"github.com/bwmarrin/discordgo"
	"time"
)

func HandleCreateAccount(s *discordgo.Session, m *discordgo.Message, b *Bank) {
	user := getUser(m)
	message := fmt.Sprintf("%s already has an account", user)
	_, err := b.CreateAccount(user)
	if err != nil {
		message = fmt.Sprintf("%s successfuly created an account", user)
	}
	s.ChannelMessageSend(m.ChannelID, message)
}

func HandleBalance(s *discordgo.Session, m *discordgo.Message, b *Bank) {
	user := getUser(m)
	account := b.getAccount(user)
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("You have a balance of %d", account.Balance))
}

func HandlePayment(s *discordgo.Session, m *discordgo.Message, b *Bank, user string, amount int) {
	if amount <= 0 {
		s.ChannelMessageSend(m.ChannelID, "Amount must be greater than 0")
		return
	}

	reciever := b.getAccount(user)
	rUser := getUser(m)
	sender := b.getAccount(rUser)

	accErr := func(user string) string {
		return fmt.Sprintf("%s does not have an account", user)
	}

	if sender == nil {
		s.ChannelMessageSend(m.ChannelID, accErr(user))
		return
	}

	if reciever == nil {
		s.ChannelMessageSend(m.ChannelID, accErr(rUser))
		return
	}

	if err := sender.Update(-amount); err != nil {
		s.ChannelMessageSend(m.ChannelID, err.Error())
		return
	}

	reciever.Update(amount)
}

func HandleGamble(s *discordgo.Session, m *discordgo.Message, b *Bank, amount int) {
	if amount <= 0 {
		s.ChannelMessageSend(m.ChannelID, "Amount must be greater than 0")
		return
	}
	user := getUser(m)
	account := b.getAccount(user)

	if account == nil {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s does not have an account", user))
		return
	}

	if account.Update(amount) != nil {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s, you cannot gamble more money than you have. smh", account.User))
		return
	}

	if coinFlip() {
		account.Update(amount * 2)
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s, won!", account.User))
		return
	}

	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s, lost :(", account.User))
	return
}

func coinFlip() bool {
	rand.Seed(time.Now().UnixNano())
	if n := rand.Intn(2); n == 0 {
		return false
	}
	return true
}

func getUser(m *discordgo.Message) string {
	return m.Author.Username
}
