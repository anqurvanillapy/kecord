package main

import (
    "os"
    "fmt"
)

const (
    BtrfsPath string = "/var/kecord"
)

func usage() {
    fmt.Println(
`Usage: kecord COMMANDS files...

Commands:
    init    Create an image
    pull    Pull an image from Docker Hub
    rm      Remove an image or container
    ls      List images
    ps      List containers
    run     Create a container
    exec    Execute a command in a running container
    logs    View logs from a container
    commit  Commit a container to an image
    help    Show usage`)
}

/* TODO */
func createImage() {}

func pullImage() {
    // fmt.Println(BtrfsPath)
}

func remove() {}
func listImages() {}
func listContainers() {}
func createContainer() {}
func execute() {}
func catLog() {}
func commit() {}

func main() {
    subcmds := map[string]func() {
        "help":     usage,
        "init":     createImage,
        "pull":     pullImage,
        "rm":       remove,
        "ls":       listImages,
        "ps":       listContainers,
        "run":      createContainer,
        "exec":     execute,
        "logs":     catLog,
        "commit":   commit,
    }

    cmd := subcmds[os.Args[1]]
    if cmd != nil {
        cmd()
    } else {
        fmt.Printf("Unrecognized command: %s\n\n", os.Args[1])
        usage()
        os.Exit(1)
    }
}
