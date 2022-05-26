package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/mjibson/go-dsp/wav"
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

		pos := time.Duration(
			(float64(i*len(f)) / float64(w.SampleRate)) * float64(time.Second),
		)
		fmt.Printf("cycle %d pos %v %v\n", i+1, pos, f)
		// time.Sleep(time.Millisecond * 50)

		// var samples []float64
		// for _, s := range f {
		// 	samples = append(samples, float64(s))
		// 	fmt.Println(s)
		// }

		// window.Apply(samples, window.Hann)

		// sfft := fft.FFTReal(samples)
		// for _, f := range sfft {
		// 	fmt.Println(cmplx.Abs(f))
		// }
	}
}
