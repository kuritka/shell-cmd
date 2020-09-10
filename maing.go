package main

func main() {

	cmd := shell.Command{
		Command: "terraform",
		Args: []string{"plan", "-input=false", "-out=tfplan"},
	}
	o, err := shell.Execute(cmd)
	if err != nil {
		fmt.Printf("%v",err)
	}
	fmt.Println(o)
	//output, _ := Plan()
	//output, _ = Apply()
	//log.Printf("AA %+v", output)
}
