module main

go 1.13

require hinimix.com/demo v0.0.0-incompatible

replace hinimix.com/demo => ../demo

require hinimix.com/ifactory v0.0.0-incompatible

replace hinimix.com/ifactory => ../ifactory

require (
	github.com/dullgiulio/pingo v0.0.0-20151111190101-8b1949e35b5a
	github.com/skip2/go-qrcode v0.0.0-20191027152451-9434209cb086 // indirect
	hinimix.com/lib5 v0.0.0-incompatible
)

replace hinimix.com/lib5 => ../lib
