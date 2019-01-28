# coat

stdout and stderr coating with color, time and muffler included

## how to
    NewConsole first parameter is the domain or name of the console, second one is the silence option
```golang
package main

func main() {
	console := NewConsole("test", false)
    console.Log("coated")
    console.Err("uncoated")
}
```