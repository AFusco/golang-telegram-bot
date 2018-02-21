package gtb

//A Handler responds to an Update.
//ServeUpdate should write reply data to the Bot then return.
type UpdateHandler interface {
	ServeUpdate(*Bot, *Update)
}
type UpdateHandlerFunc func(*Bot, *Update)

func (f UpdateHandlerFunc) ServeUpdate(bot *Bot, upd *Update) {
	return f(bot, upd)
}

// MessageHandler
type MessageHandler interface {
	ServeMessage(*Bot, *Message)
}
type MessageHandlerFunc func(*Bot, *Message)

func (f MessageHandlerFunc) ServeMessage(bot *Bot, msg *Message) {
	return f(bot, msg)
}
func (f MessageHandlerFunc) ServeUpdate(bot *Bot, upd *Update) {
	if upd.Message != nil {
		return f(bot, upd.Message)
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
	return f(bot, cb)
}
func (f CallbackHandlerFunc) ServeUpdate(bot *Bot, upd *Update) {
	if upd.Callback != nil {
		return f(bot, upd.Callback)
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
	return f(bot, q)
}
func (f QueryHandlerFunc) ServeUpdate(bot *Bot, upd *Update) {
	if upd.Query != nil {
		return f(bot, upd.Query)
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

func (f ChosenInlineResultHandlerFunc) ServeChosenInlineResult(bot *Bot, in_res ChosenInlineResult) {
	return f(bot, in_res)
}
func (f ChosenInlineResultHandlerFunc) ServeUpdate(bot *Bot, upd *Update) {
	if upd.ChosenInlineResult != nil {
		return f(bot, upd.ChosenInlineResult)
	} else {
		//TODO handle better
		panic("ChosenInlineResult handler cannot handle nil chosen inline result")
	}
}
