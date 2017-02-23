package main
import(
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	// "os"
)
import (
	// "bytes"
	// "fmt"
	stdLog "log"
	// "path/filepath"
	// "runtime"
	// "strconv"
	"strings"
	// "testing"
	"time"
)
var palette=[]color.Color{color.White,color.Black}
const(
	whiteIndex=0
	blackIndex=1
)
func init() {
	CopyStandardLogTo("INFO")
}
func main() {
	// lissajous(os.Stdout)
	test_log()
	glog.Info("Prepare to repel boarders")
	
	glog.Fatalf("Initialization failed: err")
}
func test_log() {
	setFlags()
	var err error
	defer func(previous func(error)) { logExitFunc = previous }(logExitFunc)
	logExitFunc = func(e error) {
		err = e
	}
	defer func(previous uint64) { MaxSize = previous }(MaxSize)
	MaxSize = 512

	Info("x") // Be sure we have a file.
	info, ok := logging.file[infoLog].(*syncBuffer)
	if !ok {
		t.Fatal("info wasn't created")
	}
	if err != nil {
		t.Fatalf("info has initial error: %v", err)
	}
	fname0 := info.file.Name()
	Info(strings.Repeat("x", int(MaxSize))) // force a rollover
	if err != nil {
		t.Fatalf("info has error after big write: %v", err)
	}

	// Make sure the next log file gets a file name with a different
	// time stamp.
	//
	// TODO: determine whether we need to support subsecond log
	// rotation.  C++ does not appear to handle this case (nor does it
	// handle Daylight Savings Time properly).
	time.Sleep(1 * time.Second)

	Info("x") // create a new file
	if err != nil {
		t.Fatalf("error after rotation: %v", err)
	}
	fname1 := info.file.Name()
	if fname0 == fname1 {
		t.Errorf("info.f.Name did not change: %v", fname0)
	}
	if info.nbytes >= MaxSize {
		t.Errorf("file size was not reset: %d", info.nbytes)
	}
}
func lissajous(out io.Writer) {
	const(
		cycles=5
		res=0.001
		size=100
		nframes=64
		delay=10
	)
	freq:=rand.Float64()*3.0
	anim:=gif.GIF{LoopCount: nframes}
	phase:=0.0
	for i:=0;i<nframes;i++{
		rect:=image.Rect(0,0,2*size+1,2*size+1)
		// image.NewPalletted(rect,palette)
		img:=image.NewPaletted(rect, palette)
		for t:=0.0;t<cycles*2*math.Pi;t+=res{
			x:=math.Sin(t)
			y:=math.Sin(t*freq+phase)
			img.SetColorIndex(size+int(x*size+0.5),size+int(y*size+0.5),blackIndex)
		}
		phase+=0.1
		anim.Delay=append(anim.Delay,delay)
		anim.Image=append(anim.Image,img)
	}
	gif.EncodeAll(out,&anim)
}