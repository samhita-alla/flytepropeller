package validators

import (
	"fmt"

	"github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/core"
	c "github.com/flyteorg/flytepropeller/pkg/compiler/common"
	"github.com/flyteorg/flytepropeller/pkg/compiler/errors"
)

// Validate interface has its required attributes set
func ValidateInterface(nodeID c.NodeID, iface *core.TypedInterface, errs errors.CompileErrors) (
	typedInterface *core.TypedInterface, ok bool) {

	if iface == nil {
		iface = &core.TypedInterface{}
	}

	// validate InputsRef/OutputsRef parameters required attributes are set
	if iface.Inputs != nil && iface.Inputs.Variables != nil {
		validateVariables(nodeID, iface.Inputs, errs.NewScope())
	} else {
		iface.Inputs = &core.VariableMap{Variables: map[string]*core.Variable{}}
	}

	if iface.Outputs != nil && iface.Outputs.Variables != nil {
		validateVariables(nodeID, iface.Outputs, errs.NewScope())
	} else {
		iface.Outputs = &core.VariableMap{Variables: map[string]*core.Variable{}}
	}

	return iface, !errs.HasErrors()
}

// Validates underlying interface of a node and returns the effective Typed Interface.
func ValidateUnderlyingInterface(w c.WorkflowBuilder, node c.NodeBuilder, errs errors.CompileErrors) (iface *core.TypedInterface, ok bool) {
	switch node.GetCoreNode().GetTarget().(type) {
	case *core.Node_TaskNode:
		if node.GetTaskNode().GetReferenceId() == nil {
			errs.Collect(errors.NewValueRequiredErr(node.GetId(), "TaskNode.ReferenceId"))
		} else if task, taskOk := w.GetTask(*node.GetTaskNode().GetReferenceId()); taskOk {
			iface = task.GetInterface()
			if iface == nil {
				// Default value for no interface is nil, initialize an empty interface
				iface = &core.TypedInterface{
					Inputs:  &core.VariableMap{Variables: map[string]*core.Variable{}},
					Outputs: &core.VariableMap{Variables: map[string]*core.Variable{}},
				}
			}
		} else {
			errs.Collect(errors.NewTaskReferenceNotFoundErr(node.GetId(), node.GetTaskNode().GetReferenceId().String()))
		}
	case *core.Node_WorkflowNode:
		if node.GetWorkflowNode().GetLaunchplanRef().String() == w.GetCoreWorkflow().Template.Id.String() {
			iface = w.GetCoreWorkflow().Template.Interface
			if iface == nil {
				errs.Collect(errors.NewValueRequiredErr(node.GetId(), "WorkflowNode.Interface"))
			}
		} else if node.GetWorkflowNode().GetLaunchplanRef() != nil {
			if launchPlan, launchPlanOk := w.GetLaunchPlan(*node.GetWorkflowNode().GetLaunchplanRef()); launchPlanOk {
				inputs := launchPlan.GetExpectedInputs()
				if inputs == nil {
					errs.Collect(errors.NewValueRequiredErr(node.GetId(), "WorkflowNode.ExpectedInputs"))
				}

				outputs := launchPlan.GetExpectedOutputs()
				if outputs == nil {
					errs.Collect(errors.NewValueRequiredErr(node.GetId(), "WorkflowNode.ExpectedOutputs"))
				}

				// Compute exposed inputs as the union of all required inputs and any input overwritten by the node.
				exposedInputs := map[string]*core.Variable{}
				if inputs != nil && inputs.Parameters != nil {
					for name, p := range inputs.Parameters {
						if p.GetRequired() {
							exposedInputs[name] = p.Var
						} else if containsBindingByVariableName(node.GetInputs(), name) {
							exposedInputs[name] = p.Var
						}
						// else, the param has a default value and is not being overwritten by the node
					}
				}

				iface = &core.TypedInterface{
					Inputs: &core.VariableMap{
						Variables: exposedInputs,
					},
					Outputs: outputs,
				}
			} else {
				errs.Collect(errors.NewWorkflowReferenceNotFoundErr(
					node.GetId(),
					fmt.Sprintf("%v", node.GetWorkflowNode().GetLaunchplanRef())))
			}
		} else if node.GetWorkflowNode().GetSubWorkflowRef() != nil {
			if wf, wfOk := w.GetSubWorkflow(*node.GetWorkflowNode().GetSubWorkflowRef()); wfOk {
				if wf.Template == nil {
					errs.Collect(errors.NewValueRequiredErr(node.GetId(), "WorkflowNode.Template"))
				} else {
					iface = wf.Template.Interface
					if iface == nil {
						errs.Collect(errors.NewValueRequiredErr(node.GetId(), "WorkflowNode.Template.Interface"))
					}
				}
			} else {
				errs.Collect(errors.NewWorkflowReferenceNotFoundErr(
					node.GetId(),
					fmt.Sprintf("%v", node.GetWorkflowNode().GetSubWorkflowRef())))
			}
		} else {
			errs.Collect(errors.NewWorkflowReferenceNotFoundErr(
				node.GetId(),
				fmt.Sprintf("%v/%v", node.GetWorkflowNode().GetLaunchplanRef(), node.GetWorkflowNode().GetSubWorkflowRef())))
		}
	case *core.Node_BranchNode:
		iface, _ = validateBranchInterface(w, node, errs.NewScope())
	default:
		errs.Collect(errors.NewValueRequiredErr(node.GetId(), "Target"))
	}

	if iface != nil {
		ValidateInterface(node.GetId(), iface, errs.NewScope())
	}

	if !errs.HasErrors() {
		node.SetInterface(iface)
	}

	return iface, !errs.HasErrors()
}
