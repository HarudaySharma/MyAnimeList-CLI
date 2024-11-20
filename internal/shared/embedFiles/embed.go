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
)

func init() {
    DaemonPort = strings.ReplaceAll(DaemonPort, "\n", "")
    MalApiUrl = strings.ReplaceAll(MalApiUrl, "\n", "")
    MalCodeChallenge = strings.ReplaceAll(MalCodeChallenge, "\n", "")
    MalClientId = strings.ReplaceAll(MalClientId, "\n", "")

    PreviewDataCacheDir = strings.ReplaceAll(PreviewDataCacheDir, "\n", "")
    PreviewImageCacheDir = strings.ReplaceAll(PreviewImageCacheDir, "\n", "")

}
