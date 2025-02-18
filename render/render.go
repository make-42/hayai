package render

import (
	"fmt"
	"hayai/config"
	"hayai/jmaeew"
	"hayai/utils"
	"hayai/waves"
	"image/color"
	"log"
	"math"
	"os"
	"time"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"github.com/mmcloughlin/globe"
	"github.com/pymaxion/geographiclib-go/geodesic"
	"golang.org/x/image/font"

	"github.com/hajimehoshi/ebiten/v2"

	_ "embed"
)

var gameData *Game

type Game struct {
	Event      jmaeew.JMAEEW
	ImpactTime float64
}

var Displaying bool

func (g *Game) Update() error {
	if ebiten.IsWindowBeingClosed() {
		os.Exit(0)
	}
	return nil // Add kill after timer TODO
}

func ComputeLoveImpactTime(event jmaeew.JMAEEW) float64 {
	wgs84 := geodesic.WGS84
	line := wgs84.InverseLine(event.Latitude, event.Longitude, config.Config.Latitude, config.Config.Longitude)
	return waves.PredictLoveWaveImpactTime(line.Arc())
}

func (g *Game) Draw(screen *ebiten.Image) {
	jst := time.FixedZone("JST", 9*3600)
	// fake time here for testing
	// g.Event.OriginTime = "2025/02/18 22:45:30"
	originalTimeParsed, _ := time.ParseInLocation("2006/01/02 15:04:05", g.Event.OriginTime, jst)

	currentTime := float64(time.Now().Sub(originalTimeParsed).Nanoseconds()) / 1e9
	gl := globe.New()
	gl.DrawGraticule(10.0)
	gl.DrawLandBoundaries()
	green := color.NRGBA{0x00, 0x64, 0x3c, 192}
	red := color.NRGBA{0xff, 0x42, 0x3c, 192}
	gl.DrawDot(config.Config.Latitude, config.Config.Longitude, 0.05, globe.Color(green))
	gl.DrawDot(g.Event.Latitude, g.Event.Longitude, 0.05, globe.Color(red))

	// Draw waves
	for _, wave := range waves.WaveList {
		travel, travelA, travelB, double, exists := waves.TimeToTravel(*wave, currentTime)
		if exists {
			if double {
				DrawCircle(gl, g.Event.Latitude, g.Event.Longitude, travelA, globe.Color(wave.Color))
				DrawCircle(gl, g.Event.Latitude, g.Event.Longitude, travelB, globe.Color(wave.Color))
			} else {
				DrawCircle(gl, g.Event.Latitude, g.Event.Longitude, travel, globe.Color(wave.Color))
			}
		}
	}
	gl.CenterOn(g.Event.Latitude, g.Event.Longitude)
	renderedImage := gl.Image(config.Config.RealtimeVisRenderSize)
	inverted := imaging.Invert(renderedImage)
	ctx := gg.NewContextForImage(inverted)
	ctx.SetRGBA(1, 0, 0, 1)
	ctx.SetFontFace(FontFace)
	ctx.DrawString(fmt.Sprintf("Detected a M%0.1f earthquake in %s", g.Event.Magunitude, g.Event.Title), config.Config.RealtimeVisTextPadding, config.Config.RealtimeVisTextPadding+config.Config.RealtimeVisFontSize)
	ctx.DrawString(fmt.Sprintf("%0.4f, %0.4f @ %d km", g.Event.Latitude, g.Event.Longitude, g.Event.Depth), config.Config.RealtimeVisTextPadding, config.Config.RealtimeVisTextPadding+2*config.Config.RealtimeVisFontSize)
	ctx.DrawString(fmt.Sprintf("Love waves impact in %0.1fs", g.ImpactTime-currentTime), config.Config.RealtimeVisTextPadding, config.Config.RealtimeVisTextPadding+3*config.Config.RealtimeVisFontSize)

	for i, wave := range waves.WaveList {
		travel, travelA, travelB, double, exists := waves.TimeToTravel(*wave, currentTime)
		if exists {
			ctx.SetRGBA(float64(wave.Color.R)/255, float64(wave.Color.G)/255, float64(wave.Color.B)/255, 1)
			if double {
				angle := 2 * math.Pi * float64(i) / float64(len(waves.WaveList))
				x := math.Cos(-angle)
				y := math.Sin(-angle)
				if travelA < 90 {
					ctx.DrawStringAnchored(wave.Label, float64(config.Config.RealtimeVisRenderSize)/2+x*0.99*float64(config.Config.RealtimeVisRenderSize)/2*math.Sin(travelA/180*math.Pi), float64(config.Config.RealtimeVisRenderSize)/2+y*0.99*float64(config.Config.RealtimeVisRenderSize)/2*math.Sin(travel/180*math.Pi), x, y)
				}
				if travelB < 90 {
					ctx.DrawStringAnchored(wave.Label, float64(config.Config.RealtimeVisRenderSize)/2-x*0.99*float64(config.Config.RealtimeVisRenderSize)/2*math.Sin(travelB/180*math.Pi), float64(config.Config.RealtimeVisRenderSize)/2-y*0.99*float64(config.Config.RealtimeVisRenderSize)/2*math.Sin(travel/180*math.Pi), -x, -y)
				}

			} else {
				if travel < 90 {
					angle := 2 * math.Pi * float64(i) / float64(len(waves.WaveList))
					x := math.Cos(-angle)
					y := math.Sin(-angle)
					ctx.DrawStringAnchored(wave.Label, float64(config.Config.RealtimeVisRenderSize)/2+x*0.99*float64(config.Config.RealtimeVisRenderSize)/2*math.Sin(travel/180*math.Pi), float64(config.Config.RealtimeVisRenderSize)/2+y*0.99*float64(config.Config.RealtimeVisRenderSize)/2*math.Sin(travel/180*math.Pi), x, y)
				}
			}
			ctx.Fill()
		}
	}
	screen.DrawImage(ebiten.NewImageFromImage(ctx.Image()), &ebiten.DrawImageOptions{})
}

func DrawCircle(g *globe.Globe, lat, lon float64, radius float64, color globe.Option) {
	//fmt.Println(radius)
	wgs84 := geodesic.WGS84
	pointLats := []float64{}
	pointLons := []float64{}
	for i := 0; i < config.Config.RealtimeVisCircleResolution; i++ {
		gData := wgs84.Direct(lat, lon, float64(i)/float64(config.Config.RealtimeVisCircleResolution)*360., radius/360*math.Pi*wgs84.EquatorialRadius())
		pointLats = append(pointLats, gData.Lat2)
		pointLons = append(pointLons, gData.Lon2)
	}
	for i := 0; i < config.Config.RealtimeVisCircleResolution; i++ {
		if i == 0 {
			g.DrawLine(pointLats[config.Config.RealtimeVisCircleResolution-1], pointLons[config.Config.RealtimeVisCircleResolution-1], pointLats[0], pointLons[0], color)
		} else {
			g.DrawLine(pointLats[i-1], pointLons[i-1], pointLats[i], pointLons[i], color)
			//fmt.Println(pointLats[i-1], pointLons[i-1])
		}
	}
}

//go:embed assets/NotoSansJP-VariableFont_wght.ttf
var fontBytes []byte
var FontFace font.Face

func Init() {
	font, err := truetype.Parse(fontBytes)
	utils.CheckError(err)
	FontFace = truetype.NewFace(font, &truetype.Options{Size: config.Config.RealtimeVisFontSize})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return config.Config.RealtimeVisRenderSize, config.Config.RealtimeVisRenderSize
}

func Render(event jmaeew.JMAEEW) {
	if config.Config.TestWarning {
		t := time.Now()
		_, offset := t.Zone()
		event.OriginTime = time.Now().Add(time.Second * time.Duration(9*3600-offset)).Format("2006/01/02 15:04:05")
	}
	gameData = &Game{event, ComputeLoveImpactTime(event)}
	if !Displaying {
		Displaying = true
		//ebiten.SetWindowIcon([]image.Image{icons.WindowIcon48, icons.WindowIcon32, icons.WindowIcon16})
		ebiten.SetWindowSize(config.Config.RealtimeVisRenderSize, config.Config.RealtimeVisRenderSize)
		ebiten.SetWindowTitle("hayai")
		ebiten.SetWindowMousePassthrough(true)
		ebiten.SetTPS(config.Config.RealtimeVisFPS) // target FPS todo
		ebiten.SetWindowDecorated(true)
		screenW, screenH := ebiten.Monitor().Size()
		ebiten.SetWindowPosition(screenW/2-config.Config.RealtimeVisRenderSize/2, screenH/2-config.Config.RealtimeVisRenderSize/2)
		ebiten.SetVsyncEnabled(true)
		ebiten.SetWindowClosingHandled(true)
		if err := ebiten.RunGameWithOptions(gameData, &ebiten.RunGameOptions{}); err != nil {
			log.Fatal(err)
		}
	}
}
