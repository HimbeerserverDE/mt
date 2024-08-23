package mt

type PointabilityType uint8

const (
	Through PointabilityType = iota
	Pointable
	Blocking
)

type Pointabilities struct {
	// Version.
	//mt:const uint8(0)

	Nodes      map[string]PointabilityType
	NodeGroups map[string]PointabilityType

	Objects      map[string]PointabilityType
	ObjectGroups map[string]PointabilityType
}
