package actions

type CommandInfo struct {
	Name        string
	Function    func([]string)
	Description string
	Arguments   string
}

func (cmd *CommandInfo) Print() string {
	return "- " + cmd.Name + " " + cmd.Arguments + "\n" + "  - Description: " + cmd.Description + "\n"
}
