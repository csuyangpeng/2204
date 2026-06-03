package nas

import "errors"

// Common error definition
var (
	ErrInvalidEpd            = errors.New("invalid epd in nas header")
	ErrInvalidSecHeaderType  = errors.New("invalid security header type in nas header")
	ErrInvalidMmMsgType      = errors.New("invalid mm message type in nas header")
	ErrInvalidSmMsgType      = errors.New("invalid sm message type in nas header")
	ErrInvalidPtiType        = errors.New("invalid PTI in nas header")
	ErrInvalidPsiType        = errors.New("invalid PSI in nas header")
	ErrInvalidMmNasMsgHeader = errors.New("invalid mm nas message header")
	ErrInvalidSmNasMsgHeader = errors.New("invalid sm nas message header")
	ErrInvalidMobileIdType   = errors.New("invalid mobile identity type")
	ErrInvalidIeLength       = errors.New("invalid nas ie length")
	ErrDecodeNasIeFailed     = errors.New("decode nas ie failed")
)
