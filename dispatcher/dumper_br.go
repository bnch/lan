// +build banchoreader

package dispatcher

import (
	"fmt"
	"os"

	"github.com/bnch/banchoreader/lib"
)

var _bReader = banchoreader.New()

// DumpPackets uses banchoreader to print to the standard output the contents of
// the packets passed as data.
func DumpPackets(data []byte) {
	err := _bReader.Dump(os.Stdout, data)
	if err != nil {
		fmt.Println(err)
	}
}
