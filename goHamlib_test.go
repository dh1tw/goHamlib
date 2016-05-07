package goHamlib_test

import (
	"log"
 	"testing"

	"github.com/dh1tw/goHamlib"
)

func TestDummyRig(t *testing.T){

	var p goHamlib.Port_t
	p.RigPortType = 1 
	p.Portname = "/dev/mhuxd/cat"
	p.Baudrate = 38400
	p.Databits = 8
	p.Stopbits = 1
	p.Parity = goHamlib.N
	p.Handshake = goHamlib.NO_HANDSHAKE

	var rig goHamlib.Rig
	rig.Init(128)
	rig.SetPort(p)
	rig.Open()
	//rig.SetVfo(1)
	if err:= rig.SetFreq(goHamlib.RIG_VFO_CURR, 7005000); err != nil{
		log.Println(err)
	}

	if err:= rig.SetFreq(goHamlib.RIG_VFO_CURR, -7005000); err != nil{
		log.Println(err)
	}

	log.Println("finished testing")
}
