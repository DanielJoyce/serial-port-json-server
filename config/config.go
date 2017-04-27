package config

import (
  "flag"
)

var(
	Version      = "1.94"
	VersionFloat = float32(1.94)
	Addr         = flag.String("addr", ":8989", "http service address")
	SAddr        = flag.String("saddr", ":8990", "https service address")
	SCert        = flag.String("scert", "cert.pem", "https certificate file")
	SKey         = flag.String("skey", "key.pem", "https key file")
	//assets       = flag.String("assets", defaultAssetPath(), "path to assets")
	//	verbose = flag.Bool("v", true, "show debug logging")
	Verbose = flag.Bool("v", false, "show debug logging")
	//homeTempl *template.Template
	IsLaunchSelf = flag.Bool("ls", false, "launch self 5 seconds later")
	IsAllowExec  = flag.Bool("allowexec", false, "Allow terminal commands to be executed")

	// regular expression to sort the serial port list
	// typically this wouldn't be provided, but if the user wants to clean
	// up their list with a regexp so it's cleaner inside their end-user interface
	// such as ChiliPeppr, this can make the massive list that Linux gives back
	// to you be a bit more manageable
	RegExpFilter = flag.String("regex", "", "Regular expression to filter serial port list")

	// allow garbageCollection()
	//isGC = flag.Bool("gc", false, "Is garbage collection on? Off by default.")
	//isGC = flag.Bool("gc", true, "Is garbage collection on? Off by default.")
	GCType = flag.String("gc", "std", "Type of garbage collection. std = Normal garbage collection allowing system to decide (this has been known to cause a stop the world in the middle of a CNC job which can cause lost responses from the CNC controller and thus stalled jobs. use max instead to solve.), off = let memory grow unbounded (you have to send in the gc command manually to garbage collect or you will run out of RAM eventually), max = Force garbage collection on each recv or send on a serial port (this minimizes stop the world events and thus lost serial responses, but increases CPU usage)")

	// whether to do buffer flow debugging
	BufFlowDebugType = flag.String("bufflowdebug", "off", "off = (default) We do not send back any debug JSON, on = We will send back a JSON response with debug info based on the configuration of the buffer flow that the user picked")

	// hostname. allow user to override, otherwise we look it up
	Hostname = flag.String("hostname", "unknown-hostname", "Override the hostname we get from the OS")
)