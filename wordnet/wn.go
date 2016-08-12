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

func PrintLicenses() {
	C.printlicense()
}

// initialize WordNet
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

func FindTheInfo_ds(search string, dbase, ptrtyp, whichsense int) string {
	cSearch := C.CString(search)
	defer C.free(unsafe.Pointer(cSearch))

	cDbase := C.int(dbase)
	cPtrtyp := C.int(ptrtyp)
	cWhichsense := C.int(whichsense)
	resultPtr := C.findtheinfo_ds(cSearch, cDbase, cPtrtyp, cWhichsense)
	resultChar := C.FmtSynset(resultPtr, 0)
	result := C.GoString(resultChar)
	return result
}
