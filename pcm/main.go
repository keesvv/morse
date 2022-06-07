package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"math/cmplx"
	"os"
	"time"

	"github.com/mjibson/go-dsp/fft"
	"github.com/mjibson/go-dsp/wav"
	"github.com/mjibson/go-dsp/window"
)

func main() {
	freq := flag.Float64("f", 1000, "frequency (Hz)")
	flag.Parse()

	w, _ := wav.New(os.Stdin)
	nSamples := int(float64(w.SampleRate) / *freq)

	for i := 0; ; i++ {
		f, err := w.ReadFloats(nSamples)

		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		var samples []float64
		for _, s := range f {
			samples = append(samples, float64(s))
		}

		window.Apply(samples, window.Hann)

		sfft := fft.FFTReal(samples)
		var fSfft []float64
		for _, sf := range sfft {
			fSfft = append(fSfft, cmplx.Abs(sf))
		}

		pos := time.Duration(
			(float64(i*len(samples)) / float64(w.SampleRate)) * float64(time.Second),
		)

		fmt.Printf("cycle %d pos %v samples %v fft %v\n", i+1, pos, samples, fSfft)
		// time.Sleep(time.Millisecond * 50)
	}
}
