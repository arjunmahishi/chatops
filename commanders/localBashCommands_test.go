package commanders

import (
	"reflect"
	"regexp"
	"testing"

	exoctx "bitbucket.org/exotel/exotel_code/commonix/lib/exoctx/go"
	"github.com/arjunmahishi/Chatops/mocks"
	"github.com/arjunmahishi/Chatops/payload"
)

func Test_runCommandLocally(t *testing.T) {
	var mockCtx exoctx.Ctx

	type args struct {
		ctx     exoctx.Ctx
		command string
		args    []string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name:    "Running pwd command",
			args:    args{ctx: mockCtx, command: "pwd", args: []string{}},
			want:    []byte("/home/arjunmahishi/go/src/bitbucket.org/exotel/Chatops/commanders"),
			wantErr: false,
		},
		{
			name:    "Running bad command for error",
			args:    args{ctx: mockCtx, command: "cd nodir"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Running a python script",
			args:    args{ctx: mockCtx, command: "python", args: []string{"test.py"}},
			want:    []byte("hello world"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := runCommandLocally(tt.args.command)
			if (err != nil) != tt.wantErr {
				t.Errorf("runCommandLocally() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("runCommandLocally() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocalBashCommand_Execute(t *testing.T) {
	var mockCtx exoctx.Ctx
	var mockPayloadPy, mockPayloadPyArg, mockPayloadNoArgs mocks.PayloadHandler

	mockPayloadPy.On("GetMessage").Return("testpy")
	mockPayloadPyArg.On("GetMessage").Return("testpy test.py")
	mockPayloadNoArgs.On("GetMessage").Return("list ")

	type fields struct {
		Command      string
		Args         []string
		Name         string
		Catagory     string
		OutputFormat string
		Scope        int64
		RegexPattern *regexp.Regexp
	}
	type args struct {
		ctx            exoctx.Ctx
		payloadHandler payload.Handler
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name: "Running a simple python file",
			fields: fields{
				Command:      "python",
				Args:         []string{"test.py"},
				Name:         "test-py",
				OutputFormat: "",
				RegexPattern: regexp.MustCompile(`(mi)testpy$`),
			},
			args:    args{ctx: mockCtx, payloadHandler: &mockPayloadPy},
			want:    map[string]interface{}{"text": "```hello world```"},
			wantErr: false,
		},
		{
			name: "Running a simple python file with file as a dynamic arg",
			fields: fields{
				Command:      "python",
				Args:         []string{"{{filename}}"},
				Name:         "",
				OutputFormat: "",
				RegexPattern: regexp.MustCompile(`(?ms)testpy\s+(?P<filename>\S+)$`),
			},
			args:    args{ctx: mockCtx, payloadHandler: &mockPayloadPyArg},
			want:    map[string]interface{}{"text": "```hello world```"},
			wantErr: false,
		},
		{
			name: "Running a command without args",
			fields: fields{
				Command:      "pwd",
				Args:         nil,
				Name:         "",
				OutputFormat: "",
				RegexPattern: regexp.MustCompile(`(mi)^$`),
			},
			args:    args{ctx: mockCtx, payloadHandler: &mockPayloadNoArgs},
			want:    map[string]interface{}{"text": "```/home/arjunmahishi/go/src/bitbucket.org/exotel/Chatops/commanders```"},
			wantErr: false,
		},
		{
			name: "Running a command that fails",
			fields: fields{
				Command:      "cd nodir",
				Args:         nil,
				Name:         "",
				OutputFormat: "",
				RegexPattern: regexp.MustCompile(`(mi)cd$`),
			},
			args:    args{ctx: mockCtx, payloadHandler: &mockPayloadNoArgs},
			want:    map[string]interface{}{"text": ""},
			wantErr: true,
		},
		{
			name: "Running a command that cant be parsed",
			fields: fields{
				Command:      "pwd",
				Args:         nil,
				Name:         "",
				OutputFormat: "access-log",
				RegexPattern: regexp.MustCompile(`(mi)^$`),
			},
			args:    args{ctx: mockCtx, payloadHandler: &mockPayloadNoArgs},
			want:    map[string]interface{}{"text": ""},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := LocalBashCommand{
				Command:      tt.fields.Command,
				Name:         tt.fields.Name,
				Catagory:     tt.fields.Catagory,
				OutputFormat: tt.fields.OutputFormat,
				Scope:        tt.fields.Scope,
				RegexPattern: tt.fields.RegexPattern,
			}
			_, err := bc.Execute(tt.args.payloadHandler)
			if (err != nil) != tt.wantErr {
				t.Errorf("LocalBashCommand.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want == nil {
				t.Errorf("LocalBashCommand.Execute() = %v", tt.want)
			}
		})
	}
}

func TestLocalBashCommand_OtherMethods(t *testing.T) {
	localCommand := LocalBashCommand{
		Name:         "name",
		Command:      "",
		Catagory:     "catagory",
		Scope:        3,
		RegexPattern: regexp.MustCompile(`^match$`),
	}

	if !localCommand.MatchCommand("match") {
		t.Fatal("Command should have matched")
	}
	if localCommand.MatchCommand("dont match") {
		t.Fatal("Command should not have matched")
	}
	if localCommand.GetName() != "name" {
		t.Fail()
	}
	if localCommand.GetCatagory() != "catagory" {
		t.Fail()
	}
	if localCommand.GetScope() != 3 {
		t.Fail()
	}
}
