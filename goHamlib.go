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
	RIG_MODE_AM = 1<<0
	RIG_MODE_CW = 1<<1
	RIG_MODE_USB = 1<<2
	RIG_MODE_LSB = 1<<3
	RIG_MODE_RTTY = 1<<4
	RIG_MODE_FM = 1<<5
	RIG_MODE_WFM = 1<<6
	RIG_MODE_CWR = 1<<7
	RIG_MODE_RTTYR = 1<<8
	RIG_MODE_AMS = 1<<9
	RIG_MODE_PKTLSB = 1<<10
	RIG_MODE_PKTUSB = 1<<11
	RIG_MODE_PKTFM = 1<<12
	RIG_MODE_ECSSUSB = 1<<13
	RIG_MODE_ECSSLSB = 1<<14
	RIG_MODE_FAX = 1<<15
	RIG_MODE_SAM = 1<<16
	RIG_MODE_SAL = 1<<17
	RIG_MODE_SAH = 1<<18
	RIG_MODE_DSB = 1<<19
	RIG_MODE_FMN = 1<<21
	RIG_MODE_TESTS_MAX
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
	RIG_ANT_1 = 1<<0
	RIG_ANT_2 = 1<<1
	RIG_ANT_3 = 1<<2
	RIG_ANT_4 = 1<<3
	RIG_ANT_5 = 1<<4
)

// Hamlib Split 
const (
	RIG_SPLIT_OFF = 0
	RIG_SPLIT_ON = 1
)

// Hamlib Levels
const (
	RIG_LEVEL_NONE = 0
	RIG_LEVEL_PREAMP = 1<<0
	RIG_LEVEL_ATT =	1<<1
	RIG_LEVEL_VOX =	1<<2
	RIG_LEVEL_AF =	1<<3
	RIG_LEVEL_RF = 1<<4
	RIG_LEVEL_SQL = 1<<5
	RIG_LEVEL_IF = 1<<6
	RIG_LEVEL_APF = 1<<7
	RIG_LEVEL_NR = 1<<8
	RIG_LEVEL_PBT_IN = 1<<9
	RIG_LEVEL_PBT_OUT = 1<<10
	RIG_LEVEL_CWPITCH = 1<<11
	RIG_LEVEL_RFPOWER = 1<<12
	RIG_LEVEL_MICGAIN = 1<<13
	RIG_LEVEL_KEYSPD = 1<<14
	RIG_LEVEL_NOTCHF = 1<<15
	RIG_LEVEL_COMP = 1<<16
	RIG_LEVEL_AGC = 1<<17
	RIG_LEVEL_BKINDL = 1<<18
	RIG_LEVEL_BALANCE = 1<<19
	RIG_LEVEL_METER = 1<<20
	RIG_LEVEL_VOXGAIN = 1<<21
	RIG_LEVEL_VOXDELAY = RIG_LEVEL_VOX
	RIG_LEVEL_ANTIVOX = 1<<22
	RIG_LEVEL_SLOPE_LOW = 1<<23
	RIG_LEVEL_SLOPE_HIGH = 1<<24
	RIG_LEVEL_BKIN_DLYMS = 1<<25
	RIG_LEVEL_RAWSTR = 1<<26
	RIG_LEVEL_SQLSTAT = 1<<27
	RIG_LEVEL_SWR = 1<<28
	RIG_LEVEL_ALC = 1<<29
	RIG_LEVEL_STRENGTH = 1<<30
)

//Hamlib Params
const (
	RIG_PARM_NONE = 0
	RIG_PARM_ANN = 1<<0
	RIG_PARM_APO = 1<<1
	RIG_PARM_BACKLIGHT = 1<<2
	RIG_PARM_BEEP = 1<<4
	RIG_PARM_TIME = 1<<5
	RIG_PARM_BAT = 1<<6
	RIG_PARM_KEYLIGHT = 1<<7
)

//Hamlib Functions
const (
	RIG_FUNC_NONE = 0
	RIG_FUNC_FAGC = 1<<0
	RIG_FUNC_NB = 1<<1
	RIG_FUNC_COMP = 1<<2
	RIG_FUNC_VOX = 1<<3
	RIG_FUNC_TONE = 1<<4
	RIG_FUNC_TSQL = 1<<5
	RIG_FUNC_SBKIN = 1<<6
	RIG_FUNC_FBKIN = 1<<7
	RIG_FUNC_ANF = 1<<8
	RIG_FUNC_NR = 1<<9
	RIG_FUNC_AIP = 1<<10
	RIG_FUNC_APF = 1<<11
	RIG_FUNC_MON = 1<<12
	RIG_FUNC_MN = 1<<13
	RIG_FUNC_RF = 1<<14
	RIG_FUNC_ARO = 1<<15
	RIG_FUNC_LOCK = 1<<16
	RIG_FUNC_MUTE = 1<<17
	RIG_FUNC_VSC = 1<<18
	RIG_FUNC_REV = 1<<19
	RIG_FUNC_SQL = 1<<20
	RIG_FUNC_ABM = 1<<21
	RIG_FUNC_BC = 1<<22
	RIG_FUNC_MBC = 1<<23
	RIG_FUNC_RIT = 1<<24
	RIG_FUNC_AFC = 1<<25
	RIG_FUNC_SATMODE = 1<<26
	RIG_FUNC_SCOPE = 1<<27
	RIG_FUNC_RESUME = 1<<28
	RIG_FUNC_TBURST = 1<<29
	RIG_FUNC_TUNER = 1<<30
	RIG_FUNC_XIT = 1<<31
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
