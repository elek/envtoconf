package app

import "testing"

func TestParseKeyValues(t *testing.T) {
	var envs = make(map[string]string)
	envs["CORE-SITE.XML_key.something"] = "value"
	var result = ParseKeyValues(envs)
	if len(result) != 1 {
		t.Error("There should be only one record")
	}

	if first, ok := result["core-site.xml"]; ok {
		if len(first.Entries) != 1 {
			t.Error("There should be one config entries")
		}
	} else {
		t.Error("No core-site.xml in the files")
	}
}




func TestParseKey(t *testing.T) {

	filename, format, configkey, err := ParseKey("CORE-SITE.XML_key.something")
	if err != nil {
		t.Error("Error during the config key parsing: " + err.Error())
	}
	if filename != "core-site.xml" {
		t.Error("Filename has not been parsed well")
	}

	if format != "xml" {
		t.Error("format has not been parsed well")
	}

	if configkey != "key.something" {
		t.Error("Config key has not been parsed well")
	}
}

func TestParseKeyYaml(t *testing.T) {

	filename, format, configkey, err := ParseKey("STORM.YAML_storm.zookeeper.servers.0")
	if err != nil {
		t.Error("Error during the config key parsing: " + err.Error())
	}
	if filename != "storm.yaml" {
		t.Error("Filename has not been parsed well")
	}

	if format != "yaml" {
		t.Error("format has not been parsed well")
	}

	if configkey != "storm.zookeeper.servers.0" {
		t.Error("Config key has not been parsed well")
	}
}

func TestParseKeyLog4j(t *testing.T) {

	filename, format, configkey, err := ParseKey("LOG4J.PROPERTIES_log4j.appender.stdout.layout.ConversionPattern")
	if err != nil {
		t.Error("Error during the config key parsing: " + err.Error())
	}
	if filename != "log4j.properties" {
		t.Error("Filename has not been parsed well")
	}

	if format != "properties" {
		t.Error("format has not been parsed well")
	}

	if configkey != "log4j.appender.stdout.layout.ConversionPattern" {
		t.Error("Config key has not been parsed well")
	}
}





func TestParseKeyWithOptional(t *testing.T) {

	filename, format, configkey, err := ParseKey("CORE-SITE.XML!CONF_key.something")
	if err != nil {
		t.Error("Error during the config key parsing: " + err.Error())
	}
	if filename != "core-site.xml" {
		t.Error("Filename has not been parsed well")
	}

	if format != "conf" {
		t.Error("format has not been parsed well")
	}

	if configkey != "key.something" {
		t.Error("Config key has not been parsed well")
	}
}