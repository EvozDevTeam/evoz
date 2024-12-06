package vm

import (
	"github.com/ethereum/go-ethereum/common"
)


func opVulnerableStaticCall(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	// Pop gas. The actual gas is in interpreter.evm.callGasTemp.
	stack := scope.Stack
	// We use it as a temporary value
	temp := stack.pop()
	gas := interpreter.evm.callGasTemp
	// Pop other call parameters.
	addr, inOffset, inSize, retOffset, retSize := stack.pop(), stack.pop(), stack.pop(), stack.pop(), stack.pop()
	toAddr := common.Address(addr.Bytes20())
	// Get arguments from the memory.
	args := scope.Memory.GetPtr(int64(inOffset.Uint64()), int64(inSize.Uint64()))

	ret, returnGas, err := interpreter.evm.StaticCall(scope.Contract, toAddr, args, gas)
	if err != nil {
		temp.Clear()
	} else {
		temp.SetOne()
	}
	stack.push(&temp)
	if err == nil || err == ErrExecutionReverted {
		// TX 0xe42c11a8a31f3c0a990a8264b23bf4e936b9d97f3242cfbc21e63b2b0abd09f0 would result
		// in a differenct contract runtime code with the following commented line.
		// ret = common.CopyBytes(ret)
		scope.Memory.Set(retOffset.Uint64(), retSize.Uint64(), ret)
	}
	scope.Contract.Gas += returnGas

	interpreter.returnData = common.CopyBytes(ret)
	return ret, nil
}

func UseVulnerableCalls(jt *JumpTable) {
	jt[STATICCALL].execute = opVulnerableStaticCall
}
