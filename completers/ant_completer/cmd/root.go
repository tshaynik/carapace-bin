package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/ant_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ant",
	Short: "software tool for automating software build processes",
	Long:  "https://ant.apache.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("autoproxy", "autoproxy", false, "Java1.5+: use the OS proxy settings")
	rootCmd.Flags().StringS("buildfile", "buildfile", "", "use given buildfile")
	rootCmd.Flags().BoolS("debug", "debug", false, "print debugging information")
	rootCmd.Flags().BoolS("diagnostics", "diagnostics", false, "print information that might be helpful to")
	rootCmd.Flags().BoolS("emacs", "emacs", false, "produce logging information without adornments")
	rootCmd.Flags().Bool("execdebug", false, "print ant exec line generated by this launch script")
	rootCmd.Flags().StringS("f", "f", "", "use given buildfile")
	rootCmd.Flags().StringS("file", "file", "", "use given buildfile")
	rootCmd.Flags().StringS("find", "find", "", "(s)earch for buildfile towards the root of")
	rootCmd.Flags().BoolS("help", "help", false, "print this message and exit")
	rootCmd.Flags().StringS("inputhandler", "inputhandler", "", "the class which will handle input requests")
	rootCmd.Flags().BoolS("k", "k", false, "execute all targets that do not depend on failed target(s)")
	rootCmd.Flags().BoolS("keep-going", "keep-going", false, "execute all targets that do not depend on failed target(s)")
	rootCmd.Flags().StringS("l", "l", "", "use given file for log")
	rootCmd.Flags().StringS("lib", "lib", "", "specifies a path to search for jars and classes")
	rootCmd.Flags().StringS("listener", "listener", "", "add an instance of class as a project listener")
	rootCmd.Flags().StringS("logfile", "logfile", "", "use given file for log")
	rootCmd.Flags().StringS("logger", "logger", "", "the class which is to perform logging")
	rootCmd.Flags().StringS("main", "main", "", "override Ant's normal entry point")
	rootCmd.Flags().StringS("nice", "nice", "", "A niceness value for the main thread")
	rootCmd.Flags().BoolS("noclasspath", "noclasspath", false, "Run ant without using CLASSPATH")
	rootCmd.Flags().Bool("noconfig", false, "suppress sourcing of configuration files")
	rootCmd.Flags().BoolS("noinput", "noinput", false, "do not allow interactive input")
	rootCmd.Flags().BoolS("nouserlib", "nouserlib", false, "Run ant without using the jar files from ${user.home}/.ant/lib")
	rootCmd.Flags().BoolS("projecthelp", "projecthelp", false, "print project help information and exit")
	rootCmd.Flags().StringS("propertyfile", "propertyfile", "", "load all properties from file with -D properties taking precedence")
	rootCmd.Flags().BoolS("quiet", "quiet", false, "be extra quiet")
	rootCmd.Flags().StringS("s", "s", "", "(s)earch for buildfile towards the root of")
	rootCmd.Flags().BoolS("silent", "silent", false, "print nothing but task outputs and build failures")
	rootCmd.Flags().Bool("usejikes", false, "enable use of jikes by default, unless set explicitly in configuration files")
	rootCmd.Flags().BoolS("verbose", "verbose", false, "be extra verbose")
	rootCmd.Flags().BoolS("version", "version", false, "print the version information and exit")

	// TODO class completions
	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"buildfile":    carapace.ActionFiles(".xml"),
		"f":            carapace.ActionFiles(".xml"),
		"file":         carapace.ActionFiles(".xml"),
		"l":            carapace.ActionFiles(),
		"lib":          carapace.ActionDirectories(),
		"logfile":      carapace.ActionFiles(),
		"propertyfile": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			path := ""
			if flag := rootCmd.Flag("buildfile"); flag.Changed {
				path = flag.Value.String()
			}
			if flag := rootCmd.Flag("file"); flag.Changed {
				path = flag.Value.String()
			}
			if flag := rootCmd.Flag("f"); flag.Changed {
				path = flag.Value.String()
			}
			return action.ActionTargets(path).FilterArgs()
		}),
	)
}
