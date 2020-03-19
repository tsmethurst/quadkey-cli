package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"gopkg.in/urfave/cli.v1"
)

const (
	earthRadius  = 6378137
	minLatitude  = -85.05112878
	maxLatitude  = 85.05112878
	minLongitude = -180
	maxLongitude = 180
)

func encode(c *cli.Context) error {
	detail := c.Int("level")
	coordsStr := c.String("coords")
	if coordsStr == "" {
		return cli.NewExitError("coords undefined", 1)
	}
	coordsSpl := strings.Split(coordsStr, ",")
	if len(coordsSpl) != 2 {
		return cli.NewExitError("malformed coords", 1)
	}
	lat, err := strconv.ParseFloat(coordsSpl[0], 64)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	long, err := strconv.ParseFloat(coordsSpl[1], 64)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	x, y := latLongToPixelXY(lat, long, detail)
	tileX, tileY := pixelXYToTileXY(x, y)
	quadkey := tileXYToQuadKey(tileX, tileY, detail)
	fmt.Println(quadkey)
	return nil
}

func clip(n, minValue, maxValue float64) float64 {
	return math.Min(math.Max(n, minValue), maxValue)
}

func mapSize(levelOfDetail int) uint {
	return uint(256 << levelOfDetail)
}

func latLongToPixelXY(latitude, longitude float64, levelOfDetail int) (pixelX, pixelY int) {
	latitude = clip(latitude, minLatitude, maxLatitude)
	longitude = clip(longitude, minLongitude, maxLongitude)
	x := (longitude + 180) / 360
	sinLatitude := math.Sin(latitude * math.Pi / 180)
	y := 0.5 - math.Log((1+sinLatitude)/(1-sinLatitude))/(4*math.Pi)
	mapSize := mapSize(levelOfDetail)
	pixelX = int(clip(x*float64(mapSize)+0.5, 0, float64(mapSize)-1))
	pixelY = int(clip(y*float64(mapSize)+0.5, 0, float64(mapSize)-1))
	return pixelX, pixelY
}

func pixelXYToTileXY(pixelX, pixelY int) (tileX, tileY int) {
	return pixelX / 256, pixelY / 256
}

func tileXYToQuadKey(tileX, tileY, levelOfDetail int) string {
	quadKey := ""
	for i := levelOfDetail; i > 0; i-- {
		digit := '0'
		mask := 1 << (i - 1)
		if (tileX & mask) != 0 {
			digit++
		}
		if (tileY & mask) != 0 {
			digit++
			digit++
		}
		quadKey += string(digit)
	}
	return quadKey
}
