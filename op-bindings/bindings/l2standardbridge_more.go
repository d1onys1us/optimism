// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const L2StandardBridgeStorageLayoutJSON = "{\"storage\":[{\"astId\":23577,\"contract\":\"contracts/L2/L2StandardBridge.sol:L2StandardBridge\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":23580,\"contract\":\"contracts/L2/L2StandardBridge.sol:L2StandardBridge\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":22664,\"contract\":\"contracts/L2/L2StandardBridge.sol:L2StandardBridge\",\"label\":\"messenger\",\"offset\":2,\"slot\":\"0\",\"type\":\"t_contract(CrossDomainMessenger)22047\"},{\"astId\":22668,\"contract\":\"contracts/L2/L2StandardBridge.sol:L2StandardBridge\",\"label\":\"otherBridge\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_contract(StandardBridge)23345\"},{\"astId\":22675,\"contract\":\"contracts/L2/L2StandardBridge.sol:L2StandardBridge\",\"label\":\"deposits\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_mapping(t_address,t_mapping(t_address,t_uint256))\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_contract(CrossDomainMessenger)22047\":{\"encoding\":\"inplace\",\"label\":\"contract CrossDomainMessenger\",\"numberOfBytes\":\"20\"},\"t_contract(StandardBridge)23345\":{\"encoding\":\"inplace\",\"label\":\"contract StandardBridge\",\"numberOfBytes\":\"20\"},\"t_mapping(t_address,t_mapping(t_address,t_uint256))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(address =\u003e uint256))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_address,t_uint256)\"},\"t_mapping(t_address,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_uint256\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L2StandardBridgeStorageLayout = new(solc.StorageLayout)

var L2StandardBridgeDeployedBin = "0x6080604052600436106100ec5760003560e01c8063662a633a1161008a578063af565a1311610059578063af565a131461032f578063c4d66de81461034f578063c89701a21461036f578063e11013dd1461039c57600080fd5b8063662a633a146102a357806387087623146102b65780638f601f66146102d6578063a3a795481461031c57600080fd5b806332b7006d116100c657806332b7006d146101f15780633cb747bf14610204578063540abf731461026157806354fd4d501461028157600080fd5b80630166a07a146101ab57806309fc8843146101cb5780631635f5fd146101de57600080fd5b366101a657333b15610185576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f4100000000000000000060648201526084015b60405180910390fd5b6101a433333462030d40604051806020016040528060008152506103cb565b005b600080fd5b3480156101b757600080fd5b506101a46101c636600461248e565b610567565b6101a46101d936600461253f565b6108b2565b6101a46101ec366004612592565b610989565b6101a46101ff366004612605565b610d6e565b34801561021057600080fd5b506000546102379062010000900473ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b34801561026d57600080fd5b506101a461027c366004612659565b610e13565b34801561028d57600080fd5b50610296610e23565b6040516102589190612746565b6101a46102b136600461248e565b610ec6565b3480156102c257600080fd5b506101a46102d1366004612759565b610fb3565b3480156102e257600080fd5b5061030e6102f13660046127dc565b600260209081526000928352604080842090915290825290205481565b604051908152602001610258565b6101a461032a366004612759565b611052565b34801561033b57600080fd5b506101a461034a366004612815565b611061565b34801561035b57600080fd5b506101a461036a366004612866565b611374565b34801561037b57600080fd5b506001546102379073ffffffffffffffffffffffffffffffffffffffff1681565b6101a46103aa366004612883565b61151d565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167f2849b43074093a05396b6f2a937dee8565b15a48a7b3d4bffb732a5017380af5858460405161042a9291906128e6565b60405180910390a360005460015460405173ffffffffffffffffffffffffffffffffffffffff62010000909304831692633dbb202b9287929116907f1635f5fd000000000000000000000000000000000000000000000000000000009061049b908b908b9086908a906024016128ff565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009485161790525160e086901b909216825261052e92918890600401612948565b6000604051808303818588803b15801561054757600080fd5b505af115801561055b573d6000803e3d6000fd5b50505050505050505050565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff16331480156106495750600154600054604080517f6e296e45000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff938416936201000090930490921691636e296e45916004808201926020929091908290030181865afa15801561060d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610631919061298d565b73ffffffffffffffffffffffffffffffffffffffff16145b6106fb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604160248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20746865206f7468657220627269646760648201527f6500000000000000000000000000000000000000000000000000000000000000608482015260a40161017c565b6040517faf565a1300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff808916600483015280881660248301528516604482015260648101849052309063af565a1390608401600060405180830381600087803b15801561077957600080fd5b505af192505050801561078a575060015b6108265761079f878786888760008888611560565b8473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff167f2755817676249910615f0a6a240ad225abe5343df8d527f7294c4af36a92009a8787878760405161081994939291906129f3565b60405180910390a46108a9565b8473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff167fd59c65b35445225835c83f50b6ede06a7be047d22e357073e250d9af537518cd878787876040516108a094939291906129f3565b60405180910390a45b50505050505050565b333b15610941576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f41000000000000000000606482015260840161017c565b6109843333348686868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506103cb92505050565b505050565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff1633148015610a6b5750600154600054604080517f6e296e45000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff938416936201000090930490921691636e296e45916004808201926020929091908290030181865afa158015610a2f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a53919061298d565b73ffffffffffffffffffffffffffffffffffffffff16145b610b1d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604160248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20746865206f7468657220627269646760648201527f6500000000000000000000000000000000000000000000000000000000000000608482015260a40161017c565b823414610bac576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603a60248201527f5374616e646172644272696467653a20616d6f756e742073656e7420646f657360448201527f206e6f74206d6174636820616d6f756e74207265717569726564000000000000606482015260840161017c565b3073ffffffffffffffffffffffffffffffffffffffff851603610c51576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f5374616e646172644272696467653a2063616e6e6f742073656e6420746f207360448201527f656c660000000000000000000000000000000000000000000000000000000000606482015260840161017c565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167f31b2166ff604fc5672ea5df08a78081d2bc6d746cadce880747f3643d819e83d858585604051610cb293929190612a29565b60405180910390a36000610cd7855a866040518060200160405280600081525061171d565b905080610d66576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f5374616e646172644272696467653a20455448207472616e736665722066616960448201527f6c65640000000000000000000000000000000000000000000000000000000000606482015260840161017c565b505050505050565b333b15610dfd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f41000000000000000000606482015260840161017c565b610e0c85333387878787611737565b5050505050565b6108a9878733888888888861196e565b6060610e4e7f0000000000000000000000000000000000000000000000000000000000000000611c18565b610e777f0000000000000000000000000000000000000000000000000000000000000000611c18565b610ea07f0000000000000000000000000000000000000000000000000000000000000000611c18565b604051602001610eb293929190612a4c565b604051602081830303815290604052905090565b73ffffffffffffffffffffffffffffffffffffffff8716158015610f13575073ffffffffffffffffffffffffffffffffffffffff861673deaddeaddeaddeaddeaddeaddeaddeaddead0000145b15610f2a57610f258585858585610989565b610f39565b610f3986888787878787610567565b8473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff167fb0444523268717a02698be47d0803aa7468c00acbed2f8bd93a0459cde61dd89878787876040516108a094939291906129f3565b333b15611042576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f41000000000000000000606482015260840161017c565b610d66868633338888888861196e565b610d6686338787878787611737565b3330146110f0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603360248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642062792073656c6600000000000000000000000000606482015260840161017c565b3073ffffffffffffffffffffffffffffffffffffffff851603611195576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5374616e646172644272696467653a206c6f63616c20746f6b656e2063616e6e60448201527f6f742062652073656c6600000000000000000000000000000000000000000000606482015260840161017c565b61119e84611d55565b156112ec576111ad8484611d87565b61125f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604a60248201527f5374616e646172644272696467653a2077726f6e672072656d6f746520746f6b60448201527f656e20666f72204f7074696d69736d204d696e7461626c65204552433230206c60648201527f6f63616c20746f6b656e00000000000000000000000000000000000000000000608482015260a40161017c565b6040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8381166004830152602482018390528516906340c10f1990604401600060405180830381600087803b1580156112cf57600080fd5b505af11580156112e3573d6000803e3d6000fd5b5050505061136e565b73ffffffffffffffffffffffffffffffffffffffff80851660009081526002602090815260408083209387168352929052205461132a908290612af1565b73ffffffffffffffffffffffffffffffffffffffff80861660008181526002602090815260408083209489168352939052919091209190915561136e908383611e2e565b50505050565b600054610100900460ff16158080156113945750600054600160ff909116105b806113ae5750303b1580156113ae575060005460ff166001145b61143a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161017c565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561149857600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b6114b673420000000000000000000000000000000000000783611f02565b801561151957600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b61136e3385348686868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506103cb92505050565b60005460015460405173ffffffffffffffffffffffffffffffffffffffff62010000909304831692633dbb202b9216907f0166a07a00000000000000000000000000000000000000000000000000000000906115cc908c908e908d908d908d908c908c90602401612b08565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009485161790525160e085901b909216825261165f92918890600401612948565b600060405180830381600087803b15801561167957600080fd5b505af115801561168d573d6000803e3d6000fd5b505050508573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff167f7ff126db8024424bbfd9826e8ab82ff59136289ea440b04b39a0df1b03b9cabf8888878760405161170b94939291906129f3565b60405180910390a45050505050505050565b600080600080845160208601878a8af19695505050505050565b60008773ffffffffffffffffffffffffffffffffffffffff1663c01e1bd66040518163ffffffff1660e01b8152600401602060405180830381865afa158015611784573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117a8919061298d565b90507fffffffffffffffffffffffff215221522152215221522152215221522153000073ffffffffffffffffffffffffffffffffffffffff8916016118e45784341461189c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604360248201527f4c325374616e646172644272696467653a20455448207769746864726177616c60448201527f73206d75737420696e636c7564652073756666696369656e742045544820766160648201527f6c75650000000000000000000000000000000000000000000000000000000000608482015260a40161017c565b6118df8787878787878080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506103cb92505050565b6118f4565b6118f4888289898989898961196e565b8673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167f73d170910aba9e6d50b102db522b1dbcd796216f5128b445aa2135272886497e8989888860405161170b94939291906129f3565b3073ffffffffffffffffffffffffffffffffffffffff891603611a13576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5374616e646172644272696467653a206c6f63616c20746f6b656e2063616e6e60448201527f6f742062652073656c6600000000000000000000000000000000000000000000606482015260840161017c565b611a1c88611d55565b15611b6a57611a2b8888611d87565b611add576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604a60248201527f5374616e646172644272696467653a2077726f6e672072656d6f746520746f6b60448201527f656e20666f72204f7074696d69736d204d696e7461626c65204552433230206c60648201527f6f63616c20746f6b656e00000000000000000000000000000000000000000000608482015260a40161017c565b6040517f9dc29fac00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff878116600483015260248201869052891690639dc29fac90604401600060405180830381600087803b158015611b4d57600080fd5b505af1158015611b61573d6000803e3d6000fd5b50505050611bfe565b611b8c73ffffffffffffffffffffffffffffffffffffffff891687308761200e565b73ffffffffffffffffffffffffffffffffffffffff8089166000908152600260209081526040808320938b1683529290522054611bca908590612b65565b73ffffffffffffffffffffffffffffffffffffffff808a166000908152600260209081526040808320938c16835292905220555b611c0e8888888888888888611560565b5050505050505050565b606081600003611c5b57505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b8115611c855780611c6f81612b7d565b9150611c7e9050600a83612be4565b9150611c5f565b60008167ffffffffffffffff811115611ca057611ca0612bf8565b6040519080825280601f01601f191660200182016040528015611cca576020820181803683370190505b5090505b8415611d4d57611cdf600183612af1565b9150611cec600a86612c27565b611cf7906030612b65565b60f81b818381518110611d0c57611d0c612c3b565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350611d46600a86612be4565b9450611cce565b949350505050565b6000611d81827f1d1d8b630000000000000000000000000000000000000000000000000000000061206c565b92915050565b60008273ffffffffffffffffffffffffffffffffffffffff1663c01e1bd66040518163ffffffff1660e01b8152600401602060405180830381865afa158015611dd4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611df8919061298d565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614905092915050565b60405173ffffffffffffffffffffffffffffffffffffffff83166024820152604481018290526109849084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009093169290921790915261208f565b600054610100900460ff16611f99576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161017c565b600080547fffffffffffffffffffff0000000000000000000000000000000000000000ffff166201000073ffffffffffffffffffffffffffffffffffffffff94851602179055600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001691909216179055565b60405173ffffffffffffffffffffffffffffffffffffffff8085166024830152831660448201526064810182905261136e9085907f23b872dd0000000000000000000000000000000000000000000000000000000090608401611e80565b60006120778361219b565b8015612088575061208883836121ff565b9392505050565b60006120f1826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166122ce9092919063ffffffff16565b805190915015610984578080602001905181019061210f9190612c6a565b610984576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f74207375636365656400000000000000000000000000000000000000000000606482015260840161017c565b60006121c7827f01ffc9a7000000000000000000000000000000000000000000000000000000006121ff565b8015611d8157506121f8827fffffffff000000000000000000000000000000000000000000000000000000006121ff565b1592915050565b604080517fffffffff000000000000000000000000000000000000000000000000000000008316602480830191909152825180830390910181526044909101909152602080820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f01ffc9a700000000000000000000000000000000000000000000000000000000178152825160009392849283928392918391908a617530fa92503d915060005190508280156122b7575060208210155b80156122c35750600081115b979650505050505050565b6060611d4d84846000858573ffffffffffffffffffffffffffffffffffffffff85163b612357576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161017c565b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516123809190612c8c565b60006040518083038185875af1925050503d80600081146123bd576040519150601f19603f3d011682016040523d82523d6000602084013e6123c2565b606091505b50915091506122c3828286606083156123dc575081612088565b8251156123ec5782518084602001fd5b816040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161017c9190612746565b73ffffffffffffffffffffffffffffffffffffffff8116811461244257600080fd5b50565b60008083601f84011261245757600080fd5b50813567ffffffffffffffff81111561246f57600080fd5b60208301915083602082850101111561248757600080fd5b9250929050565b600080600080600080600060c0888a0312156124a957600080fd5b87356124b481612420565b965060208801356124c481612420565b955060408801356124d481612420565b945060608801356124e481612420565b93506080880135925060a088013567ffffffffffffffff81111561250757600080fd5b6125138a828b01612445565b989b979a50959850939692959293505050565b803563ffffffff8116811461253a57600080fd5b919050565b60008060006040848603121561255457600080fd5b61255d84612526565b9250602084013567ffffffffffffffff81111561257957600080fd5b61258586828701612445565b9497909650939450505050565b6000806000806000608086880312156125aa57600080fd5b85356125b581612420565b945060208601356125c581612420565b935060408601359250606086013567ffffffffffffffff8111156125e857600080fd5b6125f488828901612445565b969995985093965092949392505050565b60008060008060006080868803121561261d57600080fd5b853561262881612420565b94506020860135935061263d60408701612526565b9250606086013567ffffffffffffffff8111156125e857600080fd5b600080600080600080600060c0888a03121561267457600080fd5b873561267f81612420565b9650602088013561268f81612420565b9550604088013561269f81612420565b9450606088013593506126b460808901612526565b925060a088013567ffffffffffffffff81111561250757600080fd5b60005b838110156126eb5781810151838201526020016126d3565b8381111561136e5750506000910152565b600081518084526127148160208601602086016126d0565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061208860208301846126fc565b60008060008060008060a0878903121561277257600080fd5b863561277d81612420565b9550602087013561278d81612420565b9450604087013593506127a260608801612526565b9250608087013567ffffffffffffffff8111156127be57600080fd5b6127ca89828a01612445565b979a9699509497509295939492505050565b600080604083850312156127ef57600080fd5b82356127fa81612420565b9150602083013561280a81612420565b809150509250929050565b6000806000806080858703121561282b57600080fd5b843561283681612420565b9350602085013561284681612420565b9250604085013561285681612420565b9396929550929360600135925050565b60006020828403121561287857600080fd5b813561208881612420565b6000806000806060858703121561289957600080fd5b84356128a481612420565b93506128b260208601612526565b9250604085013567ffffffffffffffff8111156128ce57600080fd5b6128da87828801612445565b95989497509550505050565b828152604060208201526000611d4d60408301846126fc565b600073ffffffffffffffffffffffffffffffffffffffff80871683528086166020840152508360408301526080606083015261293e60808301846126fc565b9695505050505050565b73ffffffffffffffffffffffffffffffffffffffff8416815260606020820152600061297760608301856126fc565b905063ffffffff83166040830152949350505050565b60006020828403121561299f57600080fd5b815161208881612420565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff8516815283602082015260606040820152600061293e6060830184866129aa565b838152604060208201526000612a436040830184866129aa565b95945050505050565b60008451612a5e8184602089016126d0565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551612a9a816001850160208a016126d0565b60019201918201528351612ab58160028401602088016126d0565b0160020195945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082821015612b0357612b03612ac2565b500390565b600073ffffffffffffffffffffffffffffffffffffffff808a1683528089166020840152808816604084015280871660608401525084608083015260c060a0830152612b5860c0830184866129aa565b9998505050505050505050565b60008219821115612b7857612b78612ac2565b500190565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203612bae57612bae612ac2565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600082612bf357612bf3612bb5565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082612c3657612c36612bb5565b500690565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600060208284031215612c7c57600080fd5b8151801515811461208857600080fd5b60008251612c9e8184602087016126d0565b919091019291505056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(L2StandardBridgeStorageLayoutJSON), L2StandardBridgeStorageLayout); err != nil {
		panic(err)
	}

	layouts["L2StandardBridge"] = L2StandardBridgeStorageLayout
	deployedBytecodes["L2StandardBridge"] = L2StandardBridgeDeployedBin
}
