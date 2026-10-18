package resolver

import (
	"strings"

	"github.com/kryptamine/herdr-auto-title/internal/state"
)

// Nerd Font glyphs, which sit in the Private Use Area: a terminal without such
// a font draws each as a box, so icons are only ever asked for.
const (
	iconAgent     = "\uF4BE"     // nf-oct-dependabot
	iconRemote    = "\uF233"     // nf-fa-server
	iconShell     = "\uF120"     // nf-fa-terminal
	iconProgram   = "\uF013"     // nf-fa-gear
	iconEditor    = "\uF044"     // nf-fa-pen_to_square
	iconVim       = "\uE62B"     // nf-custom-vim
	iconEmacs     = "\uE632"     // nf-custom-emacs
	iconGit       = "\uE702"     // nf-dev-git
	iconGitHub    = "\uE709"     // nf-dev-github_badge
	iconDocker    = "\uF308"     // nf-linux-docker
	iconK8s       = "\U000F10FE" // nf-md-kubernetes
	iconTerraform = "\U000F1062" // nf-md-terraform
	iconDatabase  = "\uF1C0"     // nf-fa-database
	iconPostgres  = "\uE76E"     // nf-dev-postgresql
	iconMySQL     = "\uE704"     // nf-dev-mysql
	iconRedis     = "\uE76D"     // nf-dev-redis
	iconMongo     = "\uE7A4"     // nf-dev-mongodb
	iconMonitor   = "\uF0E4"     // nf-fa-gauge
	iconPager     = "\uF02D"     // nf-fa-book
	iconFiles     = "\uF07C"     // nf-fa-folder_open
	iconBuild     = "\uF085"     // nf-fa-gears
	iconDebugger  = "\uF188"     // nf-fa-bug
	iconNetwork   = "\uF0AC"     // nf-fa-globe
	iconTmux      = "\uEBC8"     // nf-cod-terminal_tmux
	iconChat      = "\U000F0B79" // nf-md-chat
	iconPackage   = "\uEB29"     // nf-cod-package
	iconDebian    = "\uE77D"     // nf-dev-debian
	iconArch      = "\uF303"     // nf-linux-archlinux
	iconFedora    = "\uF30A"     // nf-linux-fedora
	iconNix       = "\uF313"     // nf-linux-nixos
	iconRust      = "\uE7A8"     // nf-dev-rust
	iconGo        = "\uE724"     // nf-dev-go
	iconNode      = "\uE718"     // nf-dev-nodejs_small
	iconPython    = "\uE73C"     // nf-dev-python
	iconRuby      = "\uE739"     // nf-dev-ruby
	iconPHP       = "\uE73D"     // nf-dev-php
	iconJava      = "\uE738"     // nf-dev-java
	iconScala     = "\uE737"     // nf-dev-scala
	iconElixir    = "\uE62D"     // nf-seti-elixir
	iconErlang    = "\uE7B1"     // nf-dev-erlang
	iconHaskell   = "\uE777"     // nf-dev-haskell
	iconLua       = "\uE620"     // nf-seti-lua
	iconSwift     = "\uE755"     // nf-dev-swift
	iconZig       = "\uE6A9"     // nf-seti-zig
	iconDotnet    = "\uE77F"     // nf-dev-dotnet
	iconPerl      = "\uE769"     // nf-dev-perl
	iconJulia     = "\uE624"     // nf-seti-julia
	iconC         = "\uE61E"     // nf-custom-c
)

// programIcons maps a lower-cased process name to its icon. A program missing
// here still gets iconProgram, so the list only has to name what is common.
var programIcons = map[string]string{
	"nvim": iconVim, "vim": iconVim, "vi": iconVim, "lvim": iconVim,
	"emacs": iconEmacs,
	"hx":    iconEditor, "helix": iconEditor, "nano": iconEditor, "micro": iconEditor,
	"kak": iconEditor,

	"git": iconGit, "lazygit": iconGit, "tig": iconGit, "gitui": iconGit,
	"jj": iconGit, "lazyjj": iconGit, "hg": iconGit,
	"gh": iconGitHub,

	"docker": iconDocker, "docker-compose": iconDocker, "lazydocker": iconDocker,
	"podman":  iconDocker,
	"kubectl": iconK8s, "k9s": iconK8s, "helm": iconK8s, "minikube": iconK8s,
	"kubie": iconK8s, "kind": iconK8s,
	"terraform": iconTerraform, "tofu": iconTerraform,

	"psql": iconPostgres, "pgcli": iconPostgres,
	"mysql": iconMySQL, "mycli": iconMySQL,
	"redis-cli": iconRedis, "valkey-cli": iconRedis,
	"mongosh": iconMongo, "mongo": iconMongo,
	"sqlite3": iconDatabase, "sqlite": iconDatabase, "litecli": iconDatabase,
	"duckdb": iconDatabase, "usql": iconDatabase,

	"htop": iconMonitor, "btop": iconMonitor, "top": iconMonitor, "btm": iconMonitor,
	"glances": iconMonitor, "nvtop": iconMonitor, "mactop": iconMonitor,
	"less": iconPager, "more": iconPager, "man": iconPager, "bat": iconPager,
	"glow": iconPager,
	"yazi": iconFiles, "ranger": iconFiles, "lf": iconFiles, "nnn": iconFiles,
	"vifm": iconFiles, "mc": iconFiles, "broot": iconFiles,

	"make": iconBuild, "just": iconBuild, "cmake": iconBuild, "ninja": iconBuild,
	"bazel": iconBuild, "gradle": iconBuild, "mvn": iconBuild,
	"gdb": iconDebugger, "lldb": iconDebugger, "dlv": iconDebugger,
	"valgrind": iconDebugger,
	"curl":     iconNetwork, "wget": iconNetwork, "xh": iconNetwork, "http": iconNetwork,
	"ping": iconNetwork, "gping": iconNetwork, "mtr": iconNetwork,
	"tmux": iconTmux, "screen": iconTmux, "zellij": iconTmux,
	"weechat": iconChat, "irssi": iconChat,

	"brew": iconPackage,
	"apt":  iconDebian, "apt-get": iconDebian, "dpkg": iconDebian, "nala": iconDebian,
	"pacman": iconArch, "paru": iconArch, "yay": iconArch,
	"dnf": iconFedora, "yum": iconFedora,
	"nix": iconNix,

	"cargo": iconRust, "rustc": iconRust, "rustup": iconRust, "bacon": iconRust,
	"go": iconGo, "air": iconGo,
	"node": iconNode, "npm": iconNode, "npx": iconNode, "pnpm": iconNode,
	"yarn": iconNode, "bun": iconNode, "deno": iconNode, "tsx": iconNode,
	"python": iconPython, "python3": iconPython, "ipython": iconPython,
	"uv": iconPython, "pip": iconPython, "pip3": iconPython, "pytest": iconPython,
	"poetry": iconPython,
	"ruby":   iconRuby, "irb": iconRuby, "bundle": iconRuby, "rails": iconRuby,
	"php": iconPHP, "composer": iconPHP,
	"java":  iconJava,
	"scala": iconScala, "sbt": iconScala,
	"elixir": iconElixir, "iex": iconElixir, "mix": iconElixir,
	"erl": iconErlang, "beam.smp": iconErlang,
	"ghc": iconHaskell, "ghci": iconHaskell, "stack": iconHaskell, "cabal": iconHaskell,
	"lua": iconLua, "luajit": iconLua,
	"swift":  iconSwift,
	"zig":    iconZig,
	"dotnet": iconDotnet,
	"perl":   iconPerl,
	"julia":  iconJulia,
	"gcc":    iconC, "clang": iconC,
}

// iconFor picks the icon for the program in a pane's foreground. Only that one:
// a descendant is not what the pane is for, and nvim starts language servers
// under node, git starts ssh. A runtime and a remote session still get a mark.
func iconFor(pane *state.PaneState) string {
	if pane == nil {
		return iconShell
	}

	if pane.HasAgent() {
		return iconAgent
	}

	process, ok := pane.Foreground()
	name := strings.ToLower(strings.TrimSpace(process.Name))

	if _, isShell := shellNames[name]; !ok || name == "" || isShell {
		return iconShell
	}

	if runsRemote(name) {
		return iconRemote
	}

	if icon, known := programIcons[name]; known {
		return icon
	}

	return iconProgram
}

// IconedTitles puts an icon for what a tab's pane runs in front of its title.
// It wraps a resolver for the reason Numbered does: the icon is not one of the
// things a title says, so no source supplies it and no repetition check drops it.
type IconedTitles struct {
	inner TitleResolver
	// maxLength is the bound inner was built with; the icon is counted against
	// it, as Numbered counts the position.
	maxLength int
}

var _ TitleResolver = (*IconedTitles)(nil)

// NewIconedTitles wraps inner, which must have been given the same maxLength.
// Zero or less takes the default, as New does.
func NewIconedTitles(inner TitleResolver, maxLength int) *IconedTitles {
	if maxLength <= 0 {
		maxLength = DefaultMaxLength
	}

	return &IconedTitles{inner: inner, maxLength: maxLength}
}

func (i *IconedTitles) Resolve(tab state.TabState) Decision {
	decision := i.inner.Resolve(tab)
	decision.Name = withPrefix(iconFor(tab.Context)+" ", decision.Name, i.maxLength)

	return decision
}

// IconedPanes is IconedTitles for pane labels, each pane under its own icon.
type IconedPanes struct {
	inner     PaneResolver
	maxLength int
}

var _ PaneResolver = (*IconedPanes)(nil)

// NewIconedPanes wraps inner, which must have been given the same maxLength.
// Zero or less takes the default, as New does.
func NewIconedPanes(inner PaneResolver, maxLength int) *IconedPanes {
	if maxLength <= 0 {
		maxLength = DefaultMaxLength
	}

	return &IconedPanes{inner: inner, maxLength: maxLength}
}

func (i *IconedPanes) ResolvePanes(tab state.TabState) []Decision {
	decisions := i.inner.ResolvePanes(tab)

	for n := range decisions {
		decisions[n].Name = withPrefix(iconFor(tab.Panes[n])+" ", decisions[n].Name, i.maxLength)
	}

	return decisions
}
