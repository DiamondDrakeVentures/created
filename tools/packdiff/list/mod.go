package list

import "github.com/diamonddrakeventures/created/tools/packdiff/crashassistant"

type Mod struct {
	ModID   string `json:"mod_id"`
	Jar     string `json:"jar"`
	Version string `json:"version"`
}

func ModFromCrashAssistant(mod crashassistant.Mod) Mod {
	return Mod{
		ModID:   mod.ModID,
		Jar:     mod.JarName,
		Version: mod.Version,
	}
}
