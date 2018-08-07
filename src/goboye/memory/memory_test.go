package memory

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestLoadRom(t *testing.T) {
	buf := make([]byte, ROM_SIZE)
	for i := 0; i < ROM_SIZE; i++ {
		buf[i] = byte(i % 0x100)
	}

	f, err := ioutil.TempFile("", "test.rom")
	if err != nil {
		t.Fatal(err)
	}
	f.Write(buf)
	f.Close()
	defer os.Remove(f.Name())

	m := memoryMap{
		mem: make([]byte, MEM_SIZE),
	}

	m.LoadRomImage(f.Name())

	assert.Equal(t, uint8(0x00), m.mem[0x0000])
	assert.Equal(t, uint8(0x23), m.mem[0x0023])
	assert.Equal(t, uint8(0x69), m.mem[0x4769])
}
