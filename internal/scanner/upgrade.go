package scanner

func CollectUpgradeCommand(results []ModuleResult) []string {
	var cmds []string

	for _, r := range results {
		if r.UpgradeSuggestion != "" {
			cmds = append(cmds, r.UpgradeSuggestion)
		}
	}
	return cmds
}
