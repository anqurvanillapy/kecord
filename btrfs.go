/**
 *  btrfs is just the subshell wrappers for btrfs administration tools,
 *  running via os/exec.  For further usage of the APIs from its library, please
 *  check out Docker's btrfs driver:
 *   https://github.com/moby/moby/blob/master/daemon/graphdriver/btrfs/btrfs.go
 */

package main

import (
    "bytes"
    "log"
    "os/exec"
)

const (
    Cmd string = "btrfs"
)

type btrfsSubvol struct{}

func (sv btrfsSubvol) list(path string) string {
    var out bytes.Buffer
    cmd := exec.Command(Cmd, "subvolume", "list", path);
    cmd.Stdout = &out
    err := cmd.Run()

    if err != nil { log.Fatal(err) }
    return out.String()
}

func (sv btrfsSubvol) create()   {}
func (sv btrfsSubvol) delete()   {}
func (sv btrfsSubvol) snapshot() {}
