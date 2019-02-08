/*
 * pepipost_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package PepipostClient

import(
	"pepipost/pepipost_lib/configuration_pkg"
	"pepipost/pepipost_lib/email_pkg"
)


/*
 * Interface for the PEPIPOST_IMPL
 */
type PEPIPOST interface {
     Email()                 email_pkg.EMAIL
     Configuration()         configuration_pkg.CONFIGURATION
}

/*
 * Factory for the PEPIPOST interaface returning PEPIPOST_IMPL
 */
func NewPEPIPOST() PEPIPOST {
    pepipostClient := new(PEPIPOST_IMPL)
    pepipostClient.config = configuration_pkg.NewCONFIGURATION()
    return pepipostClient
}
