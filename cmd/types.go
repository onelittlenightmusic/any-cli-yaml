package cmd

type Command struct {
	Name string
	Options []Option
}

type Option struct {
	Key string
	Values []string
}

type SerializedCommandsFile struct {
	Type string
	Spec struct {
		Cmds []SerializedCommand
	}
}

type SerializedCommand struct {
	Cmd string
	Opts []struct {
		Opt string
		Vals []struct {
			Val string
		}
	}
}