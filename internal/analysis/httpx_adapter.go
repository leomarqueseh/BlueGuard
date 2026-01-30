package analysis

import "github.com/leomarqueseh/BlueGuard/internal/httpx"

// httpxResult é apenas um alias para reaproveitar o client httpx
// sem criar dependência direta em toda a engine de análise
type httpxResult = httpx.Result
