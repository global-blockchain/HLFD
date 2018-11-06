package main
import(
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type simplechaincode struct {

}

func (t *simplechaincode) Init(stub shim.ChaincodeStubInterface) pb.Response  {
	fmt.Println("<< ===== success init it is view in docker ==========>>")
	return shim.Success([] byte("success init"))
}

func (t *simplechaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response{
	fmt.Println("======= Invoke function is call successfully ========")
	return shim.Success([] byte("Success invoke"))
}

func main(){
	err := shim.Start(new (simplechaincode))
	if err != nil{
		fmt.Println("Error starting Simple chaincode:%s", err)
	}
}
