package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/dhowden/tag"
	"github.com/eiannone/keyboard"
	"github.com/gopxl/beep/v2"
    "github.com/gopxl/beep/v2/effects"
    "github.com/gopxl/beep/v2/mp3"
    "github.com/gopxl/beep/v2/speaker"
    "github.com/gopxl/beep/v2/wav"
    "github.com/gopxl/beep/v2/flac"
)

var supportedFormats = map[string]bool{
	".mp3":  true,
	".wav":  true,
	".flac": true,
}

func main() {

	// flags
	shuffle := flag.Bool("s", false, "shuffle playback order")
	flag.Parse()

	// path to dir or file
	path := "."
	if len(flag.Args()) > 0 {
		path = flag.Args()[0]
	}

	// check if path is file or dir
	info, err := os.Stat(path)
	if err != nil {
		log.Fatalf("Error accessing path: %v", err)
	}

	var songs []string

	if info.IsDir() {
		files, err := os.ReadDir(path)
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			ext := strings.ToLower(filepath.Ext(file.Name()))
			if !file.IsDir() && supportedFormats[ext] {
				songs = append(songs, filepath.Join(path, file.Name()))
			}
		}

		if len(songs) == 0 {
			log.Fatal("No MP3 files found in the directory")
		}
	} else {
		ext := strings.ToLower(filepath.Ext(path))
		if !supportedFormats[ext] {
			log.Fatal("Unsupported file format")
		}
		songs = append(songs, path)
	}

	// shuffle
	if *shuffle {
		rand.Shuffle(len(songs), func(i, j int) {
			songs[i], songs[j] = songs[j], songs[i]
		})
	}

	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	fmt.Println("Controls: [P] Pause/Resume | [S] Skip | [Up/Down] Volume | [Esc] Quit")

	sampleRate := beep.SampleRate(44100)
	speaker.Init(sampleRate, sampleRate.N(time.Second/10))

	for _, song := range songs {
		playSong(song, sampleRate)
	}

}

func playSong(fileName string, sampleRate beep.SampleRate) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Printf("Could not open %s: %v", fileName, err)
		return
	}
	defer file.Close()

	metadata, err := tag.ReadFrom(file)
	var displayName string
	if err == nil && metadata.Title() != "" {
		artist := metadata.Artist()
		if artist != "" {
			displayName = fmt.Sprintf("%s -%s", artist, metadata.Title())
		} else {
			displayName = metadata.Title()
		}
	} else {
		displayName = filepath.Base(fileName)
	}
	fmt.Printf("\n🎵 Now playing: %s\n", displayName)

	file.Seek(0, 0)

	var streamer beep.StreamSeekCloser
	var format beep.Format

	ext := strings.ToLower(filepath.Ext(fileName))
	switch ext {
	case ".mp3":
		streamer, format, err = mp3.Decode(file)
		if err != nil {
			log.Println(err)
		}
	case ".wav":
		streamer, format, err = wav.Decode(file)
		if err != nil {
			log.Println(err)
		}
	case ".flac":
		streamer, format, err = flac.Decode(file)
		if err != nil {
			log.Println(err)
		}
	default:
		log.Printf("Unsupported format: %s", ext)
		return
	}
	defer streamer.Close()

	resampled := beep.Resample(4, format.SampleRate, sampleRate, streamer)

	ctrl := &beep.Ctrl{Streamer: resampled, Paused: false}
	volume := &effects.Volume{Streamer: ctrl, Base: 2, Volume: 0}

	done := make(chan bool)
	quit := make(chan bool)

	speaker.Play(beep.Seq(volume, beep.Callback(func() {
		done <- true
	})))

	go func() {
		for {
			select {
			case <-done:
				return
			case <-quit:
				return
			default:
				speaker.Lock()
				position := streamer.Position()
				length := streamer.Len()
				speaker.Unlock()

				pct := float64(position) / float64(length)
				bar := strings.Repeat("█", int(pct*20)) + strings.Repeat("-", 20-int(pct*20))
				fmt.Printf("\rPlaying [%s] %.0f%% | Vol: %.1f", bar, pct*100, volume.Volume)
				time.Sleep(time.Second)
			}
		}
	}()

	type keyEvent struct {
		char rune
		key  keyboard.Key
	}
	keyChan := make(chan keyEvent, 10)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-quit:
				return
			default:
				char, key, err := keyboard.GetKey()
				if err == nil {
					keyChan <- keyEvent{char, key}
				}
			}
		}
	}()

	for {
		select {
		case <-done:
			close(quit)
			fmt.Println("\nFinished song.")
			return

		case evt := <-keyChan:
			if evt.key == keyboard.KeyEsc {
				os.Exit(0)
			}

			if evt.char == 'p' || evt.char == 'P' {
				speaker.Lock()
				ctrl.Paused = !ctrl.Paused
				speaker.Unlock()
			}

			if evt.char == 's' || evt.char == 'S' {
				close(quit)
				speaker.Clear()
				return
			}

			if evt.key == keyboard.KeyArrowUp {
				speaker.Lock()
				volume.Volume += 0.2
				speaker.Unlock()
			}

			if evt.key == keyboard.KeyArrowDown {
				speaker.Lock()
				volume.Volume -= 0.2
				speaker.Unlock()
			}

		case <-time.After(100 * time.Millisecond):
			continue
		}
	}
}
