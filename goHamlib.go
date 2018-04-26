package goHamlib

import (
	/*
		#include <hamlib/rig.h>
	*/
	"C"
	"errors"
	"fmt"
)

// Parity is the serial port parity
type Parity byte

// Serial port parity constants
const (
	ParityNone Parity = 0
	ParityEven Parity = 1
	ParityOdd  Parity = 2
)

// Handshake is a serial port handshake type
type Handshake byte

// Serial port handshake constants
const (
	HandshakeNone   Handshake = 0
	HandshakeRTSCTS Handshake = 1
)

// HamlibErrorCode is an error code returned from libhamlib
type HamlibErrorCode int

//Hamlib Error Codes
const (
	HamlibErrOK        HamlibErrorCode = 0
	HamlibErrEINVAL    HamlibErrorCode = -1
	HamlibErrECONF     HamlibErrorCode = -2
	HamlibErrENOMEM    HamlibErrorCode = -3
	HamlibErrENIMPL    HamlibErrorCode = -4
	HamlibErrETIMEOUT  HamlibErrorCode = -5
	HamlibErrEIO       HamlibErrorCode = -6
	HamlibErrEINTERNAL HamlibErrorCode = -7
	HamlibErrEPROTO    HamlibErrorCode = -8
	HamlibErrERJCTED   HamlibErrorCode = -9
	HamlibErrETRUNC    HamlibErrorCode = -10
	HamlibErrENAVAIL   HamlibErrorCode = -11
	HamlibErrENTARGET  HamlibErrorCode = -12
	HamlibErrBUSERROR  HamlibErrorCode = -13
	HamlibErrBUSBUSY   HamlibErrorCode = -14
	HamlibErrEARG      HamlibErrorCode = -15
	HamlibErrEVFO      HamlibErrorCode = -16
	HamlibErrEDOM      HamlibErrorCode = -17
)

// DebugLevel is the Hamlib Debug level type
type DebugLevel byte

// Hamlib debug level constants
const (
	DebugNone    DebugLevel = 0
	DebugBug     DebugLevel = 1
	DebugErr     DebugLevel = 2
	DebugWarn    DebugLevel = 3
	DebugVerbose DebugLevel = 4
	DebugTrace   DebugLevel = 5
)

// DebugLevelValue maps strings to DebugLevel
var DebugLevelValue = map[string]DebugLevel{
	"NONE":    DebugNone,
	"BUG":     DebugBug,
	"ERROR":   DebugErr,
	"WARN":    DebugWarn,
	"VERBOSE": DebugVerbose,
	"TRACE":   DebugTrace,
}

// DebugLevelName maps DebugLevel to a name string
var DebugLevelName = map[DebugLevel]string{
	DebugNone:    "NONE",
	DebugBug:     "BUG",
	DebugErr:     "ERROR",
	DebugWarn:    "WARN",
	DebugVerbose: "VERBOSE",
	DebugTrace:   "TRACE",
}

// VFOType is used to specify a particular VFO
type VFOType int

const vfoTXFlag VFOType = 1 << 30

//Hamlib VFOs
const (
	VFONone    VFOType = 0
	VFOCurrent VFOType = 1 << 29
	VFOMemory  VFOType = 1 << 28
	VFOLastVFO VFOType = 1 << 27
	//	VFOTX_VFO 			//TBD Macro
	VFOTX   VFOType = VFOCurrent | vfoTXFlag
	VFORX           = VFOCurrent
	VFOMain VFOType = 1 << 26
	VFOSub  VFOType = 1 << 25
	VFOA    VFOType = 1 << 0
	VFOB    VFOType = 1 << 1
	VFOC    VFOType = 1 << 2
)

// VFOName is a map containing Strings for VFOs
var VFOName = map[VFOType]string{
	VFONone:    "NONE",
	VFOA:       "VFOA",
	VFOB:       "VFOB",
	VFOC:       "VFOC",
	VFOCurrent: "CURR",
	VFOMemory:  "MEM",
	VFOLastVFO: "VFO",
	VFOTX:      "TX",
	// VFORX:      "RX", same as VFOCurrent
	VFOMain: "MAIN",
	VFOSub:  "SUB",
}

// VFOValue is a map from VFO names to values
var VFOValue = map[string]VFOType{
	"NONE": VFONone,
	"VFOA": VFOA,
	"VFOB": VFOB,
	"VFOC": VFOC,
	"CURR": VFOCurrent,
	"MAIN": VFOMain,
	"SUB":  VFOSub,
	"MEM":  VFOMemory,
	"VFO":  VFOLastVFO,
	"TX":   VFOTX,
}

// VFOOp is a VFO operation
type VFOOp int

// VFO Operation constants
const (
	VFOOpNone        VFOOp = 0
	VFOOpCopy        VFOOp = 1 << 0
	VFOOpExchange    VFOOp = 1 << 1
	VFOOpFromVFO     VFOOp = 1 << 2
	VFOOpToVFO       VFOOp = 1 << 3
	VFOOpMemoryClear VFOOp = 1 << 4
	VFOOpUp          VFOOp = 1 << 5
	VFOOpDown        VFOOp = 1 << 6
	VFOOpBandUp      VFOOp = 1 << 7
	VFOOpBandDown    VFOOp = 1 << 8
	VFOOpLeft        VFOOp = 1 << 9
	VFOOpRight       VFOOp = 1 << 10
	VFOOpTune        VFOOp = 1 << 11
	VFOOpToggle      VFOOp = 1 << 12
)

// VFOOperationName is a map of VFO operation values to names
var VFOOperationName = map[VFOOp]string{
	VFOOpNone:        "",
	VFOOpCopy:        "CPY",
	VFOOpExchange:    "XCHG",
	VFOOpFromVFO:     "FROM_VFO",
	VFOOpToVFO:       "TO_VFO",
	VFOOpMemoryClear: "MCL",
	VFOOpUp:          "UP",
	VFOOpDown:        "DOWN",
	VFOOpBandUp:      "BAND_UP",
	VFOOpBandDown:    "BAND_DOWN",
	VFOOpLeft:        "LEFT",
	VFOOpRight:       "RIGHT",
	VFOOpTune:        "TUNE",
	VFOOpToggle:      "TOGGLE",
}

// VFOOperationValue is a map of VFO operation names to values
var VFOOperationValue = map[string]VFOOp{
	"":          VFOOpNone,
	"CPY":       VFOOpCopy,
	"XCHG":      VFOOpExchange,
	"FROM_VFO":  VFOOpFromVFO,
	"TO_VFO":    VFOOpToVFO,
	"MCL":       VFOOpMemoryClear,
	"UP":        VFOOpUp,
	"DOWN":      VFOOpDown,
	"BAND_UP":   VFOOpBandUp,
	"BAND_DOWN": VFOOpBandDown,
	"LEFT":      VFOOpLeft,
	"RIGHT":     VFOOpRight,
	"TUNE":      VFOOpTune,
	"TOGGLE":    VFOOpToggle,
}

// Hamlib modes
const (
	RIG_MODE_NONE    = 0
	RIG_MODE_AM      = 1 << 0
	RIG_MODE_CW      = 1 << 1
	RIG_MODE_USB     = 1 << 2
	RIG_MODE_LSB     = 1 << 3
	RIG_MODE_RTTY    = 1 << 4
	RIG_MODE_FM      = 1 << 5
	RIG_MODE_WFM     = 1 << 6
	RIG_MODE_CWR     = 1 << 7
	RIG_MODE_RTTYR   = 1 << 8
	RIG_MODE_AMS     = 1 << 9
	RIG_MODE_PKTLSB  = 1 << 10
	RIG_MODE_PKTUSB  = 1 << 11
	RIG_MODE_PKTFM   = 1 << 12
	RIG_MODE_ECSSUSB = 1 << 13
	RIG_MODE_ECSSLSB = 1 << 14
	RIG_MODE_FAX     = 1 << 15
	RIG_MODE_SAM     = 1 << 16
	RIG_MODE_SAL     = 1 << 17
	RIG_MODE_SAH     = 1 << 18
	RIG_MODE_DSB     = 1 << 19
	RIG_MODE_FMN     = 1 << 21
	RIG_MODE_TESTS_MAX
)

//Map containing Strings for Modes
var ModeName = map[int]string{
	RIG_MODE_NONE:    "",
	RIG_MODE_AM:      "AM",
	RIG_MODE_CW:      "CW",
	RIG_MODE_USB:     "USB",
	RIG_MODE_LSB:     "LSB",
	RIG_MODE_RTTY:    "RTTY",
	RIG_MODE_FM:      "FM",
	RIG_MODE_WFM:     "WFM",
	RIG_MODE_CWR:     "CWR",
	RIG_MODE_RTTYR:   "RTTYR",
	RIG_MODE_AMS:     "AMS",
	RIG_MODE_PKTLSB:  "PKTLSB",
	RIG_MODE_PKTUSB:  "PKTUSB",
	RIG_MODE_PKTFM:   "PKTFM",
	RIG_MODE_ECSSUSB: "ECSSUSB",
	RIG_MODE_ECSSLSB: "ECSSLSB",
	RIG_MODE_FAX:     "FAX",
	RIG_MODE_SAM:     "SAM",
	RIG_MODE_SAL:     "SAL",
	RIG_MODE_SAH:     "SAH",
	RIG_MODE_DSB:     "DSB",
	RIG_MODE_FMN:     "FMN",
}

//Map containing Values for Modes
var ModeValue = map[string]int{
	"":        RIG_MODE_NONE,
	"AM":      RIG_MODE_AM,
	"CW":      RIG_MODE_CW,
	"USB":     RIG_MODE_USB,
	"LSB":     RIG_MODE_LSB,
	"RTTY":    RIG_MODE_RTTY,
	"FM":      RIG_MODE_FM,
	"WFM":     RIG_MODE_WFM,
	"CWR":     RIG_MODE_CWR,
	"RTTYR":   RIG_MODE_RTTYR,
	"AMS":     RIG_MODE_AMS,
	"PKTLSB":  RIG_MODE_PKTLSB,
	"PKTUSB":  RIG_MODE_PKTUSB,
	"PKTFM":   RIG_MODE_PKTFM,
	"ECSSUSB": RIG_MODE_ECSSUSB,
	"ECSSLSB": RIG_MODE_ECSSLSB,
	"FAX":     RIG_MODE_FAX,
	"SAM":     RIG_MODE_SAM,
	"SAL":     RIG_MODE_SAL,
	"SAH":     RIG_MODE_SAH,
	"DSB":     RIG_MODE_DSB,
	"FMN":     RIG_MODE_FMN,
}

// Hamlib Powerstats
const (
	RIG_POWER_OFF     = 0
	RIG_POWER_ON      = 1
	RIG_POWER_STANDBY = 2
)

var RigPowerName = map[int]string{
	RIG_POWER_OFF:     "OFF",
	RIG_POWER_ON:      "ON",
	RIG_POWER_STANDBY: "STANDBY",
}

var RigPowerValue = map[string]int{
	"OFF":     RIG_POWER_OFF,
	"ON":      RIG_POWER_ON,
	"STANDBY": RIG_POWER_STANDBY,
}

// Hamlib PTT
const (
	RIG_PTT_OFF     = 0
	RIG_PTT_ON      = 1
	RIG_PTT_ON_MIC  = 2
	RIG_PTT_ON_DATA = 3
)

var PttName = map[int]string{
	RIG_PTT_OFF:     "OFF",
	RIG_PTT_ON:      "ON",
	RIG_PTT_ON_MIC:  "ON_MIC",
	RIG_PTT_ON_DATA: "ON_DATA",
}

var PttValue = map[string]int{
	"OFF":     RIG_PTT_OFF,
	"ON":      RIG_PTT_ON,
	"ON_MIC":  RIG_PTT_ON_MIC,
	"ON_DATA": RIG_PTT_ON_DATA,
}

// Hamlib Antennas
const (
	RIG_ANT_NONE = 0
	RIG_ANT_1    = 1 << 0
	RIG_ANT_2    = 1 << 1
	RIG_ANT_3    = 1 << 2
	RIG_ANT_4    = 1 << 3
	RIG_ANT_5    = 1 << 4
)

var AntName = map[int]string{
	RIG_ANT_1: "ANT1",
	RIG_ANT_2: "ANT2",
	RIG_ANT_3: "ANT3",
	RIG_ANT_4: "ANT4",
	RIG_ANT_5: "ANT5",
}

var AntValue = map[string]int{
	"ANT1": RIG_ANT_1,
	"ANT2": RIG_ANT_2,
	"ANT3": RIG_ANT_3,
	"ANT4": RIG_ANT_4,
	"ANT5": RIG_ANT_5,
}

// Hamlib Split
const (
	RIG_SPLIT_OFF = 0
	RIG_SPLIT_ON  = 1
)

// Hamlib Levels
const (
	RIG_LEVEL_NONE       = 0
	RIG_LEVEL_PREAMP     = 1 << 0
	RIG_LEVEL_ATT        = 1 << 1
	RIG_LEVEL_VOX        = 1 << 2
	RIG_LEVEL_AF         = 1 << 3
	RIG_LEVEL_RF         = 1 << 4
	RIG_LEVEL_SQL        = 1 << 5
	RIG_LEVEL_IF         = 1 << 6
	RIG_LEVEL_APF        = 1 << 7
	RIG_LEVEL_NR         = 1 << 8
	RIG_LEVEL_PBT_IN     = 1 << 9
	RIG_LEVEL_PBT_OUT    = 1 << 10
	RIG_LEVEL_CWPITCH    = 1 << 11
	RIG_LEVEL_RFPOWER    = 1 << 12
	RIG_LEVEL_MICGAIN    = 1 << 13
	RIG_LEVEL_KEYSPD     = 1 << 14
	RIG_LEVEL_NOTCHF     = 1 << 15
	RIG_LEVEL_COMP       = 1 << 16
	RIG_LEVEL_AGC        = 1 << 17
	RIG_LEVEL_BKINDL     = 1 << 18
	RIG_LEVEL_BALANCE    = 1 << 19
	RIG_LEVEL_METER      = 1 << 20
	RIG_LEVEL_VOXGAIN    = 1 << 21
	RIG_LEVEL_VOXDELAY   = RIG_LEVEL_VOX
	RIG_LEVEL_ANTIVOX    = 1 << 22
	RIG_LEVEL_SLOPE_LOW  = 1 << 23
	RIG_LEVEL_SLOPE_HIGH = 1 << 24
	RIG_LEVEL_BKIN_DLYMS = 1 << 25
	RIG_LEVEL_RAWSTR     = 1 << 26
	RIG_LEVEL_SQLSTAT    = 1 << 27
	RIG_LEVEL_SWR        = 1 << 28
	RIG_LEVEL_ALC        = 1 << 29
	RIG_LEVEL_STRENGTH   = 1 << 30
)

var LevelName = map[uint32]string{
	RIG_LEVEL_NONE:    "",
	RIG_LEVEL_PREAMP:  "PREAMP",
	RIG_LEVEL_ATT:     "ATT",
	RIG_LEVEL_VOX:     "VOX",
	RIG_LEVEL_AF:      "AF",
	RIG_LEVEL_RF:      "RF",
	RIG_LEVEL_SQL:     "SQL",
	RIG_LEVEL_IF:      "IF",
	RIG_LEVEL_APF:     "APF",
	RIG_LEVEL_NR:      "NR",
	RIG_LEVEL_PBT_IN:  "PBT_IN",
	RIG_LEVEL_PBT_OUT: "PBT_OUT",
	RIG_LEVEL_CWPITCH: "CWPITCH",
	RIG_LEVEL_RFPOWER: "RFPOWER",
	RIG_LEVEL_MICGAIN: "MICGAIN",
	RIG_LEVEL_KEYSPD:  "KEYSPD",
	RIG_LEVEL_NOTCHF:  "NOTCHF",
	RIG_LEVEL_COMP:    "COMP",
	RIG_LEVEL_AGC:     "AGC",
	RIG_LEVEL_BKINDL:  "BKINDL",
	RIG_LEVEL_BALANCE: "BALANCE",
	RIG_LEVEL_METER:   "METER",
	RIG_LEVEL_VOXGAIN: "VOXGAIN",
	//RIG_LEVEL_VOXDELAY = RIG_LEVEL_VOX
	RIG_LEVEL_ANTIVOX:    "ANTIVOX",
	RIG_LEVEL_SLOPE_LOW:  "SLOPE_LOW",
	RIG_LEVEL_SLOPE_HIGH: "SLOPE_HIGH",
	RIG_LEVEL_BKIN_DLYMS: "BKIN_DLYMS",
	RIG_LEVEL_RAWSTR:     "RAWSTR",
	RIG_LEVEL_SQLSTAT:    "SQLSTAT",
	RIG_LEVEL_SWR:        "SWR",
	RIG_LEVEL_ALC:        "ALC",
	RIG_LEVEL_STRENGTH:   "STRENGTH",
}

var LevelValue = map[string]uint32{
	"":        RIG_LEVEL_NONE,
	"PREAMP":  RIG_LEVEL_PREAMP,
	"ATT":     RIG_LEVEL_ATT,
	"VOX":     RIG_LEVEL_VOX,
	"AF":      RIG_LEVEL_AF,
	"RF":      RIG_LEVEL_RF,
	"SQL":     RIG_LEVEL_SQL,
	"IF":      RIG_LEVEL_IF,
	"APF":     RIG_LEVEL_APF,
	"NR":      RIG_LEVEL_NR,
	"PBT_IN":  RIG_LEVEL_PBT_IN,
	"PBT_OUT": RIG_LEVEL_PBT_OUT,
	"CWPITCH": RIG_LEVEL_CWPITCH,
	"RFPOWER": RIG_LEVEL_RFPOWER,
	"MICGAIN": RIG_LEVEL_MICGAIN,
	"KEYSPD":  RIG_LEVEL_KEYSPD,
	"NOTCHF":  RIG_LEVEL_NOTCHF,
	"COMP":    RIG_LEVEL_COMP,
	"AGC":     RIG_LEVEL_AGC,
	"BKINDL":  RIG_LEVEL_BKINDL,
	"BALANCE": RIG_LEVEL_BALANCE,
	"METER":   RIG_LEVEL_METER,
	"VOXGAIN": RIG_LEVEL_VOXGAIN,
	//RIG_LEVEL_VOXDELAY = RIG_LEVEL_VOX
	"ANTIVOX":    RIG_LEVEL_ANTIVOX,
	"SLOPE_LOW":  RIG_LEVEL_SLOPE_LOW,
	"SLOPE_HIGH": RIG_LEVEL_SLOPE_HIGH,
	"BKIN_DLYMS": RIG_LEVEL_BKIN_DLYMS,
	"RAWSTR":     RIG_LEVEL_RAWSTR,
	"SQLSTAT":    RIG_LEVEL_SQLSTAT,
	"SWR":        RIG_LEVEL_SWR,
	"ALC":        RIG_LEVEL_ALC,
	"STRENGTH":   RIG_LEVEL_STRENGTH,
}

//Hamlib Params
const (
	RIG_PARM_NONE      = 0
	RIG_PARM_ANN       = 1 << 0
	RIG_PARM_APO       = 1 << 1
	RIG_PARM_BACKLIGHT = 1 << 2
	RIG_PARM_BEEP      = 1 << 4
	RIG_PARM_TIME      = 1 << 5
	RIG_PARM_BAT       = 1 << 6
	RIG_PARM_KEYLIGHT  = 1 << 7
)

var ParmName = map[uint32]string{
	RIG_PARM_NONE:      "",
	RIG_PARM_ANN:       "ANN",
	RIG_PARM_APO:       "APO",
	RIG_PARM_BACKLIGHT: "BACKLIGHT",
	RIG_PARM_BEEP:      "BEEP",
	RIG_PARM_TIME:      "TIME",
	RIG_PARM_BAT:       "BAT",
	RIG_PARM_KEYLIGHT:  "KEYLIGHT",
}

var ParmValue = map[string]uint32{
	"":          RIG_PARM_NONE,
	"ANN":       RIG_PARM_ANN,
	"APO":       RIG_PARM_APO,
	"BACKLIGHT": RIG_PARM_BACKLIGHT,
	"BEEP":      RIG_PARM_BEEP,
	"TIME":      RIG_PARM_TIME,
	"BAT":       RIG_PARM_BAT,
	"KEYLIGHT":  RIG_PARM_KEYLIGHT,
}

//Hamlib Functions
const (
	RIG_FUNC_NONE    = 0
	RIG_FUNC_FAGC    = 1 << 0
	RIG_FUNC_NB      = 1 << 1
	RIG_FUNC_COMP    = 1 << 2
	RIG_FUNC_VOX     = 1 << 3
	RIG_FUNC_TONE    = 1 << 4
	RIG_FUNC_TSQL    = 1 << 5
	RIG_FUNC_SBKIN   = 1 << 6
	RIG_FUNC_FBKIN   = 1 << 7
	RIG_FUNC_ANF     = 1 << 8
	RIG_FUNC_NR      = 1 << 9
	RIG_FUNC_AIP     = 1 << 10
	RIG_FUNC_APF     = 1 << 11
	RIG_FUNC_MON     = 1 << 12
	RIG_FUNC_MN      = 1 << 13
	RIG_FUNC_RF      = 1 << 14
	RIG_FUNC_ARO     = 1 << 15
	RIG_FUNC_LOCK    = 1 << 16
	RIG_FUNC_MUTE    = 1 << 17
	RIG_FUNC_VSC     = 1 << 18
	RIG_FUNC_REV     = 1 << 19
	RIG_FUNC_SQL     = 1 << 20
	RIG_FUNC_ABM     = 1 << 21
	RIG_FUNC_BC      = 1 << 22
	RIG_FUNC_MBC     = 1 << 23
	RIG_FUNC_RIT     = 1 << 24
	RIG_FUNC_AFC     = 1 << 25
	RIG_FUNC_SATMODE = 1 << 26
	RIG_FUNC_SCOPE   = 1 << 27
	RIG_FUNC_RESUME  = 1 << 28
	RIG_FUNC_TBURST  = 1 << 29
	RIG_FUNC_TUNER   = 1 << 30
	RIG_FUNC_XIT     = 1 << 31
)

var FuncName = map[uint32]string{
	RIG_FUNC_NONE:    "",
	RIG_FUNC_FAGC:    "FAGC",
	RIG_FUNC_NB:      "NB",
	RIG_FUNC_COMP:    "COMP",
	RIG_FUNC_VOX:     "VOX",
	RIG_FUNC_TONE:    "TONE",
	RIG_FUNC_TSQL:    "TSQL",
	RIG_FUNC_SBKIN:   "SBKIN",
	RIG_FUNC_FBKIN:   "FBKIN",
	RIG_FUNC_ANF:     "ANF",
	RIG_FUNC_NR:      "NR",
	RIG_FUNC_AIP:     "AIP",
	RIG_FUNC_APF:     "APF",
	RIG_FUNC_MON:     "MON",
	RIG_FUNC_MN:      "MN",
	RIG_FUNC_RF:      "RF",
	RIG_FUNC_ARO:     "ARO",
	RIG_FUNC_LOCK:    "LOCK",
	RIG_FUNC_MUTE:    "MUTE",
	RIG_FUNC_VSC:     "VSC",
	RIG_FUNC_REV:     "REV",
	RIG_FUNC_SQL:     "SQL",
	RIG_FUNC_ABM:     "ABM",
	RIG_FUNC_BC:      "BC",
	RIG_FUNC_MBC:     "MBC",
	RIG_FUNC_RIT:     "RIT",
	RIG_FUNC_AFC:     "AFC",
	RIG_FUNC_SATMODE: "SATMODE",
	RIG_FUNC_SCOPE:   "SCOPE",
	RIG_FUNC_RESUME:  "RESUME",
	RIG_FUNC_TBURST:  "TBURST",
	RIG_FUNC_TUNER:   "TUNER",
	RIG_FUNC_XIT:     "XIT",
}

var FuncValue = map[string]uint32{
	"":        RIG_FUNC_NONE,
	"FAGC":    RIG_FUNC_FAGC,
	"NB":      RIG_FUNC_NB,
	"COMP":    RIG_FUNC_COMP,
	"VOX":     RIG_FUNC_VOX,
	"TONE":    RIG_FUNC_TONE,
	"TSQL":    RIG_FUNC_TSQL,
	"SBKIN":   RIG_FUNC_SBKIN,
	"FBKIN":   RIG_FUNC_FBKIN,
	"ANF":     RIG_FUNC_ANF,
	"NR":      RIG_FUNC_NR,
	"AIP":     RIG_FUNC_AIP,
	"APF":     RIG_FUNC_APF,
	"MON":     RIG_FUNC_MON,
	"MN":      RIG_FUNC_MN,
	"RF":      RIG_FUNC_RF,
	"ARO":     RIG_FUNC_ARO,
	"LOCK":    RIG_FUNC_LOCK,
	"MUTE":    RIG_FUNC_MUTE,
	"VSC":     RIG_FUNC_VSC,
	"REV":     RIG_FUNC_REV,
	"SQL":     RIG_FUNC_SQL,
	"ABM":     RIG_FUNC_ABM,
	"BC":      RIG_FUNC_BC,
	"MBC":     RIG_FUNC_MBC,
	"RIT":     RIG_FUNC_RIT,
	"AFC":     RIG_FUNC_AFC,
	"SATMODE": RIG_FUNC_SATMODE,
	"SCOPE":   RIG_FUNC_SCOPE,
	"RESUME":  RIG_FUNC_RESUME,
	"TBURST":  RIG_FUNC_TBURST,
	"TUNER":   RIG_FUNC_TUNER,
	"XIT":     RIG_FUNC_XIT,
}

//Rig Meter
const (
	RIG_METER_NONE = 0
	RIG_METER_SWR  = 1 << 0
	RIG_METER_COMP = 1 << 1
	RIG_METER_ALC  = 1 << 2
	RIG_METER_IC   = 1 << 3
	RIG_METER_DB   = 1 << 4
	RIG_METER_PO   = 1 << 5
	RIG_METER_VDD  = 1 << 6
)

//configuration Token
const (
	TOK_FAST_SET_CMD = 1
)

type Port struct {
	RigPortType RigPort
	Portname    string
	Baudrate    int
	Databits    int
	Stopbits    int
	Parity      Parity
	Handshake   Handshake
}

type Value struct {
	Name string
	Step float32
	Min  float32
	Max  float32
}

const (
	RIG_STATUS_ALPHA    = 0
	RIG_STATUS_UNTESTED = 1
	RIG_STATUS_BETA     = 2
	RIG_STATUS_STABLE   = 3
	RIG_STATUS_BUGGY    = 4
)

var RigStatusName = map[int]string{
	RIG_STATUS_ALPHA:    "ALPHA",
	RIG_STATUS_UNTESTED: "UNTESTED",
	RIG_STATUS_BETA:     "BETA",
	RIG_STATUS_STABLE:   "STABLE",
	RIG_STATUS_BUGGY:    "BUGGY",
}

var RigStatusValue = map[string]int{
	"ALPHA":    RIG_STATUS_ALPHA,
	"UNTESTED": RIG_STATUS_UNTESTED,
	"BETA":     RIG_STATUS_BETA,
	"STABLE":   RIG_STATUS_STABLE,
	"BUGGY":    RIG_STATUS_BUGGY,
}

const (
	RIG_PTT_NONE        = 0
	RIG_PTT_RIG         = 1
	RIG_PTT_SERIAL_DTR  = 2
	RIG_PTT_SERIAL_RTS  = 3
	RIG_PTT_PARALLEL    = 4
	RIG_PTT_RIG_MICDATA = 5
	RIG_PTT_CM108       = 6
)

var RigPttName = map[int]string{
	RIG_PTT_NONE:        "NONE",
	RIG_PTT_RIG:         "CAT",
	RIG_PTT_SERIAL_DTR:  "DTR",
	RIG_PTT_SERIAL_RTS:  "RTS",
	RIG_PTT_PARALLEL:    "PARALLEL",
	RIG_PTT_RIG_MICDATA: "MIC",
	RIG_PTT_CM108:       "GPIO",
}

var RigPttValue = map[string]int{
	"NONE":     RIG_PTT_NONE,
	"CAT":      RIG_PTT_RIG,
	"DTR":      RIG_PTT_SERIAL_DTR,
	"RTS":      RIG_PTT_SERIAL_RTS,
	"PARALLEL": RIG_PTT_PARALLEL,
	"MIC":      RIG_PTT_RIG_MICDATA,
	"GPIO":     RIG_PTT_CM108,
}

const (
	RIG_DCD_NONE       = 0
	RIG_DCD_RIG        = 1
	RIG_DCD_SERIAL_DSR = 2
	RIG_DCD_SERIAL_CTS = 3
	RIG_DCD_SERIAL_CAR = 4
	RIG_DCD_PARALLEL   = 5
	RIG_DCD_CM108      = 6
)

var RigDcdName = map[int]string{
	RIG_DCD_NONE:       "NONE",
	RIG_DCD_RIG:        "RIG",
	RIG_DCD_SERIAL_DSR: "DSR",
	RIG_DCD_SERIAL_CTS: "CTS",
	RIG_DCD_SERIAL_CAR: "CAR",
	RIG_DCD_PARALLEL:   "PARALLEL",
	RIG_DCD_CM108:      "GPIO",
}

var RigDcdValue = map[string]int{
	"NONE":     RIG_DCD_NONE,
	"RIG":      RIG_DCD_RIG,
	"DSR":      RIG_DCD_SERIAL_DSR,
	"CTS":      RIG_DCD_SERIAL_CTS,
	"CAR":      RIG_DCD_SERIAL_CAR,
	"PARALLEL": RIG_DCD_PARALLEL,
	"GPIO":     RIG_DCD_CM108,
}

// RigPortType is the port connection type for the rig
type RigPort byte

// Rig Port constants
const (
	RigPortNone       RigPort = 0
	RigPortSerial     RigPort = 1
	RigPortNetwork    RigPort = 2
	RigPortDevice     RigPort = 3
	RigPortPacket     RigPort = 4
	RigPortDTMF       RigPort = 5
	RigPortUltra      RigPort = 6
	RigPortRPC        RigPort = 7
	RigPortParallel   RigPort = 8
	RigPortUSB        RigPort = 9
	RigPortUDPNetwork RigPort = 10
	RigPortCM108      RigPort = 11
)

// RigPortValue is a map of RigPort to names
var RigPortName = map[RigPort]string{
	RigPortNone:       "RIG_PORT_NONE",
	RigPortSerial:     "RIG_PORT_SERIAL",
	RigPortNetwork:    "RIG_PORT_NETWORK",
	RigPortDevice:     "RIG_PORT_DEVICE",
	RigPortPacket:     "RIG_PORT_PACKET",
	RigPortDTMF:       "RIG_PORT_DTMF",
	RigPortUltra:      "RIG_PORT_ULTRA",
	RigPortRPC:        "RIG_PORT_RPC",
	RigPortParallel:   "RIG_PORT_PARALLEL",
	RigPortUSB:        "RIG_PORT_USB",
	RigPortUDPNetwork: "RIG_PORT_UDP_NETWORK",
	RigPortCM108:      "RIG_PORT_CM108",
}

// RigPortValue is a map of names to RigPort
var RigPortValue = map[string]RigPort{
	"RIG_PORT_NONE":        RigPortNone,
	"RIG_PORT_SERIAL":      RigPortSerial,
	"RIG_PORT_NETWORK":     RigPortNetwork,
	"RIG_PORT_DEVICE":      RigPortDevice,
	"RIG_PORT_PACKET":      RigPortPacket,
	"RIG_PORT_DTMF":        RigPortDTMF,
	"RIG_PORT_ULTRA":       RigPortUltra,
	"RIG_PORT_RPC":         RigPortRPC,
	"RIG_PORT_PARALLEL":    RigPortParallel,
	"RIG_PORT_USB":         RigPortUSB,
	"RIG_PORT_UDP_NETWORK": RigPortUDPNetwork,
	"RIG_PORT_CM108":       RigPortCM108,
}

type Values []Value

type ConfParams struct {
	Token    int64
	Name     string
	Label    string
	Tooltip  string
	Dflt     string
	RigConfE int
}

type ExtList struct {
	Token int64
	Value Value
}

type Channel struct {
	ChannelNum  int
	BankNum     int
	Vfo         uint32
	Ant         int
	Freq        float64
	Mode        uint32
	Width       int
	TxFreq      float64
	TxMode      uint32
	TxWidth     int
	Split       int
	TxVfo       uint32
	RptrShift   int
	RptrOffset  int
	TuningStep  int
	Rit         int
	Xit         int
	Funcs       uint32
	Levels      Value
	CtcssTone   uint
	CtcssSql    uint
	DcsCode     uint
	DcsSql      uint
	ScanGroup   int
	Flags       int
	Description string
	ExtLevels   []ExtList
}

type Caps struct {
	RigModel        int
	ModelName       string
	MfgName         string
	Version         string
	Copyright       string
	Status          int
	RigType         int
	PttType         int
	DcdType         int
	PortType        int
	SerialRateMin   int
	SerialRateMax   int
	SerialDataBits  int
	SerialStopBits  int
	SerialParity    int
	SerialHandshake int
	WriteDelay      int
	PostWriteDelay  int
	Timeout         int
	Retry           int
	Preamps         []int
	Attenuators     []int
	MaxRit          int
	MaxXit          int
	MaxIfShift      int
	Vfos            []string
	Operations      []string
	Modes           []string
	GetFunctions    []string
	SetFunctions    []string
	GetLevels       Values
	SetLevels       Values
	GetParameters   Values
	SetParameters   Values
	TargetableVfo   int
	TargetableVfos  []int
	Transceive      int
	BankQty         int
	ChannelDescSz   int
	ChannelList     []Channel
	Filters         map[string][]int //mode + List of supported filter bandwidths
	TuningSteps     map[string][]int // mode + List of supported tuning steps
	ExtParms        []ConfParams
	ExtLevels       []ConfParams
	CtcssList       []uint
	DcsList         []uint
	Announces       int
	ScanOperations  []string
	CfgParams       []ConfParams
	HasGetPowerStat bool
	HasSetPowerStat bool
	HasGetVfo       bool
	HasSetVfo       bool
	HasSetFreq      bool
	HasGetFreq      bool
	HasSetMode      bool
	HasGetMode      bool
	HasSetPtt       bool
	HasGetPtt       bool
	HasSetRit       bool
	HasGetRit       bool
	HasSetXit       bool
	HasGetXit       bool
	HasSetSplitVfo  bool
	HasGetSplitVfo  bool
	HasSetSplitMode bool
	HasGetSplitMode bool
	HasSetSplitFreq bool
	HasGetSplitFreq bool
	HasSetAnt       bool
	HasGetAnt       bool
	HasSetTs        bool
	HasGetTs        bool
	HasGetConf      bool
	HasSetConf      bool

	/* TODO */
	/*********/

	// ScanOperations
	// targetable VfoName
	// ctcss_list
	// dcs_list
	// extparms
	// extlevels
	// transceive
	// announces
	// bank_qty
	// chan_desc_sz
	// Channels
	// rx_range_list1
	// tx_range_list1
	// rx_range_list2
	// tx_range_list2
	// TuningSteps
	// StrCal
}

type Rig struct {
	port   Port
	Caps   Caps
	handle *C.RIG
}

type HamlibError struct {
	Operation   string
	Errorcode   HamlibErrorCode
	Description string
}

//implement interface for Sorting Levels
func (slice Values) Len() int {
	return len(slice)
}

//implement interface for Sorting Levels
func (slice Values) Less(i, j int) bool {
	return slice[i].Name < slice[j].Name
}

//implement interface for Sorting Levels
func (slice Values) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

// CIntToBool ....
func CIntToBool(myInt C.int) (bool, error) {
	i := int(myInt)
	if i == 0 {
		return false, nil
	} else if i == 1 {
		return true, nil
	}
	return false, errors.New("Unable to convert C.int to bool")
}

// BoolToCint ...
func BoolToCint(myBool bool) (C.int, error) {
	if myBool {
		return C.int(1), nil
	}

	return C.int(0), nil
}

func (e *HamlibError) Error() string {
	switch e.Errorcode {
	case HamlibErrOK:
		e.Description = "OK"
	case HamlibErrEINVAL:
		e.Description = "invalid parameter"
	case HamlibErrECONF:
		e.Description = "invalid configuration (serial,..)"
	case HamlibErrENOMEM:
		e.Description = "memory shortage"
	case HamlibErrENIMPL:
		e.Description = "function not implemented, but will be"
	case HamlibErrETIMEOUT:
		e.Description = "communication timed out"
	case HamlibErrEIO:
		e.Description = "IO error, including open failed"
	case HamlibErrEINTERNAL:
		e.Description = "Internal Hamlib error, huh!"
	case HamlibErrEPROTO:
		e.Description = "Protocol error"
	case HamlibErrERJCTED:
		e.Description = "Command rejected by the rig"
	case HamlibErrETRUNC:
		e.Description = "Command performed, but arg truncated"
	case HamlibErrENAVAIL:
		e.Description = "function not available"
	case HamlibErrENTARGET:
		e.Description = "VFO not targetable"
	case HamlibErrBUSERROR:
		e.Description = "Error talking on the bus"
	case HamlibErrBUSBUSY:
		e.Description = "Collision on the bus"
	case HamlibErrEARG:
		e.Description = "NULL RIG handle or any invalid pointer parameter in get arg"
	case HamlibErrEVFO:
		e.Description = "Invalid VFO"
	case HamlibErrEDOM:
		e.Description = "Argument out of domain of func"
	default:
		e.Description = "unknown Error"
	}
	return fmt.Sprintf("%s: %s (%v)", e.Operation, e.Description, e.Errorcode)
}

type Error struct {
	Operation       string
	UnderlyingError error
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %v", e.Operation, e.UnderlyingError)
}
