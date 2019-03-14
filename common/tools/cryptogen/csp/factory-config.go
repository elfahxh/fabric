package csp

import (
	"fmt"
	"os"

	"github.com/hyperledger/fabric/bccsp/factory"
)

func GetCspFactoryConfig(keystorePath, defaultBccsp string) (opts *factory.FactoryOpts) {
	fmt.Println("\n--> GetCspFactoryConfig: ", keystorePath)
	library := os.Getenv("CORE_PEER_BCCSP_PLUGIN_LIBRARY")
	if library == "" {
		library = "bccsp-grep11.so"
	}
	conn_address := os.Getenv("CORE_PEER_BCCSP_PLUGIN_CONFIG_CONNECTION_ADDRESS")
	conn_port := os.Getenv("CORE_PEER_BCCSP_PLUGIN_CONFIG_CONNECTION_PORT")
	if conn_address == "" {
		conn_address = "9.47.152.179"
	}
	if conn_port == "" {
		conn_port = "30559"
	}
	conf := make(map[string]interface{})
	conf["Security"] = 384
	conf["Direct"] = false
	conf["Hash"] = "SHA3"
	// conf["KeystorePath"] = "msp/keystore"
	conf["KeystorePath"] = keystorePath
	conf["LogLevel"] = "debug"
	conf["NoKeyImport"] = true
	conf["SoftwareVerify"] = false
	tls := make(map[string]interface{})
	tls["Enabled"] = false
	conf["TLS"] = tls
	conn := make(map[string]interface{})
	conn["Address"] = conn_address
	conn["Port"] = conn_port
	conf["Connection"] = conn

	fmt.Println("conf --> ", conf)
	if "SW" == defaultBccsp {
		opts = &factory.FactoryOpts{
			ProviderName: "SW",
			SwOpts: &factory.SwOpts{
				HashFamily: "SHA2",
				SecLevel:   256,

				FileKeystore: &factory.FileKeystoreOpts{
					KeyStorePath: keystorePath,
				},
			},
		}
	} else {
		opts = &factory.FactoryOpts{
			ProviderName: "PLUGIN",
			PluginOpts: &factory.PluginOpts{
				Library: library,
				Config:  conf,
			},
		}
	}
	return
}

func main() {
	opts := GetCspFactoryConfig("./", "PLUGIN")
	fmt.Println("opts --> ", opts, "\nSwOpt --> ", opts.SwOpts, "\nPluginOpts --> ", opts.PluginOpts)
}
