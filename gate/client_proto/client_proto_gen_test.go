package client_proto

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestClientProtoGen(t *testing.T) {
	// 生成根据proto类名获取对象实例的switch方法
	dir, err := os.ReadDir("./proto")
	if err != nil {
		panic(err)
	}
	protoObjNameList := make([]string, 0)
	for _, entry := range dir {
		if entry.IsDir() {
			continue
		}
		split := strings.Split(entry.Name(), ".")
		if len(split) < 2 || split[len(split)-1] != "proto" {
			continue
		}
		protoObjNameList = append(protoObjNameList, split[len(split)-2])
	}
	// 生成初始化cmdId和cmdName的方法
	clientCmdFile, err := os.ReadFile("./proto/client_cmd.csv")
	if err != nil {
		panic(err)
	}
	clientCmdData := string(clientCmdFile)
	clientCmdLineList := strings.Split(clientCmdData, "\n")
	// 生成代码文件
	var fileDataBuffer bytes.Buffer
	fileDataBuffer.WriteString(`package client_proto

import (
	"hk4e/gate/client_proto/proto"
)

func (c *ClientCmdProtoMap) LoadClientCmdIdAndCmdName() {
`)
	for _, clientCmdLine := range clientCmdLineList {
		// 清理空格以及换行符之类的
		clientCmdLine = strings.TrimSpace(clientCmdLine)
		if clientCmdLine == "" {
			continue
		}
		item := strings.Split(clientCmdLine, ",")
		if len(item) != 2 {
			panic("parse client cmd file error")
		}
		cmdName := item[0]
		cmdId := item[1]
		_, err = fmt.Fprintf(&fileDataBuffer, `	c.clientCmdIdCmdNameMap[uint16(%s)] = "%s"
	c.clientCmdNameCmdIdMap["%s"] = uint16(%s)
`, cmdId, cmdName, cmdName, cmdId)
		if err != nil {
			panic(err)
		}
	}
	fileDataBuffer.WriteString(`}

func (c *ClientCmdProtoMap) GetClientProtoObjByName(protoObjName string) any {
	switch protoObjName {
`)
	for _, protoObjName := range protoObjNameList {
		_, err = fmt.Fprintf(&fileDataBuffer, `	case "%s":
		return new(proto.%s)
`, protoObjName, protoObjName)
		if err != nil {
			panic(err)
		}
	}
	fileDataBuffer.WriteString(`	default:
		return nil
	}
}
`)
	err = os.WriteFile("./client_proto_gen.go", fileDataBuffer.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
	// 处理枚举
	for _, entry := range dir {
		rawFileData, err := os.ReadFile("./proto/" + entry.Name())
		if err != nil {
			panic(err)
		}
		rawFileStr := string(rawFileData)
		rawFileLine := strings.Split(rawFileStr, "\n")
		var newFileBuffer bytes.Buffer
		for i := 0; i < len(rawFileLine); i++ {
			line := rawFileLine[i]
			newFileBuffer.WriteString(line + "\n")
			if !strings.Contains(line, "enum") {
				continue
			}
			split := strings.Split(strings.TrimSpace(line), " ")
			if len(split) != 3 || split[0] != "enum" || split[2] != "{" {
				continue
			}
			enumName := split[1]
			// 从protocol/proto_hk4e下复制同名的枚举类替换掉原proto文件里的内容
			refEnum := FindEnumInDirFile("../../protocol/proto_hk4e", enumName)
			if refEnum == nil {
				continue
			}
			for _, ref := range refEnum {
				newFileBuffer.WriteString(ref + "\n")
			}
			i++
			for {
				nextLine := rawFileLine[i]
				if !strings.Contains(nextLine, "}") {
					i++
				} else {
					newFileBuffer.WriteString(nextLine + "\n")
					break
				}
			}
		}
		err = os.WriteFile("./proto/"+entry.Name(), newFileBuffer.Bytes(), 0644)
		if err != nil {
			panic(err)
		}
	}
}

func FindEnumInDirFile(path string, name string) (lineList []string) {
	dir, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, entry := range dir {
		if entry.IsDir() {
			ret := FindEnumInDirFile(path+"/"+entry.Name(), name)
			if ret != nil {
				return ret
			}
			continue
		}
		fileData, err := os.ReadFile(path + "/" + entry.Name())
		if err != nil {
			panic(err)
		}
		fileStr := string(fileData)
		fileLine := strings.Split(fileStr, "\n")
		for i := 0; i < len(fileLine); i++ {
			line := fileLine[i]
			if !strings.Contains(line, "enum") {
				continue
			}
			split := strings.Split(strings.TrimSpace(line), " ")
			if len(split) != 3 || split[0] != "enum" || split[2] != "{" {
				continue
			}
			enumName := split[1]
			if enumName != name {
				continue
			}
			i++
			lineList := make([]string, 0)
			for {
				nextLine := fileLine[i]
				if !strings.Contains(nextLine, "}") {
					lineList = append(lineList, nextLine)
				} else {
					return lineList
				}
				i++
			}
		}
	}
	return nil
}
