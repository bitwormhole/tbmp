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

//go:embed "src/client/resources"
var theClientModuleResFS embed.FS

//go:embed "src/main/resources"
var theMainModuleResFS embed.FS

//go:embed "src/server/resources"
var theServerModuleResFS embed.FS

//go:embed "src/test/resources"
var theTestModuleResFS embed.FS

const (
	theClientModuleResPath = "src/client/resources"
	theMainModuleResPath   = "src/main/resources"
	theServerModuleResPath = "src/server/resources"
	theTestModuleResPath   = "src/test/resources"
)

////////////////////////////////////////////////////////////////////////////////

// NewClientModule make module (client)
func NewClientModule() *application.ModuleBuilder {
	mb := new(application.ModuleBuilder)
	mb.Name(theModuleName + "#client")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.EmbedResources(theClientModuleResFS, theClientModuleResPath)
	return mb
}

// NewMainModule make module (main)
func NewMainModule() *application.ModuleBuilder {
	mb := new(application.ModuleBuilder)
	mb.Name(theModuleName)
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.EmbedResources(theMainModuleResFS, theMainModuleResPath)
	return mb
}

// NewServerModule make module (server)
func NewServerModule() *application.ModuleBuilder {
	mb := new(application.ModuleBuilder)
	mb.Name(theModuleName + "#server")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.EmbedResources(theServerModuleResFS, theServerModuleResPath)
	return mb
}

// NewTestModule make module (test)
func NewTestModule() *application.ModuleBuilder {
	mb := new(application.ModuleBuilder)
	mb.Name(theModuleName + "#test")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.EmbedResources(theTestModuleResFS, theTestModuleResPath)
	return mb
}

////////////////////////////////////////////////////////////////////////////////
