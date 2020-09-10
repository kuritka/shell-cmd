# shellcmd
blocking shell command



```golang
cmd := shell.Command{
	Command: "terraform",
	Args: []string{"plan", "-input=false", "-out=tfplan"},
}
o, _ := shell.Execute(cmd)
fmt.Println(o)
```
  
```golang  
cmdsh := shell.Command{
	Command: "sh",
	Args: []string{"-c", "terraform apply -auto-approve tfplan"},
}
o, _ = shell.Execute(cmdsh)
fmt.Println(o)
```
