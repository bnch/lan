package packets_test

import (
	"fmt"
	"testing"

	"github.com/bnch/lan/packets"
	"github.com/sergi/go-diff/diffmatchpatch"
)

const specFramesTest = `12000032010000531ea90012000000718308432d33ea4200000000000071830843558fec420a000000000071830843558fec421e000000000071830843558fec422c000000000071830843558fec423f000000000071830843558fec424c000000000071830843558fec4261000000000071830843558fec4271000000000071830843558fec4284000000000071830843558fec4290000000000071830843558fec42a9000000000071830843558fec42b5000000000071830843558fec42c8000000000071830843558fec42d4000000000071830843558fec42e7000000000071830843558fec42f7000000000071830843558fec4209010000000071830843558fec4217010000024522000000000000000000000000000000000000000000000000c80001000000000000000000000000000000000100`

func TestSpecFrames(t *testing.T) {
	var data string
	fmt.Sscanf(specFramesTest, "%x", &data)
	pks, err := packets.Depacketify([]byte(data))
	if err != nil {
		t.Fatal(err)
	}
	if len(pks) != 1 {
		t.Fatalf("len want 1 got %d", len(pks))
	}
	frames := pks[0].(*packets.OsuSpectateFrames)
	out, err := packets.Packetify([]packets.Packetifier{frames})
	if err != nil {
		t.Fatal(err)
	}
	res := fmt.Sprintf("%x", string(out))
	if res != specFramesTest {
		dmp := diffmatchpatch.New()
		diff := dmp.DiffMain(specFramesTest, res, false)
		t.Log(dmp.DiffPrettyText(diff))
		t.Fatalf("Spec frames in & out differ, in:\n%s\nout:\n%s", specFramesTest, res)
	}
}
