/**
 *  btrfs is just the subshell wrappers for btrfs administration tools,
 *  running via os/exec.  For further usage of the APIs from its library, please
 *  check out Docker's btrfs driver:
 *   https://github.com/moby/moby/blob/master/daemon/graphdriver/btrfs/btrfs.go
 */

package main

import (
    "fmt"
    "os/exec"
)

type btrfsSubVol struct{}

/* TODO */
func (sv btrfsSubVol) list()     {}
func (sv btrfsSubVol) create()   {}
func (sv btrfsSubVol) delete()   {}
func (sv btrfsSubVol) snapshot() {}

const (
    Cmd string = "btrfs subvolume"
)
