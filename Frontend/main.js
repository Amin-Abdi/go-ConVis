import axios from "axios";

const URL = "http://localhost:8000/operations";
let myData
//Fetching data from backend
const response = await axios(URL).catch(function (error) {
    console.log("1:",error);
})
if (response !== undefined) {
    myData = response.data
}
const operationMap = new Map()

//Correlating goroutines and channels
for (let i = 0; i < myData.length; i++) {
    /**
     * @param myObj information about object
     * @param myObj.operation information about the operation of object
     */
    let myObj = myData[i]
    //check for goroutines
    if (myObj.operation === "goroutine") {
        operationMap.set(myObj.name, "goroutine")
    }
    //check for channels
    if (myObj.operation === "send") {
        operationMap.set(myObj.destination, "channel")
    }else if (myObj.operation === "receive") {
        operationMap.set(myObj.origin, "channel")
    }

}

console.log("operation map:", operationMap)

