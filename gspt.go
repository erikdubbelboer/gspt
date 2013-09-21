
package gspt

/*
#include "setproctitle.h"

static void cSetProcTitle(char *title) {
#if HAVE_SETPROCTITLE || defined(HAVE_SETPROCTITLE_REPLACEMENT)
  setproctitle("%s", title);
#else
  (void)title;
#endif
}

static void cInit(int argc, char *arg0) {
#ifdef HAVE_SETPROCTITLE_REPLACEMENT
  spt_init(argc, arg0);
#else
  (void)argc;
  (void)argv;
#endif
}
*/
import "C"

import (
  "os"
  "unsafe"
)


func init() {
  argc := C.int(len(os.Args))
  arg0 := C.CString(os.Args[0])
  defer C.free(unsafe.Pointer(arg0))

  C.cInit(argc, arg0)
}


func SetProcTitle(title string) {
  cs := C.CString(title)
  defer C.free(unsafe.Pointer(cs))

  C.cSetProcTitle(cs)
}

