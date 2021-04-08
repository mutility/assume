// +build noverify

package assume

func True(v bool) bool {
	return v
}

func Truef(v bool, format string, a ...interface{}) bool {
	return v
}

func False(v bool) bool {
	return v
}

func Falsef(v bool, format string, a ...interface{}) bool {
	return v
}

func Ok(err error) error {
	return err
}

func Okf(err error, format string, a ...interface{}) error {
	return err
}

func Err(err error) error {
	return err
}

func Errf(err error, format string, a ...interface{}) error {
	return err
}

func Unreached() {
	// no-op
}

func Unreachedf(format string, a ...interface{}) {
	// no-op
}
