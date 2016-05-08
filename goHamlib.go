package goHamlib

import (
		"fmt"
)

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
	RIG_ENTARGET	= -12
	RIG_BUSERROR	= -13
	RIG_BUSBUSY	= -14
	RIG_EARG	= -15
	RIG_EVFO	= -16
	RIG_EDOM	= -17
)

//Hamlib Debug levels
const (
	RIG_DEBUG_NONE = 0
	RIG_DEBUG_BUG = 1
	RIG_DEBUG_ERR = 2
	RIG_DEBUG_WARN = 3
	RIG_DEBUG_VERBOSE = 4
	RIG_DEBUG_TRACE = 5
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

// Hamlib modes
const (
	RIG_MODE_NONE = 0
	RIG_MODE_AM = 1
	RIG_MODE_CW = 2
	RIG_MODE_USB = 3
	RIG_MODE_LSB = 4
	RIG_MODE_RTTY = 5
	RIG_MODE_FM = 6
	RIG_MODE_WFM = 7
	RIG_MODE_CWR = 8
	RIG_MODE_RTTYR = 9
	RIG_MODE_AMS = 10
	RIG_MODE_PKTLSB = 11
	RIG_MODE_PKTUSB = 12
	RIG_MODE_PKTFM = 13
	RIG_MODE_ECSSUSB = 14
	RIG_MODE_ECSSLSB = 15
	RIG_MODE_FAX = 16
	RIG_MODE_SAM = 17
	RIG_MODE_SAL = 18
	RIG_MODE_SAH = 18
	RIG_MODE_DSB = 19
	RIG_MODE_FMN = 20
	RIG_MODE_TESTS_MAX = 21
)

// Hamlib Powerstats
const (
	RIG_POWER_OFF = 0
	RIG_POWER_ON = 1
	RIG_POWER_STANDBY = 2
)

// Hamlib PTT
const (
	RIG_PTT_OFF = 0
	RIG_PTT_ON = 1
	RIG_PTT_ON_MIC = 2
	RIG_PTT_ON_DATA = 3
)

// Hamlib Antennas
const (
	RIG_ANT_NONE = 0
	RIG_ANT_1 = 1
	RIG_ANT_2 = 2
	RIG_ANT_3 = 3
	RIG_ANT_4 = 4
	RIG_ANT_5 = 5
)


// Hamlib Split 
const (
	RIG_SPLIT_OFF = 0
	RIG_SPLIT_ON = 1
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


type HamlibError struct{
	Operation	string
	Errorcode	int
	Description	string
}

func (e *HamlibError) Error() string{
	switch e.Errorcode{
		case RIG_OK:
			e.Description = "OK"
		case RIG_EINVAL:
			e.Description = "invalid parameter"
		case RIG_ECONF:
			e.Description = "invalid configuration (serial,..)"
		case RIG_ENOMEM:
			e.Description = "memory shortage"
		case RIG_ENIMPL:
			e.Description = "function not implemented, but will be"
		case RIG_ETIMEOUT:
			e.Description = "communication timed out"
		case RIG_EIO:
			e.Description = "IO error, including open failed"
		case RIG_EINTERNAL:
			e.Description = "Internal Hamlib error, huh!"
		case RIG_EPROTO:
			e.Description = "Protocol error"
		case RIG_ERJCTED:
			e.Description = "Command rejected by the rig"
		case RIG_ETRUNC:
			e.Description = "Command performed, but arg truncated"
		case RIG_ENAVAIL:
			e.Description = "function not available"
		case RIG_ENTARGET:
			e.Description = "VFO not targetable"
		case RIG_BUSERROR:
			e.Description = "Error talking on the bus"
		case RIG_BUSBUSY:
			e.Description = "Collision on the bus"
		case RIG_EARG:
			e.Description = "NULL RIG handle or any invalid pointer parameter in get arg"
		case RIG_EVFO:
			e.Description = "Invalid VFO"
		case RIG_EDOM:
			e.Description = "Argument out of domain of func"
		default:
			e.Description = "unkown Error"
	}
	return fmt.Sprintf("%s: %s (%v)", e.Operation, e.Description, e.Errorcode)
}

type Error struct{
	Operation	string
	UnderlyingError error
}

func (e *Error) Error() string{
	return fmt.Sprintf("%s: %v", e.Operation, e.UnderlyingError)
}
