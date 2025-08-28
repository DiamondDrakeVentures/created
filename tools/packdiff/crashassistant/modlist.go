package crashassistant

type ModList map[string]Mod

type Mod struct {
	JarName string `json:"jarName"`
	ModID   string `json:"modId"`
	Version string `json:"version"`
}
