package images

import (
	"log"
	"os"
	"github.com/davidbyttow/govips/v2/vips"
)
	
func ResizeImage(path string, target int) *vips.ImageRef {
	im, err := vips.NewImageFromFile(path)
	if err != nil {
		log.Fatal("Could not load image")
	}

	scale := float64(target) / float64(im.PageHeight())

	err = im.Resize(scale, vips.KernelLanczos3)
	if err != nil {
		log.Fatal("failed to resize")
	}

	return im
}

func SaveImage(im *vips.ImageRef, path string) {
	dat, _,  err := im.ExportJpeg(&vips.JpegExportParams{Quality: 99, SubsampleMode: vips.VipsForeignSubsampleOff})
	if err != nil {
		log.Fatal("failed to encode ", err)
	}

	err = os.WriteFile(path, dat, 0644)
	if err != nil {
		log.Fatal("failed to write file ", err)
	}
}