package goHamlib_test

import (
	"log"
	"testing"
	"time"

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

	//Set Frequency
	if err:= rig.SetFreq(goHamlib.RIG_VFO_CURR, 7005000); err != nil{
		log.Println(err)
	}

	//set invalid frequency
//	if err:= rig.SetFreq(goHamlib.RIG_VFO_CURR, -7005000); err != nil{
//		log.Println(err)
//	}

	//set mode Narrow Filter
	mode, err := rig.GetPbNarrow(goHamlib.RIG_MODE_CW)
	if err != nil{
		log.Println("Couldn't get Narrow Pb")
	}
	if mode == 0{
		log.Println("can not determin narrow Passband")
	}
	if err:= rig.SetMode(goHamlib.RIG_VFO_CURR, goHamlib.RIG_MODE_CW,mode); err != nil{
		log.Println(err)
	}

	// set mode with Normal Filter
	time.Sleep(time.Second)
	mode, _ = rig.GetPbNormal(goHamlib.RIG_MODE_CW)
	if err:= rig.SetMode(goHamlib.RIG_VFO_CURR, goHamlib.RIG_MODE_CW,mode); err != nil{
		log.Println(err)
	}

	// set mode with Wide Filter
	time.Sleep(time.Second)
	mode, _ = rig.GetPbWide(goHamlib.RIG_MODE_CW)
	if err:= rig.SetMode(goHamlib.RIG_VFO_CURR, goHamlib.RIG_MODE_CW,mode); err != nil{
		log.Println(err)
	}

	// get Frequency
	freq, _ := rig.GetFreq(goHamlib.RIG_VFO_CURR)
	log.Println("Current Frequency is: %08v Hz", freq)

	rig.Close()
	rig.Cleanup()

	log.Println("finished testing")
}
