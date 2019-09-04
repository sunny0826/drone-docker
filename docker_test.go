package docker

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"reflect"
	"strings"
	"testing"
)

func TestPlugin_Read(t *testing.T) {

	b, err := ioutil.ReadFile("git.txt")
	if err != nil {
		fmt.Println("ioutil ReadFile error: ", err)
	}
	if string(b) == "" {

	}
	fmt.Println(strings.Split(string(b), ","))
}

func TestPlugin_Exec(t *testing.T) {
	type fields struct {
		Login   Login
		Build   Build
		Daemon  Daemon
		Dryrun  bool
		Cleanup bool
		Modname string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Plugin{
				Login:   tt.fields.Login,
				Build:   tt.fields.Build,
				Daemon:  tt.fields.Daemon,
				Dryrun:  tt.fields.Dryrun,
				Cleanup: tt.fields.Cleanup,
				Modname: tt.fields.Modname,
			}
			if err := p.Exec(); (err != nil) != tt.wantErr {
				t.Errorf("Plugin.Exec() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_commandLogin(t *testing.T) {
	type args struct {
		login Login
	}
	tests := []struct {
		name string
		args args
		want *exec.Cmd
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := commandLogin(tt.args.login); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("commandLogin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isCommandPull(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isCommandPull(tt.args.args); got != tt.want {
				t.Errorf("isCommandPull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_commandPull(t *testing.T) {
	type args struct {
		repo string
	}
	tests := []struct {
		name string
		args args
		want *exec.Cmd
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := commandPull(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("commandPull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_commandLoginEmail(t *testing.T) {
	type args struct {
		login Login
	}
	tests := []struct {
		name string
		args args
		want *exec.Cmd
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := commandLoginEmail(tt.args.login); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("commandLoginEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_commandVersion(t *testing.T) {
	tests := []struct {
		name string
		want *exec.Cmd
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := commandVersion(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("commandVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_commandInfo(t *testing.T) {
	tests := []struct {
		name string
		want *exec.Cmd
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := commandInfo(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("commandInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_commandBuild(t *testing.T) {
	type args struct {
		build Build
	}
	tests := []struct {
		name string
		args args
		want *exec.Cmd
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := commandBuild(tt.args.build); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("commandBuild() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addProxyBuildArgs(t *testing.T) {
	type args struct {
		build *Build
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			addProxyBuildArgs(tt.args.build)
		})
	}
}

func Test_addProxyValue(t *testing.T) {
	type args struct {
		build *Build
		key   string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			addProxyValue(tt.args.build, tt.args.key)
		})
	}
}

func Test_getProxyValue(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getProxyValue(tt.args.key); got != tt.want {
				t.Errorf("getProxyValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hasProxyBuildArg(t *testing.T) {
	type args struct {
		build *Build
		key   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasProxyBuildArg(tt.args.build, tt.args.key); got != tt.want {
				t.Errorf("hasProxyBuildArg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_commandTag(t *testing.T) {
	type args struct {
		build Build
		tag   string
	}
	tests := []struct {
		name string
		args args
		want *exec.Cmd
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := commandTag(tt.args.build, tt.args.tag); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("commandTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_commandPush(t *testing.T) {
	type args struct {
		build Build
		tag   string
	}
	tests := []struct {
		name string
		args args
		want *exec.Cmd
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := commandPush(tt.args.build, tt.args.tag); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("commandPush() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_commandDaemon(t *testing.T) {
	type args struct {
		daemon Daemon
	}
	tests := []struct {
		name string
		args args
		want *exec.Cmd
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := commandDaemon(tt.args.daemon); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("commandDaemon() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_commandPrune(t *testing.T) {
	tests := []struct {
		name string
		want *exec.Cmd
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := commandPrune(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("commandPrune() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_trace(t *testing.T) {
	type args struct {
		cmd *exec.Cmd
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trace(tt.args.cmd)
		})
	}
}

func TestPlugin_checkModuleNmae(t *testing.T) {
	type fields struct {
		Login   Login
		Build   Build
		Daemon  Daemon
		Dryrun  bool
		Cleanup bool
		Modname string
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "test",
			fields: fields{
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Plugin{
				Login:   tt.fields.Login,
				Build:   tt.fields.Build,
				Daemon:  tt.fields.Daemon,
				Dryrun:  tt.fields.Dryrun,
				Cleanup: tt.fields.Cleanup,
				Modname: tt.fields.Modname,
			}
			if got := p.checkModuleNmae(tt.args.name); got != tt.want {
				t.Errorf("Plugin.checkModuleNmae() = %v, want %v", got, tt.want)
			}
		})
	}
}
