package mt

type AttractionKind uint8

const (
	NoAttraction    AttractionKind = iota // None
	PointAttraction                       // Point
	LineAttraction                        // Line
	PlaneAttraction                       // Plane
)

//go:generate stringer -linecomment -type AttractionKind

type ParticleSpawnerFlags uint8

const (
	AttractorKill ParticleSpawnerFlags = 1 << iota // Particle dies on contact
)

//go:generate stringer -linecomment -type ParticleSpawnerFlags

type ParticleTextureFlags uint8

const (
	Animated ParticleTextureFlags = 1 << iota // Animated
)

//go:generate stringer -linecomment -type ParticleTextureFlags

type ParticleTexture struct {
	Flags ParticleTextureFlags
	Alpha TweenF32
	Scale TweenV2F32

	//mt:if %s.Flags & Animated == Animated
	Animation TileAnim
	//mt:end
}

type ParticleParams struct {
	Pos, Vel, Acc  [3]float32
	ExpirationTime float32 // in seconds.
	Size           float32
	Collide        bool

	//mt:len32
	TextureName Texture

	Vertical    bool
	CollisionRm bool
	AnimParams  TileAnim
	Glow        uint8
	AOCollision bool
	NodeParam0  Content
	NodeParam2  uint8
	NodeTile    uint8
	Drag        [3]float32
	Jitter      RangeV3F32
	Bounce      RangeF32
	Texture     ParticleTexture
}
