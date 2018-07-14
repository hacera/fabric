// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/hyperledger/fabric/core/common/ccprovider"
	pb "github.com/hyperledger/fabric/protos/peer"
	"golang.org/x/net/context"
)

type Lifecycle struct {
	GetChaincodeDefinitionStub        func(ctx context.Context, txid string, signedProp *pb.SignedProposal, prop *pb.Proposal, chainID string, chaincodeID string) (ccprovider.ChaincodeDefinition, error)
	getChaincodeDefinitionMutex       sync.RWMutex
	getChaincodeDefinitionArgsForCall []struct {
		ctx         context.Context
		txid        string
		signedProp  *pb.SignedProposal
		prop        *pb.Proposal
		chainID     string
		chaincodeID string
	}
	getChaincodeDefinitionReturns struct {
		result1 ccprovider.ChaincodeDefinition
		result2 error
	}
	getChaincodeDefinitionReturnsOnCall map[int]struct {
		result1 ccprovider.ChaincodeDefinition
		result2 error
	}
	GetChaincodeDeploymentSpecStub        func(ctx context.Context, txid string, signedProp *pb.SignedProposal, prop *pb.Proposal, chainID string, chaincodeID string) (*pb.ChaincodeDeploymentSpec, error)
	getChaincodeDeploymentSpecMutex       sync.RWMutex
	getChaincodeDeploymentSpecArgsForCall []struct {
		ctx         context.Context
		txid        string
		signedProp  *pb.SignedProposal
		prop        *pb.Proposal
		chainID     string
		chaincodeID string
	}
	getChaincodeDeploymentSpecReturns struct {
		result1 *pb.ChaincodeDeploymentSpec
		result2 error
	}
	getChaincodeDeploymentSpecReturnsOnCall map[int]struct {
		result1 *pb.ChaincodeDeploymentSpec
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Lifecycle) GetChaincodeDefinition(ctx context.Context, txid string, signedProp *pb.SignedProposal, prop *pb.Proposal, chainID string, chaincodeID string) (ccprovider.ChaincodeDefinition, error) {
	fake.getChaincodeDefinitionMutex.Lock()
	ret, specificReturn := fake.getChaincodeDefinitionReturnsOnCall[len(fake.getChaincodeDefinitionArgsForCall)]
	fake.getChaincodeDefinitionArgsForCall = append(fake.getChaincodeDefinitionArgsForCall, struct {
		ctx         context.Context
		txid        string
		signedProp  *pb.SignedProposal
		prop        *pb.Proposal
		chainID     string
		chaincodeID string
	}{ctx, txid, signedProp, prop, chainID, chaincodeID})
	fake.recordInvocation("GetChaincodeDefinition", []interface{}{ctx, txid, signedProp, prop, chainID, chaincodeID})
	fake.getChaincodeDefinitionMutex.Unlock()
	if fake.GetChaincodeDefinitionStub != nil {
		return fake.GetChaincodeDefinitionStub(ctx, txid, signedProp, prop, chainID, chaincodeID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getChaincodeDefinitionReturns.result1, fake.getChaincodeDefinitionReturns.result2
}

func (fake *Lifecycle) GetChaincodeDefinitionCallCount() int {
	fake.getChaincodeDefinitionMutex.RLock()
	defer fake.getChaincodeDefinitionMutex.RUnlock()
	return len(fake.getChaincodeDefinitionArgsForCall)
}

func (fake *Lifecycle) GetChaincodeDefinitionArgsForCall(i int) (context.Context, string, *pb.SignedProposal, *pb.Proposal, string, string) {
	fake.getChaincodeDefinitionMutex.RLock()
	defer fake.getChaincodeDefinitionMutex.RUnlock()
	return fake.getChaincodeDefinitionArgsForCall[i].ctx, fake.getChaincodeDefinitionArgsForCall[i].txid, fake.getChaincodeDefinitionArgsForCall[i].signedProp, fake.getChaincodeDefinitionArgsForCall[i].prop, fake.getChaincodeDefinitionArgsForCall[i].chainID, fake.getChaincodeDefinitionArgsForCall[i].chaincodeID
}

func (fake *Lifecycle) GetChaincodeDefinitionReturns(result1 ccprovider.ChaincodeDefinition, result2 error) {
	fake.GetChaincodeDefinitionStub = nil
	fake.getChaincodeDefinitionReturns = struct {
		result1 ccprovider.ChaincodeDefinition
		result2 error
	}{result1, result2}
}

func (fake *Lifecycle) GetChaincodeDefinitionReturnsOnCall(i int, result1 ccprovider.ChaincodeDefinition, result2 error) {
	fake.GetChaincodeDefinitionStub = nil
	if fake.getChaincodeDefinitionReturnsOnCall == nil {
		fake.getChaincodeDefinitionReturnsOnCall = make(map[int]struct {
			result1 ccprovider.ChaincodeDefinition
			result2 error
		})
	}
	fake.getChaincodeDefinitionReturnsOnCall[i] = struct {
		result1 ccprovider.ChaincodeDefinition
		result2 error
	}{result1, result2}
}

func (fake *Lifecycle) GetChaincodeDeploymentSpec(ctx context.Context, txid string, signedProp *pb.SignedProposal, prop *pb.Proposal, chainID string, chaincodeID string) (*pb.ChaincodeDeploymentSpec, error) {
	fake.getChaincodeDeploymentSpecMutex.Lock()
	ret, specificReturn := fake.getChaincodeDeploymentSpecReturnsOnCall[len(fake.getChaincodeDeploymentSpecArgsForCall)]
	fake.getChaincodeDeploymentSpecArgsForCall = append(fake.getChaincodeDeploymentSpecArgsForCall, struct {
		ctx         context.Context
		txid        string
		signedProp  *pb.SignedProposal
		prop        *pb.Proposal
		chainID     string
		chaincodeID string
	}{ctx, txid, signedProp, prop, chainID, chaincodeID})
	fake.recordInvocation("GetChaincodeDeploymentSpec", []interface{}{ctx, txid, signedProp, prop, chainID, chaincodeID})
	fake.getChaincodeDeploymentSpecMutex.Unlock()
	if fake.GetChaincodeDeploymentSpecStub != nil {
		return fake.GetChaincodeDeploymentSpecStub(ctx, txid, signedProp, prop, chainID, chaincodeID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getChaincodeDeploymentSpecReturns.result1, fake.getChaincodeDeploymentSpecReturns.result2
}

func (fake *Lifecycle) GetChaincodeDeploymentSpecCallCount() int {
	fake.getChaincodeDeploymentSpecMutex.RLock()
	defer fake.getChaincodeDeploymentSpecMutex.RUnlock()
	return len(fake.getChaincodeDeploymentSpecArgsForCall)
}

func (fake *Lifecycle) GetChaincodeDeploymentSpecArgsForCall(i int) (context.Context, string, *pb.SignedProposal, *pb.Proposal, string, string) {
	fake.getChaincodeDeploymentSpecMutex.RLock()
	defer fake.getChaincodeDeploymentSpecMutex.RUnlock()
	return fake.getChaincodeDeploymentSpecArgsForCall[i].ctx, fake.getChaincodeDeploymentSpecArgsForCall[i].txid, fake.getChaincodeDeploymentSpecArgsForCall[i].signedProp, fake.getChaincodeDeploymentSpecArgsForCall[i].prop, fake.getChaincodeDeploymentSpecArgsForCall[i].chainID, fake.getChaincodeDeploymentSpecArgsForCall[i].chaincodeID
}

func (fake *Lifecycle) GetChaincodeDeploymentSpecReturns(result1 *pb.ChaincodeDeploymentSpec, result2 error) {
	fake.GetChaincodeDeploymentSpecStub = nil
	fake.getChaincodeDeploymentSpecReturns = struct {
		result1 *pb.ChaincodeDeploymentSpec
		result2 error
	}{result1, result2}
}

func (fake *Lifecycle) GetChaincodeDeploymentSpecReturnsOnCall(i int, result1 *pb.ChaincodeDeploymentSpec, result2 error) {
	fake.GetChaincodeDeploymentSpecStub = nil
	if fake.getChaincodeDeploymentSpecReturnsOnCall == nil {
		fake.getChaincodeDeploymentSpecReturnsOnCall = make(map[int]struct {
			result1 *pb.ChaincodeDeploymentSpec
			result2 error
		})
	}
	fake.getChaincodeDeploymentSpecReturnsOnCall[i] = struct {
		result1 *pb.ChaincodeDeploymentSpec
		result2 error
	}{result1, result2}
}

func (fake *Lifecycle) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getChaincodeDefinitionMutex.RLock()
	defer fake.getChaincodeDefinitionMutex.RUnlock()
	fake.getChaincodeDeploymentSpecMutex.RLock()
	defer fake.getChaincodeDeploymentSpecMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Lifecycle) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}