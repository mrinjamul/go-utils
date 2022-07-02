/*Package tzinit
This package is use to initialize default timezone for India.

*/
package tzinit

import "os"

func init() {
	os.Setenv("TZ", "Asia/Kolkata")
}
