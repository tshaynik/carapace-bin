package mvn

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionGoalsAndPhases completes goals and phases
//
//	clean (remove all files generated by the previous build)
//	compile (compile the source code of the project)
func ActionGoalsAndPhases(file string) carapace.Action {
	// TODO caching
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if loadedProject, err := loadProject(file); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			goals := make(map[string]string)
			for _, projectPlugin := range loadedProject.Plugins {
				if loadedPlugin := loadPlugin(projectPlugin.Location(repositoryLocation())); loadedPlugin != nil {
					for key, value := range loadedPlugin.FormattedGoals() {
						goals[key] = value
					}
				}
			}

			for key, value := range defaultGoalsAndPhases() {
				goals[key] = value
			}

			vals := make([]string, 0)
			for key, value := range goals {
				vals = append(vals, key, value)
			}

			return carapace.ActionValuesDescribed(vals...)
		}
	})
}

func defaultGoalsAndPhases() (goals map[string]string) {
	goals = map[string]string{
		// clean lifecycle
		"pre-clean":  "execute processes needed prior to the actual project cleaning",
		"clean":      "remove all files generated by the previous build",
		"post-clean": "execute processes needed to finalize the project cleanin",

		// default lifecycle
		"validate":                "validate the project is correct and all necessary information is available",
		"initialize":              "initialize build state, e.g. set properties or create directories",
		"generate-sources":        "generate any source code for inclusion in compilation",
		"process-sources":         "process the source code, for example to filter any values",
		"generate-resources":      "generate resources for inclusion in the package",
		"process-resources":       "copy and process the resources into the destination directory, ready for packaging",
		"compile":                 "compile the source code of the project",
		"process-classes":         "post-process the generated files from compilation",
		"generate-test-sources":   "generate any test source code for inclusion in compilation",
		"process-test-sources":    "process the test source code, for example to filter any values",
		"generate-test-resources": "create resources for testing",
		"process-test-resources":  "copy and process the resources into the test destination directory",
		"test-compile":            "compile the test source code into the test destination directory",
		"process-test-classes":    "post-process the generated files from test compilation",
		"test":                    "run tests using a suitable unit testing framework",
		"prepare-package":         "perform any operations necessary to prepare a package before the actual packaging",
		"package":                 "take the compiled code and package it in its distributable format, such as a JAR",
		"pre-integration-test":    "perform actions required before integration tests are executed",
		"integration-test":        "process and deploy the package into an environment where integration tests can be run",
		"post-integration-test":   "perform actions required after integration tests have been executed",
		"verify":                  "run any checks to verify the package is valid and meets quality criteria",
		"install":                 "install the package into the local repository",
		"deploy":                  "copies the final package to the remote repository",

		// site lifecycle
		"pre-site":    "execute processes needed prior to the actual project site generation",
		"site":        "generate the project's site documentation",
		"post-site":   "execute processes needed to finalize the site generation",
		"site-deploy": "deploy the generated site documentation to the specified web server",
	}

	// TODO this traverses over all versions of a plugin while only the latest should be used
	_ = filepath.Walk(repositoryLocation()+"/org/apache/maven/plugins/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".jar") {
			if loadedPlugin := loadPlugin(path); loadedPlugin != nil {
				for key, value := range loadedPlugin.FormattedGoals() {
					goals[key] = value
				}
			}
		}
		return nil
	})
	return
}
