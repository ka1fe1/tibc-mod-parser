package tibc

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/tibc-mod-parser/modules"
)

type TibcClient struct {
}

func NewClient() TibcClient {
	return TibcClient{}
}
func (ibc TibcClient) HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {
	var (
		msgDocInfo MsgDocInfo
	)
	ok := true
	switch msg := v.(type) {
	case *MsgTIBCNftTransfer:
		docMsg := DocMsgTIBCNftTransfer{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgTIBCUpdateClient:
		docMsg := DocMsgTIBCUpdateClient{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgTIBCRecvPacket:
		docMsg := DocMsgTIBCRecvPacket{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgTIBCAcknowledgement:
		docMsg := DocMsgTIBCAcknowledgement{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgCleanPacket:
		docMsg := DocMsgCleanPacket{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgRecvCleanPacket:
		docMsg := DocMsgRecvCleanPacket{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	default:
		ok = false
	}
	return msgDocInfo, ok
}
