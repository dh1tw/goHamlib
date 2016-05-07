// +build linux

package goHamlib

/*
#cgo CFLAGS: -I /usr/local/lib
#cgo LDFLAGS: -L /usr/local/lib -lhamlib 

#include <hamlib/rig.h>

extern int set_port(int rig_port_type, char* portname, int baudrate, int databits, int stopbits, int parity, int handshake);
extern int init_rig(int rig_model);
extern int open_rig();
extern int set_vfo(int vfo);
extern int set_freq(int vfo, double freq);
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
