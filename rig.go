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
extern int test1(int var);
extern int test2();
*/
import "C"

import (
	"log"
)

func (rig *Rig) SetPort(p Port_t) error{

	_, err := C.set_port(C.int(p.RigPortType), C.CString(p.Portname) , C.int(p.Baudrate), C.int(p.Databits), C.int(p.Stopbits), C.int(p.Parity), C.int(p.Handshake))

	if err != nil{
		return err
	}

	return nil
}

func (rig *Rig) Init(rigModel int) error{
	myRig, err := C.init_rig(C.int(rigModel))
	if err != nil{
		log.Println("could not initialize rig")
		log.Println(err)
		return err
	}
	log.Println(myRig)
	return nil
}

func (rig *Rig) Open() error{
	res, err := C.open_rig()
	if  err != nil{
		log.Println("could not open port")
		log.Println(err)
		return err
	}
	log.Println("open res: ", res)
	return nil
}

func (rig *Rig) SetVfo(vfo int) error{
	if _, err := C.set_vfo(C.int(vfo)); err != nil{
		log.Println("Could not set VFO")
		return err
	}
	return nil
}

func (rig *Rig) SetFreq(vfo int, freq float64) error{
	res, err := C.set_freq(C.int(vfo), C.double(freq))
	if err != nil{
		log.Println("Could not set freq")
		return err
	}
	log.Println(res)
	return nil
}

func SetVar(num int){
	res, _ := C.test1(C.int(num))
	log.Println("Res setvar: ", res)
}

func ReadVar(){
	res, _ := C.test2()
	log.Println("Res readvar: ", res)
}
