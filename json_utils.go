package adguard

import "encoding/json"

// JSONMarshalFunc is a function signature compatible with json.Marshal.
type JSONMarshalFunc func(v any) ([]byte, error)

// JSONMarshal is a package-level variable that points to the real
// json.Marshal by default. Tests can override this to inject errors.
var JSONMarshal JSONMarshalFunc = json.Marshal
