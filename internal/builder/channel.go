package builder

// Channel .
type Channel struct {
    ChannelId          string `json:"channelId"`
    ChannelName        string `json:"channelName"`
    NotifyType         int32  `json:"notifyType"`
    ChannelDescription string `json:"channelDescription,omitempty"`
    SoundUri           string `json:"soundUri,omitempty"`
}
