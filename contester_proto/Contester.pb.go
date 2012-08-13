// Code generated by protoc-gen-go from "Contester.proto"
// DO NOT EDIT!

package contester_proto

import proto "code.google.com/p/goprotobuf/proto"
import "encoding/json"
import "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type TestRun_Code int32

const (
	TestRun_Accepted            TestRun_Code = 10
	TestRun_TimeLimitExceeded   TestRun_Code = 11
	TestRun_RuntimeError        TestRun_Code = 12
	TestRun_WrongAnswer         TestRun_Code = 13
	TestRun_PresentationError   TestRun_Code = 14
	TestRun_MemoryLimitExceeded TestRun_Code = 15
	TestRun_TestingError        TestRun_Code = 16
)

var TestRun_Code_name = map[int32]string{
	10: "Accepted",
	11: "TimeLimitExceeded",
	12: "RuntimeError",
	13: "WrongAnswer",
	14: "PresentationError",
	15: "MemoryLimitExceeded",
	16: "TestingError",
}
var TestRun_Code_value = map[string]int32{
	"Accepted":            10,
	"TimeLimitExceeded":   11,
	"RuntimeError":        12,
	"WrongAnswer":         13,
	"PresentationError":   14,
	"MemoryLimitExceeded": 15,
	"TestingError":        16,
}

func (x TestRun_Code) Enum() *TestRun_Code {
	p := new(TestRun_Code)
	*p = x
	return p
}
func (x TestRun_Code) String() string {
	return proto.EnumName(TestRun_Code_name, int32(x))
}
func (x TestRun_Code) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *TestRun_Code) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TestRun_Code_value, data, "TestRun_Code")
	if err != nil {
		return err
	}
	*x = TestRun_Code(value)
	return nil
}

type Compilation_Code int32

const (
	Compilation_Success Compilation_Code = 1
	Compilation_Failure Compilation_Code = 2
)

var Compilation_Code_name = map[int32]string{
	1: "Success",
	2: "Failure",
}
var Compilation_Code_value = map[string]int32{
	"Success": 1,
	"Failure": 2,
}

func (x Compilation_Code) Enum() *Compilation_Code {
	p := new(Compilation_Code)
	*p = x
	return p
}
func (x Compilation_Code) String() string {
	return proto.EnumName(Compilation_Code_name, int32(x))
}
func (x Compilation_Code) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *Compilation_Code) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Compilation_Code_value, data, "Compilation_Code")
	if err != nil {
		return err
	}
	*x = Compilation_Code(value)
	return nil
}

type TestRun struct {
	Solution         *LocalExecution `protobuf:"bytes,1,opt,name=solution" json:"solution,omitempty"`
	Tester           *LocalExecution `protobuf:"bytes,2,opt,name=tester" json:"tester,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (this *TestRun) Reset()         { *this = TestRun{} }
func (this *TestRun) String() string { return proto.CompactTextString(this) }
func (*TestRun) ProtoMessage()       {}

func (this *TestRun) GetSolution() *LocalExecution {
	if this != nil {
		return this.Solution
	}
	return nil
}

func (this *TestRun) GetTester() *LocalExecution {
	if this != nil {
		return this.Tester
	}
	return nil
}

type ComputedTestResult struct {
	Result           *TestRun_Code `protobuf:"varint,1,opt,name=result,enum=contester.proto.TestRun_Code" json:"result,omitempty"`
	Time             *uint32       `protobuf:"varint,2,opt,name=time" json:"time,omitempty"`
	Memory           *uint32       `protobuf:"varint,3,opt,name=memory" json:"memory,omitempty"`
	ReturnCode       *uint32       `protobuf:"varint,4,opt,name=return_code" json:"return_code,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (this *ComputedTestResult) Reset()         { *this = ComputedTestResult{} }
func (this *ComputedTestResult) String() string { return proto.CompactTextString(this) }
func (*ComputedTestResult) ProtoMessage()       {}

func (this *ComputedTestResult) GetResult() TestRun_Code {
	if this != nil && this.Result != nil {
		return *this.Result
	}
	return 0
}

func (this *ComputedTestResult) GetTime() uint32 {
	if this != nil && this.Time != nil {
		return *this.Time
	}
	return 0
}

func (this *ComputedTestResult) GetMemory() uint32 {
	if this != nil && this.Memory != nil {
		return *this.Memory
	}
	return 0
}

func (this *ComputedTestResult) GetReturnCode() uint32 {
	if this != nil && this.ReturnCode != nil {
		return *this.ReturnCode
	}
	return 0
}

type Compilation struct {
	Failure          *bool                 `protobuf:"varint,1,opt,name=failure" json:"failure,omitempty"`
	ResultSteps      []*Compilation_Result `protobuf:"bytes,2,rep,name=result_steps" json:"result_steps,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (this *Compilation) Reset()         { *this = Compilation{} }
func (this *Compilation) String() string { return proto.CompactTextString(this) }
func (*Compilation) ProtoMessage()       {}

func (this *Compilation) GetFailure() bool {
	if this != nil && this.Failure != nil {
		return *this.Failure
	}
	return false
}

type Compilation_Result struct {
	StepName         *string         `protobuf:"bytes,1,opt,name=step_name" json:"step_name,omitempty"`
	Execution        *LocalExecution `protobuf:"bytes,2,opt,name=execution" json:"execution,omitempty"`
	Failure          *bool           `protobuf:"varint,3,opt,name=failure" json:"failure,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (this *Compilation_Result) Reset()         { *this = Compilation_Result{} }
func (this *Compilation_Result) String() string { return proto.CompactTextString(this) }
func (*Compilation_Result) ProtoMessage()       {}

func (this *Compilation_Result) GetStepName() string {
	if this != nil && this.StepName != nil {
		return *this.StepName
	}
	return ""
}

func (this *Compilation_Result) GetExecution() *LocalExecution {
	if this != nil {
		return this.Execution
	}
	return nil
}

func (this *Compilation_Result) GetFailure() bool {
	if this != nil && this.Failure != nil {
		return *this.Failure
	}
	return false
}

type CompileRequest struct {
	SourceModule     *Module   `protobuf:"bytes,1,req,name=source_module" json:"source_module,omitempty"`
	ExtraModules     []*Module `protobuf:"bytes,2,rep,name=extra_modules" json:"extra_modules,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (this *CompileRequest) Reset()         { *this = CompileRequest{} }
func (this *CompileRequest) String() string { return proto.CompactTextString(this) }
func (*CompileRequest) ProtoMessage()       {}

func (this *CompileRequest) GetSourceModule() *Module {
	if this != nil {
		return this.SourceModule
	}
	return nil
}

type CompileResponse struct {
	CompileResult    *Compilation `protobuf:"bytes,1,opt,name=compile_result" json:"compile_result,omitempty"`
	CompiledModule   *Module      `protobuf:"bytes,2,opt,name=compiled_module" json:"compiled_module,omitempty"`
	GeneralFailure   *string      `protobuf:"bytes,3,opt,name=general_failure" json:"general_failure,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (this *CompileResponse) Reset()         { *this = CompileResponse{} }
func (this *CompileResponse) String() string { return proto.CompactTextString(this) }
func (*CompileResponse) ProtoMessage()       {}

func (this *CompileResponse) GetCompileResult() *Compilation {
	if this != nil {
		return this.CompileResult
	}
	return nil
}

func (this *CompileResponse) GetCompiledModule() *Module {
	if this != nil {
		return this.CompiledModule
	}
	return nil
}

func (this *CompileResponse) GetGeneralFailure() string {
	if this != nil && this.GeneralFailure != nil {
		return *this.GeneralFailure
	}
	return ""
}

type ProblemHandle struct {
	Url              *string `protobuf:"bytes,1,req,name=url" json:"url,omitempty"`
	Revision         *uint32 `protobuf:"varint,2,opt,name=revision" json:"revision,omitempty"`
	Host             *string `protobuf:"bytes,3,opt,name=host" json:"host,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *ProblemHandle) Reset()         { *this = ProblemHandle{} }
func (this *ProblemHandle) String() string { return proto.CompactTextString(this) }
func (*ProblemHandle) ProtoMessage()       {}

func (this *ProblemHandle) GetUrl() string {
	if this != nil && this.Url != nil {
		return *this.Url
	}
	return ""
}

func (this *ProblemHandle) GetRevision() uint32 {
	if this != nil && this.Revision != nil {
		return *this.Revision
	}
	return 0
}

func (this *ProblemHandle) GetHost() string {
	if this != nil && this.Host != nil {
		return *this.Host
	}
	return ""
}

type SanitizeRequest struct {
	Problem          *ProblemHandle `protobuf:"bytes,1,opt,name=problem" json:"problem,omitempty"`
	Pdb              *string        `protobuf:"bytes,2,opt,name=pdb" json:"pdb,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (this *SanitizeRequest) Reset()         { *this = SanitizeRequest{} }
func (this *SanitizeRequest) String() string { return proto.CompactTextString(this) }
func (*SanitizeRequest) ProtoMessage()       {}

func (this *SanitizeRequest) GetProblem() *ProblemHandle {
	if this != nil {
		return this.Problem
	}
	return nil
}

func (this *SanitizeRequest) GetPdb() string {
	if this != nil && this.Pdb != nil {
		return *this.Pdb
	}
	return ""
}

type ProblemManifest struct {
	ProblemXml       *Blob                        `protobuf:"bytes,1,opt,name=problemXml" json:"problemXml,omitempty"`
	CheckerHash      []byte                       `protobuf:"bytes,2,opt,name=checker_hash" json:"checker_hash,omitempty"`
	TestHashes       []*ProblemManifest_TestEntry `protobuf:"bytes,3,rep,name=test_hashes" json:"test_hashes,omitempty"`
	XXX_unrecognized []byte                       `json:"-"`
}

func (this *ProblemManifest) Reset()         { *this = ProblemManifest{} }
func (this *ProblemManifest) String() string { return proto.CompactTextString(this) }
func (*ProblemManifest) ProtoMessage()       {}

func (this *ProblemManifest) GetProblemXml() *Blob {
	if this != nil {
		return this.ProblemXml
	}
	return nil
}

func (this *ProblemManifest) GetCheckerHash() []byte {
	if this != nil {
		return this.CheckerHash
	}
	return nil
}

type ProblemManifest_TestEntry struct {
	TestId           *uint32 `protobuf:"varint,1,opt,name=test_id" json:"test_id,omitempty"`
	InputHash        []byte  `protobuf:"bytes,2,opt,name=input_hash" json:"input_hash,omitempty"`
	AnswerHash       []byte  `protobuf:"bytes,3,opt,name=answer_hash" json:"answer_hash,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *ProblemManifest_TestEntry) Reset()         { *this = ProblemManifest_TestEntry{} }
func (this *ProblemManifest_TestEntry) String() string { return proto.CompactTextString(this) }
func (*ProblemManifest_TestEntry) ProtoMessage()       {}

func (this *ProblemManifest_TestEntry) GetTestId() uint32 {
	if this != nil && this.TestId != nil {
		return *this.TestId
	}
	return 0
}

func (this *ProblemManifest_TestEntry) GetInputHash() []byte {
	if this != nil {
		return this.InputHash
	}
	return nil
}

func (this *ProblemManifest_TestEntry) GetAnswerHash() []byte {
	if this != nil {
		return this.AnswerHash
	}
	return nil
}

type TestHandle struct {
	Pdb              *string        `protobuf:"bytes,1,opt,name=pdb" json:"pdb,omitempty"`
	Problem          *ProblemHandle `protobuf:"bytes,2,opt,name=problem" json:"problem,omitempty"`
	TestId           *uint32        `protobuf:"varint,3,opt,name=test_id" json:"test_id,omitempty"`
	TimeLimitMs      *uint32        `protobuf:"varint,4,opt,name=time_limit_ms" json:"time_limit_ms,omitempty"`
	MemoryLimitBytes *uint32        `protobuf:"varint,5,opt,name=memory_limit_bytes" json:"memory_limit_bytes,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (this *TestHandle) Reset()         { *this = TestHandle{} }
func (this *TestHandle) String() string { return proto.CompactTextString(this) }
func (*TestHandle) ProtoMessage()       {}

func (this *TestHandle) GetPdb() string {
	if this != nil && this.Pdb != nil {
		return *this.Pdb
	}
	return ""
}

func (this *TestHandle) GetProblem() *ProblemHandle {
	if this != nil {
		return this.Problem
	}
	return nil
}

func (this *TestHandle) GetTestId() uint32 {
	if this != nil && this.TestId != nil {
		return *this.TestId
	}
	return 0
}

func (this *TestHandle) GetTimeLimitMs() uint32 {
	if this != nil && this.TimeLimitMs != nil {
		return *this.TimeLimitMs
	}
	return 0
}

func (this *TestHandle) GetMemoryLimitBytes() uint32 {
	if this != nil && this.MemoryLimitBytes != nil {
		return *this.MemoryLimitBytes
	}
	return 0
}

type TestData struct {
	Input            *Blob   `protobuf:"bytes,1,opt,name=input" json:"input,omitempty"`
	Output           *Blob   `protobuf:"bytes,2,opt,name=output" json:"output,omitempty"`
	Checked          *Module `protobuf:"bytes,3,opt,name=checked" json:"checked,omitempty"`
	TimeLimitMs      *uint32 `protobuf:"varint,4,opt,name=time_limit_ms" json:"time_limit_ms,omitempty"`
	MemoryLimitBytes *uint32 `protobuf:"varint,5,opt,name=memory_limit_bytes" json:"memory_limit_bytes,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *TestData) Reset()         { *this = TestData{} }
func (this *TestData) String() string { return proto.CompactTextString(this) }
func (*TestData) ProtoMessage()       {}

func (this *TestData) GetInput() *Blob {
	if this != nil {
		return this.Input
	}
	return nil
}

func (this *TestData) GetOutput() *Blob {
	if this != nil {
		return this.Output
	}
	return nil
}

func (this *TestData) GetChecked() *Module {
	if this != nil {
		return this.Checked
	}
	return nil
}

func (this *TestData) GetTimeLimitMs() uint32 {
	if this != nil && this.TimeLimitMs != nil {
		return *this.TimeLimitMs
	}
	return 0
}

func (this *TestData) GetMemoryLimitBytes() uint32 {
	if this != nil && this.MemoryLimitBytes != nil {
		return *this.MemoryLimitBytes
	}
	return 0
}

type TestRequest struct {
	Module           *Module     `protobuf:"bytes,1,req,name=module" json:"module,omitempty"`
	TestHandle       *TestHandle `protobuf:"bytes,2,opt,name=test_handle" json:"test_handle,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (this *TestRequest) Reset()         { *this = TestRequest{} }
func (this *TestRequest) String() string { return proto.CompactTextString(this) }
func (*TestRequest) ProtoMessage()       {}

func (this *TestRequest) GetModule() *Module {
	if this != nil {
		return this.Module
	}
	return nil
}

func (this *TestRequest) GetTestHandle() *TestHandle {
	if this != nil {
		return this.TestHandle
	}
	return nil
}

type TestResponse struct {
	TestResult       *TestRun `protobuf:"bytes,1,opt,name=test_result" json:"test_result,omitempty"`
	TestNotFound     *bool    `protobuf:"varint,2,opt,name=test_not_found" json:"test_not_found,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (this *TestResponse) Reset()         { *this = TestResponse{} }
func (this *TestResponse) String() string { return proto.CompactTextString(this) }
func (*TestResponse) ProtoMessage()       {}

func (this *TestResponse) GetTestResult() *TestRun {
	if this != nil {
		return this.TestResult
	}
	return nil
}

func (this *TestResponse) GetTestNotFound() bool {
	if this != nil && this.TestNotFound != nil {
		return *this.TestNotFound
	}
	return false
}

type JudgeRequest struct {
	Module           *Module `protobuf:"bytes,1,opt,name=module" json:"module,omitempty"`
	ContestId        *int32  `protobuf:"varint,2,opt,name=contest_id" json:"contest_id,omitempty"`
	ProblemId        *string `protobuf:"bytes,3,opt,name=problem_id" json:"problem_id,omitempty"`
	SchoolMode       *bool   `protobuf:"varint,4,opt,name=school_mode" json:"school_mode,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *JudgeRequest) Reset()         { *this = JudgeRequest{} }
func (this *JudgeRequest) String() string { return proto.CompactTextString(this) }
func (*JudgeRequest) ProtoMessage()       {}

func (this *JudgeRequest) GetModule() *Module {
	if this != nil {
		return this.Module
	}
	return nil
}

func (this *JudgeRequest) GetContestId() int32 {
	if this != nil && this.ContestId != nil {
		return *this.ContestId
	}
	return 0
}

func (this *JudgeRequest) GetProblemId() string {
	if this != nil && this.ProblemId != nil {
		return *this.ProblemId
	}
	return ""
}

func (this *JudgeRequest) GetSchoolMode() bool {
	if this != nil && this.SchoolMode != nil {
		return *this.SchoolMode
	}
	return false
}

type JudgeResponse struct {
	ProblemNotFound  *bool  `protobuf:"varint,1,opt,name=problem_not_found" json:"problem_not_found,omitempty"`
	CompileFailed    *bool  `protobuf:"varint,2,opt,name=compile_failed" json:"compile_failed,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (this *JudgeResponse) Reset()         { *this = JudgeResponse{} }
func (this *JudgeResponse) String() string { return proto.CompactTextString(this) }
func (*JudgeResponse) ProtoMessage()       {}

func (this *JudgeResponse) GetProblemNotFound() bool {
	if this != nil && this.ProblemNotFound != nil {
		return *this.ProblemNotFound
	}
	return false
}

func (this *JudgeResponse) GetCompileFailed() bool {
	if this != nil && this.CompileFailed != nil {
		return *this.CompileFailed
	}
	return false
}

type ContestMapping struct {
	LocalId          *uint32 `protobuf:"varint,1,opt,name=local_id" json:"local_id,omitempty"`
	PolygonId        *uint32 `protobuf:"varint,2,opt,name=polygon_id" json:"polygon_id,omitempty"`
	SchoolMode       *bool   `protobuf:"varint,3,opt,name=school_mode" json:"school_mode,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *ContestMapping) Reset()         { *this = ContestMapping{} }
func (this *ContestMapping) String() string { return proto.CompactTextString(this) }
func (*ContestMapping) ProtoMessage()       {}

func (this *ContestMapping) GetLocalId() uint32 {
	if this != nil && this.LocalId != nil {
		return *this.LocalId
	}
	return 0
}

func (this *ContestMapping) GetPolygonId() uint32 {
	if this != nil && this.PolygonId != nil {
		return *this.PolygonId
	}
	return 0
}

func (this *ContestMapping) GetSchoolMode() bool {
	if this != nil && this.SchoolMode != nil {
		return *this.SchoolMode
	}
	return false
}

func init() {
	proto.RegisterEnum("contester.proto.TestRun_Code", TestRun_Code_name, TestRun_Code_value)
	proto.RegisterEnum("contester.proto.Compilation_Code", Compilation_Code_name, Compilation_Code_value)
}