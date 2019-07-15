package NtmXmpp

import "github.com/alfredyang1986/blackmirror/bmxmpp"

type NtmXmpp struct {
	//GroupRoomID string
	XmppConfig bmxmpp.BmXmppConfig
}

func (r NtmXmpp) NewNtmXmppBDaemon(args map[string]string) *NtmXmpp {
	//env := os.Getenv("BM_XMPP_CONF_HOME") + "/resource/xmppconfig.json"
	//os.Setenv("BM_XMPP_CONF_HOME", env)
	bxc, _ := bmxmpp.GetConfigInstance()
	ins := NtmXmpp {
		//GroupRoomID: args["room"],
		XmppConfig: *bxc,
	}
	return &ins
}

func (r NtmXmpp) SendGroupMsg(room, msg string) error {
	return r.XmppConfig.Forward2Group(room, msg)
}