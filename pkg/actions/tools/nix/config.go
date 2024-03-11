package nix

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionConfigKeys completes config keys
//
//	access-tokens (Access tokens used to access protected GitHub, GitLab, or other locations)
//	allow-dirty (Whether to allow dirty Git/Mercurial trees)
func ActionConfigKeys() carapace.Action {
	return carapace.ActionValuesDescribed(
		"accept-flake-config", "Whether to accept nix configuration from a flake without prompting",
		"access-tokens", "Access tokens used to access protected GitHub, GitLab, or other locations",
		"allow-dirty", "Whether to allow dirty Git/Mercurial trees",
		"allow-import-from-derivation", "allowing building at evaluation time",
		"allow-new-privileges", "allow privilege aquirement using setuid/segid",
		"allow-symlinked-store", "allow symlink components in store directory",
		"allow-unsafe-native-code-during-evaluation", "Whether builtin functions that allow executing native code should be enabled",
		"allowed-impure-host-deps", "Which prefixes to allow derivations to ask for access to",
		"allowed-uris", "A list of URI prefixes to which access is allowed in restricted evaluation mode",
		"allowed-users", "A list of names of users that are allowed to connect to the Nix daemon",
		"auto-allocate-uids", "Whether to select UIDs for builds automatically",
		"auto-optimise-store", "automatically detect files in the store that have identical contents",
		"bash-prompt", "The bash prompt (PS1) in nix develop shells.",
		"bash-prompt-prefix", "Prefix prepended to the PS1 environment variable in nix develop shells.",
		"bash-prompt-suffix", "Suffix appended to the PS1 environment variable in nix develop shells.",
		"build-hook", "The path of the helper program that executes builds to remote machines.",
		"build-poll-interval", "How often (in seconds) to poll for locks.",
		"build-users-group", "This  options  specifies  the Unix group containing the Nix build user accounts",
		"builders", "A semicolon-separated list of build machines",
		"builders-use-substitutes", "instruct remote build machines to use their own binary substitutes if available",
		"commit-lockfile-summary", "The commit summary to use when committing changed flake lock files",
		"compress-build-log", "build logs written to /nix/var/log/nix/drvs will be compressed on the fly using bzip2",
		"connect-timeout", "The  timeout (in seconds) for establishing connections in the binary cache substituter",
		"cores", "Sets the value of the NIX_BUILD_CORES environment variable in the invocation of builders",
		"diff-hook", "Absolute  path  to  an executable capable of diffing build results",
		"download-attempts", "How often Nix will attempt to download a file before giving up.",
		"download-speed", "Specify the maximum transfer rate in kilobytes per second you want Nix to use for downloads.",
		"eval-cache", "Whether to use the flake evaluation cache.",
		"experimental-features", "Experimental Nix features to enable.",
		"extra-platforms", "Platforms other than the native one which this machine is capable of building for",
		"fallback", "fall back to building from source if a binary substitute fails",
		"filter-syscalls", "Whether to prevent certain dangerous system calls",
		"flake-registry", "Path or URI of the global flake registry.",
		"fsync-metadata", "changes to the Nix store metadata are synchronously flushed to disk",
		"gc-reserved-space", "Amount of reserved disk space for the garbage collector",
		"hashed-mirrors", "A  list of web servers used by builtins.fetchurl to obtain files by hash",
		"http-connections", "The maximum number of parallel TCP connections used to fetch files",
		"http2", "Whether to enable HTTP/2 support.",
		"id-count", "The number of UIDs/GIDs to use for dynamic ID allocation.",
		"ignore-try", "ignore exceptions inside ‘tryEval’ calls when evaluating nix expressions in debug mode",
		"ignored-acls", "A list of ACLs that should be ignored",
		"impersonate-linux-26", "Whether to impersonate a Linux 2.6 machine on newer kernels.",
		"keep-build-log", "write the build log of a derivation",
		"keep-derivations", "the garbage collector will keep the derivations from which non-garbage store paths were built",
		"keep-env-derivations", "If  false (default), derivations are not stored in Nix user environments",
		"keep-failed", "Whether to keep temporary directories of failed builds.",
		"keep-going", "Whether to keep building derivations when another build fails.",
		"keep-outputs", "If  true,  the  garbage collector will keep the outputs of non-garbage derivations",
		"log-lines", "The number of lines of the tail of the log to show if a build fails.",
		"max-build-log-size", "This option defines the maximum number of bytes that a builder can write to its stdout/stderr",
		"max-free", "When a garbage collection is triggered by the min-free option",
		"max-jobs", "This option defines the maximum number of jobs that Nix will try to build in parallel",
		"max-silent-time", "maximum number of seconds that a builder can go without producing any data on standard output",
		"min-free", "When  free  disk  space  in /nix/store drops below min-free during a build, Nix performs a garbage-collection",
		"min-free-check-interval", "Number of seconds between checking free disk space",
		"nar-buffer-size", "Maximum size of NARs before spilling them to disk",
		"narinfo-cache-negative-ttl", "The TTL in seconds for negative lookups",
		"narinfo-cache-positive-ttl", "The  TTL in seconds for positive lookups",
		"netrc-file", "Nix will use the HTTP authentication credentials in this file when trying to download",
		"nix-path", "List of directories to be searched for <...> file references.",
		"plugin-files", "A  list  of  plugin  files to be loaded by Nix",
		"post-build-hook", "Optional. The path to a program to execute after each build",
		"pre-build-hook", "the path to a program that can set extra derivation-specific settings for this system",
		"preallocate-contents", "Whether to preallocate files when writing objects with known size.",
		"print-missing", "Whether to print what paths need to be built or downloaded.",
		"pure-eval", "Whether to restrict file system and network access to files specified by cryptographic hash.",
		"require-sigs", "any non-content-addressed path added or copied to the Nix store must have a signature by a trusted key",
		"restrict-eval", "If set to true, the Nix evaluator will not allow access to any files outside of the Nix search path",
		"run-diff-hook", "If true, enable the execution of the diff-hook program.",
		"sandbox", "If  set to true, builds will be performed in a sandboxed environment",
		"sandbox-build-dir", "The build directory inside the sandbox.",
		"sandbox-dev-shm-size", "This option determines the maximum size of the tmpfs filesystem",
		"sandbox-fallback", "Whether to disable sandboxing when the kernel doesn’t allow it.",
		"sandbox-paths", "A  list of paths bind-mounted into Nix sandbox environments",
		"secret-key-files", "A whitespace-separated list of files containing secret (private) keys",
		"show-trace", "Whether Nix should print out a stack trace in case of Nix expression evaluation errors",
		"stalled-download-timeout", "The timeout (in seconds) for receiving data from servers during download",
		"start-id", "The first UID and GID to use for dynamic ID allocation.",
		"store", "The default Nix store to use.",
		"substitute", "If set to true (default), Nix will use binary substitutes if available",
		"substituters", "A list of URLs of substituters, separated by whitespace",
		"sync-before-registering", "Whether to call sync() before registering a path as valid.",
		"system", "This  option  specifies  the  canonical Nix system name of the current installation",
		"system-features", "A  set  of  system  “features” supported by this machine, e.g. kvm",
		"tarball-ttl", "The  number of seconds a downloaded tarball is considered fresh",
		"timeout", "This option defines the maximum number of seconds that a builder can run",
		"trace-function-calls", "If set to true, the Nix evaluator will trace every function call",
		"trace-verbose", "Whether builtins.traceVerbose should trace its first argument when evaluated.",
		"trusted-public-keys", "A whitespace-separated list of public keys",
		"trusted-substituters", "A list of URLs of substituters, separated by whitespace",
		"trusted-users", "A list of names of users that have additional rights when connecting to the Nix daemon",
		"use-case-hack", "Whether to enable a Darwin-specific hack for dealing with file name collisions.",
		"use-cgroups", "Whether to execute builds inside cgroups",
		"use-registries", "Whether to use flake registries to resolve flake references.",
		"use-sqlite-wal", "Whether SQLite should use WAL mode.",
		"user-agent-suffix", "String appended to the user agent in HTTP requests.",
		"warn-dirty", "Whether to warn about dirty Git/Mercurial trees.",
	)
}

// ActionConfigValues completes config values for given key
//
//	false
//	true
func ActionConfigValues(key string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		_bool := carapace.ActionValues("true", "false").StyleF(style.ForKeyword)
		// TODO add missing completions
		return map[string]carapace.Action{
			"accept-flake-config":                        _bool,
			"allow-dirty":                                _bool,
			"allow-import-from-derivation":               _bool,
			"allow-new-privileges":                       _bool,
			"allow-symlinked-store":                      _bool,
			"allow-unsafe-native-code-during-evaluation": _bool,
			"allowed-users":                              os.ActionUsers().UniqueList(" "),
			"auto-allocate-uids":                         _bool,
			"auto-optimise-store":                        _bool,
			"builders-use-substitutes":                   _bool,
			"compress-build-log":                         _bool,
			"eval-cache":                                 _bool,
			"fallback":                                   _bool,
			"filter-syscalls":                            _bool,
			"fsync-metadata":                             _bool,
			"http2":                                      _bool,
			"ignore-try":                                 _bool,
			"impersonate-linux-26":                       _bool,
			"keep-build-log":                             _bool,
			"keep-derivations":                           _bool,
			"keep-env-derivations":                       _bool,
			"keep-failed":                                _bool,
			"keep-going":                                 _bool,
			"keep-outputs":                               _bool,
			"netrc-file":                                 carapace.ActionFiles(),
			"nix-path":                                   carapace.ActionDirectories().List(","),
			"plugin-files":                               carapace.ActionFiles().List(","),
			"print-missing":                              _bool,
			"pure-eval":                                  _bool,
			"require-sigs":                               _bool,
			"restrict-eval":                              _bool,
			"run-diff-hook":                              _bool,
			"sandbox":                                    _bool,
			"sandbox-fallback":                           _bool,
			"secret-key-files":                           carapace.ActionFiles().List(","),
			"show-trace":                                 _bool,
			"sync-before-registering":                    _bool,
			"trace-function-calls":                       _bool,
			"trace-verbose":                              _bool,
			"trusted-users":                              os.ActionUsers().UniqueList(" "),
			"use-case-hack":                              _bool,
			"use-cgroups":                                _bool,
			"use-registries":                             _bool,
			"use-sqlite-wal":                             _bool,
			"warn-dirty":                                 _bool,
		}[key]
	})
}
