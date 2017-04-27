package appui

import (
	"github.com/moncho/dry/docker"
	"github.com/moncho/dry/ui"
)

const (
	//DownArrow character
	DownArrow = string('\U00002193')
	//RightArrow character
	RightArrow = string('\U00002192')

	//MainScreenHeaderSize is the number of lines the header of the main screen uses
	MainScreenHeaderSize = 5
	//MainScreenFooterSize is the number of lines the footer of the main screen uses
	MainScreenFooterSize = 1

	imageTableStartPos     = MainScreenHeaderSize + 5 //5 its the number of lines in the image table header
	containerTableStartPos = MainScreenHeaderSize + 5
	networkTableStartPos   = MainScreenHeaderSize + 5
	//DefaultColumnSpacing defines the minimun space between columns in pixels
	DefaultColumnSpacing = 1
	//IDColumnWidth defines a fixed width for ID columns
	IDColumnWidth = docker.ShortLen
)

//CalcItemWidth calculates the width of each item for the given total width and item count
func CalcItemWidth(width, items int) int {
	spacing := DefaultColumnSpacing * items
	return (width - spacing) / items
}

//MainScreenAvailableHeight returns how many lines in the main screen are available for rendering
func MainScreenAvailableHeight() int {
	return ui.ActiveScreen.Dimensions.Height - MainScreenHeaderSize - MainScreenFooterSize - 5
}
