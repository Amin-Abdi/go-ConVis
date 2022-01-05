import * as PIXI from "pixi.js";

import axios from "axios";
import { insert } from "./utils";

const URL = "http://localhost:8000/operations";
let myData;
//Fetching data from backend
const response = await axios(URL).catch(function (error) {
  console.log("1:", error);
});
if (response !== undefined) {
  myData = response.data;
}
//This will be used to arrange the channels and goroutines in proper order
let sequenceMsg = [];

//This map will be used later when drawing the lines for threads and channels
const operationMap = new Map();
//Correlating goroutines and channels
for (let i = 0; i < myData.length; i++) {
  /**
   * @param myObj information about object
   * @param myObj.operation information about the operation of object
   */
  let myObj = myData[i];
  //check for goroutines
  if (myObj.operation === "goroutine") {
    operationMap.set(myObj.name, "goroutine");
    sequenceMsg.push(myObj.name);
  }
  //check for channels
  if (myObj.operation === "send") {
    operationMap.set(myObj.destination, "channel");
  } else if (myObj.operation === "receive") {
    operationMap.set(myObj.origin, "channel");
  }
}
console.log("operation map:", operationMap);
console.log("Sequence:", sequenceMsg);

//Getting the placement of channels
for (let i = 0; i < myData.length; i++) {
  let myObj = myData[i];

  if (myObj.operation === "send") {
    //console.log(myObj.origin, " : index: ", sequenceMsg.indexOf(myObj.origin))
    let sendIndex = sequenceMsg.indexOf(myObj.origin);
    sequenceMsg = insert(sequenceMsg, sendIndex, myObj.destination);
  }
}
console.log("Result:", sequenceMsg);

const app = new PIXI.Application({
  width: innerWidth,
  height: innerHeight,
  backgroundColor: 0xffffff,
  antialias: true,
});
window.devicePixelRatio = 2;
