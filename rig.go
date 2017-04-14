// +build linux,cgo darwin,cgo

package goHamlib

/*
#cgo CFLAGS: -I /usr/local/include
#cgo LDFLAGS: -L /usr/local/lib -lhamlib

#include <stdio.h>
#include <stdlib.h>
#include <hamlib/rig.h>

extern int set_port(int rig_port_type, char* portname, int baudrate, int databits, int stopbits, int parity, int handshake);
extern int init_rig(int rig_model);
extern int open_rig();
extern int has_set_vfo();
extern int set_vfo(int vfo);
extern int has_get_vfo();
extern int get_vfo(int *vfo);
extern int has_set_freq();
extern int set_freq(int vfo, double freq);
extern int has_set_mode();
extern int set_mode(int vfo, int mode, long pb_width);
extern int get_passband_narrow(int mode);
extern int get_passband_normal(int mode);
extern int get_passband_wide(int mode);
extern int has_get_freq();
extern int get_freq(int vfo, double *freq);
extern int has_get_mode();
extern int get_mode(int vfo, int *mode, long *pb_width);
extern int has_set_ptt();
extern int set_ptt(int vfo, int ptt);
extern int has_get_ptt();
extern int get_ptt(int vfo, int *ptt);
extern int has_set_rit();
extern int set_rit(int vfo, int offset);
extern int has_get_rit();
extern int get_rit(int vfo, int *offset);
extern int has_set_xit();
extern int set_xit(int vfo, int offset);
extern int has_get_xit();
extern int get_xit(int vfo, int *offset);
extern int has_set_split_freq();
extern int set_split_freq(int vfo, double tx_freq);
extern int has_get_split_freq();
extern int get_split_freq(int vfo, double *tx_freq);
extern int has_set_split_mode();
extern int set_split_mode(int vfo, int tx_mode, int tx_width);
extern int has_get_split_mode();
extern int get_split_mode(int vfo, int *tx_mode, long *tx_width);
extern int has_set_split_vfo();
extern int set_split_vfo(int vfo, int split, int tx_vfo);
extern int has_get_split_vfo();
extern int get_split_vfo(int vfo, int *split, int *tx_vfo);
extern int has_set_powerstat();
extern int set_powerstat(int status);
extern int has_get_powerstat();
extern int get_powerstat(int *status);
extern const char* get_info();
extern int has_set_ant();
extern int set_ant(int vfo, int ant);
extern int has_get_ant();
extern int get_ant(int vfo, int *ant);
extern int has_set_ts();
extern int set_ts(int vfo, int ts);
extern int has_get_ts();
extern int get_ts(int vfo, long *ts);
extern signed long get_rig_resolution(int mode);
extern unsigned long has_get_level(unsigned long level);
extern unsigned long has_set_level(unsigned long level);
extern unsigned long has_get_func(unsigned long function);
extern unsigned long has_set_func(unsigned long function);
extern unsigned long has_get_parm(unsigned long parm);
extern unsigned long has_set_parm(unsigned long parm);
extern int has_token(char *token);
extern int has_get_conf();
extern int get_conf(char *token, char* val);
extern int has_set_conf();
extern int set_conf(char *token, char *val);
extern int get_level(int vfo, unsigned long level, float *value);
extern int set_level(int vfo, unsigned long level, float value);
extern int get_level_gran(unsigned long level, float *step, float *min, float *max);
extern int get_func(int vfo, unsigned long function, int *value);
extern int set_func(int vfo, unsigned long function, int value);
extern int get_parm(unsigned long parm, float *value);
extern int set_parm(unsigned long parm, float value);
extern int get_parm_gran(unsigned long parm, float *step, float *min, float *max);
extern int vfo_op(int vfo, int op);
extern int get_caps_max_rit(int *rit);
extern int get_caps_max_xit(int *xit);
extern int get_caps_max_if_shift(int *if_shift);
extern int* get_caps_attenuator_list_pointer_and_length(int *length);
extern int* get_caps_preamp_list_pointer_and_length(int *length);
extern int get_supported_vfos(int *vfo_list);
extern int get_supported_vfo_operations(int *vfo_ops);
extern int get_supported_modes(int *modes);
extern int get_filter_count(int *filter_count);
extern int get_filter_mode_width(int filter, int *mode, signed long *width);
extern int get_ts_count(int *ts_count);
extern int get_tuning_steps(int el, int *mode, signed long *ts);
extern void get_model_name(char *model_name);
extern void get_version(char *version);
extern void get_mfg_name(char *mfg_name);
extern int get_status();
extern int get_int_from_array(int *array, int *el, int index);
extern void set_debug_level(int debug_level);
extern int close_rig();
extern int cleanup_rig();

*/
import "C"

import (
	"errors"
	"log"
	"sort"
	"strings"
	"unsafe"
	//"encoding/hex"
)

// Initialize Rig
func (rig *Rig) Init(rigModel int) error {

	if rigModel <= 0 {
		return checkError(RIG_EINVAL, errors.New("invalid rig model"), "init_rig")
	}

	res, err := C.init_rig(C.int(rigModel))
	if err != nil {
		return checkError(res, err, "init_rig")
	}
	err = rig.getCaps()

	rig.Caps.RigModel = rigModel

	return checkError(res, err, "init_rig")
}

// Set Port of Rig
func (rig *Rig) SetPort(p Port) error {
	res, err := C.set_port(C.int(p.RigPortType), C.CString(p.Portname), C.int(p.Baudrate), C.int(p.Databits), C.int(p.Stopbits), C.int(p.Parity), C.int(p.Handshake))
	return checkError(res, err, "set_port")
}

// Open Radio / Port
func (rig *Rig) Open() error {
	res, err := C.open_rig()
	return checkError(res, err, "open_rig")
}

// Set default VFO
func (rig *Rig) SetVfo(vfo int) error {
	res, err := C.set_vfo(C.int(vfo))
	return checkError(res, err, "set_vfo")
}

// Get default VFO
func (rig *Rig) GetVfo() (vfo int, err error) {
	var v C.int
	res, err := C.get_vfo(&v)
	vfo = int(v)
	return vfo, checkError(res, err, "get_vfo")
}

// Set Frequency for a VFO
func (rig *Rig) SetFreq(vfo int, freq float64) error {
	res, err := C.set_freq(C.int(vfo), C.double(freq))
	return checkError(res, err, "set_freq")
}

// Set Mode for a VFO
func (rig *Rig) SetMode(vfo int, mode int, pb_width int) error {
	res, err := C.set_mode(C.int(vfo), C.int(mode), C.long(pb_width))
	return checkError(res, err, "set_freq")
}

// Find the next suitable narrow available filter
func (rig *Rig) GetPbNarrow(mode int) (int, error) {
	pb, err := C.get_passband_narrow(C.int(mode))
	pb_width := int(pb)

	return pb_width, err
}

// Find the next suitable normal available filter
func (rig *Rig) GetPbNormal(mode int) (int, error) {
	pb, err := C.get_passband_normal(C.int(mode))
	pb_width := int(pb)

	return pb_width, err
}

// Find the next suitable wide available filter
func (rig *Rig) GetPbWide(mode int) (int, error) {
	pb, err := C.get_passband_wide(C.int(mode))
	pb_width := int(pb)

	return pb_width, err
}

// Get Frequency from a VFO
func (rig *Rig) GetFreq(vfo int) (freq float64, err error) {
	var f C.double
	var res C.int
	res, err = C.get_freq(C.int(vfo), &f)
	freq = float64(f)
	return freq, checkError(res, err, "get_freq")
}

// Get Mode and Passband width for a VFO
func (rig *Rig) GetMode(vfo int) (mode int, pb_width int, err error) {
	var m C.int
	var pb C.long
	var res C.int
	res, err = C.get_mode(C.int(vfo), &m, &pb)
	pb_width = int(pb)
	mode = int(m)
	return mode, pb_width, checkError(res, err, "get_mode")
}

// Set Ptt
func (rig *Rig) SetPtt(vfo int, ptt int) error {
	res, err := C.set_ptt(C.int(vfo), C.int(ptt))
	return checkError(res, err, "set_ptt")
}

// Get Ptt state
func (rig *Rig) GetPtt(vfo int) (ptt int, err error) {
	var p C.int
	res, err := C.get_ptt(C.int(vfo), &p)
	ptt = int(p)
	return ptt, checkError(res, err, "get_ptt")
}

// Set Rit offset value
func (rig *Rig) SetRit(vfo int, offset int) error {
	res, err := C.set_rit(C.int(vfo), C.int(offset))
	return checkError(res, err, "set_rit")
}

// Get Rit offset value
func (rig *Rig) GetRit(vfo int) (offset int, err error) {
	var o C.int
	res, err := C.get_rit(C.int(vfo), &o)
	offset = int(o)
	return offset, checkError(res, err, "get_rit")
}

// Set Xit offset value
func (rig *Rig) SetXit(vfo int, offset int) error {
	res, err := C.set_xit(C.int(vfo), C.int(offset))
	return checkError(res, err, "set_xit")
}

// Get Xit offset value
func (rig *Rig) GetXit(vfo int) (offset int, err error) {
	var o C.int
	res, err := C.get_xit(C.int(vfo), &o)
	offset = int(o)
	return offset, checkError(res, err, "get_xit")
}

// Set Split Frequency
func (rig *Rig) SetSplitFreq(vfo int, txFreq float64) error {
	res, err := C.set_split_freq(C.int(vfo), C.double(txFreq))
	return checkError(res, err, "set_split_freq")
}

// Get Split Frequency
func (rig *Rig) GetSplitFreq(vfo int) (txFreq float64, err error) {
	var f C.double
	res, err := C.get_split_freq(C.int(vfo), &f)
	txFreq = float64(f)
	return txFreq, checkError(res, err, "get_split_freq")
}

// Set Split Mode
func (rig *Rig) SetSplitMode(vfo int, txMode int, txWidth int) error {
	res, err := C.set_split_mode(C.int(vfo), C.int(txMode), C.int(txWidth))
	return checkError(res, err, "set_split_mode")
}

// Get Split Mode
func (rig *Rig) GetSplitMode(vfo int) (txMode int, txWidth int, err error) {
	var m C.int
	var w C.long
	res, err := C.get_split_mode(C.int(vfo), &m, &w)
	txMode = int(m)
	txWidth = int(w)
	return txMode, txWidth, checkError(res, err, "get_split_mode")
}

// Set Split Vfo
func (rig *Rig) SetSplitVfo(vfo int, split int, txVfo int) error {
	res, err := C.set_split_vfo(C.int(vfo), C.int(split), C.int(txVfo))
	return checkError(res, err, "set_split_vfo")
}

// Get Split Vfo
func (rig *Rig) GetSplitVfo(vfo int) (split int, txVfo int, err error) {
	var s C.int
	var v C.int
	res, err := C.get_split_vfo(C.int(vfo), &s, &v)
	split = int(s)
	txVfo = int(v)
	return split, txVfo, checkError(res, err, "get_split_vfo")
}

// Get Split (shortcut for GetSplitVfo)
func (rig *Rig) GetSplit(vfo int) (split int, txVfo int, err error) {
	var s C.int
	var t C.int
	res, err := C.get_split_vfo(C.int(vfo), &s, &t)
	split = int(s)
	txVfo = int(t)
	return split, txVfo, checkError(res, err, "get_split")
}

// Set Rig Power On/Off/Standby
func (rig *Rig) SetPowerStat(status int) error {
	res, err := C.set_powerstat(C.int(status))
	return checkError(res, err, "set_powerstat")
}

// Get Rig Power On/Off/Standby
func (rig *Rig) GetPowerStat() (status int, err error) {
	var s C.int
	var res C.int
	res, err = C.get_powerstat(&s)
	status = int(s)
	return status, checkError(res, err, "get_powerstat")
}

// Get Rig info
func (rig *Rig) GetInfo() (info string, err error) {
	i, err := C.get_info()
	info = C.GoString(i)
	return info, checkError(C.int(0), err, "get_info")
}

// Set Antenna
func (rig *Rig) SetAnt(vfo int, ant int) error {
	res, err := C.set_ant(C.int(vfo), C.int(ant))
	return checkError(res, err, "set_ant")
}

// Get Antenna
func (rig *Rig) GetAnt(vfo int) (ant int, err error) {
	var a C.int
	res, err := C.get_ant(C.int(vfo), &a)
	ant = int(a)
	return ant, checkError(res, err, "get_ant")
}

// Set Tuning step
func (rig *Rig) SetTs(vfo int, ts int) error {
	res, err := C.set_ts(C.int(vfo), C.int(ts))
	return checkError(res, err, "set_ts")
}

// Get Tuning step
func (rig *Rig) GetTs(vfo int) (ts int, err error) {
	var t C.long
	res, err := C.get_ts(C.int(vfo), &t)
	ts = int(t)
	return ts, checkError(res, err, "get_ts")
}

// has supports getting a specific level
func (rig *Rig) HasGetLevel(level uint32) (res uint32, err error) {
	var c C.ulong
	c, err = C.has_get_level(C.ulong(level))
	res = uint32(c)
	return res, checkError(0, err, "has_get_level")
}

// get the best frequency resolution for this rig (minimum step size)
func (rig *Rig) GetRigResolution(mode int) (resolution int, err error) {
	var r C.long
	r, err = C.get_rig_resolution(C.int(mode))
	resolution = int(r)
	return resolution, checkError(0, err, "get_rig_resolution")
}

// has supports setting a specific level
func (rig *Rig) HasSetLevel(level uint32) (res uint32, err error) {
	var c C.ulong
	c, err = C.has_set_level(C.ulong(level))
	res = uint32(c)
	return res, checkError(0, err, "has_set_level")
}

// has supports getting a specific function
func (rig *Rig) HasGetFunc(function uint32) (res uint32, err error) {
	var c C.ulong
	c, err = C.has_get_func(C.ulong(function))
	res = uint32(c)
	return res, checkError(0, err, "has_get_func")
}

// has supports setting a specific function
func (rig *Rig) HasSetFunc(function uint32) (res uint32, err error) {
	var c C.ulong
	c, err = C.has_set_func(C.ulong(function))
	res = uint32(c)
	return res, checkError(0, err, "has_set_func")
}

// has supports getting a specific parameter
func (rig *Rig) HasGetParm(parm uint32) (res uint32, err error) {
	var c C.ulong
	c, err = C.has_get_parm(C.ulong(parm))
	res = uint32(c)
	return res, checkError(0, err, "has_get_parm")
}

// has supports setting a specific parameter
func (rig *Rig) HasSetParm(parm uint32) (res uint32, err error) {
	var c C.ulong
	c, err = C.has_set_parm(C.ulong(parm))
	res = uint32(c)
	return res, checkError(0, err, "has_set_parm")
}

//get Level
func (rig *Rig) GetLevel(vfo int, level uint32) (value float32, err error) {
	var v C.float
	var res C.int
	res, err = C.get_level(C.int(vfo), C.ulong(level), &v)
	value = float32(v)
	return value, checkError(res, err, "get_level")
}

//set Level
func (rig *Rig) SetLevel(vfo int, level uint32, value float32) error {
	res, err := C.set_level(C.int(vfo), C.ulong(level), C.float(value))
	return checkError(res, err, "set_level")
}

//Get granularity (stepsize, minimum, maximum) for a Level
func (rig *Rig) GetLevelGran(level uint32) (step float32, min float32, max float32, err error) {
	var cStep C.float
	var cMin C.float
	var cMax C.float

	res, err := C.get_level_gran(C.ulong(level), &cStep, &cMin, &cMax)
	if checkError(res, err, "get_level_gran") != nil {
		return 0, 0, 0, err
	}

	return float32(cStep), float32(cMin), float32(cMax), nil
}

//get Function
func (rig *Rig) GetFunc(vfo int, function uint32) (value bool, err error) {
	var v C.int
	var res C.int
	res, err = C.get_func(C.int(vfo), C.ulong(function), &v)
	value, err2 := CIntToBool(v)
	if err2 != nil { //not so nice...
		return value, checkError(0, err2, "get_func")
	}
	return value, checkError(res, err, "get_func")
}

//set Function
func (rig *Rig) SetFunc(vfo int, function uint32, value bool) error {
	var v C.int
	v, err := BoolToCint(value)
	if err != nil {
		return checkError(0, err, "set_func")
	}
	res, err := C.set_func(C.int(vfo), C.ulong(function), v)
	return checkError(res, err, "set_func")
}

//get Parameter
func (rig *Rig) GetParm(vfo int, parm uint32) (value float32, err error) {
	var v C.float
	var res C.int
	res, err = C.get_parm(C.ulong(parm), &v)
	value = float32(v)
	return value, checkError(res, err, "get_parm")
}

//set Parameter
func (rig *Rig) SetParm(vfo int, parm uint32, value float32) error {
	res, err := C.set_parm(C.ulong(parm), C.float(value))
	return checkError(res, err, "set_parm")
}

//Get granularity (stepsize, minimum, maximum) for a Parameter
func (rig *Rig) GetParmGran(parm uint32) (step float32, min float32, max float32, err error) {
	var cStep C.float
	var cMin C.float
	var cMax C.float

	res, err := C.get_parm_gran(C.ulong(parm), &cStep, &cMin, &cMax)
	if checkError(res, err, "get_parm_gran") != nil {
		return 0, 0, 0, err
	}

	return float32(cStep), float32(cMin), float32(cMax), nil
}

//Set configuration token
func (rig *Rig) SetConf(token string, val string) error {
	res, err := C.set_conf(C.CString(token), C.CString(val))
	return checkError(res, err, "set_conf")
}

//HasToken checks if the rig supports a given token
func (rig *Rig) HasToken(token string) bool {
	//dirty hack - provide fix length char*
	//there should be a better way

	res, _ := C.has_token(C.CString(token))

	if res == RIG_OK {
		return true
	}

	return false
}

//Get configuration token
func (rig *Rig) GetConf(token string) (val string, err error) {
	//dirty hack - provide fix length char*
	//there should be a better way
	v := C.CString("                                                          ")

	res, err := C.get_conf(C.CString(token), v)
	val = C.GoString(v)
	C.free(unsafe.Pointer(v))

	return val, checkError(res, err, "get_conf")
}

//Execute VFO Operation
func (rig *Rig) VfoOp(vfo int, op int) error {
	res, err := C.vfo_op(C.int(vfo), C.int(op))
	return checkError(res, err, "vfo_op")
}

//Copy capabilities into Rig->Caps struct
func (rig *Rig) getCaps() error {
	if err := rig.getMaxRit(); err != nil {
		log.Println(err)
	}
	if err := rig.getMaxXit(); err != nil {
		log.Println(err)
	}
	if err := rig.getMaxIfShift(); err != nil {
		log.Println(err)
	}
	if err := rig.getAttenuators(); err != nil {
		log.Println(err)
	}
	if err := rig.getPreamps(); err != nil {
		log.Println(err)
	}
	if err := rig.getVfos(); err != nil {
		log.Println(err)
	}
	if err := rig.getOperations(); err != nil {
		log.Println(err)
	}
	if err := rig.getModes(); err != nil {
		log.Println(err)
	}
	if err := rig.getGetFunctions(); err != nil {
		log.Println(err)
	}
	if err := rig.getSetFunctions(); err != nil {
		log.Println(err)
	}
	if err := rig.getGetLevels(); err != nil {
		log.Println(err)
	}
	if err := rig.getSetLevels(); err != nil {
		log.Println(err)
	}
	if err := rig.getGetParameter(); err != nil {
		log.Println(err)
	}
	if err := rig.getSetParameter(); err != nil {
		log.Println(err)
	}
	if err := rig.getFilters(); err != nil {
		log.Println(err)
	}
	if err := rig.getTuningSteps(); err != nil {
		log.Println(err)
	}
	rig.getModelName()
	rig.getVersion()
	rig.getMfgName()
	rig.getStatus()
	rig.hasGetPowerStat()
	rig.hasSetPowerStat()
	rig.hasGetVfo()
	rig.hasSetVfo()
	rig.hasSetFreq()
	rig.hasGetFreq()
	rig.hasGetMode()
	rig.hasSetMode()
	rig.hasSetPtt()
	rig.hasSetPtt()
	rig.hasGetRit()
	rig.hasSetRit()
	rig.hasGetXit()
	rig.hasSetXit()
	rig.hasSetSplitVfo()
	rig.hasGetSplitVfo()
	rig.hasGetSplitFreq()
	rig.hasSetSplitVfo()
	rig.hasSetSplitMode()
	rig.hasGetSplitMode()
	rig.hasGetAnt()
	rig.hasSetAnt()
	rig.hasSetTs()
	rig.hasGetTs()
	rig.hasGetConf()
	rig.hasSetConf()

	return nil

}

//get Capabilities > Max Rit
func (rig *Rig) getMaxRit() error {
	var rit C.int
	res, err := C.get_caps_max_rit(&rit)
	if checkError(res, err, "get_caps_max_rit") != nil {
		return checkError(res, err, "get_caps_max_rit")
	}
	rig.Caps.MaxRit = int(rit)
	return nil
}

//get Capabilities > Max Xit
func (rig *Rig) getMaxXit() error {
	var xit C.int
	res, err := C.get_caps_max_xit(&xit)
	if checkError(res, err, "get_caps_max_xit") != nil {
		return checkError(res, err, "get_caps_max_xit")
	}
	rig.Caps.MaxXit = int(xit)
	return nil
}

//get Capabilities > Max IF Shift
func (rig *Rig) getMaxIfShift() error {
	var ifShift C.int
	res, err := C.get_caps_max_if_shift(&ifShift)
	if checkError(res, err, "get_caps_max_if_shift") != nil {
		return checkError(res, err, "get_caps_max_if_shift")
	}
	rig.Caps.MaxIfShift = int(ifShift)
	return nil
}

//get Capabilities > List of supported Attenuators
func (rig *Rig) getAttenuators() error {

	var att_array *C.int
	var length C.int
	var el C.int

	att_array, err := C.get_caps_attenuator_list_pointer_and_length(&length)
	if att_array == nil {
		return &HamlibError{"getAttenuators", int(RIG_EINTERNAL), "invalid pointer"}
	}
	if err != nil {
		return &Error{"getAttenuators", err}
	}

	var att []int
	for i := 0; i < int(length); i++ {
		C.get_int_from_array(att_array, &el, C.int(i))

		if int(el) == 0 {
			break
		}

		att = append(att, int(el))
	}

	rig.Caps.Attenuators = att
	return nil
}

//get Capabilities > List of supported Preamp Levels
func (rig *Rig) getPreamps() error {

	var preamp_array *C.int
	var length C.int
	var el C.int

	preamp_array, err := C.get_caps_preamp_list_pointer_and_length(&length)
	if preamp_array == nil {
		return &HamlibError{"getPreamp", int(RIG_EINTERNAL), "invalid pointer"}
	}
	if err != nil {
		return &Error{"getPreamp", err}
	}

	var preamps []int
	for i := 0; i < int(length); i++ {
		C.get_int_from_array(preamp_array, &el, C.int(i))

		if int(el) == 0 {
			break
		}

		preamps = append(preamps, int(el))
	}

	rig.Caps.Preamps = preamps
	return nil
}

//get Capabilities > List of supported VFOs
func (rig *Rig) getVfos() error {
	var vfoClist C.int
	var vfoList []string

	res, err := C.get_supported_vfos(&vfoClist)
	if checkError(res, err, "get_supported_vfos") != nil {
		return checkError(res, err, "get_supported_vfos")
	}

	for vfo, vfoStr := range VfoName {
		if int(vfoClist)&vfo > 0 {
			vfoList = append(vfoList, vfoStr)
		}
	}
	sort.Strings(vfoList)
	rig.Caps.Vfos = vfoList
	return nil
}

//get Capabilities > List of supported VFO Operations
func (rig *Rig) getOperations() error {
	var vfoOpClist C.int
	var vfoOpList []string

	res, err := C.get_supported_vfo_operations(&vfoOpClist)
	if checkError(res, err, "get_supported_vfo_operations") != nil {
		return checkError(res, err, "get_supported_vfo_operations")
	}

	for op, opStr := range OperationName {
		if int(vfoOpClist)&op > 0 {
			vfoOpList = append(vfoOpList, opStr)
		}
	}
	sort.Strings(vfoOpList)
	rig.Caps.Operations = vfoOpList
	return nil
}

//get Capabilities > List of supported Modes
func (rig *Rig) getModes() error {
	var modesClist C.int
	var modesList []string

	res, err := C.get_supported_modes(&modesClist)
	if checkError(res, err, "get_supported_modes") != nil {
		return checkError(res, err, "get_supported_modes")
	}

	for mode, modeStr := range ModeName {
		if int(modesClist)&mode > 0 {
			modesList = append(modesList, modeStr)
		}
	}
	sort.Strings(modesList)
	rig.Caps.Modes = modesList
	return nil
}

//get Capabilities > List of supported Functions that can be read
func (rig *Rig) getGetFunctions() error {
	var funcList []string

	for f, fStr := range FuncName {
		if res, err := rig.HasGetFunc(f); err != nil {
			return err
		} else {
			if res > 0 {
				funcList = append(funcList, fStr)
			}
		}
	}
	sort.Strings(funcList)
	rig.Caps.GetFunctions = funcList
	return nil
}

//get Capabilities > List of supported Functions that can be set
func (rig *Rig) getSetFunctions() error {
	var funcList []string

	for f, fStr := range FuncName {
		if res, err := rig.HasSetFunc(f); err != nil {
			return err
		} else {
			if res > 0 {
				funcList = append(funcList, fStr)
			}
		}
	}
	sort.Strings(funcList)
	rig.Caps.SetFunctions = funcList
	return nil
}

//get Capabilities > List of supported Levels that can be read
func (rig *Rig) getGetLevels() error {
	var levelList Values

	for l, lStr := range LevelName {

		res, err := rig.HasGetLevel(l)
		if err != nil {
			return err
		}

		if res > 0 {
			var level Value
			level.Step, level.Min, level.Max, err = rig.GetLevelGran(l)
			if err != nil {
				return err
			}
			level.Name = lStr
			levelList = append(levelList, level)
		}
	}
	sort.Sort(levelList)
	rig.Caps.GetLevels = levelList
	return nil
}

//get Capabilities > List of supported Levels that can be set
func (rig *Rig) getSetLevels() error {
	var levelList Values

	for l, lStr := range LevelName {
		res, err := rig.HasSetLevel(l)
		if err != nil {
			return err
		}

		if res > 0 {
			var level Value
			level.Step, level.Min, level.Max, err = rig.GetLevelGran(l)
			if err != nil {
				return err
			}
			level.Name = lStr
			levelList = append(levelList, level)
		}
	}
	sort.Sort(levelList)
	rig.Caps.SetLevels = levelList
	return nil
}

//get Capabilities > List of supported Parameters that can be read
func (rig *Rig) getGetParameter() error {
	var parmList Values

	for p, pStr := range ParmName {
		if res, err := rig.HasGetParm(p); err != nil {
			return err
		} else {
			if res > 0 {
				var parm Value
				parm.Step, parm.Min, parm.Max, err = rig.GetParmGran(p)
				if err != nil {
					return err
				}
				parm.Name = pStr
				parmList = append(parmList, parm)
			}
		}
	}
	sort.Sort(parmList)
	rig.Caps.GetParameters = parmList
	return nil
}

//get Capabilities > List of supported Parameters that can be set
func (rig *Rig) getSetParameter() error {
	var parmList Values

	for p, pStr := range ParmName {
		if res, err := rig.HasSetParm(p); err != nil {
			return err
		} else {
			if res > 0 {
				var parm Value
				parm.Step, parm.Min, parm.Max, err = rig.GetParmGran(p)
				if err != nil {
					return err
				}
				parm.Name = pStr
				parmList = append(parmList, parm)
			}
		}
	}
	sort.Sort(parmList)
	rig.Caps.SetParameters = parmList
	return nil
}

// get Capabilities > Rig Model Name
func (rig *Rig) getModelName() {
	cModelName := C.CString("                         ")

	C.get_model_name(cModelName)
	rig.Caps.ModelName = strings.TrimSpace(C.GoString(cModelName))
	C.free(unsafe.Pointer(cModelName))
}

// get Capabilities > Version
func (rig *Rig) getVersion() {
	cVersion := C.CString("                         ")

	C.get_version(cVersion)
	rig.Caps.Version = strings.TrimSpace(C.GoString(cVersion))
	C.free(unsafe.Pointer(cVersion))
}

// get Capabilities > Manufacturer Name
func (rig *Rig) getMfgName() {
	cMfGName := C.CString("                         ")

	C.get_mfg_name(cMfGName)
	rig.Caps.MfgName = strings.TrimSpace(C.GoString(cMfGName))
	C.free(unsafe.Pointer(cMfGName))
}

// get Capabilities > Status
func (rig *Rig) getStatus() {

	cStatus := C.get_status()
	rig.Caps.Status = int(cStatus)
}

func (rig *Rig) getFilters() error {
	var cfc C.int
	var cWidth C.long
	var cMode C.int

	var filterMap map[string][]int
	filterMap = make(map[string][]int)

	res, err := C.get_filter_count(&cfc)
	if checkError(res, err, "get_filter_count") != nil {
		return checkError(res, err, "get_filter_count")
	}

	for i := 0; i < int(cfc); i++ {
		res, err = C.get_filter_mode_width(C.int(i), &cMode, &cWidth)
		if checkError(res, err, "") != nil {
			return checkError(res, err, "get_filter_mode_width")
		}
		for mode, modeStr := range ModeName {
			if int(cMode)&mode > 0 {
				filterMap[modeStr] = append(filterMap[modeStr], int(cWidth))
			}
		}
	}

	rig.Caps.Filters = filterMap
	return nil
}

//get supported Tuning steps
func (rig *Rig) getTuningSteps() error {

	var tsc C.int
	var cMode C.int
	var cTs C.long
	var tsMap map[string][]int
	tsMap = make(map[string][]int)

	res, err := C.get_ts_count(&tsc)
	if checkError(res, err, "get_ts_count") != nil {
		return checkError(res, err, "get_ts_count")
	}

	for i := 0; i < int(tsc); i++ {
		res, err = C.get_tuning_steps(C.int(i), &cMode, &cTs)
		if checkError(res, err, "") != nil {
			return checkError(res, err, "get_tuning_steps")
		}
		for mode, modeStr := range ModeName {
			if int(cMode)&mode > 0 {
				tsMap[modeStr] = append(tsMap[modeStr], int(cTs))
			}
		}
	}

	rig.Caps.TuningSteps = tsMap
	return nil
}

func (rig *Rig) hasSetPowerStat() {
	res := C.has_set_powerstat()
	if res == RIG_OK {
		rig.Caps.HasSetPowerStat = true
		return
	}
	rig.Caps.HasSetPowerStat = false
}

func (rig *Rig) hasGetPowerStat() {
	res := C.has_get_powerstat()
	if res == RIG_OK {
		rig.Caps.HasGetPowerStat = true
		return
	}
	rig.Caps.HasGetPowerStat = false
}

func (rig *Rig) hasSetVfo() {
	res := C.has_set_vfo()
	if res == RIG_OK {
		rig.Caps.HasSetVfo = true
		return
	}
	rig.Caps.HasSetVfo = false
}

func (rig *Rig) hasGetVfo() {
	res := C.has_get_vfo()
	if res == RIG_OK {
		rig.Caps.HasGetVfo = true
		return
	}
	rig.Caps.HasGetVfo = false
}

func (rig *Rig) hasSetFreq() {
	res := C.has_set_freq()
	if res == RIG_OK {
		rig.Caps.HasSetFreq = true
		return
	}
	rig.Caps.HasSetFreq = false
}

func (rig *Rig) hasGetFreq() {
	res := C.has_get_freq()
	if res == RIG_OK {
		rig.Caps.HasGetFreq = true
		return
	}
	rig.Caps.HasGetFreq = false
}

func (rig *Rig) hasSetMode() {
	res := C.has_set_mode()
	if res == RIG_OK {
		rig.Caps.HasSetMode = true
		return
	}
	rig.Caps.HasSetMode = false
}

func (rig *Rig) hasGetMode() {
	res := C.has_get_mode()
	if res == RIG_OK {
		rig.Caps.HasGetMode = true
		return
	}
	rig.Caps.HasGetMode = false
}

func (rig *Rig) hasSetPtt() {
	res := C.has_set_ptt()
	if res == RIG_OK {
		rig.Caps.HasSetPtt = true
		return
	}
	rig.Caps.HasSetPtt = false
}

func (rig *Rig) hasGetPtt() {
	res := C.has_get_ptt()
	if res == RIG_OK {
		rig.Caps.HasGetPtt = true
		return
	}
	rig.Caps.HasGetPtt = false
}

func (rig *Rig) hasSetRit() {
	res := C.has_set_rit()
	if res == RIG_OK {
		rig.Caps.HasSetRit = true
		return
	}
	rig.Caps.HasSetRit = false
}

func (rig *Rig) hasGetRit() {
	res := C.has_get_rit()
	if res == RIG_OK {
		rig.Caps.HasGetRit = true
		return
	}
	rig.Caps.HasGetRit = false
}

func (rig *Rig) hasSetXit() {
	res := C.has_set_xit()
	if res == RIG_OK {
		rig.Caps.HasSetXit = true
		return
	}
	rig.Caps.HasSetXit = false
}

func (rig *Rig) hasGetXit() {
	res := C.has_get_xit()
	if res == RIG_OK {
		rig.Caps.HasGetXit = true
		return
	}
	rig.Caps.HasGetXit = false
}

func (rig *Rig) hasSetSplitVfo() {
	res := C.has_set_split_vfo()
	if res == RIG_OK {
		rig.Caps.HasSetSplitVfo = true
		return
	}
	rig.Caps.HasSetSplitVfo = false
}

func (rig *Rig) hasGetSplitVfo() {
	res := C.has_get_split_vfo()
	if res == RIG_OK {
		rig.Caps.HasGetSplitVfo = true
		return
	}
	rig.Caps.HasGetSplitVfo = false
}

func (rig *Rig) hasSetSplitMode() {
	res := C.has_set_split_mode()
	if res == RIG_OK {
		rig.Caps.HasSetSplitMode = true
		return
	}
	rig.Caps.HasSetSplitMode = false
}

func (rig *Rig) hasGetSplitMode() {
	res := C.has_get_split_mode()
	if res == RIG_OK {
		rig.Caps.HasGetSplitMode = true
		return
	}
	rig.Caps.HasGetSplitMode = false
}

func (rig *Rig) hasSetSplitFreq() {
	res := C.has_set_split_freq()
	if res == RIG_OK {
		rig.Caps.HasSetSplitFreq = true
		return
	}
	rig.Caps.HasSetSplitFreq = false
}

func (rig *Rig) hasGetSplitFreq() {
	res := C.has_get_split_freq()
	if res == RIG_OK {
		rig.Caps.HasGetSplitFreq = true
		return
	}
	rig.Caps.HasGetSplitFreq = false
}

func (rig *Rig) hasSetAnt() {
	res := C.has_set_ant()
	if res == RIG_OK {
		rig.Caps.HasSetAnt = true
		return
	}
	rig.Caps.HasSetAnt = false
}

func (rig *Rig) hasGetAnt() {
	res := C.has_get_ant()
	if res == RIG_OK {
		rig.Caps.HasGetAnt = true
		return
	}
	rig.Caps.HasGetAnt = false
}

func (rig *Rig) hasSetTs() {
	res := C.has_set_ts()
	if res == RIG_OK {
		rig.Caps.HasSetTs = true
		return
	}
	rig.Caps.HasSetTs = false
}

func (rig *Rig) hasGetTs() {
	res := C.has_get_ts()
	if res == RIG_OK {
		rig.Caps.HasGetTs = true
		return
	}
	rig.Caps.HasGetTs = false
}

func (rig *Rig) hasSetConf() {
	res := C.has_set_conf()
	if res == RIG_OK {
		rig.Caps.HasSetConf = true
		return
	}
	rig.Caps.HasSetConf = false
}

func (rig *Rig) hasGetConf() {
	res := C.has_get_conf()
	if res == RIG_OK {
		rig.Caps.HasGetConf = true
		return
	}
	rig.Caps.HasGetConf = false
}

// Set Debug level
func (rig *Rig) SetDebugLevel(dbgLevel int) {
	C.set_debug_level(C.int(dbgLevel))
}

//Close the Communication with the Radio
func (rig *Rig) Close() error {
	res, err := C.close_rig()
	return checkError(res, err, "close_rig")
}

//Grabage collect Radio and free up memory
func (rig *Rig) Cleanup() error {
	res, err := C.cleanup_rig()
	return checkError(res, err, "cleanup_rig")
}

// Check Errors from Hamlib C calls. C Errors have a higher priority.
// Additional Information is provided for better debugging
func checkError(res C.int, e error, operation string) error {

	if e != nil {
		return &Error{operation, e}
	}
	if int(res) != RIG_OK {
		return &HamlibError{operation, int(res), ""}
	}

	return nil
}
