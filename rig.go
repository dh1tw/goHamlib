package goHamlib

/*
#cgo pkg-config: hamlib

#include <stdio.h>
#include <stdlib.h>
#include <hamlib/rig.h>

extern int set_port(RIG *myrig,int rig_port_type, char* portname, int baudrate, int databits, int stopbits, int parity, int handshake);
extern RIG * init_rig(int rig_model);
extern int open_rig(RIG *myrig);
extern int has_set_vfo(RIG *myrig);
extern int set_vfo(RIG *myrig,int vfo);
extern int has_get_vfo(RIG *myrig);
extern int get_vfo(RIG *myrig,int *vfo);
extern int has_set_freq(RIG *myrig);
extern int set_freq(RIG *myrig,int vfo, double freq);
extern int has_set_mode(RIG *myrig);
extern int set_mode(RIG *myrig, int vfo, int mode, int pb_width);
extern int get_passband_narrow(RIG *myrig, int mode);
extern int get_passband_normal(RIG *myrig, int mode);
extern int get_passband_wide(RIG *myrig, int mode);
extern int has_get_freq(RIG *myrig);
extern int get_freq(RIG *myrig, int vfo, double *freq);
extern int has_get_mode(RIG *myrig);
extern int get_mode(RIG *myrig, int vfo, int *mode, long *pb_width);
extern int has_set_ptt(RIG *myrig);
extern int set_ptt(RIG *myrig, int vfo, int ptt);
extern int has_get_ptt(RIG *myrig);
extern int get_ptt(RIG *myrig, int vfo, int *ptt);
extern int has_set_rit(RIG *myrig);
extern int set_rit(RIG *myrig, int vfo, int offset);
extern int has_get_rit(RIG *myrig);
extern int get_rit(RIG *myrig, int vfo, long *offset);
extern int has_set_xit(RIG *myrig);
extern int set_xit(RIG *myrig, int vfo, int offset);
extern int has_get_xit(RIG *myrig);
extern int get_xit(RIG *myrig, int vfo, long *offset);
extern int has_set_split_freq(RIG *myrig);
extern int set_split_freq(RIG *myrig, int vfo, double tx_freq);
extern int has_get_split_freq(RIG *myrig);
extern int get_split_freq(RIG *myrig, int vfo, double *tx_freq);
extern int has_set_split_mode(RIG *myrig);
extern int set_split_mode(RIG *myrig, int vfo, int tx_mode, int tx_width);
extern int has_get_split_mode(RIG *myrig);
extern int get_split_mode(RIG *myrig, int vfo, int *tx_mode, long *tx_width);
extern int has_set_split_vfo(RIG *myrig);
extern int set_split_vfo(RIG *myrig, int vfo, int split, int tx_vfo);
extern int has_get_split_vfo(RIG *myrig);
extern int get_split_vfo(RIG *myrig, int vfo, int *split, int *tx_vfo);
extern int has_set_powerstat(RIG *myrig);
extern int set_powerstat(RIG *myrig, int status);
extern int has_get_powerstat(RIG *myrig);
extern int get_powerstat(RIG *myrig, int *status);
extern const char* get_info(RIG *myrig);
extern void get_model_name(RIG *myrig, char *rig_name);
extern void get_version(RIG *myrig, char *version);
extern void get_mfg_name(RIG *myrig, char *mfg_name);
extern int get_status(RIG *myrig);
extern int has_set_ant(RIG *myrig);
extern int set_ant(RIG *myrig, int vfo, int ant);
extern int has_get_ant(RIG *myrig);
extern int get_ant(RIG *myrig, int vfo, int *ant);
extern int has_set_ts(RIG *myrig);
extern int set_ts(RIG *myrig, int vfo, int ts);
extern int has_get_ts(RIG *myrig);
extern int get_ts(RIG *myrig, int vfo, long *ts);
extern signed long get_rig_resolution(RIG *myrig, int mode);
extern unsigned long has_get_level(RIG *myrig, unsigned long level);
extern unsigned long has_set_level(RIG *myrig, unsigned long level);
extern unsigned long has_get_func(RIG *myrig, unsigned long func);
extern unsigned long has_set_func(RIG *myrig, unsigned long func);
extern unsigned long has_get_parm(RIG *myrig, unsigned long parm);
extern unsigned long has_set_parm(RIG *myrig, unsigned long parm);
extern int has_token(RIG *myrig, char *token);
extern int has_get_conf(RIG *myrig);
extern int get_conf(RIG *myrig, char* token, char* val);
extern int has_set_conf(RIG *myrig);
extern int set_conf(RIG *myrig, char* token, char* val);
extern int get_level(RIG *myrig, int vfo, unsigned long level, float *value);
extern int get_level_gran(RIG *myrig, unsigned long level, float *step, float *min, float *max);
extern int set_level(RIG *myrig, int vfo, unsigned long level, float value);
extern int get_func(RIG *myrig, int vfo, unsigned long function, int *value);
extern int set_func(RIG *myrig, int vfo, unsigned long function, int value);
extern int get_parm(RIG *myrig, unsigned long parm, float *value);
extern int get_parm_gran(RIG *myrig, unsigned long parm, float *step, float *min, float *max);
extern int set_parm(RIG *myrig, unsigned long parm, float value);
extern int vfo_op(RIG *myrig, int vfo, int op);
extern int get_caps_max_rit(RIG *myrig, int *rit);
extern int get_caps_max_xit(RIG *myrig, int *xit);
extern int get_caps_max_if_shift(RIG *myrig, int *if_shift);
extern int* get_caps_attenuator_list_pointer_and_length(RIG *myrig, int *length);
extern int* get_caps_preamp_list_pointer_and_length(RIG *myrig, int *length);
extern int get_supported_vfos(RIG *myrig, int *vfo_list);
extern int get_supported_vfo_operations(RIG *myrig, int *vfo_ops);
extern int get_supported_modes(RIG *myrig, int *modes);
extern int get_filter_count(RIG *myrig, int *filter_count);
extern int get_filter_mode_width(RIG *myrig, int filter, int *mode, signed long *width);
extern int get_ts_count(RIG *myrig, int *ts_count);
extern int get_tuning_steps(RIG *myrig, int el, int *mode, signed long *ts);
extern int get_int_from_array(RIG *myrig, int *array, int *el, int index);
extern void set_debug_level(int debug_level);
extern int close_rig(RIG *myrig);
extern int cleanup_rig(RIG *myrig);

typedef struct rig_caps rig_caps_t;
extern int rig_list_foreach_wrap(void *list);
extern void set_debug_callback();
extern int internal_debug_cb(enum rig_debug_level_e debug_level, rig_ptr_t user_data, const char *fmt, va_list ap);
*/
import "C"

import (
	"errors"
	"log"
	"sort"
	"strings"
	"unsafe"
)

// RigModelID is a unique ID that identifies a particular rig driver
type RigModelID int

// RigModel is a supported rig in hamlib
type RigModel struct {
	Manufacturer string
	Model        string
	ModelID      RigModelID
}

// ListModels enumerates all of the hamlib supported rigs
func ListModels() []RigModel {
	C.rig_load_all_backends()
	knownRigs := []RigModel{}
	C.rig_list_foreach_wrap(unsafe.Pointer(&knownRigs))
	return knownRigs
}

//export go_rig_list_callback
func go_rig_list_callback(p unsafe.Pointer, d unsafe.Pointer) C.int {
	caps := (*C.rig_caps_t)(p)
	knownRigs := (*[]RigModel)(d)
	*knownRigs = append(*knownRigs, RigModel{
		Manufacturer: C.GoString(caps.mfg_name),
		Model:        C.GoString(caps.model_name),
		ModelID:      RigModelID(caps.rig_model),
	})
	return 1
}

// Set Debug level
func SetDebugLevel(dbgLevel DebugLevel) {
	C.set_debug_level(C.int(dbgLevel))
}

//export go_debug_print
func go_debug_print(lvl DebugLevel, msg *C.char) {
	if debug_cb_fn != nil {
		debug_cb_fn(lvl, C.GoString(msg))
	} else {
		log.Println(lvl, C.GoString(msg))
	}
}

var debug_cb_fn func(level DebugLevel, msg string)

// SetDebugCallback replaces the function that handles debug messages, preventing
// them from being written to STDOUT
func SetDebugCallback(fn func(level DebugLevel, msg string)) {
	debug_cb_fn = fn
	C.set_debug_callback()
}

// Initialize Rig
func (rig *Rig) Init(rigModel RigModelID) error {

	if rigModel <= 0 {
		return checkError(C.int(HamlibErrEINVAL), errors.New("invalid rig model"), "init_rig")
	}

	rig.handle = C.init_rig(C.int(rigModel))
	if rig.handle == nil {
		return checkError(0, errors.New("error initializing rig"), "init_rig")
	}
	err := rig.getCaps()

	rig.Caps.RigModel = int(rigModel)

	return checkError(0, err, "init_rig")
}

// Set Port of Rig
func (rig *Rig) SetPort(p Port) error {
	res, err := C.set_port(rig.handle, C.int(p.RigPortType), C.CString(p.Portname), C.int(p.Baudrate), C.int(p.Databits), C.int(p.Stopbits), C.int(p.Parity), C.int(p.Handshake))
	return checkError(res, err, "set_port")
}

// Open Radio / Port
func (rig *Rig) Open() error {
	res, err := C.open_rig(rig.handle)
	return checkError(res, err, "open_rig")
}

// SetVfo sets the default VFO
func (rig *Rig) SetVfo(vfo VFOType) error {
	res, err := C.set_vfo(rig.handle, C.int(vfo))
	return checkError(res, err, "set_vfo")
}

// GetVfo returns the default VFO
func (rig *Rig) GetVfo() (VFOType, error) {
	var v C.int
	res, err := C.get_vfo(rig.handle, &v)
	return VFOType(v), checkError(res, err, "get_vfo")
}

// SetFreq sets the Frequency for a VFO
func (rig *Rig) SetFreq(vfo VFOType, freq float64) error {
	res, err := C.set_freq(rig.handle, C.int(vfo), C.double(freq))
	return checkError(res, err, "set_freq")
}

// SetMode sets the Mode for a VFO
func (rig *Rig) SetMode(vfo VFOType, mode Mode, pbWidth int) error {
	res, err := C.set_mode(rig.handle, C.int(vfo), C.int(mode), C.int(pbWidth))
	return checkError(res, err, "set_mode")
}

// Find the next suitable narrow available filter
func (rig *Rig) GetPbNarrow(mode Mode) (int, error) {
	pb, err := C.get_passband_narrow(rig.handle, C.int(mode))
	pb_width := int(pb)

	return pb_width, err
}

// Find the next suitable normal available filter
func (rig *Rig) GetPbNormal(mode Mode) (int, error) {
	pb, err := C.get_passband_normal(rig.handle, C.int(mode))
	pb_width := int(pb)

	return pb_width, err
}

// Find the next suitable wide available filter
func (rig *Rig) GetPbWide(mode Mode) (int, error) {
	pb, err := C.get_passband_wide(rig.handle, C.int(mode))
	pb_width := int(pb)

	return pb_width, err
}

// Get Frequency from a VFO
func (rig *Rig) GetFreq(vfo VFOType) (freq float64, err error) {
	var f C.double
	var res C.int
	res, err = C.get_freq(rig.handle, C.int(vfo), &f)
	freq = float64(f)
	return freq, checkError(res, err, "get_freq")
}

// GetMode gets the Mode and Passband width for a VFO
func (rig *Rig) GetMode(vfo VFOType) (mode Mode, pb_width int, err error) {
	var m C.int
	var pb C.long
	var res C.int
	res, err = C.get_mode(rig.handle, C.int(vfo), &m, &pb)
	pb_width = int(pb)
	mode = Mode(m)
	return mode, pb_width, checkError(res, err, "get_mode")
}

// Set Ptt
func (rig *Rig) SetPtt(vfo VFOType, ptt int) error {
	res, err := C.set_ptt(rig.handle, C.int(vfo), C.int(ptt))
	return checkError(res, err, "set_ptt")
}

// Get Ptt state
func (rig *Rig) GetPtt(vfo VFOType) (ptt int, err error) {
	var p C.int
	res, err := C.get_ptt(rig.handle, C.int(vfo), &p)
	ptt = int(p)
	return ptt, checkError(res, err, "get_ptt")
}

// Set Rit offset value
func (rig *Rig) SetRit(vfo VFOType, offset int) error {
	res, err := C.set_rit(rig.handle, C.int(vfo), C.int(offset))
	return checkError(res, err, "set_rit")
}

// Get Rit offset value
func (rig *Rig) GetRit(vfo VFOType) (offset int, err error) {
	var o C.long
	res, err := C.get_rit(rig.handle, C.int(vfo), &o)
	offset = int(o)
	return offset, checkError(res, err, "get_rit")
}

// Set Xit offset value
func (rig *Rig) SetXit(vfo VFOType, offset int) error {
	res, err := C.set_xit(rig.handle, C.int(vfo), C.int(offset))
	return checkError(res, err, "set_xit")
}

// Get Xit offset value
func (rig *Rig) GetXit(vfo VFOType) (offset int, err error) {
	var o C.long
	res, err := C.get_xit(rig.handle, C.int(vfo), &o)
	offset = int(o)
	return offset, checkError(res, err, "get_xit")
}

// Set Split Frequency
func (rig *Rig) SetSplitFreq(vfo VFOType, txFreq float64) error {
	res, err := C.set_split_freq(rig.handle, C.int(vfo), C.double(txFreq))
	return checkError(res, err, "set_split_freq")
}

// Get Split Frequency
func (rig *Rig) GetSplitFreq(vfo VFOType) (txFreq float64, err error) {
	var f C.double
	res, err := C.get_split_freq(rig.handle, C.int(vfo), &f)
	txFreq = float64(f)
	return txFreq, checkError(res, err, "get_split_freq")
}

// Set Split Mode
func (rig *Rig) SetSplitMode(vfo VFOType, txMode Mode, txWidth int) error {
	res, err := C.set_split_mode(rig.handle, C.int(vfo), C.int(txMode), C.int(txWidth))
	return checkError(res, err, "set_split_mode")
}

// Get Split Mode
func (rig *Rig) GetSplitMode(vfo VFOType) (txMode Mode, txWidth int, err error) {
	var m C.int
	var w C.long
	res, err := C.get_split_mode(rig.handle, C.int(vfo), &m, &w)
	txMode = Mode(m)
	txWidth = int(w)
	return txMode, txWidth, checkError(res, err, "get_split_mode")
}

// Set Split Vfo
func (rig *Rig) SetSplitVfo(vfo VFOType, split int, txVfo VFOType) error {
	res, err := C.set_split_vfo(rig.handle, C.int(vfo), C.int(split), C.int(txVfo))
	return checkError(res, err, "set_split_vfo")
}

// Get Split Vfo
func (rig *Rig) GetSplitVfo(vfo VFOType) (split int, txVfo VFOType, err error) {
	var s C.int
	var v C.int
	res, err := C.get_split_vfo(rig.handle, C.int(vfo), &s, &v)
	split = int(s)
	txVfo = VFOType(v)
	return split, txVfo, checkError(res, err, "get_split_vfo")
}

// Get Split (shortcut for GetSplitVfo)
func (rig *Rig) GetSplit(vfo VFOType) (split int, txVfo VFOType, err error) {
	var s C.int
	var t C.int
	res, err := C.get_split_vfo(rig.handle, C.int(vfo), &s, &t)
	split = int(s)
	txVfo = VFOType(t)
	return split, txVfo, checkError(res, err, "get_split")
}

// SetPowerState sets the Rig Power On/Off/Standby state
func (rig *Rig) SetPowerState(status Power) error {
	res, err := C.set_powerstat(rig.handle, C.int(status))
	return checkError(res, err, "set_powerstat")
}

// GetPowerStat gets the Rig Power On/Off/Standby state
func (rig *Rig) GetPowerState() (status Power, err error) {
	var s C.int
	var res C.int
	res, err = C.get_powerstat(rig.handle, &s)
	status = Power(s)
	return status, checkError(res, err, "get_powerstat")
}

// Get Rig info
func (rig *Rig) GetInfo() (info string, err error) {
	i, err := C.get_info(rig.handle)
	info = C.GoString(i)
	return info, checkError(C.int(0), err, "get_info")
}

// Set Antenna
func (rig *Rig) SetAnt(vfo VFOType, ant int) error {
	res, err := C.set_ant(rig.handle, C.int(vfo), C.int(ant))
	return checkError(res, err, "set_ant")
}

// Get Antenna
func (rig *Rig) GetAnt(vfo VFOType) (ant int, err error) {
	var a C.int
	res, err := C.get_ant(rig.handle, C.int(vfo), &a)
	ant = int(a)
	return ant, checkError(res, err, "get_ant")
}

// Set Tuning step
func (rig *Rig) SetTs(vfo VFOType, ts int) error {
	res, err := C.set_ts(rig.handle, C.int(vfo), C.int(ts))
	return checkError(res, err, "set_ts")
}

// Get Tuning step
func (rig *Rig) GetTs(vfo VFOType) (ts int, err error) {
	var t C.long
	res, err := C.get_ts(rig.handle, C.int(vfo), &t)
	ts = int(t)
	return ts, checkError(res, err, "get_ts")
}

// has supports getting a specific level
func (rig *Rig) HasGetLevel(level uint32) (res uint32, err error) {
	var c C.ulong
	c, err = C.has_get_level(rig.handle, C.ulong(level))
	res = uint32(c)
	return res, checkError(0, err, "has_get_level")
}

// get the best frequency resolution for this rig (minimum step size)
func (rig *Rig) GetRigResolution(mode Mode) (resolution int, err error) {
	var r C.long
	r, err = C.get_rig_resolution(rig.handle, C.int(mode))
	resolution = int(r)
	return resolution, checkError(0, err, "get_rig_resolution")
}

// has supports setting a specific level
func (rig *Rig) HasSetLevel(level uint32) (res uint32, err error) {
	var c C.ulong
	c, err = C.has_set_level(rig.handle, C.ulong(level))
	res = uint32(c)
	return res, checkError(0, err, "has_set_level")
}

// has supports getting a specific function
func (rig *Rig) HasGetFunc(function uint32) (res uint32, err error) {
	var c C.ulong
	c, err = C.has_get_func(rig.handle, C.ulong(function))
	res = uint32(c)
	return res, checkError(0, err, "has_get_func")
}

// has supports setting a specific function
func (rig *Rig) HasSetFunc(function uint32) (res uint32, err error) {
	var c C.ulong
	c, err = C.has_set_func(rig.handle, C.ulong(function))
	res = uint32(c)
	return res, checkError(0, err, "has_set_func")
}

// has supports getting a specific parameter
func (rig *Rig) HasGetParm(parm uint32) (res uint32, err error) {
	var c C.ulong
	c, err = C.has_get_parm(rig.handle, C.ulong(parm))
	res = uint32(c)
	return res, checkError(0, err, "has_get_parm")
}

// has supports setting a specific parameter
func (rig *Rig) HasSetParm(parm uint32) (res uint32, err error) {
	var c C.ulong
	c, err = C.has_set_parm(rig.handle, C.ulong(parm))
	res = uint32(c)
	return res, checkError(0, err, "has_set_parm")
}

//get Level
func (rig *Rig) GetLevel(vfo VFOType, level uint32) (value float32, err error) {
	var v C.float
	var res C.int
	res, err = C.get_level(rig.handle, C.int(vfo), C.ulong(level), &v)
	value = float32(v)
	return value, checkError(res, err, "get_level")
}

//set Level
func (rig *Rig) SetLevel(vfo VFOType, level uint32, value float32) error {
	res, err := C.set_level(rig.handle, C.int(vfo), C.ulong(level), C.float(value))
	return checkError(res, err, "set_level")
}

//Get granularity (stepsize, minimum, maximum) for a Level
func (rig *Rig) GetLevelGran(level uint32) (step float32, min float32, max float32, err error) {
	var cStep C.float
	var cMin C.float
	var cMax C.float

	res, err := C.get_level_gran(rig.handle, C.ulong(level), &cStep, &cMin, &cMax)
	if checkError(res, err, "get_level_gran") != nil {
		return 0, 0, 0, err
	}

	return float32(cStep), float32(cMin), float32(cMax), nil
}

//get Function
func (rig *Rig) GetFunc(vfo VFOType, function uint32) (value bool, err error) {
	var v C.int
	var res C.int
	res, err = C.get_func(rig.handle, C.int(vfo), C.ulong(function), &v)
	value, err2 := CIntToBool(v)
	if err2 != nil { //not so nice...
		return value, checkError(0, err2, "get_func")
	}
	return value, checkError(res, err, "get_func")
}

//set Function
func (rig *Rig) SetFunc(vfo VFOType, function uint32, value bool) error {
	var v C.int
	v, err := BoolToCint(value)
	if err != nil {
		return checkError(0, err, "set_func")
	}
	res, err := C.set_func(rig.handle, C.int(vfo), C.ulong(function), v)
	return checkError(res, err, "set_func")
}

//get Parameter
func (rig *Rig) GetParm(vfo VFOType, parm uint32) (value float32, err error) {
	var v C.float
	var res C.int
	res, err = C.get_parm(rig.handle, C.ulong(parm), &v)
	value = float32(v)
	return value, checkError(res, err, "get_parm")
}

//set Parameter
func (rig *Rig) SetParm(vfo VFOType, parm uint32, value float32) error {
	res, err := C.set_parm(rig.handle, C.ulong(parm), C.float(value))
	return checkError(res, err, "set_parm")
}

//Get granularity (stepsize, minimum, maximum) for a Parameter
func (rig *Rig) GetParmGran(parm uint32) (step float32, min float32, max float32, err error) {
	var cStep C.float
	var cMin C.float
	var cMax C.float

	res, err := C.get_parm_gran(rig.handle, C.ulong(parm), &cStep, &cMin, &cMax)
	if checkError(res, err, "get_parm_gran") != nil {
		return 0, 0, 0, err
	}

	return float32(cStep), float32(cMin), float32(cMax), nil
}

//Set configuration token
func (rig *Rig) SetConf(token string, val string) error {
	res, err := C.set_conf(rig.handle, C.CString(token), C.CString(val))
	return checkError(res, err, "set_conf")
}

//HasToken checks if the rig supports a given token
func (rig *Rig) HasToken(token string) bool {
	//dirty hack - provide fix length char*
	//there should be a better way

	res, _ := C.has_token(rig.handle, C.CString(token))

	if HamlibErrorCode(res) == HamlibErrOK {
		return true
	}

	return false
}

//Get configuration token
func (rig *Rig) GetConf(token string) (val string, err error) {
	//dirty hack - provide fix length char*
	//there should be a better way
	v := C.CString("                                                          ")

	res, err := C.get_conf(rig.handle, C.CString(token), v)
	val = C.GoString(v)
	C.free(unsafe.Pointer(v))

	return val, checkError(res, err, "get_conf")
}

//Execute VFO Operation
func (rig *Rig) VfoOp(vfo VFOType, op VFOOp) error {
	res, err := C.vfo_op(rig.handle, C.int(vfo), C.int(op))
	return checkError(res, err, "vfo_op")
}

//Copy capabilities into Rig->Caps struct
func (rig *Rig) getCaps() error {
	if err := rig.getMaxRit(); err != nil {
		return err
	}
	if err := rig.getMaxXit(); err != nil {
		return err
	}
	if err := rig.getMaxIfShift(); err != nil {
		return err
	}
	if err := rig.getAttenuators(); err != nil {
		return err
	}
	if err := rig.getPreamps(); err != nil {
		return err
	}
	if err := rig.getVfos(); err != nil {
		return err
	}
	if err := rig.getOperations(); err != nil {
		return err
	}
	if err := rig.getModes(); err != nil {
		return err
	}
	if err := rig.getGetFunctions(); err != nil {
		return err
	}
	if err := rig.getSetFunctions(); err != nil {
		return err
	}
	if err := rig.getGetLevels(); err != nil {
		return err
	}
	if err := rig.getSetLevels(); err != nil {
		return err
	}
	if err := rig.getGetParameter(); err != nil {
		return err
	}
	if err := rig.getSetParameter(); err != nil {
		return err
	}
	if err := rig.getFilters(); err != nil {
		return err
	}
	if err := rig.getTuningSteps(); err != nil {
		return err
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
	rig.hasGetPtt()
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
	res, err := C.get_caps_max_rit(rig.handle, &rit)
	if checkError(res, err, "get_caps_max_rit") != nil {
		return checkError(res, err, "get_caps_max_rit")
	}
	rig.Caps.MaxRit = int(rit)
	return nil
}

//get Capabilities > Max Xit
func (rig *Rig) getMaxXit() error {
	var xit C.int
	res, err := C.get_caps_max_xit(rig.handle, &xit)
	if checkError(res, err, "get_caps_max_xit") != nil {
		return checkError(res, err, "get_caps_max_xit")
	}
	rig.Caps.MaxXit = int(xit)
	return nil
}

//get Capabilities > Max IF Shift
func (rig *Rig) getMaxIfShift() error {
	var ifShift C.int
	res, err := C.get_caps_max_if_shift(rig.handle, &ifShift)
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

	att_array, err := C.get_caps_attenuator_list_pointer_and_length(rig.handle, &length)
	if att_array == nil {
		return &HamlibError{"getAttenuators", HamlibErrEINTERNAL, "invalid pointer"}
	}
	if err != nil {
		return &Error{"getAttenuators", err}
	}

	var att []int
	for i := 0; i < int(length); i++ {
		C.get_int_from_array(rig.handle, att_array, &el, C.int(i))

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

	preamp_array, err := C.get_caps_preamp_list_pointer_and_length(rig.handle, &length)
	if preamp_array == nil {
		return &HamlibError{"getPreamp", HamlibErrEINTERNAL, "invalid pointer"}
	}
	if err != nil {
		return &Error{"getPreamp", err}
	}

	var preamps []int
	for i := 0; i < int(length); i++ {
		C.get_int_from_array(rig.handle, preamp_array, &el, C.int(i))

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

	res, err := C.get_supported_vfos(rig.handle, &vfoClist)
	if checkError(res, err, "get_supported_vfos") != nil {
		return checkError(res, err, "get_supported_vfos")
	}

	for vfo, vfoStr := range VFOName {
		if VFOType(vfoClist)&vfo > 0 {
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

	res, err := C.get_supported_vfo_operations(rig.handle, &vfoOpClist)
	if checkError(res, err, "get_supported_vfo_operations") != nil {
		return checkError(res, err, "get_supported_vfo_operations")
	}

	for op, opStr := range VFOOperationName {
		if VFOOp(vfoOpClist)&op > 0 {
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

	res, err := C.get_supported_modes(rig.handle, &modesClist)
	if checkError(res, err, "get_supported_modes") != nil {
		return checkError(res, err, "get_supported_modes")
	}

	for mode, modeStr := range ModeName {
		if int(modesClist)&int(mode) > 0 {
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

	C.get_model_name(rig.handle, cModelName)
	rig.Caps.ModelName = strings.TrimSpace(C.GoString(cModelName))
	C.free(unsafe.Pointer(cModelName))
}

// get Capabilities > Version
func (rig *Rig) getVersion() {
	cVersion := C.CString("                         ")

	C.get_version(rig.handle, cVersion)
	rig.Caps.Version = strings.TrimSpace(C.GoString(cVersion))
	C.free(unsafe.Pointer(cVersion))
}

// get Capabilities > Manufacturer Name
func (rig *Rig) getMfgName() {
	cMfGName := C.CString("                         ")

	C.get_mfg_name(rig.handle, cMfGName)
	rig.Caps.MfgName = strings.TrimSpace(C.GoString(cMfGName))
	C.free(unsafe.Pointer(cMfGName))
}

// get Capabilities > Status
func (rig *Rig) getStatus() {

	cStatus := C.get_status(rig.handle)
	rig.Caps.Status = int(cStatus)
}

func (rig *Rig) getFilters() error {
	var cfc C.int
	var cWidth C.long
	var cMode C.int

	var filterMap map[string][]int
	filterMap = make(map[string][]int)

	res, err := C.get_filter_count(rig.handle, &cfc)
	if checkError(res, err, "get_filter_count") != nil {
		return checkError(res, err, "get_filter_count")
	}

	for i := 0; i < int(cfc); i++ {
		res, err = C.get_filter_mode_width(rig.handle, C.int(i), &cMode, &cWidth)
		if checkError(res, err, "") != nil {
			return checkError(res, err, "get_filter_mode_width")
		}
		for mode, modeStr := range ModeName {
			if int(cMode)&int(mode) > 0 {
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

	res, err := C.get_ts_count(rig.handle, &tsc)
	if checkError(res, err, "get_ts_count") != nil {
		return checkError(res, err, "get_ts_count")
	}

	for i := 0; i < int(tsc); i++ {
		res, err = C.get_tuning_steps(rig.handle, C.int(i), &cMode, &cTs)
		if checkError(res, err, "") != nil {
			return checkError(res, err, "get_tuning_steps")
		}
		for mode, modeStr := range ModeName {
			if int(cMode)&int(mode) > 0 {
				tsMap[modeStr] = append(tsMap[modeStr], int(cTs))
			}
		}
	}

	rig.Caps.TuningSteps = tsMap
	return nil
}

func (rig *Rig) hasSetPowerStat() {
	res := C.has_set_powerstat(rig.handle)
	if HamlibErrorCode(res) == HamlibErrOK {
		rig.Caps.HasSetPowerStat = true
		return
	}
	rig.Caps.HasSetPowerStat = false
}

func (rig *Rig) hasGetPowerStat() {
	res := C.has_get_powerstat(rig.handle)
	if HamlibErrorCode(res) == HamlibErrOK {
		rig.Caps.HasGetPowerStat = true
		return
	}
	rig.Caps.HasGetPowerStat = false
}

func (rig *Rig) hasSetVfo() {
	res := C.has_set_vfo(rig.handle)
	if HamlibErrorCode(res) == HamlibErrOK {
		rig.Caps.HasSetVfo = true
		return
	}
	rig.Caps.HasSetVfo = false
}

func (rig *Rig) hasGetVfo() {
	res := C.has_get_vfo(rig.handle)
	if HamlibErrorCode(res) == HamlibErrOK {
		rig.Caps.HasGetVfo = true
		return
	}
	rig.Caps.HasGetVfo = false
}

func (rig *Rig) hasSetFreq() {
	res := C.has_set_freq(rig.handle)
	if HamlibErrorCode(res) == HamlibErrOK {
		rig.Caps.HasSetFreq = true
		return
	}
	rig.Caps.HasSetFreq = false
}

func (rig *Rig) hasGetFreq() {
	res := C.has_get_freq(rig.handle)
	if HamlibErrorCode(res) == HamlibErrOK {
		rig.Caps.HasGetFreq = true
		return
	}
	rig.Caps.HasGetFreq = false
}

func (rig *Rig) hasSetMode() {
	res := C.has_set_mode(rig.handle)
	if HamlibErrorCode(res) == HamlibErrOK {
		rig.Caps.HasSetMode = true
		return
	}
	rig.Caps.HasSetMode = false
}

func (rig *Rig) hasGetMode() {
	res := C.has_get_mode(rig.handle)
	if HamlibErrorCode(res) == HamlibErrOK {
		rig.Caps.HasGetMode = true
		return
	}
	rig.Caps.HasGetMode = false
}

func (rig *Rig) hasSetPtt() {
	res := C.has_set_ptt(rig.handle)
	if HamlibErrorCode(res) == HamlibErrOK {
		rig.Caps.HasSetPtt = true
		return
	}
	rig.Caps.HasSetPtt = false
}

func (rig *Rig) hasGetPtt() {
	res := C.has_get_ptt(rig.handle)
	if HamlibErrorCode(res) == HamlibErrOK {
		rig.Caps.HasGetPtt = true
		return
	}
	rig.Caps.HasGetPtt = false
}

func (rig *Rig) hasSetRit() {
	res := C.has_set_rit(rig.handle)
	if HamlibErrorCode(res) == HamlibErrOK {
		rig.Caps.HasSetRit = true
		return
	}
	rig.Caps.HasSetRit = false
}

func (rig *Rig) hasGetRit() {
	res := C.has_get_rit(rig.handle)
	if HamlibErrorCode(res) == HamlibErrOK {
		rig.Caps.HasGetRit = true
		return
	}
	rig.Caps.HasGetRit = false
}

func (rig *Rig) hasSetXit() {
	res := C.has_set_xit(rig.handle)
	if HamlibErrorCode(res) == HamlibErrOK {
		rig.Caps.HasSetXit = true
		return
	}
	rig.Caps.HasSetXit = false
}

func (rig *Rig) hasGetXit() {
	res := C.has_get_xit(rig.handle)
	if HamlibErrorCode(res) == HamlibErrOK {
		rig.Caps.HasGetXit = true
		return
	}
	rig.Caps.HasGetXit = false
}

func (rig *Rig) hasSetSplitVfo() {
	res := C.has_set_split_vfo(rig.handle)
	if HamlibErrorCode(res) == HamlibErrOK {
		rig.Caps.HasSetSplitVfo = true
		return
	}
	rig.Caps.HasSetSplitVfo = false
}

func (rig *Rig) hasGetSplitVfo() {
	res := C.has_get_split_vfo(rig.handle)
	if HamlibErrorCode(res) == HamlibErrOK {
		rig.Caps.HasGetSplitVfo = true
		return
	}
	rig.Caps.HasGetSplitVfo = false
}

func (rig *Rig) hasSetSplitMode() {
	res := C.has_set_split_mode(rig.handle)
	if HamlibErrorCode(res) == HamlibErrOK {
		rig.Caps.HasSetSplitMode = true
		return
	}
	rig.Caps.HasSetSplitMode = false
}

func (rig *Rig) hasGetSplitMode() {
	res := C.has_get_split_mode(rig.handle)
	if HamlibErrorCode(res) == HamlibErrOK {
		rig.Caps.HasGetSplitMode = true
		return
	}
	rig.Caps.HasGetSplitMode = false
}

func (rig *Rig) hasSetSplitFreq() {
	res := C.has_set_split_freq(rig.handle)
	if HamlibErrorCode(res) == HamlibErrOK {
		rig.Caps.HasSetSplitFreq = true
		return
	}
	rig.Caps.HasSetSplitFreq = false
}

func (rig *Rig) hasGetSplitFreq() {
	res := C.has_get_split_freq(rig.handle)
	if HamlibErrorCode(res) == HamlibErrOK {
		rig.Caps.HasGetSplitFreq = true
		return
	}
	rig.Caps.HasGetSplitFreq = false
}

func (rig *Rig) hasSetAnt() {
	res := C.has_set_ant(rig.handle)
	if HamlibErrorCode(res) == HamlibErrOK {
		rig.Caps.HasSetAnt = true
		return
	}
	rig.Caps.HasSetAnt = false
}

func (rig *Rig) hasGetAnt() {
	res := C.has_get_ant(rig.handle)
	if HamlibErrorCode(res) == HamlibErrOK {
		rig.Caps.HasGetAnt = true
		return
	}
	rig.Caps.HasGetAnt = false
}

func (rig *Rig) hasSetTs() {
	res := C.has_set_ts(rig.handle)
	if HamlibErrorCode(res) == HamlibErrOK {
		rig.Caps.HasSetTs = true
		return
	}
	rig.Caps.HasSetTs = false
}

func (rig *Rig) hasGetTs() {
	res := C.has_get_ts(rig.handle)
	if HamlibErrorCode(res) == HamlibErrOK {
		rig.Caps.HasGetTs = true
		return
	}
	rig.Caps.HasGetTs = false
}

func (rig *Rig) hasSetConf() {
	res := C.has_set_conf(rig.handle)
	if HamlibErrorCode(res) == HamlibErrOK {
		rig.Caps.HasSetConf = true
		return
	}
	rig.Caps.HasSetConf = false
}

func (rig *Rig) hasGetConf() {
	res := C.has_get_conf(rig.handle)
	if HamlibErrorCode(res) == HamlibErrOK {
		rig.Caps.HasGetConf = true
		return
	}
	rig.Caps.HasGetConf = false
}

//Close the Communication with the Radio
func (rig *Rig) Close() error {
	res, err := C.close_rig(rig.handle)
	return checkError(res, err, "close_rig")
}

//Grabage collect Radio and free up memory
func (rig *Rig) Cleanup() error {
	res, err := C.cleanup_rig(rig.handle)
	return checkError(res, err, "cleanup_rig")
}

// Check Errors from Hamlib C calls. C Errors have a higher priority.
// Additional Information is provided for better debugging
func checkError(res C.int, e error, operation string) error {
	if e != nil {
		return &Error{operation, e}
	}
	herr := HamlibErrorCode(res)
	if herr != HamlibErrOK {
		return &HamlibError{operation, herr, ""}
	}

	return nil
}
