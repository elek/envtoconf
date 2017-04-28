package main

import (
	"os"
	"strings"
	"github.com/elek/envtoconf/app"
	"io/ioutil"
	"path"
	"flag"
)

func main() {
	var outputDir = flag.String("outputdir", "/tmp", "Directory where the configuration files are generated.")
	flag.Parse()
	environments := make(map[string]string)
	for _, e := range os.Environ(){
		pair := strings.SplitN(e, "=", 2)
		environments[pair[0]] = pair[1]
	}
	configfiles := app.ParseKeyValues(environments)
	for key, configfile := range configfiles {
		content, err := app.TransformToString(configfile.Entries, configfile.Format)
		println(key)
		if (err == nil) {
			if error := ioutil.WriteFile(path.Join(*outputDir, configfile.File), []byte(content), 0644); error != nil {
				println("Error on writing out file " + configfile.File + " " + error.Error())
			} else {
				println("File " + configfile.File + " has been written out successfullly.")
			}
		} else {
			println(err.Error())
		}
	}
}
