package telegram

import (
	"log"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	cmds "gitlab.ozon.dev/DrompiX/homework-2/botman/internal/commands"
	"gitlab.ozon.dev/DrompiX/homework-2/botman/internal/states"
	tm "gitlab.ozon.dev/DrompiX/homework-2/botman/internal/termit"
	"gitlab.ozon.dev/DrompiX/homework-2/botman/internal/utils"
)

type tgBotClient struct {
	Bot      *tg.BotAPI
	Updates  tg.UpdatesChannel
	ClientDB *states.MemDB
	Termit   tm.TermitClient
}

func NewTGBotClient(token string, tc tm.TermitClient, debug bool) *tgBotClient {
	bot, err := tg.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = debug

	u := tg.NewUpdate(0)
	u.Timeout = 60

	db := states.NewMemDB()
	updates := bot.GetUpdatesChan(u)

	return &tgBotClient{Bot: bot, Updates: updates, ClientDB: db, Termit: tc}
}

func (c *tgBotClient) PollForUpdates() {
	for update := range c.Updates {
		if update.Message == nil {
			continue
		}
		
		if resp := c.processUpdate(update); resp != nil {
			for _, msg := range resp.Messages {
				tgMsg := tg.NewMessage(update.Message.Chat.ID, msg)
				c.Bot.Send(tgMsg)
			}
		}
	}
}

func (b *tgBotClient) processUpdate(u tg.Update) (resp *botResponse) {
	resp = newBotResponse()
	chatId := u.Message.Chat.ID
	chatState := b.ClientDB.GetState(chatId)

	upd, err := parseTgUpdate(&u)
	if err != nil {
		resp.addMessage(utils.ErrorToCapital(err))
		return
	}

	result, err := chatState.ProcessResponse(b.Termit, upd)
	if err != nil {
		resp.addMessage(utils.ErrorToCapital(err))
		return
	}
	
	resp.addMessage(result.Msg)
	if chatState != result.New {
		err := b.ClientDB.UpdateState(chatId, result.New)
		if err != nil {
			log.Printf("Could not update state for chatId %v\n", chatId)
			resp.addMessage("Bot error: could not save conversation state")
			return
		}
		resp.addMessage(result.New.PrepareMessage())
	}

	return
}

func parseTgUpdate(u *tg.Update) (*states.Update, error) {
	upd := states.NewUpdate(u.Message.From.ID)
	upd.Text = u.Message.Text
	c, err := parseTgCommand(u.Message)
	if err != nil {
		return upd, err
	}
	upd.Cmd = c
	return upd, nil
}

func parseTgCommand(m *tg.Message) (cmds.Command, error) {
	if !m.IsCommand() {
		return cmds.Unknown, nil
	}

	c, err := cmds.ParseCommand(m.Command())
	if err != nil {
		log.Printf("Error in command parsing: %s\n", err)
		return cmds.Unknown, err
	}
	log.Printf("Received command: %v", c)
	return c, nil
}
