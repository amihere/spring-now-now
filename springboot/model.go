package springboot

import (
	"fmt"
	"io"
	"net/http"
)

// these are models that store the response from the spring initializr api

// SpringBoot
type SpringBoot struct {
	GroupID      ArtifactID             `json:"groupId"`
	PackageName  ArtifactID             `json:"packageName"`
	Description  ArtifactID             `json:"description"`
	Name         ArtifactID             `json:"name"`
	Version      ArtifactID             `json:"version"`
	ArtifactID   ArtifactID             `json:"artifactId"`
	Links        SpringBootLinks        `json:"_links"`
	Packaging    BootVersion            `json:"packaging"`
	BootVersion  BootVersion            `json:"bootVersion"`
	Language     BootVersion            `json:"language"`
	JavaVersion  BootVersion            `json:"javaVersion"`
	Type         Type                   `json:"type"`
	Dependencies SpringBootDependencies `json:"dependencies"`
}

// ArtifactID
type ArtifactID struct {
	Type    string `json:"type"`
	Default string `json:"default"`
}

// BootVersion
type BootVersion struct {
	Type    string             `json:"type"`
	Default string             `json:"default"`
	Values  []BootVersionValue `json:"values"`
}

// BootVersionValue
type BootVersionValue struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// SpringBootDependencies
type SpringBootDependencies struct {
	Type   string              `json:"type"`
	Values []DependenciesValue `json:"values"`
}

// DependenciesValue
type DependenciesValue struct {
	Name   string       `json:"name"`
	Values []ValueValue `json:"values"`
}

// ValueValue
type ValueValue struct {
	VersionRange *string     `json:"versionRange,omitempty"`
	Links        *ValueLinks `json:"_links,omitempty"`
	ID           string      `json:"id"`
	Name         string      `json:"name"`
	Description  string      `json:"description"`
}

// ReferenceUnion
type ReferenceUnion struct {
	ReferenceClass *ReferenceClass
	HomeArray      []Home
}

// ValueLinks
type ValueLinks struct {
	Reference *ReferenceUnion `json:"reference"`
	Guide     *Guide          `json:"guide"`
	Home      *Home           `json:"home,omitempty"`
	Sample    *Home           `json:"sample,omitempty"`
}

// Home
type Home struct {
	Title *string `json:"title,omitempty"`
	Href  string  `json:"href"`
}

// ReferenceClass
type ReferenceClass struct {
	Templated *bool   `json:"templated,omitempty"`
	Title     *string `json:"title,omitempty"`
	Href      string  `json:"href"`
}

// SpringBootLinks
type SpringBootLinks struct {
	GradleProject       GradleBuildClass `json:"gradle-project"`
	GradleProjectKotlin GradleBuildClass `json:"gradle-project-kotlin"`
	GradleBuild         GradleBuildClass `json:"gradle-build"`
	MavenProject        GradleBuildClass `json:"maven-project"`
	MavenBuild          GradleBuildClass `json:"maven-build"`
	Dependencies        GradleBuildClass `json:"dependencies"`
}

// GradleBuildClass
type GradleBuildClass struct {
	Href      string `json:"href"`
	Templated bool   `json:"templated"`
}

// Type
type Type struct {
	Type    string      `json:"type"`
	Default string      `json:"default"`
	Values  []TypeValue `json:"values"`
}

// TypeValue
type TypeValue struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Action      string `json:"action"`
	Tags        Tags   `json:"tags"`
}

// Tags
type Tags struct {
	Build   string  `json:"build"`
	Dialect *string `json:"dialect,omitempty"`
	Format  string  `json:"format"`
}

// Guide
type Guide struct {
	Home      *Home
	HomeArray []Home
}

// New will fetch all the spring metadata in one request!
func New() (SpringBoot, error) {
	client := http.Client{Timeout: timeout}
	req, err := http.NewRequest("GET", initializrURL, nil)
	req.Header.Set("Accept", "application/json")
	if err != nil {
		return SpringBoot{}, fmt.Errorf("error creating spring boot request: %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return SpringBoot{}, fmt.Errorf("error sending spring boot request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return SpringBoot{}, fmt.Errorf("error reading spring boot response: %v", err)
	}

	springBoot, err := UnmarshalSpringBoot(body)
	if err != nil {
		return SpringBoot{}, fmt.Errorf("error unmarshalling spring boot response: %v", err)
	}

	return springBoot, nil
}
