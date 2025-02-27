package validators

import (
	"testing"

	"github.com/flyteorg/flyteidl/clients/go/coreutils"
	"github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/core"
	c "github.com/flyteorg/flytepropeller/pkg/compiler/common"
	"github.com/flyteorg/flytepropeller/pkg/compiler/common/mocks"
	"github.com/flyteorg/flytepropeller/pkg/compiler/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestValidateInterface(t *testing.T) {
	t.Run("Happy path", func(t *testing.T) {
		errs := errors.NewCompileErrors()
		iface, ok := ValidateInterface(
			c.NodeID("node1"),
			&core.TypedInterface{
				Inputs: &core.VariableMap{
					Variables: map[string]*core.Variable{},
				},
				Outputs: &core.VariableMap{
					Variables: map[string]*core.Variable{},
				},
			},
			errs.NewScope(),
		)

		assertNonEmptyInterface(t, iface, ok, errs)
	})

	t.Run("Empty Inputs/Outputs", func(t *testing.T) {
		errs := errors.NewCompileErrors()
		iface, ok := ValidateInterface(
			c.NodeID("node1"),
			&core.TypedInterface{},
			errs.NewScope(),
		)

		assertNonEmptyInterface(t, iface, ok, errs)
	})

	t.Run("Empty Interface", func(t *testing.T) {
		errs := errors.NewCompileErrors()
		iface, ok := ValidateInterface(
			c.NodeID("node1"),
			nil,
			errs.NewScope(),
		)

		assertNonEmptyInterface(t, iface, ok, errs)
	})
}

func assertNonEmptyInterface(t testing.TB, iface *core.TypedInterface, ifaceOk bool, errs errors.CompileErrors) {
	assert.True(t, ifaceOk)
	assert.NotNil(t, iface)
	assert.False(t, errs.HasErrors())
	if !ifaceOk {
		t.Fatal(errs)
	}

	assert.NotNil(t, iface.Inputs)
	assert.NotNil(t, iface.Inputs.Variables)
	assert.NotNil(t, iface.Outputs)
	assert.NotNil(t, iface.Outputs.Variables)
}

func TestValidateUnderlyingInterface(t *testing.T) {
	t.Run("Invalid empty node", func(t *testing.T) {
		wfBuilder := mocks.WorkflowBuilder{}
		nodeBuilder := mocks.NodeBuilder{}
		nodeBuilder.On("GetCoreNode").Return(&core.Node{})
		nodeBuilder.On("GetId").Return("node_1")
		errs := errors.NewCompileErrors()
		iface, ifaceOk := ValidateUnderlyingInterface(&wfBuilder, &nodeBuilder, errs.NewScope())
		assert.False(t, ifaceOk)
		assert.Nil(t, iface)
		assert.True(t, errs.HasErrors())
	})

	t.Run("Task Node", func(t *testing.T) {
		task := mocks.Task{}
		task.On("GetInterface").Return(nil)

		wfBuilder := mocks.WorkflowBuilder{}
		wfBuilder.On("GetTask", mock.MatchedBy(func(id core.Identifier) bool {
			return id.String() == (&core.Identifier{
				Name: "Task_1",
			}).String()
		})).Return(&task, true)

		taskNode := &core.TaskNode{
			Reference: &core.TaskNode_ReferenceId{
				ReferenceId: &core.Identifier{
					Name: "Task_1",
				},
			},
		}

		nodeBuilder := mocks.NodeBuilder{}
		nodeBuilder.On("GetCoreNode").Return(&core.Node{
			Target: &core.Node_TaskNode{
				TaskNode: taskNode,
			},
		})

		nodeBuilder.On("GetTaskNode").Return(taskNode)
		nodeBuilder.On("GetId").Return("node_1")
		nodeBuilder.On("SetInterface", mock.Anything).Return()

		errs := errors.NewCompileErrors()
		iface, ifaceOk := ValidateUnderlyingInterface(&wfBuilder, &nodeBuilder, errs.NewScope())
		assertNonEmptyInterface(t, iface, ifaceOk, errs)
	})

	t.Run("Workflow Node", func(t *testing.T) {
		wfBuilder := mocks.WorkflowBuilder{}
		wfBuilder.On("GetCoreWorkflow").Return(&core.CompiledWorkflow{
			Template: &core.WorkflowTemplate{
				Id: &core.Identifier{
					Name: "Ref_1",
				},
			},
		})
		workflowNode := &core.WorkflowNode{
			Reference: &core.WorkflowNode_LaunchplanRef{
				LaunchplanRef: &core.Identifier{
					Name: "Ref_1",
				},
			},
		}

		nodeBuilder := mocks.NodeBuilder{}
		nodeBuilder.On("GetCoreNode").Return(&core.Node{
			Target: &core.Node_WorkflowNode{
				WorkflowNode: workflowNode,
			},
		})

		nodeBuilder.On("GetWorkflowNode").Return(workflowNode)
		nodeBuilder.On("GetId").Return("node_1")
		nodeBuilder.On("SetInterface", mock.Anything).Return()
		nodeBuilder.On("GetInputs").Return([]*core.Binding{})

		t.Run("Self", func(t *testing.T) {
			errs := errors.NewCompileErrors()
			_, ifaceOk := ValidateUnderlyingInterface(&wfBuilder, &nodeBuilder, errs.NewScope())
			assert.False(t, ifaceOk)

			wfBuilder := mocks.WorkflowBuilder{}
			wfBuilder.On("GetCoreWorkflow").Return(&core.CompiledWorkflow{
				Template: &core.WorkflowTemplate{
					Id: &core.Identifier{
						Name: "Ref_1",
					},
					Interface: &core.TypedInterface{
						Inputs: &core.VariableMap{
							Variables: map[string]*core.Variable{},
						},
						Outputs: &core.VariableMap{
							Variables: map[string]*core.Variable{},
						},
					},
				},
			})

			errs = errors.NewCompileErrors()
			iface, ifaceOk := ValidateUnderlyingInterface(&wfBuilder, &nodeBuilder, errs.NewScope())
			assertNonEmptyInterface(t, iface, ifaceOk, errs)
		})

		t.Run("LP_Ref", func(t *testing.T) {
			lp := mocks.InterfaceProvider{}
			lp.On("GetID").Return(&core.Identifier{Name: "Ref_1"})
			lp.On("GetExpectedInputs").Return(&core.ParameterMap{
				Parameters: map[string]*core.Parameter{
					"required": {
						Var: &core.Variable{
							Type: &core.LiteralType{Type: &core.LiteralType_Simple{Simple: core.SimpleType_INTEGER}},
						},
						Behavior: &core.Parameter_Required{
							Required: true,
						},
					},
					"default_value": {
						Var: &core.Variable{
							Type: &core.LiteralType{Type: &core.LiteralType_Simple{Simple: core.SimpleType_INTEGER}},
						},
						Behavior: &core.Parameter_Default{
							Default: coreutils.MustMakeLiteral(5),
						},
					},
				},
			})
			lp.On("GetExpectedOutputs").Return(&core.VariableMap{})

			wfBuilder := mocks.WorkflowBuilder{}
			wfBuilder.On("GetCoreWorkflow").Return(&core.CompiledWorkflow{
				Template: &core.WorkflowTemplate{
					Id: &core.Identifier{
						Name: "Ref_2",
					},
				},
			})

			wfBuilder.On("GetLaunchPlan", mock.Anything).Return(nil, false)

			errs := errors.NewCompileErrors()
			_, ifaceOk := ValidateUnderlyingInterface(&wfBuilder, &nodeBuilder, errs.NewScope())
			assert.False(t, ifaceOk)

			wfBuilder = mocks.WorkflowBuilder{}
			wfBuilder.On("GetCoreWorkflow").Return(&core.CompiledWorkflow{
				Template: &core.WorkflowTemplate{
					Id: &core.Identifier{
						Name: "Ref_2",
					},
				},
			})

			wfBuilder.On("GetLaunchPlan", matchIdentifier(core.Identifier{Name: "Ref_1"})).Return(&lp, true)

			errs = errors.NewCompileErrors()
			iface, ifaceOk := ValidateUnderlyingInterface(&wfBuilder, &nodeBuilder, errs.NewScope())
			assertNonEmptyInterface(t, iface, ifaceOk, errs)
		})

		t.Run("Subwf", func(t *testing.T) {
			subWf := core.CompiledWorkflow{
				Template: &core.WorkflowTemplate{
					Interface: &core.TypedInterface{
						Inputs:  &core.VariableMap{},
						Outputs: &core.VariableMap{},
					},
				},
			}

			wfBuilder := mocks.WorkflowBuilder{}
			wfBuilder.On("GetCoreWorkflow").Return(&core.CompiledWorkflow{
				Template: &core.WorkflowTemplate{
					Id: &core.Identifier{
						Name: "Ref_2",
					},
				},
			})

			wfBuilder.On("GetLaunchPlan", mock.Anything).Return(nil, false)

			errs := errors.NewCompileErrors()
			_, ifaceOk := ValidateUnderlyingInterface(&wfBuilder, &nodeBuilder, errs.NewScope())
			assert.False(t, ifaceOk)

			wfBuilder = mocks.WorkflowBuilder{}
			wfBuilder.On("GetCoreWorkflow").Return(&core.CompiledWorkflow{
				Template: &core.WorkflowTemplate{
					Id: &core.Identifier{
						Name: "Ref_2",
					},
				},
			})

			wfBuilder.On("GetSubWorkflow", matchIdentifier(core.Identifier{Name: "Ref_1"})).Return(&subWf, true)

			workflowNode.Reference = &core.WorkflowNode_SubWorkflowRef{
				SubWorkflowRef: &core.Identifier{Name: "Ref_1"},
			}

			errs = errors.NewCompileErrors()
			iface, ifaceOk := ValidateUnderlyingInterface(&wfBuilder, &nodeBuilder, errs.NewScope())
			assertNonEmptyInterface(t, iface, ifaceOk, errs)
		})
	})
}

func matchIdentifier(id core.Identifier) interface{} {
	return mock.MatchedBy(func(arg core.Identifier) bool {
		return arg.String() == id.String()
	})
}
