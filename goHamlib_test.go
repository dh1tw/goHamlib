package goHamlib_test

import (
	"log"
	"testing"
	"time"

	"github.com/dh1tw/goHamlib"
)

//Tests against an FT950
func TestFT950(t *testing.T){

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

	if stat, err := rig.GetPowerStat(); err != nil {
		log.Println(err)
	} else {
		if stat == goHamlib.RIG_POWER_OFF{
			if err = rig.SetPowerStat(goHamlib.RIG_POWER_ON); err != nil{
				log.Println(err)
				// HERE WE SHOULD PANIC!!!
			}
			time.Sleep(time.Second * 5)
		}
	}

	//Set Frequency
	if err:= rig.SetFreq(goHamlib.RIG_VFO_MAIN, 7005000); err != nil{
		log.Println(err)
	}

	//set invalid frequency
	if err:= rig.SetFreq(goHamlib.RIG_VFO_MAIN, -3580000); err != nil{
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
	time.Sleep(time.Second)

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

	// set Rit on invalid VFO
	if err := rig.SetRit(goHamlib.RIG_VFO_C, 20000); err != nil{
		log.Println(err)
	}

	// get Rit from invalid VFO
	if rit, err := rig.GetRit(goHamlib.RIG_VFO_C); err != nil{
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

	time.Sleep(time.Second)

	// rig.SetVfo(goHamlib.RIG_VFO_MAIN)

	// set split on VFOB - same mode and pb as VFO A
	if mode, pb, err := rig.GetMode(goHamlib.RIG_VFO_MAIN); err != nil{
		log.Println(err)
	} else {
		if err := rig.SetSplitVfo(goHamlib.RIG_VFO_MAIN, goHamlib.RIG_SPLIT_ON, goHamlib.RIG_VFO_SUB); err != nil{
			log.Println(err)
		}
		if err := rig.SetSplitFreq(goHamlib.RIG_VFO_SUB, 7020000); err != nil{
			log.Println(err)
		}
		mode = mode
		pb = pb
		if err = rig.SetSplitMode(goHamlib.RIG_VFO_SUB, mode, pb); err != nil{
			log.Println(err)
		}
	}

	// check status on Split
	if split, txVfo, err := rig.GetSplit(goHamlib.RIG_VFO_MAIN); err != nil{
		log.Println(err)
	} else {
		log.Printf("Split is: %v on VFO: %v", split, txVfo)
	}

	time.Sleep(time.Second)

	// get Rig Info
	if info, err := rig.GetInfo(); err != nil{
		log.Println(err)
	} else {
		log.Printf("Rig Info: %s", info)
	}

	// Set and check Antenna
	if err := rig.SetAnt(goHamlib.RIG_VFO_CURR, goHamlib.RIG_ANT_1); err != nil{
		log.Println(err)
	}

	if ant, err := rig.GetAnt(goHamlib.RIG_VFO_CURR); err != nil{
		log.Println(err)
	} else {
		log.Printf("Selected antenna: %v", ant)
	}

	// Get current tuning step
	if ts, err := rig.GetTs(goHamlib.RIG_VFO_CURR); err != nil{
		log.Println(err)
	} else {
		log.Printf("Tuning step: %vHz", ts)
	}

	// Set tuning step to 100Hz
	if err := rig.SetTs(goHamlib.RIG_VFO_CURR, 100); err != nil{
		log.Println(err)
	}

        // Verify that tuning step was set accordingly
        if ts, err := rig.GetTs(goHamlib.RIG_VFO_CURR); err != nil{
                log.Println(err)
        } else {
                log.Printf("Tuning step: %vHz", ts)
        }

	// Has Set / Get Levels
	if level, err := rig.HasGetLevel(goHamlib.RIG_LEVEL_ATT); err != nil{
		log.Println(err)
	} else {
		log.Printf("can get Level ATT: %t", level==goHamlib.RIG_LEVEL_ATT)
	}
 	time.Sleep(time.Second*2)

	if res, err := rig.HasSetLevel(goHamlib.RIG_LEVEL_IF); err != nil{
		log.Println(err)
	} else {
		log.Printf("can set Level IF: %t", res==goHamlib.RIG_LEVEL_IF)
	}
 	time.Sleep(time.Second*2)

	// Has set / get Functions
        if res, err := rig.HasGetFunc(goHamlib.RIG_FUNC_LOCK); err != nil{
                log.Println(err)
        } else {
                log.Printf("can get Function LOCK: %t", res==goHamlib.RIG_FUNC_LOCK)
        }
        time.Sleep(time.Second*2)

        if res, err := rig.HasSetFunc(goHamlib.RIG_FUNC_ABM); err != nil{
                log.Println(err)
        } else {
                log.Printf("can set Function ABM: %t", res==goHamlib.RIG_FUNC_ABM)
        }
        time.Sleep(time.Second*2)

        // Has set / get Parameters
        if res, err := rig.HasGetParm(goHamlib.RIG_PARM_APO); err != nil{
                log.Println(err)
        } else {
                log.Printf("can get Parameter APO: %t", res==goHamlib.RIG_PARM_APO)
        }
        time.Sleep(time.Second*2)

        if res, err := rig.HasSetParm(goHamlib.RIG_PARM_TIME); err != nil{
                log.Println(err)
        } else {
                log.Printf("can set Parameter TIME: %t", res==goHamlib.RIG_PARM_TIME)
        }
        time.Sleep(time.Second*2)



//	rig.SetPowerStat(goHamlib.RIG_POWER_OFF);

	//Shutdown & Cleanup
	rig.Close()
	rig.Cleanup()

	log.Println("finished testing")
}
