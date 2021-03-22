package excepts

import "fmt"

type InvalidPointFormat struct {
	File string
}
type InvalidPointCo struct {
	File string
}
type PointNotFound struct {
	File string
	Err  string
}

func (e *PointNotFound) Error() string {
	return fmt.Sprintf("Point file not found: %s (%s)", e.Err, e.File)
}
func (e *InvalidPointFormat) Error() string {
	return fmt.Sprintf("Point is in invalid format: %s (Must be X,Y format)", e.File)
}
func (e *InvalidPointCo) Error() string {
	return fmt.Sprintf("Point has invalid axis value: %s (0 & upper)", e.File)
}
