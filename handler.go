package gtb

//A Handler responds to an Update.
//ServeUpdate should write reply data to the Bot then return.
type UpdateHandler interface {
	ServeUpdate(*Bot, *Update)
}
type UpdateHandlerFunc func(*Bot, *Update)

func (f UpdateHandlerFunc) ServeUpdate(bot *Bot, upd *Update) {
	f(bot, upd)
}

// MessageHandler
type MessageHandler interface {
	ServeMessage(*Bot, *Message)
}
type MessageHandlerFunc func(*Bot, *Message)

func (f MessageHandlerFunc) ServeMessage(bot *Bot, msg *Message) {
	f(bot, msg)
}
func (f MessageHandlerFunc) ServeUpdate(bot *Bot, upd *Update) {
	if upd.Message != nil {
		f(bot, upd.Message)
	} else {
		//TODO handle better
		panic("Message handler cannot handle nil message")
	}
}

type MigrationHandler interface {
	ServeMigration(*Bot, int64, int64)
}
type MigrationHandlerFunc func(*Bot, int64, int64)

func (f MigrationHandlerFunc) ServeMigration(bot *Bot, from, to int64) {
	f(bot, from, to)
}

func (f MigrationHandlerFunc) ServeMessage(bot *Bot, msg *Message) {
	if msg.MigrateTo != 0 {
		f(bot, msg.MigrateFrom, msg.MigrateTo)
	} else {
		//TODO handle better
		panic("Message handler cannot handle nil message")
	}
}

// Callback Handler
type CallbackHandler interface {
	ServeCallback(*Bot, *Callback)
}
type CallbackHandlerFunc func(*Bot, *Callback)

func (f CallbackHandlerFunc) ServeCallback(bot *Bot, cb *Callback) {
	f(bot, cb)
}
func (f CallbackHandlerFunc) ServeUpdate(bot *Bot, upd *Update) {
	if upd.Callback != nil {
		f(bot, upd.Callback)
	} else {
		//TODO handle better
		panic("Callback handler cannot handle nil callback")
	}
}

// QueryHandler
type QueryHandler interface {
	ServeQuery(*Bot, *Query)
}
type QueryHandlerFunc func(*Bot, *Query)

func (f QueryHandlerFunc) ServeQuery(bot *Bot, q *Query) {
	f(bot, q)
}
func (f QueryHandlerFunc) ServeUpdate(bot *Bot, upd *Update) {
	if upd.Query != nil {
		f(bot, upd.Query)
	} else {
		//TODO handle better
		panic("Callback handler cannot handle nil callback")
	}
}

// ChosenInlineResultHandler
type ChosenInlineResultHandler interface {
	ServeChosenInlineResult(*Bot, *ChosenInlineResult)
}
type ChosenInlineResultHandlerFunc func(*Bot, *ChosenInlineResult)

func (f ChosenInlineResultHandlerFunc) ServeChosenInlineResult(bot *Bot, in_res *ChosenInlineResult) {
	f(bot, in_res)
}
func (f ChosenInlineResultHandlerFunc) ServeUpdate(bot *Bot, upd *Update) {
	if upd.ChosenInlineResult != nil {
		f(bot, upd.ChosenInlineResult)
	} else {
		//TODO handle better
		panic("ChosenInlineResult handler cannot handle nil chosen inline result")
	}
}
