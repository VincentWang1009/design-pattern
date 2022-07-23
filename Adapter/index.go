package adapter

type MediaPlayer interface {
	Play(filename string) string
}

type AdvancedMediaPlayer interface {
	PlayFile(filename string) string
}

type MP3Player struct{}

func (m *MP3Player) Play(filename string) string {
	return "playing mp3 file" + filename
}

type MP4Player struct{}

func (m *MP4Player) PlayFile(filename string) string {
	return "playing mp4 file" + filename
}

type VLCPlayer struct{}

func (v *VLCPlayer) PlayFile(filename string) string {
	return "playing vlc file" + filename
}

type MediaPlayerAdapter struct {
	InnerAdvancedMediaPlayer AdvancedMediaPlayer
}

func (m *MediaPlayerAdapter) Play(filename string) string {
	return "playing MediaPlayerAdapter" + filename
}
