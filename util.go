package xenditfasthttp

import (
	"encoding/base64"
	"fmt"
)

const (
	// VAMandiri is ..
	VAMandiri = "MANDIRI"
	// VABNI is ..
	VABNI = "BNI"
	// VAPermata is ..
	VAPermata = "PERMATA"
	// VABCA is ..
	VABCA = "BCA"

	// VABRI is ...
	VABRI = "BRI"

	// VABNISyariah is ....
	VABNISyariah = "BNI_SYARIAH"

	// OutletIndomaret ...
	OutletIndomaret = "INDOMARET"

	// OutletAlfamart ...
	OutletAlfamart = "ALFAMART"
)

// BasicAuth is ..
func BasicAuth(usr, pwd string) string {
	return fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", usr, pwd))))
}
