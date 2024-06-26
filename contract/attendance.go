// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// AddressInfo is an auto generated low-level Go binding around an user-defined struct.
type AddressInfo struct {
	Label   string
	Details string
	Lat     uint32
	Long    uint32
}

// AttendanceData is an auto generated low-level Go binding around an user-defined struct.
type AttendanceData struct {
	Date            uint32
	CheckInTime     uint32
	CheckOutTime    uint32
	ReleventDetails RelevantDetails
}

// History is an auto generated low-level Go binding around an user-defined struct.
type History struct {
	CheckInTime  uint32
	CheckOutTime uint32
	Reason       string
}

// RelevantDetails is an auto generated low-level Go binding around an user-defined struct.
type RelevantDetails struct {
	AddressInfo AddressInfo
	Imgs        []string
	Note        string
}

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"employeeID\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"date\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"checkInTime\",\"type\":\"uint32\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"lat\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"long\",\"type\":\"uint32\"}],\"internalType\":\"structAddressInfo\",\"name\":\"addressInfo\",\"type\":\"tuple\"},{\"internalType\":\"string[]\",\"name\":\"imgs\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"note\",\"type\":\"string\"}],\"internalType\":\"structRelevantDetails\",\"name\":\"relevantDetails\",\"type\":\"tuple\"}],\"name\":\"checkin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"employeeID\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"date\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"checkOutTime\",\"type\":\"uint32\"}],\"name\":\"checkout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"employeeID\",\"type\":\"uint256\"}],\"name\":\"getAllAttendanceByEmployeeID\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"date\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"checkInTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"checkOutTime\",\"type\":\"uint32\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"lat\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"long\",\"type\":\"uint32\"}],\"internalType\":\"structAddressInfo\",\"name\":\"addressInfo\",\"type\":\"tuple\"},{\"internalType\":\"string[]\",\"name\":\"imgs\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"note\",\"type\":\"string\"}],\"internalType\":\"structRelevantDetails\",\"name\":\"releventDetails\",\"type\":\"tuple\"}],\"internalType\":\"structAttendanceData[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"employeeID\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"startDate\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"endDate\",\"type\":\"uint32\"}],\"name\":\"getAttendanceByRangeDate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"date\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"checkInTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"checkOutTime\",\"type\":\"uint32\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"lat\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"long\",\"type\":\"uint32\"}],\"internalType\":\"structAddressInfo\",\"name\":\"addressInfo\",\"type\":\"tuple\"},{\"internalType\":\"string[]\",\"name\":\"imgs\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"note\",\"type\":\"string\"}],\"internalType\":\"structRelevantDetails\",\"name\":\"releventDetails\",\"type\":\"tuple\"}],\"internalType\":\"structAttendanceData[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"employeeID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getAttendanceDetail\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"date\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"checkInTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"checkOutTime\",\"type\":\"uint32\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"lat\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"long\",\"type\":\"uint32\"}],\"internalType\":\"structAddressInfo\",\"name\":\"addressInfo\",\"type\":\"tuple\"},{\"internalType\":\"string[]\",\"name\":\"imgs\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"note\",\"type\":\"string\"}],\"internalType\":\"structRelevantDetails\",\"name\":\"releventDetails\",\"type\":\"tuple\"}],\"internalType\":\"structAttendanceData\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"checkInTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"checkOutTime\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"internalType\":\"structHistory[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"employeeID\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"date\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"checkInTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"checkOutTime\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"updateAttendanceRecord\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b505f80546001600160a01b0319163390811782556040519091907f342827c97908e5e2f71151c08502a66d44b6f758e3ac2f1de95f02eb95f0a735908290a36123e68061005a5f395ff3fe608060405234801561000f575f80fd5b5060043610610085575f3560e01c80639c81ca25116100585780639c81ca2514610102578063a6f9dae114610123578063fcf9199c14610136578063fd25569d14610149575f80fd5b8063132a2ba2146100895780631d848671146100b257806353f7c5ef146100c7578063893d20e8146100e8575b5f80fd5b61009c610097366004611a58565b61015c565b6040516100a99190611c13565b60405180910390f35b6100c56100c0366004611ddb565b6104bc565b005b6100da6100d5366004611f22565b6107d5565b6040516100a9929190611f42565b5f546040516001600160a01b0390911681526020016100a9565b610115610110366004611fd9565b610b99565b6040516100a9929190612012565b6100c5610131366004612032565b61116d565b6100c5610144366004612058565b6111ef565b6100c5610157366004611fd9565b61146d565b5f818152600160209081526040808320805482518185028101850190935280835260609493849084015b828210156104b0575f8481526020908190206040805160808101825260068602909201805463ffffffff8082168552600160201b8204811695850195909552600160401b900490931682820152805160e0810190915260018301805492939260608086019392918391820190839082908290610201906120cc565b80601f016020809104026020016040519081016040528092919081815260200182805461022d906120cc565b80156102785780601f1061024f57610100808354040283529160200191610278565b820191905f5260205f20905b81548152906001019060200180831161025b57829003601f168201915b50505050508152602001600182018054610291906120cc565b80601f01602080910402602001604051908101604052809291908181526020018280546102bd906120cc565b80156103085780601f106102df57610100808354040283529160200191610308565b820191905f5260205f20905b8154815290600101906020018083116102eb57829003601f168201915b50505091835250506002919091015463ffffffff808216602080850191909152600160201b909204166040928301529183526003840180548251818502810185019093528083529383019391929091905f9084015b82821015610405578382905f5260205f2001805461037a906120cc565b80601f01602080910402602001604051908101604052809291908181526020018280546103a6906120cc565b80156103f15780601f106103c8576101008083540402835291602001916103f1565b820191905f5260205f20905b8154815290600101906020018083116103d457829003601f168201915b50505050508152602001906001019061035d565b50505050815260200160048201805461041d906120cc565b80601f0160208091040260200160405190810160405280929190818152602001828054610449906120cc565b80156104945780601f1061046b57610100808354040283529160200191610494565b820191905f5260205f20905b81548152906001019060200180831161047757829003601f168201915b5050505050815250508152505081526020019060010190610186565b50929695505050505050565b5f546001600160a01b031633146104ee5760405162461bcd60e51b81526004016104e590612104565b60405180910390fd5b6305f5ba5563ffffffff841611156105185760405162461bcd60e51b81526004016104e590612131565b61277563ffffffff841610156105405760405162461bcd60e51b81526004016104e59061217b565b6201517f63ffffffff831611156105b55760405162461bcd60e51b815260206004820152603360248201527f436865636b496e2074696d65206d757374206265206c657373206f7220657175604482015272616c20746f203234682832342a36302a36302960681b60648201526084016104e5565b5f84815260026020908152604080832063ffffffff871684529091529020541561062c5760405162461bcd60e51b815260206004820152602260248201527f417474656e64616e63652068617320616c726561647920636865636b656420696044820152611b9d60f21b60648201526084016104e5565b6040805160808101825263ffffffff808616825284811660208084019182525f848601818152606086018881528b835260018085529783208054808a0182559084529390922086516006909402018054945191518616600160401b026bffffffff000000000000000019928716600160201b0267ffffffffffffffff19909616949096169390931793909317929092169290921782555180518051939485949084019190829081906106de908261220c565b50602082015160018201906106f3908261220c565b5060408201516002909101805460609093015163ffffffff908116600160201b0267ffffffffffffffff1990941692169190911791909117905560208281015180516107459260038501920190611939565b506040820151600482019061075a908261220c565b5050505f878152600160208181526040808420546002835281852063ffffffff8c1686528352908420819055928a90528190526107ce935088925061079e916122dc565b855f6040518060400160405280601081526020016f14915054d3d397d0d21150d2d7d3d55560821b81525061177e565b5050505050565b6107dd61198d565b5f8381526001602052604090205460609061084c5760405162461bcd60e51b815260206004820152602960248201527f456d706c6f79656520494420646f6573206e6f7420657869737420696e20617460448201526874656e64616e63657360b81b60648201526084016104e5565b5f84815260016020526040812080548590811061086b5761086b6122ef565b5f918252602091829020604080516080810182526006909302909101805463ffffffff8082168552600160201b8204811695850195909552600160401b900490931682820152805160e08101909152600183018054929392606080860193929183918201908390829082906108df906120cc565b80601f016020809104026020016040519081016040528092919081815260200182805461090b906120cc565b80156109565780601f1061092d57610100808354040283529160200191610956565b820191905f5260205f20905b81548152906001019060200180831161093957829003601f168201915b5050505050815260200160018201805461096f906120cc565b80601f016020809104026020016040519081016040528092919081815260200182805461099b906120cc565b80156109e65780601f106109bd576101008083540402835291602001916109e6565b820191905f5260205f20905b8154815290600101906020018083116109c957829003601f168201915b50505091835250506002919091015463ffffffff808216602080850191909152600160201b909204166040928301529183526003840180548251818502810185019093528083529383019391929091905f9084015b82821015610ae3578382905f5260205f20018054610a58906120cc565b80601f0160208091040260200160405190810160405280929190818152602001828054610a84906120cc565b8015610acf5780601f10610aa657610100808354040283529160200191610acf565b820191905f5260205f20905b815481529060010190602001808311610ab257829003601f168201915b505050505081526020019060010190610a3b565b505050508152602001600482018054610afb906120cc565b80601f0160208091040260200160405190810160405280929190818152602001828054610b27906120cc565b8015610b725780601f10610b4957610100808354040283529160200191610b72565b820191905f5260205f20905b815481529060010190602001808311610b5557829003601f168201915b5050505050815250508152505090505f610b8c8686611813565b9196919550909350505050565b5f606061277563ffffffff85161015610c0b5760405162461bcd60e51b815260206004820152602e60248201527f53746172742044617465206d757374206265206c61726765206f72206571756160448201526d6c207468616e20312f30312f303160901b60648201526084016104e5565b6305f5ba5563ffffffff84161115610c7c5760405162461bcd60e51b815260206004820152602e60248201527f456e642044617465206d757374206265206c657373206f7220657175616c207460448201526d68616e20393939392f30312f303160901b60648201526084016104e5565b8263ffffffff168463ffffffff161115610cef5760405162461bcd60e51b815260206004820152602e60248201527f53746172742064617465206d757374206265206265666f7265206f722065717560448201526d616c20746f20656e64206461746560901b60648201526084016104e5565b5f85815260016020908152604080832080548251818502810185019093528083529192909190849084015b82821015611044575f8481526020908190206040805160808101825260068602909201805463ffffffff8082168552600160201b8204811695850195909552600160401b900490931682820152805160e0810190915260018301805492939260608086019392918391820190839082908290610d95906120cc565b80601f0160208091040260200160405190810160405280929190818152602001828054610dc1906120cc565b8015610e0c5780601f10610de357610100808354040283529160200191610e0c565b820191905f5260205f20905b815481529060010190602001808311610def57829003601f168201915b50505050508152602001600182018054610e25906120cc565b80601f0160208091040260200160405190810160405280929190818152602001828054610e51906120cc565b8015610e9c5780601f10610e7357610100808354040283529160200191610e9c565b820191905f5260205f20905b815481529060010190602001808311610e7f57829003601f168201915b50505091835250506002919091015463ffffffff808216602080850191909152600160201b909204166040928301529183526003840180548251818502810185019093528083529383019391929091905f9084015b82821015610f99578382905f5260205f20018054610f0e906120cc565b80601f0160208091040260200160405190810160405280929190818152602001828054610f3a906120cc565b8015610f855780601f10610f5c57610100808354040283529160200191610f85565b820191905f5260205f20905b815481529060010190602001808311610f6857829003601f168201915b505050505081526020019060010190610ef1565b505050508152602001600482018054610fb1906120cc565b80601f0160208091040260200160405190810160405280929190818152602001828054610fdd906120cc565b80156110285780601f10610fff57610100808354040283529160200191611028565b820191905f5260205f20905b81548152906001019060200180831161100b57829003601f168201915b5050505050815250508152505081526020019060010190610d1a565b5050505090505f815167ffffffffffffffff81111561106557611065611c44565b60405190808252806020026020018201604052801561109e57816020015b61108b61198d565b8152602001906001900390816110835790505b5090505f805b8351811015611160578763ffffffff168482815181106110c6576110c66122ef565b60200260200101515f015163ffffffff161015801561110d57508663ffffffff168482815181106110f9576110f96122ef565b60200260200101515f015163ffffffff1611155b1561115857838181518110611124576111246122ef565b602002602001015183838151811061113e5761113e6122ef565b6020026020010181905250818061115490612303565b9250505b6001016110a4565b5097909650945050505050565b5f546001600160a01b031633146111965760405162461bcd60e51b81526004016104e590612104565b5f80546040516001600160a01b03808516939216917f342827c97908e5e2f71151c08502a66d44b6f758e3ac2f1de95f02eb95f0a73591a35f80546001600160a01b0319166001600160a01b0392909216919091179055565b5f546001600160a01b031633146112185760405162461bcd60e51b81526004016104e590612104565b6305f5ba5563ffffffff851611156112425760405162461bcd60e51b81526004016104e590612131565b61277563ffffffff8516101561126a5760405162461bcd60e51b81526004016104e59061217b565b6201517f63ffffffff831611156112935760405162461bcd60e51b81526004016104e59061231b565b8163ffffffff168363ffffffff1611156113035760405162461bcd60e51b815260206004820152602b60248201527f436865636b496e742074696d65206d757374206265206c65737320746f20436860448201526a65636b6f75742074696d6560a81b60648201526084016104e5565b5f85815260026020908152604080832063ffffffff8816845290915281205490036113405760405162461bcd60e51b81526004016104e59061236f565b5f85815260026020908152604080832063ffffffff8816845290915281205461136b906001906122dc565b5f878152600160205260409020805491925086918390811061138f5761138f6122ef565b905f5260205f2090600602015f015f6101000a81548163ffffffff021916908363ffffffff1602179055508360015f8881526020019081526020015f2082815481106113dd576113dd6122ef565b905f5260205f2090600602015f0160046101000a81548163ffffffff021916908363ffffffff1602179055508260015f8881526020019081526020015f20828154811061142c5761142c6122ef565b905f5260205f2090600602015f0160086101000a81548163ffffffff021916908363ffffffff160217905550611465868286868661177e565b505050505050565b5f546001600160a01b031633146114965760405162461bcd60e51b81526004016104e590612104565b6305f5ba5563ffffffff831611156114c05760405162461bcd60e51b81526004016104e590612131565b61277563ffffffff831610156114e85760405162461bcd60e51b81526004016104e59061217b565b6201517f63ffffffff821611156115115760405162461bcd60e51b81526004016104e59061231b565b5f83815260026020908152604080832063ffffffff86168452909152812054900361154e5760405162461bcd60e51b81526004016104e59061236f565b5f83815260026020908152604080832063ffffffff86168452909152812054611579906001906122dc565b5f8581526001602052604090208054919250908290811061159c5761159c6122ef565b5f91825260209091206006909102015463ffffffff600160201b9091048116908316116116215760405162461bcd60e51b815260206004820152602d60248201527f436865636b6f75742074696d65206d757374206265206c61726765207468616e60448201526c20636865636b696e2074696d6560981b60648201526084016104e5565b5f848152600160205260409020805482908110611640576116406122ef565b5f918252602090912060069091020154600160401b900463ffffffff16156116b55760405162461bcd60e51b815260206004820152602260248201527f417474656e64616e63652068617320616c726561647920636865636b6564206f6044820152611d5d60f21b60648201526084016104e5565b5f8481526001602052604090208054839190839081106116d7576116d76122ef565b905f5260205f2090600602015f0160086101000a81548163ffffffff021916908363ffffffff160217905550611778848260015f8881526020019081526020015f20848154811061172a5761172a6122ef565b905f5260205f2090600602015f0160049054906101000a900463ffffffff16856040518060400160405280601081526020016f14915054d3d397d0d21150d2d7d3d55560821b81525061177e565b50505050565b6040805160608101825263ffffffff808616825284811660208084019182528385018681525f8b8152600383528681208b82528352958620805460018181018355918852929096208551600290930201805493518516600160201b0267ffffffffffffffff199094169290941691909117919091178255519192839290820190611808908261220c565b505050505050505050565b5f8281526003602090815260408083208484528252808320805482518185028101850190935280835260609492939192909184015b8282101561192c575f8481526020908190206040805160608101825260028602909201805463ffffffff8082168552600160201b90910416938301939093526001830180549293929184019161189d906120cc565b80601f01602080910402602001604051908101604052809291908181526020018280546118c9906120cc565b80156119145780601f106118eb57610100808354040283529160200191611914565b820191905f5260205f20905b8154815290600101906020018083116118f757829003601f168201915b50505050508152505081526020019060010190611848565b5050505090505b92915050565b828054828255905f5260205f2090810192821561197d579160200282015b8281111561197d578251829061196d908261220c565b5091602001919060010190611957565b506119899291506119ee565b5090565b604080516080810182525f8082526020820181905291810191909152606081016119e96040805160e0810182526060808201818152608083018290525f60a0840181905260c08401528252602082018190529181019190915290565b905290565b80821115611989575f611a018282611a0a565b506001016119ee565b508054611a16906120cc565b5f825580601f10611a25575050565b601f0160209004905f5260205f2090810190611a419190611a44565b50565b5b80821115611989575f8155600101611a45565b5f60208284031215611a68575f80fd5b5035919050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b5f8282518085526020808601955060208260051b840101602086015f5b84811015611ae857601f19868403018952611ad6838351611a6f565b98840198925090830190600101611aba565b5090979650505050505050565b5f63ffffffff8083511684528060208401511660208501528060408401511660408501526060830151608060608601528051606060808701528051608060e0880152611b45610160880182611a6f565b9050602082015160df1988830301610100890152611b638282611a6f565b915050836040830151166101208801528360608301511661014088015260208301519350607f199150818782030160a0880152611ba08185611a9d565b93505060408201519150808684030160c087015250611bbf8282611a6f565b95945050505050565b5f8282518085526020808601955060208260051b840101602086015f5b84811015611ae857601f19868403018952611c01838351611af5565b98840198925090830190600101611be5565b602081525f611c256020830184611bc8565b9392505050565b803563ffffffff81168114611c3f575f80fd5b919050565b634e487b7160e01b5f52604160045260245ffd5b6040516060810167ffffffffffffffff81118282101715611c7b57611c7b611c44565b60405290565b6040516080810167ffffffffffffffff81118282101715611c7b57611c7b611c44565b604051601f8201601f1916810167ffffffffffffffff81118282101715611ccd57611ccd611c44565b604052919050565b5f82601f830112611ce4575f80fd5b813567ffffffffffffffff811115611cfe57611cfe611c44565b611d11601f8201601f1916602001611ca4565b818152846020838601011115611d25575f80fd5b816020850160208301375f918101602001919091529392505050565b5f82601f830112611d50575f80fd5b8135602067ffffffffffffffff80831115611d6d57611d6d611c44565b8260051b611d7c838201611ca4565b9384528581018301938381019088861115611d95575f80fd5b84880192505b85831015611dcf57823584811115611db1575f80fd5b611dbf8a87838c0101611cd5565b8352509184019190840190611d9b565b98975050505050505050565b5f805f8060808587031215611dee575f80fd5b84359350611dfe60208601611c2c565b9250611e0c60408601611c2c565b9150606085013567ffffffffffffffff80821115611e28575f80fd5b9086019060608289031215611e3b575f80fd5b611e43611c58565b823582811115611e51575f80fd5b83016080818b031215611e62575f80fd5b611e6a611c81565b813584811115611e78575f80fd5b611e848c828501611cd5565b825250602082013584811115611e98575f80fd5b611ea48c828501611cd5565b602083015250611eb660408301611c2c565b6040820152611ec760608301611c2c565b6060820152825250602083013582811115611ee0575f80fd5b611eec8a828601611d41565b602083015250604083013582811115611f03575f80fd5b611f0f8a828601611cd5565b6040830152509598949750929550505050565b5f8060408385031215611f33575f80fd5b50508035926020909101359150565b5f6040808352611f556040840186611af5565b6020848203818601528186518084528284019150828160051b8501018389015f5b83811015611fc957868303601f190185528151805163ffffffff90811685528782015116878501528801516060898501819052611fb581860183611a6f565b968801969450505090850190600101611f76565b50909a9950505050505050505050565b5f805f60608486031215611feb575f80fd5b83359250611ffb60208501611c2c565b915061200960408501611c2c565b90509250925092565b828152604060208201525f61202a6040830184611bc8565b949350505050565b5f60208284031215612042575f80fd5b81356001600160a01b0381168114611c25575f80fd5b5f805f805f60a0868803121561206c575f80fd5b8535945061207c60208701611c2c565b935061208a60408701611c2c565b925061209860608701611c2c565b9150608086013567ffffffffffffffff8111156120b3575f80fd5b6120bf88828901611cd5565b9150509295509295909350565b600181811c908216806120e057607f821691505b6020821081036120fe57634e487b7160e01b5f52602260045260245ffd5b50919050565b60208082526013908201527221b0b63632b91034b9903737ba1037bbb732b960691b604082015260600190565b6020808252602a908201527f44617465206d757374206265206c657373206f7220657175616c207468616e20604082015269393939392f30312f303160b01b606082015260800190565b60208082526028908201527f44617465206d757374206265206c61726765206f7220657175616c207468616e60408201526720312f30312f303160c01b606082015260800190565b601f82111561220757805f5260205f20601f840160051c810160208510156121e85750805b601f840160051c820191505b818110156107ce575f81556001016121f4565b505050565b815167ffffffffffffffff81111561222657612226611c44565b61223a8161223484546120cc565b846121c3565b602080601f83116001811461226d575f84156122565750858301515b5f19600386901b1c1916600185901b178555611465565b5f85815260208120601f198616915b8281101561229b5788860151825594840194600190910190840161227c565b50858210156122b857878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b5f52601160045260245ffd5b81810381811115611933576119336122c8565b634e487b7160e01b5f52603260045260245ffd5b5f60018201612314576123146122c8565b5060010190565b60208082526034908201527f436865636b6f75742074696d65206d757374206265206c657373206f7220657160408201527375616c20746f203234682832342a36302a36302960601b606082015260800190565b60208082526021908201527f417474656e64616e6365206f66206461746520646f6573206e6f7420657869736040820152601d60fa1b60608201526080019056fea26469706673582212207c0491d8e8a67cb5faef74b5f6b9e8573b1278279b4888e3f57b12ec7e8392d464736f6c63430008190033",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// GetAllAttendanceByEmployeeID is a free data retrieval call binding the contract method 0x132a2ba2.
//
// Solidity: function getAllAttendanceByEmployeeID(uint256 employeeID) view returns((uint32,uint32,uint32,((string,string,uint32,uint32),string[],string))[])
func (_Contract *ContractCaller) GetAllAttendanceByEmployeeID(opts *bind.CallOpts, employeeID *big.Int) ([]AttendanceData, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getAllAttendanceByEmployeeID", employeeID)

	if err != nil {
		return *new([]AttendanceData), err
	}

	out0 := *abi.ConvertType(out[0], new([]AttendanceData)).(*[]AttendanceData)

	return out0, err

}

// GetAllAttendanceByEmployeeID is a free data retrieval call binding the contract method 0x132a2ba2.
//
// Solidity: function getAllAttendanceByEmployeeID(uint256 employeeID) view returns((uint32,uint32,uint32,((string,string,uint32,uint32),string[],string))[])
func (_Contract *ContractSession) GetAllAttendanceByEmployeeID(employeeID *big.Int) ([]AttendanceData, error) {
	return _Contract.Contract.GetAllAttendanceByEmployeeID(&_Contract.CallOpts, employeeID)
}

// GetAllAttendanceByEmployeeID is a free data retrieval call binding the contract method 0x132a2ba2.
//
// Solidity: function getAllAttendanceByEmployeeID(uint256 employeeID) view returns((uint32,uint32,uint32,((string,string,uint32,uint32),string[],string))[])
func (_Contract *ContractCallerSession) GetAllAttendanceByEmployeeID(employeeID *big.Int) ([]AttendanceData, error) {
	return _Contract.Contract.GetAllAttendanceByEmployeeID(&_Contract.CallOpts, employeeID)
}

// GetAttendanceByRangeDate is a free data retrieval call binding the contract method 0x9c81ca25.
//
// Solidity: function getAttendanceByRangeDate(uint256 employeeID, uint32 startDate, uint32 endDate) view returns(uint256, (uint32,uint32,uint32,((string,string,uint32,uint32),string[],string))[])
func (_Contract *ContractCaller) GetAttendanceByRangeDate(opts *bind.CallOpts, employeeID *big.Int, startDate uint32, endDate uint32) (*big.Int, []AttendanceData, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getAttendanceByRangeDate", employeeID, startDate, endDate)

	if err != nil {
		return *new(*big.Int), *new([]AttendanceData), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([]AttendanceData)).(*[]AttendanceData)

	return out0, out1, err

}

// GetAttendanceByRangeDate is a free data retrieval call binding the contract method 0x9c81ca25.
//
// Solidity: function getAttendanceByRangeDate(uint256 employeeID, uint32 startDate, uint32 endDate) view returns(uint256, (uint32,uint32,uint32,((string,string,uint32,uint32),string[],string))[])
func (_Contract *ContractSession) GetAttendanceByRangeDate(employeeID *big.Int, startDate uint32, endDate uint32) (*big.Int, []AttendanceData, error) {
	return _Contract.Contract.GetAttendanceByRangeDate(&_Contract.CallOpts, employeeID, startDate, endDate)
}

// GetAttendanceByRangeDate is a free data retrieval call binding the contract method 0x9c81ca25.
//
// Solidity: function getAttendanceByRangeDate(uint256 employeeID, uint32 startDate, uint32 endDate) view returns(uint256, (uint32,uint32,uint32,((string,string,uint32,uint32),string[],string))[])
func (_Contract *ContractCallerSession) GetAttendanceByRangeDate(employeeID *big.Int, startDate uint32, endDate uint32) (*big.Int, []AttendanceData, error) {
	return _Contract.Contract.GetAttendanceByRangeDate(&_Contract.CallOpts, employeeID, startDate, endDate)
}

// GetAttendanceDetail is a free data retrieval call binding the contract method 0x53f7c5ef.
//
// Solidity: function getAttendanceDetail(uint256 employeeID, uint256 index) view returns((uint32,uint32,uint32,((string,string,uint32,uint32),string[],string)), (uint32,uint32,string)[])
func (_Contract *ContractCaller) GetAttendanceDetail(opts *bind.CallOpts, employeeID *big.Int, index *big.Int) (AttendanceData, []History, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getAttendanceDetail", employeeID, index)

	if err != nil {
		return *new(AttendanceData), *new([]History), err
	}

	out0 := *abi.ConvertType(out[0], new(AttendanceData)).(*AttendanceData)
	out1 := *abi.ConvertType(out[1], new([]History)).(*[]History)

	return out0, out1, err

}

// GetAttendanceDetail is a free data retrieval call binding the contract method 0x53f7c5ef.
//
// Solidity: function getAttendanceDetail(uint256 employeeID, uint256 index) view returns((uint32,uint32,uint32,((string,string,uint32,uint32),string[],string)), (uint32,uint32,string)[])
func (_Contract *ContractSession) GetAttendanceDetail(employeeID *big.Int, index *big.Int) (AttendanceData, []History, error) {
	return _Contract.Contract.GetAttendanceDetail(&_Contract.CallOpts, employeeID, index)
}

// GetAttendanceDetail is a free data retrieval call binding the contract method 0x53f7c5ef.
//
// Solidity: function getAttendanceDetail(uint256 employeeID, uint256 index) view returns((uint32,uint32,uint32,((string,string,uint32,uint32),string[],string)), (uint32,uint32,string)[])
func (_Contract *ContractCallerSession) GetAttendanceDetail(employeeID *big.Int, index *big.Int) (AttendanceData, []History, error) {
	return _Contract.Contract.GetAttendanceDetail(&_Contract.CallOpts, employeeID, index)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Contract *ContractCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Contract *ContractSession) GetOwner() (common.Address, error) {
	return _Contract.Contract.GetOwner(&_Contract.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Contract *ContractCallerSession) GetOwner() (common.Address, error) {
	return _Contract.Contract.GetOwner(&_Contract.CallOpts)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_Contract *ContractTransactor) ChangeOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "changeOwner", newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_Contract *ContractSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.ChangeOwner(&_Contract.TransactOpts, newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_Contract *ContractTransactorSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.ChangeOwner(&_Contract.TransactOpts, newOwner)
}

// Checkin is a paid mutator transaction binding the contract method 0x1d848671.
//
// Solidity: function checkin(uint256 employeeID, uint32 date, uint32 checkInTime, ((string,string,uint32,uint32),string[],string) relevantDetails) returns()
func (_Contract *ContractTransactor) Checkin(opts *bind.TransactOpts, employeeID *big.Int, date uint32, checkInTime uint32, relevantDetails RelevantDetails) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "checkin", employeeID, date, checkInTime, relevantDetails)
}

// Checkin is a paid mutator transaction binding the contract method 0x1d848671.
//
// Solidity: function checkin(uint256 employeeID, uint32 date, uint32 checkInTime, ((string,string,uint32,uint32),string[],string) relevantDetails) returns()
func (_Contract *ContractSession) Checkin(employeeID *big.Int, date uint32, checkInTime uint32, relevantDetails RelevantDetails) (*types.Transaction, error) {
	return _Contract.Contract.Checkin(&_Contract.TransactOpts, employeeID, date, checkInTime, relevantDetails)
}

// Checkin is a paid mutator transaction binding the contract method 0x1d848671.
//
// Solidity: function checkin(uint256 employeeID, uint32 date, uint32 checkInTime, ((string,string,uint32,uint32),string[],string) relevantDetails) returns()
func (_Contract *ContractTransactorSession) Checkin(employeeID *big.Int, date uint32, checkInTime uint32, relevantDetails RelevantDetails) (*types.Transaction, error) {
	return _Contract.Contract.Checkin(&_Contract.TransactOpts, employeeID, date, checkInTime, relevantDetails)
}

// Checkout is a paid mutator transaction binding the contract method 0xfd25569d.
//
// Solidity: function checkout(uint256 employeeID, uint32 date, uint32 checkOutTime) returns()
func (_Contract *ContractTransactor) Checkout(opts *bind.TransactOpts, employeeID *big.Int, date uint32, checkOutTime uint32) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "checkout", employeeID, date, checkOutTime)
}

// Checkout is a paid mutator transaction binding the contract method 0xfd25569d.
//
// Solidity: function checkout(uint256 employeeID, uint32 date, uint32 checkOutTime) returns()
func (_Contract *ContractSession) Checkout(employeeID *big.Int, date uint32, checkOutTime uint32) (*types.Transaction, error) {
	return _Contract.Contract.Checkout(&_Contract.TransactOpts, employeeID, date, checkOutTime)
}

// Checkout is a paid mutator transaction binding the contract method 0xfd25569d.
//
// Solidity: function checkout(uint256 employeeID, uint32 date, uint32 checkOutTime) returns()
func (_Contract *ContractTransactorSession) Checkout(employeeID *big.Int, date uint32, checkOutTime uint32) (*types.Transaction, error) {
	return _Contract.Contract.Checkout(&_Contract.TransactOpts, employeeID, date, checkOutTime)
}

// UpdateAttendanceRecord is a paid mutator transaction binding the contract method 0xfcf9199c.
//
// Solidity: function updateAttendanceRecord(uint256 employeeID, uint32 date, uint32 checkInTime, uint32 checkOutTime, string reason) returns()
func (_Contract *ContractTransactor) UpdateAttendanceRecord(opts *bind.TransactOpts, employeeID *big.Int, date uint32, checkInTime uint32, checkOutTime uint32, reason string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateAttendanceRecord", employeeID, date, checkInTime, checkOutTime, reason)
}

// UpdateAttendanceRecord is a paid mutator transaction binding the contract method 0xfcf9199c.
//
// Solidity: function updateAttendanceRecord(uint256 employeeID, uint32 date, uint32 checkInTime, uint32 checkOutTime, string reason) returns()
func (_Contract *ContractSession) UpdateAttendanceRecord(employeeID *big.Int, date uint32, checkInTime uint32, checkOutTime uint32, reason string) (*types.Transaction, error) {
	return _Contract.Contract.UpdateAttendanceRecord(&_Contract.TransactOpts, employeeID, date, checkInTime, checkOutTime, reason)
}

// UpdateAttendanceRecord is a paid mutator transaction binding the contract method 0xfcf9199c.
//
// Solidity: function updateAttendanceRecord(uint256 employeeID, uint32 date, uint32 checkInTime, uint32 checkOutTime, string reason) returns()
func (_Contract *ContractTransactorSession) UpdateAttendanceRecord(employeeID *big.Int, date uint32, checkInTime uint32, checkOutTime uint32, reason string) (*types.Transaction, error) {
	return _Contract.Contract.UpdateAttendanceRecord(&_Contract.TransactOpts, employeeID, date, checkInTime, checkOutTime, reason)
}

// ContractOwnerSetIterator is returned from FilterOwnerSet and is used to iterate over the raw logs and unpacked data for OwnerSet events raised by the Contract contract.
type ContractOwnerSetIterator struct {
	Event *ContractOwnerSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractOwnerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOwnerSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractOwnerSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractOwnerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOwnerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOwnerSet represents a OwnerSet event raised by the Contract contract.
type ContractOwnerSet struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerSet is a free log retrieval operation binding the contract event 0x342827c97908e5e2f71151c08502a66d44b6f758e3ac2f1de95f02eb95f0a735.
//
// Solidity: event OwnerSet(address indexed oldOwner, address indexed newOwner)
func (_Contract *ContractFilterer) FilterOwnerSet(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*ContractOwnerSetIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OwnerSet", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOwnerSetIterator{contract: _Contract.contract, event: "OwnerSet", logs: logs, sub: sub}, nil
}

// WatchOwnerSet is a free log subscription operation binding the contract event 0x342827c97908e5e2f71151c08502a66d44b6f758e3ac2f1de95f02eb95f0a735.
//
// Solidity: event OwnerSet(address indexed oldOwner, address indexed newOwner)
func (_Contract *ContractFilterer) WatchOwnerSet(opts *bind.WatchOpts, sink chan<- *ContractOwnerSet, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OwnerSet", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOwnerSet)
				if err := _Contract.contract.UnpackLog(event, "OwnerSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnerSet is a log parse operation binding the contract event 0x342827c97908e5e2f71151c08502a66d44b6f758e3ac2f1de95f02eb95f0a735.
//
// Solidity: event OwnerSet(address indexed oldOwner, address indexed newOwner)
func (_Contract *ContractFilterer) ParseOwnerSet(log types.Log) (*ContractOwnerSet, error) {
	event := new(ContractOwnerSet)
	if err := _Contract.contract.UnpackLog(event, "OwnerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
