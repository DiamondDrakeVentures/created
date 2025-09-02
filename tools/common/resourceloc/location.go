package resourceloc

import (
	"path/filepath"
	"strings"
)

const (
	ResourcePackBaseDir = "resourcepacks"
	DataPackBaseDir     = "datapacks"
)

func ResourcePackPath(packID, resourceID string) string {
	namespace, resourcePath := ResourcePath(resourceID)

	return filepath.Join(ResourcePackBaseDir, packID, "assets", namespace, resourcePath)
}

func DataPackPath(packID, resourceID string) string {
	namespace, resourcePath := ResourcePath(resourceID)

	return filepath.Join(DataPackBaseDir, packID, "data", namespace, resourcePath)
}

func ResourcePath(resourceID string) (namespace string, path string) {
	if strings.Contains(resourceID, ":") {
		fragments := strings.SplitN(resourceID, ":", 2)
		namespace = fragments[0]
		path = fragments[1]
	} else {
		namespace = "minecraft"
		path = resourceID
	}
	return
}

func SubPathToDir(path string) string {
	return strings.ReplaceAll(path, ":", string(filepath.Separator))
}

func SubPathToUnderscore(path string) string {
	return strings.ReplaceAll(path, ":", "_")
}
