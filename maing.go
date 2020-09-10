package main

func main() {
	//cmd := shell.Command{
	//	Command: "sh",
	//	Args: []string{"-c", "plan -input=false -out=tfplan"},
	//}

	cmd := shell.Command{
		Command: "terraform",
		Args: []string{"plan", "-input=false", "-out=tfplan"},
	}
	o, err := shell.Execute(cmd)
	if err != nil {
		fmt.Printf("%v",err)
	}
	fmt.Println(o)
}
