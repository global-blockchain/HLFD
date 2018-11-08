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

	_,args := stub.GetFunctionAndParameters();
	var a_parm = args[0]
	var b_parm = args[1]
	var c_parm = args[2]
	var logger = shim.NewLogger("simplechaincode")
	logger.Info("================================================just for test================================")
	/*logger.info("args0:%s",a_parm)
	logger.info("args1:%s",b_parm)
	logger.info("args2:%s",c_parm)*/

	return shim.Success([] byte("Success invoke"))
}

func main(){
	err := shim.Start(new (simplechaincode))
	if err != nil{
		fmt.Println("Error starting Simple chaincode:%s", err)
	}
}
