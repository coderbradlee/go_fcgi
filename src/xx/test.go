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
	// stdLog "log"
	// "path/filepath"
	// "runtime"
	// "strconv"
	// "strings"
	// "testing"
	// "time"
	"glog"
)
var palette=[]color.Color{color.White,color.Black}
const(
	whiteIndex=0
	blackIndex=1
)
func init() {
	glog.CopyStandardLogTo("INFO")
}
func main() {
	// lissajous(os.Stdout)
	test_log()
	glog.Info("Prepare to repel boarders")
	
	glog.Fatalf("Initialization failed: err")
}
func test_log() {
	logging.toStderr = false
	defer logging.swap(logging.newBuffers())
	stdLog.Print("test")
	if !contains(infoLog, "I", t) {
		t.Errorf("Info has wrong character: %q", contents(infoLog))
	}
	if !contains(infoLog, "test", t) {
		t.Error("Info failed")
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