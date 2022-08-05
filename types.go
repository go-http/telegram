package telegram

type User struct {
	Id                      int64
	IsBot                   bool   `json:"is_bot"`
	FirstName               string `json:"first_name"`
	LastName                string `json:"last_name"`
	Username                string
	LanguageCode            string `json:"language_code"`
	CanJoinGroups           bool   `json:"can_join_groups"`
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"`
	SupportsInlineQueries   bool   `json:"supports_inline_queries"`
}
type Chat struct {
	Id                    int64            `json:"id"`                       //Unique identifier for this chat. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
	Type                  string           `json:"type"`                     //Type of chat, can be either “private”, “group”, “supergroup” or “channel”
	Title                 string           `json:"title"`                    //Optional. Title, for supergroups, channels and group chats
	Username              string           `json:"username"`                 //Optional. Username, for private chats, supergroups and channels if available
	FirstName             string           `json:"first_name"`               //Optional. First name of the other party in a private chat
	LastName              string           `json:"last_name"`                //Optional. Last name of the other party in a private chat
	Photo                 *ChatPhoto       `json:"photo"`                    //Optional. Chat photo. Returned only in getChat.
	Bio                   string           `json:"bio"`                      //Optional. Bio of the other party in a private chat. Returned only in getChat.
	HasPrivateForwards    bool             `json:"has_private_forwards"`     //Optional. True, if privacy settings of the other party in the private chat allows to use tg://user?id=<user_id> links only in chats with the user. Returned only in getChat.
	Description           string           `json:"description"`              //Optional. Description, for groups, supergroups and channel chats. Returned only in getChat.
	InviteLink            string           `json:"invite_link"`              //Optional. Primary invite link, for groups, supergroups and channel chats. Returned only in getChat.
	PinnedMessage         Message          `json:"pinned_message"`           //Optional. The most recent pinned message (by sending date). Returned only in getChat.
	Permissions           *ChatPermissions `json:"permissions"`              //Optional. Default chat member permissions, for groups and supergroups. Returned only in getChat.
	SlowModeDelay         int              `json:"slow_mode_delay"`          //Optional. For supergroups, the minimum allowed delay between consecutive messages sent by each unpriviledged user; in seconds. Returned only in getChat.
	MessageAutoDeleteTime int              `json:"message_auto_delete_time"` //Optional. The time after which all messages sent to the chat will be automatically deleted; in seconds. Returned only in getChat.
	HasProtectedContent   bool             `json:"has_protected_content"`    //Optional. True, if messages from the chat can't be forwarded to other chats. Returned only in getChat.
	StickerSetName        string           `json:"sticker_set_name"`         //Optional. For supergroups, name of group sticker set. Returned only in getChat.
	CanSetStickerSet      bool             `json:"can_set_sticker_set"`      //Optional. True, if the bot can change the group sticker set. Returned only in getChat.
	LinkedChatId          int              `json:"linked_chat_id"`           //Optional. Unique identifier for the linked chat, i.e. the discussion group identifier for a channel and vice versa; for supergroups and channel chats. This identifier may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier. Returned only in getChat.
	Location              *ChatLocation    `json:"location"`                 //Optional. For supergroups, the location to which the supergroup is connected. Returned only in getChat.
}
type ChatPhoto interface{}
type ChatPermissions interface{}
type ChatLocation interface{}
type Animation interface{}
type Audio interface{}
type Document interface{}
type PhotoSize interface{}
type Sticker interface{}
type Video interface{}
type VideoNote interface{}
type Voice interface{}
type MessageEntity interface{}
type Contact interface{}
type Dice interface{}
type Game interface{}
type Poll interface{}
type Venue interface{}
type Location interface{}
type MessageAutoDeleteTimerChanged interface{}
type Invoice interface{}
type SuccessfulPayment interface{}
type PassportData interface{}
type ProximityAlertTriggered interface{}
type VideoChatScheduled interface{}
type VideoChatStarted interface{}
type VideoChatEnded interface{}
type VideoChatParticipantsInvited interface{}
type WebAppData interface{}
type InlineKeyboardMarkup interface{}

type Message struct {
	MessageId                     int `json:"message_id"`
	From                          *User
	SenderChat                    *Chat `json:"sender_chat"`
	Date                          int
	Chat                          *Chat
	ForwardFrom                   *User                         `json:"forwardFrom"`                       //Optional. For forwarded messages, sender of the original message
	ForwardFromChat               *Chat                         `json:"forwardFromChat"`                   //Optional. For messages forwarded from channels or from anonymous administrators, information about the original sender chat
	ForwardFromMessageId          *int                          `json:"forwardFrom_message_id"`            //Optional. For messages forwarded from channels, identifier of the original message in the channel
	ForwardSignature              string                        `json:"forward_signature"`                 //Optional. For forwarded messages that were originally sent in channels or by an anonymous chat administrator, signature of the message sender if present
	ForwardSenderName             string                        `json:"forward_sender_name"`               //Optional. Sender's name for messages forwarded from users who disallow adding a link to their account in forwarded messages
	ForwardDate                   *int                          `json:"forward_date"`                      //Optional. For forwarded messages, date the original message was sent in Unix time
	IsAutomaticForward            bool                          `json:"is_automatic_forward"`              //Optional. bool, if the message is a channel post that was automatically forwarded to the connected discussion group
	ReplyToMessage                *Message                      `json:"reply_to_message"`                  //Optional. For replies, the original message. Note that the Message object in this field will not contain further reply_to_message fields even if it itself is a reply.
	ViaBot                        *User                         `json:"via_bot"`                           //Optional. Bot through which the message was sent
	EditDate                      *int                          `json:"edit_date"`                         //Optional. Date the message was last edited in Unix time
	HasProtectedContent           bool                          `json:"has_protected_content"`             //Optional. bool, if the message can't be forwarded
	MediaGroupId                  string                        `json:"media_group_id"`                    //Optional. The unique identifier of a media message group this message belongs to
	AuthorSignature               string                        `json:"author_signature"`                  //Optional. Signature of the post author for messages in channels, or the custom title of an anonymous group administrator
	Text                          string                        `json:"text"`                              //Optional. For text messages, the actual UTF-8 text of the message, 0-4096 characters
	Entities                      *[]MessageEntity              `json:"entities"`                          //Optional. For text messages, special entities like usernames, URLs, bot commands, etc. that appear in the text
	Animation                     *Animation                    `json:"animation"`                         //Optional. Message is an animation, information about the animation. For backward compatibility, when this field is set, the document field will also be set
	Audio                         *Audio                        `json:"audio"`                             //Optional. Message is an audio file, information about the file
	Document                      *Document                     `json:"document"`                          //Optional. Message is a general file, information about the file
	Photo                         *[]PhotoSize                  `json:"photo"`                             //Optional. Message is a photo, available sizes of the photo
	Sticker                       *Sticker                      `json:"sticker"`                           //Optional. Message is a sticker, information about the sticker
	Video                         *Video                        `json:"video"`                             //Optional. Message is a video, information about the video
	VideoNote                     *VideoNote                    `json:"video_note"`                        //Optional. Message is a video note, information about the video message
	Voice                         *Voice                        `json:"voice"`                             //Optional. Message is a voice message, information about the file
	Caption                       string                        `json:"caption"`                           //Optional. Caption for the animation, audio, document, photo, video or voice, 0-1024 characters
	CaptionEntities               *[]MessageEntity              `json:"caption_entities"`                  //Optional. For messages with a caption, special entities like usernames, URLs, bot commands, etc. that appear in the caption
	Contact                       *Contact                      `json:"contact"`                           //Optional. Message is a shared contact, information about the contact
	Dice                          *Dice                         `json:"dice"`                              //Optional. Message is a dice with random value
	Game                          *Game                         `json:"game"`                              //Optional. Message is a game, information about the game. More about games »
	Poll                          *Poll                         `json:"poll"`                              //Optional. Message is a native poll, information about the poll
	Venue                         *Venue                        `json:"venue"`                             //Optional. Message is a venue, information about the venue. For backward compatibility, when this field is set, the location field will also be set
	Location                      *Location                     `json:"location"`                          //Optional. Message is a shared location, information about the location
	NewChatMembers                *[]User                       `json:"newChat_members"`                   //Optional. New members that were added to the group or supergroup and information about them (the bot itself may be one of these members)
	LeftChatMember                *User                         `json:"leftChat_member"`                   //Optional. A member was removed from the group, information about them (this member may be the bot itself)
	NewChatTitle                  string                        `json:"newChat_title"`                     //Optional. A chat title was changed to this value
	NewChatPhoto                  *[]PhotoSize                  `json:"newChatPhoto"`                      //Optional. A chat photo was change to this value
	DeleteChatPhoto               bool                          `json:"deleteChatPhoto"`                   //Optional. Service message: the chat photo was deleted
	GroupChatCreated              bool                          `json:"groupChat_created"`                 //Optional. Service message: the group has been created
	SupergroupChatCreated         bool                          `json:"supergroupChat_created"`            //Optional. Service message: the supergroup has been created. This field can't be received in a message coming through updates, because bot can't be a member of a supergroup when it is created. It can only be found in replyToMessage if someone replies to a very first message in a directly created supergroup.
	ChannelChatCreated            bool                          `json:"channelChat_created"`               //Optional. Service message: the channel has been created. This field can't be received in a message coming through updates, because bot can't be a member of a channel when it is created. It can only be found in replyToMessage if someone replies to a very first message in a channel.
	MessageAutoDeleteTimerChanged MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed"` //Optional. Service message: auto-delete timer settings changed in the chat
	MigrateToChatId               *int                          `json:"migrate_toChat_id"`                 //Optional. The group has been migrated to a supergroup with the specified identifier. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
	MigrateFromChatId             *int                          `json:"migrateFromChat_id"`                //Optional. The supergroup has been migrated from a group with the specified identifier. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
	PinnedMessage                 *Message                      `json:"pinned_message"`                    //Optional. Specified message was pinned. Note that the Message object in this field will not contain further replyTo_message fields even if it is itself a reply.
	Invoice                       *Invoice                      `json:"invoice"`                           //Optional. Message is an invoice for a payment, information about the invoice. More about payments »
	SuccessfulPayment             *SuccessfulPayment            `json:"successful_payment"`                //Optional. Message is a service message about a successful payment, information about the payment. More about payments »
	ConnectedWebsite              string                        `json:"connected_website"`                 //Optional. The domain name of the website on which the user has logged in. More about Telegram Login »
	PassportData                  *PassportData                 `json:"passport_data"`                     //Optional. Telegram Passport data
	ProximityAlertTriggered       *ProximityAlertTriggered      `json:"proximityAlert_triggered"`          //Optional. Service message. A user in the chat triggered another user's proximity alert while sharing Live Location.
	VideoChatScheduled            *VideoChatScheduled           `json:"videoChatScheduled"`                //Optional. Service message: video chat scheduled
	VideoChatStarted              *VideoChatStarted             `json:"videoChatStarted"`                  //Optional. Service message: video chat started
	VideoChatEnded                *VideoChatEnded               `json:"videoChatEnded"`                    //Optional. Service message: video chat ended
	VideoChatParticipantsInvited  VideoChatParticipantsInvited  `json:"videoChatParticipants_invited"`     //Optional. Service message: new participants invited to a video chat
	WebAppData                    *WebAppData                   `json:"web_app_data"`                      //Optional. Service message: data sent by a Web App
	ReplyMarkup                   *InlineKeyboardMarkup         `json:"reply_markup"`                      //Optional. Inline keyboard attached to the message. login_url buttons are represented as ordinary url buttons.
}

type Update struct {
	UpdateId           int `json:"update_id"`
	Message            *Message
	EditedMessage      *Message    `json:"edited_message"`
	ChannelPost        *Message    `json:"channel_post"`
	EditedChannelPost  *Message    `json:"edited_channel_post"`
	InlineQuery        interface{} `json:"inline_query"`
	ChosenInlineResult interface{} `json:"chosen_inline_result"`
	CallbackQuery      interface{} `json:"callback_query"`
	ShippingQuery      interface{} `json:"shipping_query"`
	PreCheckoutQuery   interface{} `json:"pre_checkout_query"`
	Poll               interface{} `json:"poll"`
	PollAnswer         interface{} `json:"poll_answer"`
	MyChatMember       interface{} `json:"my_chat_member"`
	ChatMember         interface{} `json:"chat_member"`
	ChatJoinRequest    interface{} `json:"chat_join_request"`
}
