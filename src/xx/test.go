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
	"logger"
	"strconv"
	// "testing"
	"time"
)
var palette=[]color.Color{color.White,color.Black}
const(
	whiteIndex=0
	blackIndex=1
)
func init() {
	logger.SetConsole(true)
	//指定日志文件备份方式为文件大小的方式
	//第一个参数为日志文件存放目录
	//第二个参数为日志文件命名
	//第三个参数为备份文件最大数量
	//第四个参数为备份文件大小
	//第五个参数为文件大小的单位
	logger.SetRollingFile("log", "test.log", 10, 5, logger.KB)

	//指定日志文件备份方式为日期的方式
	//第一个参数为日志文件存放目录
	//第二个参数为日志文件命名
	// logger.SetRollingDaily("d:/logtest", "test.log")

	//指定日志级别  ALL，DEBUG，INFO，WARN，ERROR，FATAL，OFF 级别由低到高
	//一般习惯是测试阶段为debug，生成环境为info以上
	logger.SetLevel(logger.ERROR)
}
func main() {
	// lissajous(os.Stdout)
	test_log()
}
func test_log() {
	
	for i := 10000; i > 0; i-- {
		go func() {
			logger.Debug("Debug>>>>>>>>>>>>>>>>>>>>>>" + strconv.Itoa(i))
			logger.Info("Info>>>>>>>>>>>>>>>>>>>>>>>>>" + strconv.Itoa(i))
			logger.Warn("Warn>>>>>>>>>>>>>>>>>>>>>>>>>" + strconv.Itoa(i))
			logger.Error("Error>>>>>>>>>>>>>>>>>>>>>>>>>" + strconv.Itoa(i))
			logger.Fatal("Fatal>>>>>>>>>>>>>>>>>>>>>>>>>" + strconv.Itoa(i))
		}()
		time.Sleep(100 * time.Millisecond)
	}
	time.Sleep(15 * time.Second)
	
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