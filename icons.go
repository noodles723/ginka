package ginka

var (
	// Icons is a default SpecsIcons
	Icons = newSpecsIcons()
)

// SpecsIcons contains unicode "icons" for specs status.
type SpecsIcons struct {
	passed  string
	failed  string
	pending string
	skipped string
	panicked string
}

func newSpecsIcons() SpecsIcons {
	return SpecsIcons{
		passed:  `✔`,
		failed:  `✘`,
		pending: `❗`,
		skipped: `✱`,
		panicked: `💀`,
	}
}
