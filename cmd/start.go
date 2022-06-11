package main

import (
	"k8s.io/klog"
)

func main() {
	defer klog.Flush()
	klog.Info("this is the first info")
	klog.Warning("this is the first warning")
	klog.Error("this is the first error")
}
