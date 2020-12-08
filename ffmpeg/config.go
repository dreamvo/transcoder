package ffmpeg

// Config represent the high-level config options
type Config struct {
	FFmpegBinPath   string
	FFprobeBinPath  string
	ProgressEnabled bool
	Verbose         bool
}
