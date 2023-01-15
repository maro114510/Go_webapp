//!/usr/local/go/bin/go
package clock

import (
	"time"
)


type Clocker interface {
	Now() time.Time
} /* Clocker */

type RealClocker struct {}

func ( r RealClocker ) Now() time.Time {
	return time.Now()
} /* Now */

type FixedClocker struct{}

func ( fc FixedClocker ) Now() time.Time {
	return time.Date( 2022, 5, 19, 12, 34, 56, 0, time.UTC )
} /* Now */



// End_Of_Script