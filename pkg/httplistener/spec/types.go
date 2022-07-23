package spec

type HttpListenerConfig struct {
	Host     string `json:"host"`
	Port     int32  `json:"port"`
	Basepath string `json:"basepath"`
}
