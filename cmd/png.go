package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/fredericlemoine/gotree/draw"
	"github.com/spf13/cobra"
)

var pngwidth int
var pngheight int
var pngradial bool

// pngCmd represents the png command
var pngCmd = &cobra.Command{
	Use:   "png",
	Short: "Draw trees in png files",
	Long:  `Draw trees in png files.`,
	Run: func(cmd *cobra.Command, args []string) {
		var d draw.TreeDrawer
		var l draw.TreeLayout
		ntree := 0
		for tr := range readTrees(intreefile) {
			fname := outtreefile
			if ntree > 0 {
				extension := filepath.Ext(fname)
				if extension == ".png" {
					fname = fname[0 : len(fname)-len(extension)]
				}
				fname = fmt.Sprintf(fname+"_%03d.png", ntree)
			}
			f := openWriteFile(fname)
			d = draw.NewPngTreeDrawer(f, pngwidth, pngheight, 30, 30, 30, 30)
			if pngradial {
				l = draw.NewRadialLayout(d, true)
			} else {
				l = draw.NewNormalLayout(d)
			}
			l.DrawTree(tr.Tree)
			f.Close()
			ntree++
		}
	},
}

func init() {
	drawCmd.AddCommand(pngCmd)
	pngCmd.PersistentFlags().IntVarP(&pngwidth, "width", "w", 200, "Width of png image in pixels")
	pngCmd.PersistentFlags().IntVarP(&pngheight, "height", "H", 200, "Height of png image in pixels")
	pngCmd.PersistentFlags().BoolVarP(&pngradial, "radial", "r", false, "Radial layout (default : normal)")
}