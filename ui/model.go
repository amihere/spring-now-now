package ui

import (
	"github.com/amihere/spring-now-now/springboot"
	springlist "github.com/amihere/spring-now-now/springlist"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
)

type state int

const (
	chooseProjectType state = iota
	chooseBootVersion
	chooseLanguage
	inputMetaData
	choosePackaging
	chooseJavaVersion
	chooseDependencies
	downloadFile
)

const (
	groupID = iota
	artifactID
	applicationName
	description
	packageName
	version
)

// Model
type Model struct {
	list         list.Model
	Type         string
	BootVersion  string
	Language     string
	GroupID      string
	ArtifactID   string
	Name         string
	Description  string
	PackageName  string
	Version      string
	Packaging    string
	JavaVersion  string
	message      string
	Dependencies []string
	inputs       []textinput.Model
	springBoot   springboot.SpringBoot
	focused      int
	state        state
	failed       bool
	quitting     bool
}

// NewModel Gets Called first in the Main Thread
func NewModel() *Model {
	sp, err := springboot.New()
	if err != nil {
		panic(err)
	}

	// First menu item shown
	l := springlist.NewNormalListModel(
		"Project Type",
		getProjectTypes(sp),
		sp.Type.Default,
		0, 0,
	)

	inputs := make([]textinput.Model, 6)

	inputs[groupID] = textinput.New()
	inputs[groupID].Placeholder = sp.GroupID.Default
	inputs[groupID].SetValue(sp.GroupID.Default)
	inputs[groupID].Focus()
	inputs[groupID].Width = 60
	inputs[groupID].Prompt = ""

	inputs[artifactID] = textinput.New()
	inputs[artifactID].Placeholder = sp.ArtifactID.Default
	inputs[artifactID].SetValue(sp.ArtifactID.Default)
	inputs[artifactID].Focus()
	inputs[artifactID].Width = 60
	inputs[artifactID].Prompt = ""

	inputs[applicationName] = textinput.New()
	inputs[applicationName].Placeholder = sp.Name.Default
	inputs[applicationName].SetValue(sp.Name.Default)
	inputs[applicationName].Focus()
	inputs[applicationName].Width = 60
	inputs[applicationName].Prompt = ""

	inputs[description] = textinput.New()
	inputs[description].Placeholder = sp.Description.Default
	inputs[description].SetValue(sp.Description.Default)
	inputs[description].Focus()
	inputs[description].Width = 60
	inputs[description].Prompt = ""

	inputs[packageName] = textinput.New()
	inputs[packageName].Placeholder = sp.PackageName.Default
	inputs[packageName].SetValue(sp.PackageName.Default)
	inputs[packageName].Focus()
	inputs[packageName].Width = 60
	inputs[packageName].Prompt = ""

	inputs[version] = textinput.New()
	inputs[version].Placeholder = sp.Version.Default
	inputs[version].SetValue(sp.Version.Default)
	inputs[version].Focus()
	inputs[version].Width = 60
	inputs[version].Prompt = ""

	return &Model{
		Packaging:    sp.Packaging.Default,
		JavaVersion:  sp.JavaVersion.Default,
		Language:     sp.Language.Default,
		BootVersion:  sp.BootVersion.Default,
		GroupID:      sp.GroupID.Default,
		ArtifactID:   sp.ArtifactID.Default,
		Name:         sp.Name.Default,
		Description:  sp.Description.Default,
		PackageName:  sp.PackageName.Default,
		Version:      sp.Version.Default,
		Type:         sp.Type.Default,
		Dependencies: []string{},
		springBoot:   sp,
		state:        chooseProjectType,
		list:         l,
		inputs:       inputs,
		focused:      0,
		failed:       false,
		quitting:     false,
	}
}
