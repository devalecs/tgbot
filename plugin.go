package hugo

import "gopkg.in/devalecs/tgo.v1"

const (
	PluginTagAdmin PluginTag = "admin"
)

type Plugin interface {
	Reply(b B, update *tgo.Update)
	Match(b B, update *tgo.Update) bool
	Settings() PluginSettings
}

type PluginSettings struct {
	Tag  PluginTag
	Help string
}

type PluginTag string

type B interface {
	GetUpdatesChan(params tgo.GetUpdatesParams) chan *tgo.Update
	GetUpdates(params tgo.GetUpdatesParams) (*[]tgo.Update, error)
	GetMe() (*tgo.User, error)
	SendMessage(params tgo.SendMessageParams) (*tgo.Message, error)
	ForwardMessage(params tgo.ForwardMessageParams) (*tgo.Message, error)
	SendPhoto(params tgo.SendPhotoParams) (*tgo.Message, error)
	SendAudio(params tgo.SendAudioParams) (*tgo.Message, error)
	SendDocument(params tgo.SendDocumentParams) (*tgo.Message, error)
	SendSticker(params tgo.SendStickerParams) (*tgo.Message, error)
	SendVideo(params tgo.SendVideoParams) (*tgo.Message, error)
	SendVoice(params tgo.SendVoiceParams) (*tgo.Message, error)
	SendLocation(params tgo.SendLocationParams) (*tgo.Message, error)
	SendVenue(params tgo.SendVenueParams) (*tgo.Message, error)
	SendContact(params tgo.SendContactParams) (*tgo.Message, error)
	SendChatAction(params tgo.SendChatActionParams) (*bool, error)
	GetUserProfilePhotos(params tgo.GetUserProfilePhotosParams) (*tgo.UserProfilePhotos, error)
	KickChatMember(params tgo.KickChatMemberParams) (*bool, error)
	LeaveChat(params tgo.LeaveChatParams) (*bool, error)
	UnbanChatMember(params tgo.UnbanChatMemberParams) (*bool, error)
	GetChat(params tgo.GetChatParams) (*tgo.Chat, error)
	GetChatAdministrators(params tgo.GetChatParams) (*[]tgo.ChatMember, error)
	GetChatMembersCount(params tgo.GetChatMembersCountParams) (*uint, error)
	GetChatMember(params tgo.GetChatMemberParams) (*tgo.ChatMember, error)
	AnswerCallbackQuery(params tgo.AnswerCallbackQueryParams) (*bool, error)
	AnswerInlineQuery(params tgo.AnswerInlineQueryParams) (*bool, error)

	MessageTextMatch(update *tgo.Update, pattern string) bool
	InlineQueryMatch(update *tgo.Update, pattern string) bool
	MessageTextMatchParams(update *tgo.Update, pattern string) (bool, []interface{})
	InlineQueryMatchParams(update *tgo.Update, pattern string) (bool, []interface{})

	UserIsAdmin(user tgo.User) bool
	GetPlugins() []Plugin
}
