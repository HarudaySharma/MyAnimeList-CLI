package embedfiles

import (
	_ "embed"
	"strings"
)

var (
	//go:embed DAEMON_PORT
	DaemonPort string

	//go:embed MAL_API_URL
	MalApiUrl string

	//go:embed MAL_CLIENT_ID
	MalClientId string

	//go:embed CODE_CHALLENGE
	MalCodeChallenge string

	//go:embed PREVIEW_IMAGE_CACHE_DIR
	PreviewImageCacheDir string

	//go:embed PREVIEW_DATA_CACHE_DIR
	PreviewDataCacheDir string

	//go:embed CLIENT_CACHE_DIR
	ClientCacheDir string
)

func init() {
	const defaultDaemonPort = "42069"
	const defaultMalApiUrl = "https://api.myanimelist.net/v2"
	const defaultMalCodeChallenge = ""
	const defaultMalClientId = ""
	const defaultClientCacheDir = "/tmp/mal-cli"
	const defaultPreviewDataCacheDir = "/tmp/mal-cli/data"
	const defaultPreviewImageCacheDir = "/tmp/mal-cli/images"

	DaemonPort = strings.ReplaceAll(DaemonPort, "\n", "")
	if DaemonPort == "" {
		DaemonPort = defaultDaemonPort
	}

	MalApiUrl = strings.ReplaceAll(MalApiUrl, "\n", "")
	if MalApiUrl == "" {
		MalApiUrl = defaultMalApiUrl
	}

	MalCodeChallenge = strings.ReplaceAll(MalCodeChallenge, "\n", "")
	if MalCodeChallenge == "" {
		MalCodeChallenge = defaultMalCodeChallenge
	}

	MalClientId = strings.ReplaceAll(MalClientId, "\n", "")
	if MalClientId == "" {
		MalClientId = defaultMalClientId
	}

	PreviewDataCacheDir = strings.ReplaceAll(PreviewDataCacheDir, "\n", "")
	if PreviewDataCacheDir == "" {
		PreviewDataCacheDir = defaultPreviewDataCacheDir
	}

	PreviewImageCacheDir = strings.ReplaceAll(PreviewImageCacheDir, "\n", "")
	if PreviewImageCacheDir == "" {
		PreviewImageCacheDir = defaultPreviewImageCacheDir
	}

	ClientCacheDir = strings.ReplaceAll(ClientCacheDir, "\n", "")
	if ClientCacheDir == "" {
		ClientCacheDir = defaultClientCacheDir
	}

}
