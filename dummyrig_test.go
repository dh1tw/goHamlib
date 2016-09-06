package goHamlib_test

import (
	"fmt"
	"testing"

	"github.com/dh1tw/goHamlib"
)

func getDummyRig() (*goHamlib.Rig, error) {
	rig := goHamlib.Rig{}
	rig.SetDebugLevel(goHamlib.RIG_DEBUG_VERBOSE)
	// rig.SetDebugLevel(goHamlib.RIG_DEBUG_ERR)
	if err := rig.Init(1); err != nil {
		return nil, err
	}
	if err := rig.Open(); err != nil {
		return nil, err
	}

	return &rig, nil
}

//Initialize rig with invalid data
// func TestInitializeRigWithInvalidData(t *testing.T) {

// 	rig := goHamlib.Rig{}
// 	rig.SetDebugLevel(goHamlib.RIG_DEBUG_NONE)

// 	//rig model must be > 0
// 	if err := rig.Init(0); fmt.Sprint(err) != "init_rig: invalid rig model" {
// 		t.Fatal("error must be raised on invalid rig model")
// 	}

// 	var p goHamlib.Port_t
// 	p.RigPortType = 1
// 	p.Portname = "/dev/myport" // invalid
// 	p.Baudrate = 38400
// 	p.Databits = 8
// 	p.Stopbits = 1
// 	p.Parity = goHamlib.N
// 	p.Handshake = goHamlib.NO_HANDSHAKE

// 	rig.SetPort(p)

// 	if err := rig.Init(1); fmt.Sprint(err) != "init_rig: invalid parameter" {
// 		t.Fatal("error must be raised when port_t struct is supplied with invalid data")
// 	}

// 	rig.Close()
// 	rig.Cleanup()
// }

//Tests Frequency Get & Sep of the dummyRig
func TestDummyRigSetGetFreq(t *testing.T) {

	rig, err := getDummyRig()
	if err != nil {
		t.Fatal(err)
	}

	freq, err := rig.GetFreq(goHamlib.RIG_VFO_A)
	if err != nil {
		t.Fatal(err)
	}

	if freq != 145000000 {
		t.Fatal("frequency of Dummyrig should be 145.000.000 Hz")
	}

	var testFreq float64
	testFreq = 7005000

	// Test Set & Get Frequency on all available VFOs
	for _, vfo := range rig.Caps.Vfos {
		vfoValue := goHamlib.VfoValue[vfo]
		if err := rig.SetFreq(vfoValue, testFreq); err != nil {
			t.Fatal(err)
		}

		freq, err = rig.GetFreq(vfoValue)
		if err != nil {
			t.Fatal(err)
		}

		if freq != testFreq {
			t.Fatalf("Could not set/get Frequency for %s", vfo)
		}
	}

	rig.Close()
	rig.Cleanup()
}

//Tests vfo set & go of the dummyRig
func TestDummyRigSetGetVfo(t *testing.T) {

	rig, err := getDummyRig()
	if err != nil {
		t.Fatal(err)
	}

	// Test Set & Get Frequency on all available VFOs
	for _, vfo := range rig.Caps.Vfos {
		vfoValue := goHamlib.VfoValue[vfo]
		if err := rig.SetVfo(vfoValue); err != nil {
			t.Fatal(err)
		}

		gvfo, err := rig.GetVfo()
		if err != nil {
			t.Fatal(err)
		}

		if gvfo != vfoValue {
			t.Fatalf("Could not set/get Vfo: %s", vfo)
		}
	}

	rig.Close()
	rig.Cleanup()
}

func TestDummyRigModeAndFilters(t *testing.T) {

	rig, err := getDummyRig()
	if err != nil {
		t.Fatal(err)
	}

	// for _, mode := range rig.Caps.Modes {
	// 	fmt.Println("Mode:", mode)
	// 	fmt.Println("Filter:", rig.Caps.Filters[mode])
	// }

	// interate over all VFOs
	for _, vfo := range rig.Caps.Vfos {
		vfoValue := goHamlib.VfoValue[vfo]

		// iterate over all modes
		for _, mode := range rig.Caps.Modes {
			modeValue := goHamlib.ModeValue[mode]

			filterAmount := len(rig.Caps.Filters[mode])
			if filterAmount > 0 {
				// iterate over all available filters
				for _, filter := range rig.Caps.Filters[mode] {

					if err := rig.SetMode(vfoValue, modeValue, filter); err != nil {
						t.Fatal(err)
					}
					m, pb, err := rig.GetMode(vfoValue)
					if err != nil {
						t.Fatal(err)
					}
					if m != modeValue {
						t.Fatalf("got mode %s which is inconsistent with set mode %s on vfo %s\n", goHamlib.ModeName[m], mode, vfo)
					}
					if pb != filter {
						t.Fatalf("got filter %dHz which is inconsistent with set filter %dHz for mode %s on vfo %s\n", pb, filter, mode, vfo)
					}
				}
			} else {
				// Not sure if this is the desired behaviour
				// How should a real rig respond to a filter it does not have?
				filter := 500 // Hz
				rig.SetMode(vfoValue, modeValue, filter)
				m, pb, err := rig.GetMode(vfoValue)
				if err != nil {
					t.Fatal(err)
				}
				if m != modeValue {
					t.Fatalf("got mode %s which is inconsistent with set mode %s on vfo %s\n", goHamlib.ModeName[m], mode, vfo)
				}
				if pb != filter {
					t.Fatalf("got filter %dHz which is inconsistent with set filter %dHz for mode %s on vfo %s\n", pb, filter, mode, vfo)
				}
			}
		}
	}

	modeValue := goHamlib.RIG_MODE_CW
	modeName := goHamlib.ModeName[modeValue]

	// Get Narrow Filter for CW
	filter, err := rig.GetPbNarrow(modeValue)
	if err != nil {
		t.Fatal("Could not determin Narrow Passband filter for mode:", modeName)
	}
	if filter != 500 {
		t.Fatalf("Expected for Narrow Passband filter in %s 500Hz; got: %d", modeName, filter)
	}

	// Get Normal Filter for CW
	filter, err = rig.GetPbNormal(modeValue)
	if err != nil {
		t.Fatal("Could not determin Normal Passband filter for mode:", modeName)
	}
	if filter != 2400 {
		t.Fatalf("Expected for Normal Passband filter in %s 2400Hz; got: %d", modeName, filter)
	}

	// Get Wide Filter for CW
	filter, err = rig.GetPbNarrow(modeValue)
	if err != nil {
		t.Fatal("Could not determin Narrow Passband filter for mode:", modeName)
	}
	if filter != 500 {
		t.Fatalf("Expected for Wide Passband filter in %s 500Hz; got: %d", modeName, filter)
	}

	rig.Close()
	rig.Cleanup()
}

func TestDummyRigPTT(t *testing.T) {

	rig, err := getDummyRig()
	if err != nil {
		t.Fatal(err)
	}

	// Test Set & Get PTT on all available VFOs
	for _, vfo := range rig.Caps.Vfos {
		vfoValue := goHamlib.VfoValue[vfo]

		// iterate over all PTT possibilities
		for pttValue, pttName := range goHamlib.PttName {

			if err := rig.SetPtt(vfoValue, pttValue); err != nil {
				t.Fatal(err)
			}

			ptt, err := rig.GetPtt(vfoValue)
			if err != nil {
				t.Fatal(err)
			}

			if pttValue == goHamlib.RIG_PTT_OFF {
				if ptt != goHamlib.RIG_PTT_OFF {
					t.Fatalf("inconsisted values! Set: %s, Should be: %s", pttName, goHamlib.PttName[goHamlib.RIG_PTT_OFF])
				}
			} else {
				// Dummy Rig returns just "ON" also for "ON_DATA" and "ON_MIC"
				if ptt != goHamlib.RIG_PTT_ON {
					t.Fatalf("inconsisted values! Set: %s, Got: %s", pttName, goHamlib.PttName[goHamlib.RIG_PTT_ON])
				}
			}
		}
	}

	rig.Close()
	rig.Cleanup()
}

func TestDummyRigRit(t *testing.T) {

	rig, err := getDummyRig()
	if err != nil {
		t.Fatal(err)
	}

	// Test Set & Get RIT on all available VFOs
	for _, vfo := range rig.Caps.Vfos {
		vfoValue := goHamlib.VfoValue[vfo]

		ritTestValues := []int{-9999, -5330, -1, 0, 1, 100, 9999}
		for _, ritTV := range ritTestValues {
			if err := rig.SetRit(vfoValue, ritTV); err != nil {
				t.Fatal(err)
			}

			fmt.Println("") // without this command, the test will cause an application crash ?!
			rit, err := rig.GetRit(vfoValue)
			if err != nil {
				t.Fatal(err)
			}

			if rit != ritTV {
				t.Fatalf("rit value set (%dHz) does not correspond with read rit value (%dHz)", ritTV, rit)
			}
		}
	}

	rig.Close()
	rig.Cleanup()
}

func TestDummyRigXit(t *testing.T) {

	rig, err := getDummyRig()
	if err != nil {
		t.Fatal(err)
	}

	// Test Set & Get XIT on all available VFOs
	for _, vfo := range rig.Caps.Vfos {
		vfoValue := goHamlib.VfoValue[vfo]

		xitTestValues := []int{-9999, -5330, -1, 0, 1, 100, 9999}
		for _, xitTV := range xitTestValues {

			if err := rig.SetXit(vfoValue, xitTV); err != nil {
				t.Fatal(err)
			}

			fmt.Println("") // without this command, the test will cause an application crash ?!
			xit, err := rig.GetXit(vfoValue)
			if err != nil {
				t.Fatal(err)
			}

			if xit != xitTV {
				t.Fatalf("xit value set (%dHz) does not correspond with read xit value (%dHz)", xitTV, xit)
			}
		}
	}

	rig.Close()
	rig.Cleanup()
}

func TestDummyRigSplitOperations(t *testing.T) {

	rig, err := getDummyRig()
	if err != nil {
		t.Fatal(err)
	}

	// TBD Write test case for split operations

	rig.Close()
	rig.Cleanup()
}

func TestDummyRigInfo(t *testing.T) {

	rig, err := getDummyRig()
	if err != nil {
		t.Fatal(err)
	}

	info, err := rig.GetInfo()
	fmt.Println(info)

	info_expected := "Nothing much (dummy)"
	if info != info_expected {
		t.Fatalf("info string does not match! got: %s; but expected: %s", info, info_expected)
	}

	rig.Close()
	rig.Cleanup()
}

func TestDummyRigPowerStat(t *testing.T) {

	rig, err := getDummyRig()
	if err != nil {
		t.Fatal(err)
	}

	rig.Close()
	rig.Cleanup()

}
