//Package goid exports a private function sourced from github.com/tylerb/gls
package goid

//ID returns the ID the the current Go routine
func ID() uint64 {
	return curGoroutineID()
}
