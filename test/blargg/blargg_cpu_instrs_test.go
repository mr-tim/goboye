// +build blargg

package blargg

import (
	"flag"
	"fmt"
	"github.com/mr-tim/goboye/internal/pkg/goboye"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

// run tests with eg:
// go test -tags=blargg ./test/blargg -args -blargg_roms=/path/to/gb-test-roms
var blarggRomsPath = flag.String("blargg_roms", "", "Path to blargg roms")

func TestBlarggCpuInstrs01(t *testing.T) {
	doBlargTest(t, "/cpu_instrs/individual/01-special.gb")
}

func TestBlarggCpuInstrs02(t *testing.T) {
	doBlargTest(t, "/cpu_instrs/individual/02-interrupts.gb")
}

func TestBlarggCpuInstrs03(t *testing.T) {
	doBlargTest(t, "/cpu_instrs/individual/03-op sp,hl.gb")
}

func TestBlarggCpuInstrs04(t *testing.T) {
	doBlargTest(t, "/cpu_instrs/individual/04-op r,imm.gb")
}

func TestBlarggCpuInstrs05(t *testing.T) {
	doBlargTest(t, "/cpu_instrs/individual/05-op rp.gb")
}

func TestBlarggCpuInstrs06(t *testing.T) {
	doBlargTest(t, "/cpu_instrs/individual/06-ld r,r.gb")
}

func TestBlarggCpuInstrs07(t *testing.T) {
	doBlargTest(t, "/cpu_instrs/individual/07-jr,jp,call,ret,rst.gb")
}

func TestBlarggCpuInstrs08(t *testing.T) {
	doBlargTest(t, "/cpu_instrs/individual/08-misc instrs.gb")
}

func TestBlarggCpuInstrs09(t *testing.T) {
	doBlargTest(t, "/cpu_instrs/individual/09-op r,r.gb")
}

func TestBlarggCpuInstrs10(t *testing.T) {
	doBlargTest(t, "/cpu_instrs/individual/10-bit ops.gb")
}

func TestBlarggCpuInstrs11(t *testing.T) {
	doBlargTest(t, "/cpu_instrs/individual/11-op a,(hl).gb")
}

func doBlargTest(t *testing.T, rom string) {
	if *blarggRomsPath == "" {
		t.Fatal("Path to blargg roms not specified!")
	}
	pathToRom := *blarggRomsPath + rom

	e := goboye.NewEmulator()
	e.LoadRomImage(pathToRom)
	e.SetDebug(true)
	e.ContinueDebugging(false)

	fmt.Printf("\nSerial output:\n\n")
	fmt.Printf(e.SerialOutput())
	if strings.HasSuffix(e.SerialOutput(), "Failed\n") {
		t.Fail()
	} else {
		assert.True(t, true)
	}
}
