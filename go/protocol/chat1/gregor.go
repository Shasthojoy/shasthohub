// Auto-generated by avdl-compiler v1.3.22 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/chat1/gregor.avdl

package chat1

import (
	gregor1 "github.com/keybase/client/go/protocol/gregor1"
	keybase1 "github.com/keybase/client/go/protocol/keybase1"
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
)

type GenericPayload struct {
	Action       string         `codec:"Action" json:"Action"`
	InboxVers    InboxVers      `codec:"inboxVers" json:"inboxVers"`
	ConvID       ConversationID `codec:"convID" json:"convID"`
	UnreadUpdate *UnreadUpdate  `codec:"unreadUpdate,omitempty" json:"unreadUpdate,omitempty"`
}

func (o GenericPayload) DeepCopy() GenericPayload {
	return GenericPayload{
		Action:    o.Action,
		InboxVers: o.InboxVers.DeepCopy(),
		ConvID:    o.ConvID.DeepCopy(),
		UnreadUpdate: (func(x *UnreadUpdate) *UnreadUpdate {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.UnreadUpdate),
	}
}

type NewConversationPayload struct {
	Action       string         `codec:"Action" json:"Action"`
	ConvID       ConversationID `codec:"convID" json:"convID"`
	InboxVers    InboxVers      `codec:"inboxVers" json:"inboxVers"`
	UnreadUpdate *UnreadUpdate  `codec:"unreadUpdate,omitempty" json:"unreadUpdate,omitempty"`
}

func (o NewConversationPayload) DeepCopy() NewConversationPayload {
	return NewConversationPayload{
		Action:    o.Action,
		ConvID:    o.ConvID.DeepCopy(),
		InboxVers: o.InboxVers.DeepCopy(),
		UnreadUpdate: (func(x *UnreadUpdate) *UnreadUpdate {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.UnreadUpdate),
	}
}

type NewMessagePayload struct {
	Action       string           `codec:"Action" json:"Action"`
	ConvID       ConversationID   `codec:"convID" json:"convID"`
	Message      MessageBoxed     `codec:"message" json:"message"`
	InboxVers    InboxVers        `codec:"inboxVers" json:"inboxVers"`
	UnreadUpdate *UnreadUpdate    `codec:"unreadUpdate,omitempty" json:"unreadUpdate,omitempty"`
	MaxMsgs      []MessageSummary `codec:"maxMsgs" json:"maxMsgs"`
}

func (o NewMessagePayload) DeepCopy() NewMessagePayload {
	return NewMessagePayload{
		Action:    o.Action,
		ConvID:    o.ConvID.DeepCopy(),
		Message:   o.Message.DeepCopy(),
		InboxVers: o.InboxVers.DeepCopy(),
		UnreadUpdate: (func(x *UnreadUpdate) *UnreadUpdate {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.UnreadUpdate),
		MaxMsgs: (func(x []MessageSummary) []MessageSummary {
			if x == nil {
				return nil
			}
			var ret []MessageSummary
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.MaxMsgs),
	}
}

type ReadMessagePayload struct {
	Action       string         `codec:"Action" json:"Action"`
	ConvID       ConversationID `codec:"convID" json:"convID"`
	MsgID        MessageID      `codec:"msgID" json:"msgID"`
	InboxVers    InboxVers      `codec:"inboxVers" json:"inboxVers"`
	UnreadUpdate *UnreadUpdate  `codec:"unreadUpdate,omitempty" json:"unreadUpdate,omitempty"`
}

func (o ReadMessagePayload) DeepCopy() ReadMessagePayload {
	return ReadMessagePayload{
		Action:    o.Action,
		ConvID:    o.ConvID.DeepCopy(),
		MsgID:     o.MsgID.DeepCopy(),
		InboxVers: o.InboxVers.DeepCopy(),
		UnreadUpdate: (func(x *UnreadUpdate) *UnreadUpdate {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.UnreadUpdate),
	}
}

type SetStatusPayload struct {
	Action       string             `codec:"Action" json:"Action"`
	ConvID       ConversationID     `codec:"convID" json:"convID"`
	Status       ConversationStatus `codec:"status" json:"status"`
	InboxVers    InboxVers          `codec:"inboxVers" json:"inboxVers"`
	UnreadUpdate *UnreadUpdate      `codec:"unreadUpdate,omitempty" json:"unreadUpdate,omitempty"`
}

func (o SetStatusPayload) DeepCopy() SetStatusPayload {
	return SetStatusPayload{
		Action:    o.Action,
		ConvID:    o.ConvID.DeepCopy(),
		Status:    o.Status.DeepCopy(),
		InboxVers: o.InboxVers.DeepCopy(),
		UnreadUpdate: (func(x *UnreadUpdate) *UnreadUpdate {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.UnreadUpdate),
	}
}

type TeamTypePayload struct {
	Action       string         `codec:"Action" json:"Action"`
	ConvID       ConversationID `codec:"convID" json:"convID"`
	TeamType     TeamType       `codec:"teamType" json:"teamType"`
	InboxVers    InboxVers      `codec:"inboxVers" json:"inboxVers"`
	UnreadUpdate *UnreadUpdate  `codec:"unreadUpdate,omitempty" json:"unreadUpdate,omitempty"`
}

func (o TeamTypePayload) DeepCopy() TeamTypePayload {
	return TeamTypePayload{
		Action:    o.Action,
		ConvID:    o.ConvID.DeepCopy(),
		TeamType:  o.TeamType.DeepCopy(),
		InboxVers: o.InboxVers.DeepCopy(),
		UnreadUpdate: (func(x *UnreadUpdate) *UnreadUpdate {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.UnreadUpdate),
	}
}

type SetAppNotificationSettingsPayload struct {
	Action       string                       `codec:"Action" json:"Action"`
	ConvID       ConversationID               `codec:"convID" json:"convID"`
	InboxVers    InboxVers                    `codec:"inboxVers" json:"inboxVers"`
	Settings     ConversationNotificationInfo `codec:"settings" json:"settings"`
	UnreadUpdate *UnreadUpdate                `codec:"unreadUpdate,omitempty" json:"unreadUpdate,omitempty"`
}

func (o SetAppNotificationSettingsPayload) DeepCopy() SetAppNotificationSettingsPayload {
	return SetAppNotificationSettingsPayload{
		Action:    o.Action,
		ConvID:    o.ConvID.DeepCopy(),
		InboxVers: o.InboxVers.DeepCopy(),
		Settings:  o.Settings.DeepCopy(),
		UnreadUpdate: (func(x *UnreadUpdate) *UnreadUpdate {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.UnreadUpdate),
	}
}

type UnreadUpdate struct {
	ConvID                  ConversationID              `codec:"convID" json:"convID"`
	UnreadMessages          int                         `codec:"unreadMessages" json:"unreadMessages"`
	UnreadNotifyingMessages map[keybase1.DeviceType]int `codec:"unreadNotifyingMessages" json:"unreadNotifyingMessages"`
	CompatUnreadMessages    int                         `codec:"UnreadMessages" json:"UnreadMessages"`
	Diff                    bool                        `codec:"diff" json:"diff"`
}

func (o UnreadUpdate) DeepCopy() UnreadUpdate {
	return UnreadUpdate{
		ConvID:         o.ConvID.DeepCopy(),
		UnreadMessages: o.UnreadMessages,
		UnreadNotifyingMessages: (func(x map[keybase1.DeviceType]int) map[keybase1.DeviceType]int {
			if x == nil {
				return nil
			}
			ret := make(map[keybase1.DeviceType]int)
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v
				ret[kCopy] = vCopy
			}
			return ret
		})(o.UnreadNotifyingMessages),
		CompatUnreadMessages: o.CompatUnreadMessages,
		Diff:                 o.Diff,
	}
}

type TLFFinalizeUpdate struct {
	FinalizeInfo ConversationFinalizeInfo `codec:"finalizeInfo" json:"finalizeInfo"`
	ConvIDs      []ConversationID         `codec:"convIDs" json:"convIDs"`
	InboxVers    InboxVers                `codec:"inboxVers" json:"inboxVers"`
}

func (o TLFFinalizeUpdate) DeepCopy() TLFFinalizeUpdate {
	return TLFFinalizeUpdate{
		FinalizeInfo: o.FinalizeInfo.DeepCopy(),
		ConvIDs: (func(x []ConversationID) []ConversationID {
			if x == nil {
				return nil
			}
			var ret []ConversationID
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.ConvIDs),
		InboxVers: o.InboxVers.DeepCopy(),
	}
}

type TLFResolveUpdate struct {
	ConvID    ConversationID `codec:"convID" json:"convID"`
	InboxVers InboxVers      `codec:"inboxVers" json:"inboxVers"`
}

func (o TLFResolveUpdate) DeepCopy() TLFResolveUpdate {
	return TLFResolveUpdate{
		ConvID:    o.ConvID.DeepCopy(),
		InboxVers: o.InboxVers.DeepCopy(),
	}
}

type RemoteUserTypingUpdate struct {
	Uid      gregor1.UID      `codec:"uid" json:"uid"`
	DeviceID gregor1.DeviceID `codec:"deviceID" json:"deviceID"`
	ConvID   ConversationID   `codec:"convID" json:"convID"`
	Typing   bool             `codec:"typing" json:"typing"`
}

func (o RemoteUserTypingUpdate) DeepCopy() RemoteUserTypingUpdate {
	return RemoteUserTypingUpdate{
		Uid:      o.Uid.DeepCopy(),
		DeviceID: o.DeviceID.DeepCopy(),
		ConvID:   o.ConvID.DeepCopy(),
		Typing:   o.Typing,
	}
}

type UpdateConversationMembership struct {
	InboxVers     InboxVers            `codec:"inboxVers" json:"inboxVers"`
	Joined        []ConversationMember `codec:"joined" json:"joined"`
	Removed       []ConversationMember `codec:"removed" json:"removed"`
	Reset         []ConversationMember `codec:"reset" json:"reset"`
	Previewed     []ConversationID     `codec:"previewed" json:"previewed"`
	UnreadUpdate  *UnreadUpdate        `codec:"unreadUpdate,omitempty" json:"unreadUpdate,omitempty"`
	UnreadUpdates []UnreadUpdate       `codec:"unreadUpdates" json:"unreadUpdates"`
}

func (o UpdateConversationMembership) DeepCopy() UpdateConversationMembership {
	return UpdateConversationMembership{
		InboxVers: o.InboxVers.DeepCopy(),
		Joined: (func(x []ConversationMember) []ConversationMember {
			if x == nil {
				return nil
			}
			var ret []ConversationMember
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Joined),
		Removed: (func(x []ConversationMember) []ConversationMember {
			if x == nil {
				return nil
			}
			var ret []ConversationMember
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Removed),
		Reset: (func(x []ConversationMember) []ConversationMember {
			if x == nil {
				return nil
			}
			var ret []ConversationMember
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Reset),
		Previewed: (func(x []ConversationID) []ConversationID {
			if x == nil {
				return nil
			}
			var ret []ConversationID
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Previewed),
		UnreadUpdate: (func(x *UnreadUpdate) *UnreadUpdate {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.UnreadUpdate),
		UnreadUpdates: (func(x []UnreadUpdate) []UnreadUpdate {
			if x == nil {
				return nil
			}
			var ret []UnreadUpdate
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.UnreadUpdates),
	}
}

type TeamChannelUpdate struct {
	TeamID TLFID `codec:"teamID" json:"teamID"`
}

func (o TeamChannelUpdate) DeepCopy() TeamChannelUpdate {
	return TeamChannelUpdate{
		TeamID: o.TeamID.DeepCopy(),
	}
}

type SetConvRetentionUpdate struct {
	InboxVers InboxVers       `codec:"inboxVers" json:"inboxVers"`
	ConvID    ConversationID  `codec:"convID" json:"convID"`
	Policy    RetentionPolicy `codec:"policy" json:"policy"`
}

func (o SetConvRetentionUpdate) DeepCopy() SetConvRetentionUpdate {
	return SetConvRetentionUpdate{
		InboxVers: o.InboxVers.DeepCopy(),
		ConvID:    o.ConvID.DeepCopy(),
		Policy:    o.Policy.DeepCopy(),
	}
}

type SetTeamRetentionUpdate struct {
	InboxVers InboxVers       `codec:"inboxVers" json:"inboxVers"`
	TeamID    keybase1.TeamID `codec:"teamID" json:"teamID"`
	Policy    RetentionPolicy `codec:"policy" json:"policy"`
}

func (o SetTeamRetentionUpdate) DeepCopy() SetTeamRetentionUpdate {
	return SetTeamRetentionUpdate{
		InboxVers: o.InboxVers.DeepCopy(),
		TeamID:    o.TeamID.DeepCopy(),
		Policy:    o.Policy.DeepCopy(),
	}
}

type KBFSImpteamUpgradeUpdate struct {
	ConvID    ConversationID `codec:"convID" json:"convID"`
	InboxVers InboxVers      `codec:"inboxVers" json:"inboxVers"`
}

func (o KBFSImpteamUpgradeUpdate) DeepCopy() KBFSImpteamUpgradeUpdate {
	return KBFSImpteamUpgradeUpdate{
		ConvID:    o.ConvID.DeepCopy(),
		InboxVers: o.InboxVers.DeepCopy(),
	}
}

type GregorInterface interface {
}

func GregorProtocol(i GregorInterface) rpc.Protocol {
	return rpc.Protocol{
		Name:    "chat.1.gregor",
		Methods: map[string]rpc.ServeHandlerDescription{},
	}
}

type GregorClient struct {
	Cli rpc.GenericClient
}
