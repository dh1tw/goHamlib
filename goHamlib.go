package goHamlib

const (
	N = 0
	E = 1
	O = 2
)

const (
	NO_HANDSHAKE = 0
	RTSCTS_HANDSHAKE = 1
)

//Hamlib Error Codes
const (
	RIG_OK		= 0
	RIG_EINVAL	= -1
	RIG_ECONF	= -2
	RIG_ENOMEM	= -3
	RIG_ENIMPL	= -4
	RIG_ETIMEOUT	= -5
	RIG_EIO		= -6
	RIG_EINTERNAL	= -7
	RIG_EPROTO	= -8
	RIG_ERJCTED	= -9
	RIG_ETRUNC	= -10
	RIG_ENAVAIL	= -11
	RIG_ETARGET	= -12
	RIG_BUSYERROR	= -13
	RIG_BUSYBUSY	= -14
	RIG_EARG	= -15
	RIG_EVFO	= -16
	RIG_EDOM	= -17
)

//Hamlib VFOs
const (
	RIG_VFO_NONE = 0
	RIG_VFO_TX_FLAG = 1<<30
	RIG_VFO_CURR = 1<<29
	RIG_VFO_MEM = 1<<28
	RIG_VFO_VFO = 1<<27
	RIG_VFO_RX = RIG_VFO_CURR
	RIG_VFO_MAIN = 1<<26
	RIG_VFO_SUB = 1<<25
	RIG_VFO_A = 1<<1
	RIG_VFO_B = 1<<2
	RIG_VFO_C = 1<<3
)

type Port_t struct{
	RigPortType	int
	Portname	string
	Baudrate	int
	Databits	int
	Stopbits	int
	Parity		int
	Handshake	int
}

type Rig struct{
	port	Port_t
}

