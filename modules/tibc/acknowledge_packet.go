package tibc

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	"github.com/kaifei-bianjie/common-parser/utils"
	. "github.com/kaifei-bianjie/tibc-mod-parser/modules"
)

type DocMsgTIBCAcknowledgement struct {
	Packet          Packet `bson:"packet"`
	Acknowledgement string `bson:"acknowledgement"`
	ProofAcked      string `bson:"proof_acked"`
	ProofHeight     Height `bson:"proof_height"`
	Signer          string `bson:"signer"`
}

func (m *DocMsgTIBCAcknowledgement) GetType() string {
	return MsgTypeTIBCAcknowledgement
}

func (m *DocMsgTIBCAcknowledgement) BuildMsg(v interface{}) {
	msg := v.(*MsgTIBCAcknowledgement)
	m.Signer = msg.Signer
	m.ProofHeight = loadHeight(msg.ProofHeight)
	m.Acknowledgement = UnmarshalTibcAcknowledgement(msg.Acknowledgement)
	m.ProofAcked = utils.MarshalJsonIgnoreErr(msg.ProofAcked)
	m.Packet = loadPacket(msg.Packet)

}

func (m *DocMsgTIBCAcknowledgement) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgTIBCAcknowledgement
	)

	packetData := UnmarshalPacketData(msg.Packet.GetData())
	addrs = append(addrs, msg.Signer, packetData.Receiver, packetData.Sender)
	handler := func() (Msg, []string) {
		return m, addrs
	}
	return CreateMsgDocInfo(v, handler)
}
