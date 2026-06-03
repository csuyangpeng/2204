package main

import (
	"bufio"
	"flag"
	"fmt"
	"lite5gc/udm/sidf/conf/genkey"
	"os"
	"strings"
)

// Public-private key pair of profile A using the curve25519 algorithm
func main() {
	// 命令行
	commandOpt()

}

func handlerCommand(cmd string) error {
	switch cmd {
	case "create":
		if sum != 0 {
			//fmt.Println(sum, filename)
			err := genkey.CreateJsonFile(sum, filename)
			//fmt.Println(err)
			//fmt.Println(os.Getwd())
			if err != nil {
				return err
			}
			sum = 0
			filename = ""
		} else {
			return fmt.Errorf("Please enter parameters")
		}
	default:
		return fmt.Errorf("Generated command:\ncreate -sum 10 -filename pki.key")
	}
	return nil
}

var CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

func commandOpt() {
	fmt.Println("Please specify the number of records and file name.")
	fmt.Println("Generated command:\ncreate -sum 10 -filename pki.key")

	// 定义参数，不使用默认函数

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// 循环获取命令行
		line := scanner.Text()
		if line == "exit" {
			os.Exit(0)
			//return
		}

		if line == "h" || line == "help" {
			fmt.Println("Generated command:\ncreate -sum 10 -filename pki.key")
			continue
		}
		// fmt.Println(line)
		// 分割命令
		Agrs := strings.Fields(line)
		//fmt.Println(Agrs[0])
		//fmt.Println(Agrs[1:])

		//fmt.Println(len(Agrs))
		//fmt.Println(Agrs)
		fmt.Println("Current parameter value: sum      ", sum)
		fmt.Println("Current parameter value: filename ", filename)
		//fmt.Println(os.Stdout.Seek(0,0))
		//fmt.Println(" #:")
		if len(Agrs) > 0 {
			cmd := Agrs[0]
			//解析命令行
			err := CommandLine.Parse(Agrs[1:]) // 如果错误，打印默认help，继续等待
			if err != nil {
				fmt.Println("Command line flag syntax fault")
				//CommandLine.PrintDefaults()
			}
			if cmd == "" {
				fmt.Println("Please enter a command")
			}
			err = handlerCommand(cmd)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Please enter a command")
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

// 设置参数
var sum uint
var filename string

func init() {
	// Override generic FlagSet default Usage with call to global Usage.
	// Note: This is not CommandLine.Usage = Usage,
	// because we want any eventual call to use any updated value of Usage,
	// not the value it has when this line is run.
	CommandLine.Usage = commandLineUsage

	// 方式1，传入变量的地址，获得对应参数的赋值
	//CommandLine.StringVar(&N3IpAddr, "n3-ip", "127.0.0.1", "specify ip address to use.  defaults to 127.0.0.1.")
	CommandLine.UintVar(&sum, "sum", 0, "specify the number of PKI records generated to use.  defaults to 0.")
	CommandLine.StringVar(&filename, "filename", "", "specify the file name of PKI records generated to use.  defaults to NULL.")
}

/*// 方式2，返回对应类型指针
var (
	Port     = CommandLine.Int("p", 8000, "specify port to use.  defaults to 8000.")
	N3IpAddr string
)*/

func commandLineUsage() {
	Usage()
}

var Usage = func() {
	fmt.Fprintf(CommandLine.Output(), "Usage of %s:\n", os.Args[0])
	CommandLine.PrintDefaults()
	fmt.Println("Generated command:\ncreate -sum 10 -filename pki.key")
}
