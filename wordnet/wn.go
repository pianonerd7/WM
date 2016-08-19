package wordnet

//#cgo CFLAGS: -I/usr/local/WordNet-3.0/include
//#cgo LDFLAGS: /usr/local/WordNet-3.0/lib/libWN.3.dylib
/*
 #include <stdio.h>
 #include <stdlib.h>
 #include <string.h>
 #include "wn.h"
static void printlicense() {
    printf("WordNet License %s\n\n%s", dblicense, license);
}
*/
import "C"
import (
	"errors"
	"unsafe"
)

// PrintLicenses is a mandatory function that must be included here
// or else wordnet won't work
func PrintLicenses() {
	C.printlicense()
}

// InitWN initializes WordNet
func InitWN() (err error) {
	status := C.wninit()
	if int(status) != 0 {
		return errors.New("WN Fatal: can't open database")
	}
	return nil
}

// search is the item to search for, e.g. 'house'.
// dbase maps the the 'pos' or Part Of Speech, e.g. Noun, Verb, Adjective or Adverb
// ptrtyp is the search type. Example are ANTPTR, CAUSETO, et..
// whichsense should be set to 0 (ALLSENSES) to get all meanings.
func FindTheInfo(search string, dbase, ptrtyp, whichsense int) string {
	cSearch := C.CString(search)
	defer C.free(unsafe.Pointer(cSearch))

	cDbase := C.int(dbase)
	cPtrtyp := C.int(ptrtyp)
	cWhichsense := C.int(whichsense)
	resultChar := C.findtheinfo(cSearch, cDbase, cPtrtyp, cWhichsense)
	result := C.GoString(resultChar)
	return result
}

// FindTheInfoDs queries wordnet for the word, parts of sppech and the sensenumber
// and returns the synset for that word in the form of a string
func FindTheInfoDs(search string, dbase, ptrtyp, whichsense int) string {
	cSearch := C.CString(search)
	defer C.free(unsafe.Pointer(cSearch))

	cDbase := C.int(dbase)
	cPtrtyp := C.int(ptrtyp)
	cWhichsense := C.int(whichsense)
	resultPtr := C.findtheinfo_ds(cSearch, cDbase, cPtrtyp, cWhichsense)
	if resultPtr == nil {
		return ""
	}
	resultChar := C.FmtSynset(resultPtr, 0)
	result := C.GoString(resultChar)
	return result
}

// GetSenseLength queries wordnet for the string, and parts of Speech
// and returns how many senses there are
func GetSenseLength(search string, dbase, ptrtyp int) int {
	count := 1

	for count < 50 {
		if FindTheInfo(search, dbase, ptrtyp, count) == "" {
			return count - 1
		}
		count++
	}

	return count
}
