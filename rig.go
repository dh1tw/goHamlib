// +build linux

package goHamlib

/*
#cgo CFLAGS: -I /usr/local/lib
#cgo LDFLAGS: -L /usr/local/lib -lhamlib 

#include <stdio.h>
#include <stdlib.h>
#include <hamlib/rig.h>

extern int set_port(int rig_port_type, char* portname, int baudrate, int databits, int stopbits, int parity, int handshake);
extern int init_rig(int rig_model);
extern int open_rig();
extern int set_vfo(int vfo);
extern int set_freq(int vfo, double freq);
extern int set_mode(int vfo, int mode, int pb_width);
extern int get_passband_narrow(int mode);
extern int get_passband_normal(int mode);
extern int get_passband_wide(int mode);
extern int get_freq(int vfo, double *freq);
extern int get_mode(int vfo, int *mode, int *pb_width);
extern int set_ptt(int vfo, int ptt);
extern int get_ptt(int vfo, int *ptt);
extern int set_rit(int vfo, int offset);
extern int get_rit(int vfo, int *offset);
extern int set_xit(int vfo, int offset);
extern int get_xit(int vfo, int *offset);
extern int set_split_freq(int vfo, double tx_freq);
extern int get_split_freq(int vfo, double *tx_freq);
extern int set_split_mode(int vfo, int tx_mode, int tx_width);
extern int get_split_mode(int vfo, int *tx_mode, int *tx_width);
extern int set_split_vfo(int vfo, int split, int tx_vfo);
extern int get_split_vfo(int vfo, int *split, int *tx_vfo);
extern int set_powerstat(int status);
extern int get_powerstat(int *status);
extern const char* get_info();
extern int set_ant(int vfo, int ant);
extern int get_ant(int vfo, int *ant);
extern int set_ts(int vfo, int ts);
extern int get_ts(int vfo, int *ts);
extern unsigned long has_get_level(unsigned long level);
extern unsigned long has_set_level(unsigned long level);
extern unsigned long has_get_func(unsigned long function);
extern unsigned long has_set_func(unsigned long function);
extern unsigned long has_get_parm(unsigned long parm);
extern unsigned long has_set_parm(unsigned long parm);
extern void set_debug_level(int debug_level);
extern int close_rig();
extern int cleanup_rig();

*/
import "C"

import (
	//"log"
)

// Initialize Rig
func (rig *Rig) Init(rigModel int) error{
	res, err := C.init_rig(C.int(rigModel))
	return checkError(res, err, "open_rig")
}

// Set Port of Rig
func (rig *Rig) SetPort(p Port_t) error{
	res, err := C.set_port(C.int(p.RigPortType), C.CString(p.Portname) , C.int(p.Baudrate), C.int(p.Databits), C.int(p.Stopbits), C.int(p.Parity), C.int(p.Handshake))
	return checkError(res, err, "set_port")
}

// Open Radio / Port
func (rig *Rig) Open() error{
	res, err := C.open_rig()
	return checkError(res, err, "open_rig")
}

// Set default VFO
func (rig *Rig) SetVfo(vfo int) error{
	res, err := C.set_vfo(C.int(vfo))
	return checkError(res, err, "set_vfo")
}

// Set Frequency for a VFO
func (rig *Rig) SetFreq(vfo int, freq float64) error{
	res, err := C.set_freq(C.int(vfo), C.double(freq))
	return checkError(res, err, "set_freq")
}

// Set Mode for a VFO
func (rig *Rig) SetMode(vfo int, mode int, pb_width int) error{
	res, err := C.set_mode(C.int(vfo), C.int(mode), C.int(pb_width))
	return checkError(res, err, "set_freq")
}

// Find the next suitable narrow available filter
func (rig *Rig) GetPbNarrow(mode int) (int, error){
	pb, err := C.get_passband_narrow(C.int(mode))
	pb_width := int(pb)

	return pb_width, err
}

// Find the next suitable normal available filter
func (rig *Rig) GetPbNormal(mode int) (int, error){
	pb, err := C.get_passband_normal(C.int(mode))
	pb_width := int(pb)

	return pb_width, err
}

// Find the next suitable wide available filter
func (rig *Rig) GetPbWide(mode int) (int, error){
	pb, err := C.get_passband_wide(C.int(mode))
	pb_width := int(pb)

	return pb_width, err
}

// Get Frequency from a VFO
func (rig *Rig) GetFreq(vfo int) (freq float64, err error){
	var f C.double
	var res C.int
	res, err = C.get_freq(C.int(vfo), &f)
	freq = float64(f)
	return freq, checkError(res, err, "get_freq")
}

// Get Mode and Passband width for a VFO
func (rig *Rig) GetMode(vfo int) (mode int, pb_width int, err error){
	var m C.int
	var pb C.int
	var res C.int
	res, err = C.get_mode(C.int(vfo), &m, &pb)
	pb_width = int(pb)
	mode = int(m)
	return mode, pb_width, checkError(res, err, "get_mode")
}

// Set Ptt
func (rig *Rig) SetPtt(vfo int, ptt int) error{
	res, err := C.set_ptt(C.int(vfo), C.int(ptt))
	return checkError(res, err, "set_ptt")
}

// Get Ptt state
func (rig *Rig) GetPtt(vfo int) (ptt int, err error){
	var p C.int
	res, err := C.get_ptt(C.int(vfo), &p)
	ptt = int(p)
	return ptt, checkError(res, err, "get_ptt")
}

// Set Rit offset value
func (rig *Rig) SetRit(vfo int, offset int) error{
	res, err := C.set_rit(C.int(vfo), C.int(offset))
	return checkError(res, err, "set_rit")
}

// Get Rit offset value
func (rig *Rig) GetRit(vfo int) (offset int, err error){
	var o C.int
	res, err := C.get_rit(C.int(vfo), &o)
	offset = int(o)
	return offset, checkError(res, err, "get_rit")
}

// Set Xit offset value
func (rig *Rig) SetXit(vfo int, offset int) error{
	res, err := C.set_xit(C.int(vfo), C.int(offset))
	return checkError(res, err, "set_xit")
}

// Get Xit offset value
func (rig *Rig) GetXit(vfo int) (offset int, err error){
	var o C.int
	res, err := C.get_xit(C.int(vfo), &o)
	offset = int(o)
	return offset, checkError(res, err, "get_xit")
}

// Set Split Frequency
func (rig *Rig) SetSplitFreq(vfo int, txFreq float64) error{
	res, err := C.set_split_freq(C.int(vfo), C.double(txFreq))
	return checkError(res, err, "set_split_freq")
}

// Get Split Frequency
func (rig *Rig) GetSplitFreq(vfo int) (txFreq float64, err error){
        var f C.double
        res, err := C.get_split_freq(C.int(vfo), &f)
        txFreq = float64(f)
        return txFreq, checkError(res, err, "get_split_freq")
}

// Set Split Mode
func (rig *Rig) SetSplitMode(vfo int, txMode int, txWidth int) error{
        res, err := C.set_split_mode(C.int(vfo), C.int(txMode), C.int(txWidth))
        return checkError(res, err, "set_split_mode")
}

// Get Split Mode
func (rig *Rig) GetSplitMode(vfo int) (txMode int, txWidth int, err error){
        var m C.int
	var w C.int
        res, err := C.get_split_mode(C.int(vfo), &m, &w)
        txMode = int(m)
	txWidth = int(w)
        return txMode, txWidth, checkError(res, err, "get_split_mode")
}

// Set Split Vfo
func (rig *Rig) SetSplitVfo(vfo int, split int, txVfo int) error{
        res, err := C.set_split_vfo(C.int(vfo), C.int(split), C.int(txVfo))
        return checkError(res, err, "set_split_vfo")
}

// Get Split Vfo
func (rig *Rig) GetSplitVfo(vfo int) (split int, txVfo int, err error){
        var s C.int
        var v C.int
        res, err := C.get_split_mode(C.int(vfo), &s, &v)
        split = int(s)
        txVfo = int(v)
        return split, txVfo, checkError(res, err, "get_split_vfo")
}

// Set Split (shortcut for SetSplitVfo)
func (rig *Rig) SetSplit(vfo int, split int) error{
	res, err := C.set_split_vfo(C.int(vfo), C.int(split), C.int(RIG_VFO_CURR))
	return checkError(res, err, "set_split")
}

// Get Split (shortcut for GetSplitVfo)
func (rig *Rig) GetSplit(vfo int) (split int, txVfo int, err error){
	var s C.int
	var t C.int
	res, err := C.get_split_vfo(C.int(vfo), &s, &t)
	split = int(s)
	txVfo = int(t)
	return split, txVfo, checkError(res, err, "get_split")
}

// Set Rig Power On/Off/Standby
func (rig *Rig) SetPowerStat(status int) error{
	res, err := C.set_powerstat(C.int(status))
	return checkError(res, err, "set_powerstat")
}

// Get Rig Power On/Off/Standby
func (rig *Rig) GetPowerStat() (status int, err error){
	var s C.int
	var res C.int
	res, err = C.get_powerstat(&s)
	status = int(s)
	return status, checkError(res, err, "get_powerstat")
}

// Get Rig info
func (rig *Rig) GetInfo() (info string, err error){
	i, err := C.get_info()
	info = C.GoString(i)
	return info, checkError(C.int(0), err, "get_info")
}

// Set Antenna
func (rig *Rig) SetAnt(vfo int, ant int) error{
	res, err := C.set_ant(C.int(vfo), C.int(ant))
	return checkError(res, err, "set_ant") 
}

// Get Antenna
func (rig *Rig) GetAnt(vfo int) (ant int, err error){
	var a C.int
	res, err := C.get_ant(C.int(vfo), &a)
	ant = int(a)
	return ant, checkError(res, err, "get_ant")
}

// Set Tuning step
func (rig *Rig) SetTs(vfo int, ts int) error{
	res, err := C.set_ts(C.int(vfo), C.int(ts))
	return checkError(res, err, "set_ts")
}

// Get Tuning step
func (rig *Rig) GetTs(vfo int) (ts int, err error){
	var t C.int
	res, err := C.get_ts(C.int(vfo), &t)
	ts = int(t)
	return ts, checkError(res, err, "get_ts")
}

// has supports getting a specific level
func (rig *Rig) HasGetLevel(level uint32) (res uint32, err error){
	var c C.ulong
	c, err = C.has_get_level(C.ulong(level))
	res = uint32(c)
	return res, checkError(0, err, "has_get_level")
}

// has supports setting a specific level
func (rig *Rig) HasSetLevel(level uint32) (res uint32, err error){
	var c C.ulong
	c, err = C.has_set_level(C.ulong(level))
	res = uint32(c)
	return res, checkError(0, err, "has_set_level")
}

// has supports getting a specific function
func (rig *Rig) HasGetFunc(function uint32) (res uint32, err error){
	var c C.ulong
	c, err = C.has_get_func(C.ulong(function))
	res = uint32(c)
	return res, checkError(0, err, "has_get_func")
}

// has supports setting a specific function
func (rig *Rig) HasSetFunc(function uint32) (res uint32, err error){
	var c C.ulong
	c, err = C.has_set_func(C.ulong(function))
	res = uint32(c)
	return res, checkError(0, err, "has_set_func")
}


// has supports getting a specific parameter
func (rig *Rig) HasGetParm(parm uint32) (res uint32, err error){
	var c C.ulong
	c, err = C.has_get_parm(C.ulong(parm))
	res = uint32(c)
	return res, checkError(0, err, "has_get_parm")
}

// has supports setting a specific parameter
func (rig *Rig) HasSetParm(parm uint32) (res uint32, err error){
	var c C.ulong
	c, err = C.has_set_parm(C.ulong(parm))
	res = uint32(c)
	return res, checkError(0, err, "has_set_parm")
}



// Set Debug level
func (rig *Rig) SetDebugLevel(dbgLevel int){
	C.set_debug_level(C.int(dbgLevel))
}

//Close the Communication with the Radio
func (rig *Rig) Close() error{
	res, err := C.close_rig()
	return checkError(res, err, "close_rig")
}

//Grabage collect Radio and free up memory
func (rig *Rig) Cleanup() error{
	res, err := C.cleanup_rig()
	return checkError(res, err, "cleanup_rig")
} 

// Check Errors from Hamlib C calls. C Errors have a higher priority.
// Additional Information is provided for better debugging
func checkError(res C.int, e error, operation string) error{

        if e != nil {
                return &Error{operation, e}
        }
        if int(res) != RIG_OK{
                return &HamlibError{operation, int(res), ""}
        }

        return nil
}


