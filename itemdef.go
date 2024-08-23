package mt

import "image/color"

type ItemType uint8

const (
	_ ItemType = iota
	NodeItem
	CraftItem
	ToolItem
)

//go:generate stringer -type ItemType

// An ItemDef defines the properties of an item.
type ItemDef struct {
	//mt:lenhdr 16

	// Version.
	//mt:const uint8(6)

	Type ItemType

	Name, Desc string

	InvImg, WieldImg Texture
	WieldScale       [3]float32

	StackMax uint16

	Usable          bool
	CanPointLiquids bool

	ToolCaps ToolCaps

	Groups []Group

	PlacePredict string

	PlaceSnd, PlaceFailSnd SoundDef

	PointRange float32

	// Set index in Palette with "palette_index" item meta field,
	// this overrides Color.
	Palette Texture
	Color   color.NRGBA

	// Texture overlays.
	InvOverlay, WieldOverlay Texture

	ShortDesc string

	SoundUse    SoundDef
	SoundUseAir SoundDef

	HasPlaceParam2 bool

	//mt:if %s.HasPlaceParam2
	PlaceParam2 uint8
	//mt:end

	WallmountedRotateVertical bool
	TouchInteraction          TouchInteraction

	Pointabilities string

	HasWearBarParams bool

	//mt:if %s.HasWearBarParams
	WearBarParams WearBarParams
	//mt:end

	//mt:end
}

type TouchInteractionMode uint8

const (
	UserDefined TouchInteractionMode = iota
)

// A TouchInteraction defines how touchscreen taps should be handled
// depending on what the player is pointing at.
type TouchInteraction struct {
	Nothing, Node, Object TouchInteractionMode
}

type BlendMode uint8

const (
	Constant BlendMode = iota
	Linear
)

// A WearBarParams specifies the wear bar colors for certain wear levels.
type WearBarParams struct {
	// Version.
	//mt:const uint8(1)

	Blend BlendMode

	ColorStops map[float32]color.NRGBA
}
