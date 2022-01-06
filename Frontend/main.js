import * as PIXI from "pixi.js";

import axios from "axios";
import { getLineColor, insert } from "./utils";
import { BLACK, BLUE } from "./constants";

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
const sendRecSet = [];

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
    sendRecSet.push(myObj);
  }
  if (myObj.operation === "receive") {
    operationMap.set(myObj.origin, "channel");
    sendRecSet.push(myObj);
  }
}
console.log("operation map:", operationMap);
//console.log("Initial Sequence:", sequenceMsg);

//Getting the placement of channels
for (let i = 0; i < myData.length; i++) {
  let myObj = myData[i];

  if (myObj.operation === "send") {
    let sendIndex = sequenceMsg.indexOf(myObj.origin);
    //console.log(myObj.origin, " is at index ", sendIndex);
    sequenceMsg = insert(sequenceMsg, sendIndex, myObj.destination);
  }
}

console.log("Result:", sequenceMsg);

//Drawing part
const app = new PIXI.Application({
  width: innerWidth,
  height: innerHeight,
  backgroundColor: 0xffffff,
  antialias: true,
});

window.devicePixelRatio = 2;
app.renderer.view.style.position = "absolute";
document.body.appendChild(app.view);

const textStyle = new PIXI.TextStyle({
  fontFamily: "Monteserrat",
  fontSize: 25, //Make this variable as the number of lines is unkown
});

const Graphics = PIXI.Graphics;

const numOfLines = sequenceMsg.length;
let divisions = Math.floor((innerWidth - 30) / numOfLines);
var initialLength = -40;

//Drawing the vertical Lines i.e Goroutines and channels
for (let i = 0; i < numOfLines; i++) {
  initialLength = initialLength + divisions;
  let lineName = new PIXI.Text(sequenceMsg[i], textStyle);
  lineName.resolution = 1;
  lineName.position.set(initialLength - 22, 55);
  app.stage.addChild(lineName);

  let linetype = operationMap.get(sequenceMsg[i]);
  let lineColour = getLineColor(linetype);

  var goLine = new Graphics();
  goLine
    .lineStyle(3, lineColour, 1)
    .moveTo(initialLength, 80)
    .lineTo(initialLength, 600);
  app.stage.addChild(goLine);
}

//console.log(sendRecSet);

//let chanLine = new Graphics()
