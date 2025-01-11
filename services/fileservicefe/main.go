package main

import (
	"Linda/baselibs/abstractions/xlog"
	"Linda/baselibs/abstractions/xos"
	"Linda/baselibs/multifs"
	"flag"
	"os"
	"path"
)

var (
	port = flag.Int("port", 5555, "the port of file service FE")
)

func main() {
	flag.Parse()
	xlog.Initial()
	s := multifs.New(multifs.NewFileServiceInput{
		Port:    *port,
		BaseDir: path.Join(xos.CurrentPath(), "fileBasePath"),
		Type:    multifs.FileServiceType_Local,
	})
	s.Start()
	signal := xos.WaitForSignal(os.Kill, os.Interrupt)
	xlog.Infof("end due to %v", signal)
}
