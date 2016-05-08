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

	rig.SetDebugLevel(goHamlib.RIG_DEBUG_NONE)
	rig.Init(128)
	rig.SetPort(p)
	rig.Open()
	//rig.SetVfo(1)

	//Set Frequency
	if err:= rig.SetFreq(goHamlib.RIG_VFO_CURR, 7005000); err != nil{
		log.Println(err)
	}

	//set invalid frequency
	if err:= rig.SetFreq(goHamlib.RIG_VFO_CURR, -3580000); err != nil{
		log.Println(err)
	}

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
	log.Printf("Current Frequency is: %08v Hz", freq)

	// get mode
	mode, pb_width, err := rig.GetMode(goHamlib.RIG_VFO_CURR)
	log.Printf("Current Mode: %v, Passband: %v", mode, pb_width)

	// set Ptt true
	if err := rig.SetPtt(goHamlib.RIG_VFO_CURR, goHamlib.RIG_PTT_ON); err != nil{
		log.Println(err)
	}
	time.Sleep(time.Second)

	// get Ptt state
	if ptt, err:= rig.GetPtt(goHamlib.RIG_VFO_CURR); err != nil{
		log.Println(err)
	} else {
		log.Printf("Ptt state: %v", ptt)
	}

	// set Ptt false
	if err := rig.SetPtt(goHamlib.RIG_VFO_CURR, goHamlib.RIG_PTT_OFF); err != nil{
		log.Println(err)
	}
	time.Sleep(time.Millisecond * 200)

	// get Ptt state
	if ptt, err:= rig.GetPtt(goHamlib.RIG_VFO_CURR); err != nil{
		log.Println(err)
	} else {
		log.Printf("Ptt state: %v", ptt)
	}

	// set Rit
	if err := rig.SetRit(goHamlib.RIG_VFO_CURR, -500); err != nil{
		log.Println(err)
	}

	// get Rit
	if rit, err := rig.GetRit(goHamlib.RIG_VFO_CURR); err != nil{
		log.Println(err)
	} else {
		log.Printf("Rit offset: %v Hz", rit)
	}

	// set invalid Rit (will be set to maximum)
	if err := rig.SetRit(goHamlib.RIG_VFO_CURR, 20000); err != nil{
		log.Println(err)
	}

	// get Rit(should be at 9.999kHz)
	if rit, err := rig.GetRit(goHamlib.RIG_VFO_CURR); err != nil{
		log.Println(err)
	} else {
		log.Printf("Rit offset: %v Hz", rit)
	}

	// set Xit
	if err := rig.SetXit(goHamlib.RIG_VFO_CURR, 5555); err != nil{
		log.Println(err)
	}

	// get Rit
	if xit, err := rig.GetXit(goHamlib.RIG_VFO_CURR); err != nil{
		log.Println(err)
	} else {
		log.Printf("Xit offset: %v Hz", xit)
	}



	//Shutdown & Cleanup
	rig.Close()
	rig.Cleanup()

	log.Println("finished testing")
}
