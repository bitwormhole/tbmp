package tbmp

import (
	"embed"

	"github.com/starter-go/application"
)

const (
	theModuleName     = "github.com/bitwormhole/tbmp"
	theModuleVersion  = "v0.0.1"
	theModuleRevision = 1
)

////////////////////////////////////////////////////////////////////////////////

//go:embed "src/main/resources"
var theMainModuleResFS embed.FS

//go:embed "src/test/resources"
var theTestModuleResFS embed.FS

const (
	theMainModuleResPath = "src/main/resources"
	theTestModuleResPath = "src/test/resources"
)

////////////////////////////////////////////////////////////////////////////////

// NewMainModule make module (main)
func NewMainModule() *application.ModuleBuilder {
	mb := new(application.ModuleBuilder)
	mb.Name(theModuleName)
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.EmbedResources(theMainModuleResFS, theMainModuleResPath)
	return mb
}

////////////////////////////////////////////////////////////////////////////////
