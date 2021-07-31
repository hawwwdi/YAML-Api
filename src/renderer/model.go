package renderer

type templateTwo struct {
	Command     string `json: "Command" binding:"required"`
	Sub_command string `json: "sub_command" binding:"required"`
}

type templateOne struct {
	Docker_file string `json: "docker_file" binding:"required"`
}
