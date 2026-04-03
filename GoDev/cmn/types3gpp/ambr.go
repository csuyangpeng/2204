package types3gpp

import (
	"strconv"
	"strings"

	"fmt"
)

type Ambr struct {
	Uplink   uint64
	Downlink uint64
}

func (p Ambr) String() string {
	return fmt.Sprintf("uplink(%d), downlink(%d)",
		p.Uplink, p.Downlink)
}

func StoreAmbrBitRate(br string) (uint64, error) {
	elms := strings.Split(br, " ")
	if len(elms) != 2 {
		return 0, fmt.Errorf("invalid bitrate formate (%s)", br)
	}
	val, err := strconv.ParseInt(elms[0], 0, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid bitrate value error(%s)", err)
	}
	var unit int
	switch elms[1] {
	case "bps":
		unit = 1
	case "Kbps":
		unit = 1024
	case "Mbps":
		unit = 1024 * 1024
	case "Gbps":
		unit = 1024 * 1024 * 1024
	case "Tbps":
		unit = 1024 * 1024 * 1024 * 1024
	default:
		err = fmt.Errorf("invalid bit rate unit (%s)", elms[1])
	}

	return uint64(int(val) * unit), nil
}
