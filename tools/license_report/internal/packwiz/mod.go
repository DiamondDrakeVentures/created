package packwiz

type Mod struct {
	Name     string `toml:"name"`
	Filename string `toml:"filename"`
	Side     string `toml:"side"`

	Option   Option   `toml:"option"`
	Download Download `toml:"download"`
	Update   Update   `toml:"update"`
}

type Download struct {
	HashFormat string `toml:"hash-format"`
	Hash       string `toml:"hash"`
	URL        string `toml:"url"`
}

type Option struct {
	Optional    bool   `toml:"optional"`
	Default     bool   `toml:"default"`
	Description string `toml:"description"`
}

type Update struct {
	CurseForge CurseForge `toml:"curseforge"`
	Modrinth   Modrinth   `toml:"modrinth"`
}

type CurseForge struct {
	FileID    int `toml:"file-id"`
	ProjectID int `toml:"project-id"`
}

type Modrinth struct {
	ModID   string `toml:"mod-id"`
	Version string `toml:"version"`
}
