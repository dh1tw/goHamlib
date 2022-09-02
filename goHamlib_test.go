package goHamlib_test

import (
	"testing"

	"github.com/dl1igc/goHamlib"
)

// Test consistency of Vfo Value and Name maps
func TestVfoMaps(t *testing.T) {

	// Test VFOValue map
	for vfoName, VFOValue := range goHamlib.VFOValue {
		_, ok := goHamlib.VFOName[VFOValue]
		if !ok {
			t.Fatalf("VFO %d does not exist in VfoName map", VFOValue)
		}
		if vfoName != goHamlib.VFOName[VFOValue] {
			t.Fatalf("Name of VFO inconsisted: %s", vfoName)
		}
	}

	// Test VfoName map
	for VFOValue, vfoName := range goHamlib.VFOName {
		_, ok := goHamlib.VFOValue[vfoName]
		if !ok {
			t.Fatalf("VFO %s does not exist in VFOValue map", vfoName)
		}
		if VFOValue != goHamlib.VFOValue[vfoName] {
			t.Fatalf("Value of VFO inconsisted: %s", vfoName)
		}
	}
}

// Test consistency of OperationValue and OperationName maps
func TestOperationMaps(t *testing.T) {

	// Test OperationValue map
	for opName, opValue := range goHamlib.VFOOperationValue {
		_, ok := goHamlib.VFOOperationName[opValue]
		if !ok {
			t.Fatalf("Operation %d does not exist in OperationName map", opValue)
		}
		if opName != goHamlib.VFOOperationName[opValue] {
			t.Fatalf("Name of Operation inconsisted: %s", opName)
		}
	}

	// Test OperationName map
	for opValue, opName := range goHamlib.VFOOperationName {
		_, ok := goHamlib.VFOOperationValue[opName]
		if !ok {
			t.Fatalf("Operation %s does not exist in OperationValue map", opName)
		}
		if opValue != goHamlib.VFOOperationValue[opName] {
			t.Fatalf("Value of Operation inconsisted: %s", opName)
		}
	}
}

// Test consistency of ModeValue and ModeName maps
func TestModeMaps(t *testing.T) {

	// Test ModeValue map
	for modeName, modeValue := range goHamlib.ModeValue {
		_, ok := goHamlib.ModeName[modeValue]
		if !ok {
			t.Fatalf("Mode %d does not exist in ModeName map", modeValue)
		}
		if modeName != goHamlib.ModeName[modeValue] {
			t.Fatalf("Name of Mode inconsisted: %s", modeName)
		}
	}

	// Test ModeName map
	for modeValue, modeName := range goHamlib.ModeName {
		_, ok := goHamlib.ModeValue[modeName]
		if !ok {
			t.Fatalf("Mode %s does not exist in ModeValue map", modeName)
		}
		if modeValue != goHamlib.ModeValue[modeName] {
			t.Fatalf("Value of Mode inconsisted: %s", modeName)
		}
	}
}

// Test consistency of RigPowerValue and RigPowerName maps
func TestRigPowerMaps(t *testing.T) {

	// Test ModeValue map
	for rpName, rpValue := range goHamlib.RigPowerValue {
		_, ok := goHamlib.RigPowerName[rpValue]
		if !ok {
			t.Fatalf("RigPower %d does not exist in RigPowerName map", rpValue)
		}
		if rpName != goHamlib.RigPowerName[rpValue] {
			t.Fatalf("Name of RigPower inconsisted: %s", rpName)
		}
	}

	// Test RigPowerName map
	for rpValue, rpName := range goHamlib.RigPowerName {
		_, ok := goHamlib.RigPowerValue[rpName]
		if !ok {
			t.Fatalf("RigPower %s does not exist in RigPowerValue map", rpName)
		}
		if rpValue != goHamlib.RigPowerValue[rpName] {
			t.Fatalf("Value of RigPower inconsisted: %s", rpName)
		}
	}
}

// Test consistency of LevelValue and LevelName maps
func TestLevelMaps(t *testing.T) {

	// Test LevelValue map
	for lName, lValue := range goHamlib.LevelValue {
		_, ok := goHamlib.LevelName[lValue]
		if !ok {
			t.Fatalf("Level %d does not exist in LevelName map", lValue)
		}
		if lName != goHamlib.LevelName[lValue] {
			t.Fatalf("Name of Level inconsisted: %s", lName)
		}
	}

	// Test LevelName map
	for lValue, lName := range goHamlib.LevelName {
		_, ok := goHamlib.LevelValue[lName]
		if !ok {
			t.Fatalf("Level %s does not exist in LevelValue map", lName)
		}
		if lValue != goHamlib.LevelValue[lName] {
			t.Fatalf("Value of Level inconsisted: %s", lName)
		}
	}
}

// Test consistency of ParmValue and ParmName maps
func TestParmMaps(t *testing.T) {

	// Test ParmValue map
	for pName, pValue := range goHamlib.ParmValue {
		_, ok := goHamlib.ParmName[pValue]
		if !ok {
			t.Fatalf("Parm %d does not exist in ParmName map", pValue)
		}
		if pName != goHamlib.ParmName[pValue] {
			t.Fatalf("Name of Parm inconsisted: %s", pName)
		}
	}

	// Test ParmName map
	for pValue, pName := range goHamlib.ParmName {
		_, ok := goHamlib.ParmValue[pName]
		if !ok {
			t.Fatalf("Parm %s does not exist in ParmValue map", pName)
		}
		if pValue != goHamlib.ParmValue[pName] {
			t.Fatalf("Value of Parm inconsisted: %s", pName)
		}
	}
}

// Test consistency of FuncValue and FuncName maps
func TestFuncMaps(t *testing.T) {

	// Test FuncValue map
	for fName, fValue := range goHamlib.FuncValue {
		_, ok := goHamlib.FuncName[fValue]
		if !ok {
			t.Fatalf("Func %d does not exist in FuncName map", fValue)
		}
		if fName != goHamlib.FuncName[fValue] {
			t.Fatalf("Name of Func inconsisted: %s", fName)
		}
	}

	// Test FuncName map
	for fValue, fName := range goHamlib.FuncName {
		_, ok := goHamlib.FuncValue[fName]
		if !ok {
			t.Fatalf("Func %s does not exist in FuncValue map", fName)
		}
		if fValue != goHamlib.FuncValue[fName] {
			t.Fatalf("Value of Func inconsisted: %s", fName)
		}
	}
}

// Test consistency of PttValue and PttName maps
func TestPttMaps(t *testing.T) {

	// Test PttValue map
	for pName, pValue := range goHamlib.PttValue {
		_, ok := goHamlib.PttName[pValue]
		if !ok {
			t.Fatalf("Ptt %d does not exist in PttName map", pValue)
		}
		if pName != goHamlib.PttName[pValue] {
			t.Fatalf("Name of Ptt inconsisted: %s", pName)
		}
	}

	// Test PttName map
	for pValue, pName := range goHamlib.PttName {
		_, ok := goHamlib.PttValue[pName]
		if !ok {
			t.Fatalf("Ptt %s does not exist in PttValue map", pName)
		}
		if pValue != goHamlib.PttValue[pName] {
			t.Fatalf("Value of Ptt inconsisted: %s", pName)
		}
	}
}

// Test consistency of AntValue and AntName maps
func TestAntennaMaps(t *testing.T) {

	// Test AntValue map
	for aName, aValue := range goHamlib.AntValue {
		_, ok := goHamlib.AntName[aValue]
		if !ok {
			t.Fatalf("Antenna %d does not exist in AntName map", aValue)
		}
		if aName != goHamlib.AntName[aValue] {
			t.Fatalf("Name of Antenna inconsisted: %s", aName)
		}
	}

	// Test AntName map
	for aValue, aName := range goHamlib.AntName {
		_, ok := goHamlib.AntValue[aName]
		if !ok {
			t.Fatalf("Antenna %s does not exist in AntValue map", aName)
		}
		if aValue != goHamlib.AntValue[aName] {
			t.Fatalf("Value of Antenna inconsisted: %s", aName)
		}
	}
}

// Test consistency of DebugLevelValue and DebugLevelName maps
func TestDebugLevelMaps(t *testing.T) {

	// Test DebugLevelValue map
	for dName, dValue := range goHamlib.DebugLevelValue {
		_, ok := goHamlib.DebugLevelName[dValue]
		if !ok {
			t.Fatalf("Debug Level %d does not exist in DebugLevelName map", dValue)
		}
		if dName != goHamlib.DebugLevelName[dValue] {
			t.Fatalf("Name of Debug Level inconsisted: %s", dName)
		}
	}

	// Test DebugLevelName map
	for dValue, dName := range goHamlib.DebugLevelName {
		_, ok := goHamlib.DebugLevelValue[dName]
		if !ok {
			t.Fatalf("DebugLevel %s does not exist in DebugLevelValue map", dName)
		}
		if dValue != goHamlib.DebugLevelValue[dName] {
			t.Fatalf("Value of DebugLevel inconsisted: %s", dName)
		}
	}
}

func TestCIntToBool(t *testing.T) {
	x, err := goHamlib.CIntToBool(0)
	if err != nil || x != false {
		t.Fatalf("0 should result in false")
	}

	x, err = goHamlib.CIntToBool(1)
	if err != nil || x != true {
		t.Fatalf("1 should result in true")
	}

	_, err = goHamlib.CIntToBool(-1)
	if err == nil {
		t.Fatalf("-1 should result in err")
	}

	_, err = goHamlib.CIntToBool(2)
	if err == nil {
		t.Fatalf("2 should result in err")
	}
}

func TestBoolToInt(t *testing.T) {
	x, err := goHamlib.BoolToCint(true)
	if err != nil || x != 1 {
		t.Fatal("true should result in 1")
	}

	x, err = goHamlib.BoolToCint(false)
	if err != nil || x != 0 {
		t.Fatal("false should result in 0")
	}

	x, err = goHamlib.BoolToCint(false)
	if err != nil || x != 0 {
		t.Fatal("false should result in 0")
	}

}

func TestListModels(t *testing.T) {

	supported := goHamlib.ListModels()
	if len(supported) < 10 {
		t.Errorf("expected at least 10 supported hamlib models")
	}
	foundDummy := false
	for _, v := range supported {
		if v.Manufacturer == "Hamlib" && v.Model == "Dummy" {
			foundDummy = true
		}
	}
	if !foundDummy {
		t.Errorf("did not find the hamlib standard dummy driver")
	}
}

func TestDebugCallback(t *testing.T) {
	called := 0
	goHamlib.SetDebugCallback(func(lvl goHamlib.DebugLevel, msg string) {
		called++
	})
	goHamlib.SetDebugLevel(goHamlib.DebugTrace)
	goHamlib.ListModels()
	goHamlib.SetDebugLevel(goHamlib.DebugWarn)
	if called == 0 {
		t.Errorf("expected a debug message, got none")
	}
}
