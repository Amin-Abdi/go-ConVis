import * as PIXI from "pixi.js";
import axios from "axios";
import { getLineColor, insert } from "./utils";

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
const sendRecArr = [];

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
    sendRecArr.push(myObj);
  }
  if (myObj.operation === "receive") {
    operationMap.set(myObj.origin, "channel");
    sendRecArr.push(myObj);
  }
}
//console.log("operation map:", operationMap);
console.log("Initial Sequence:", sequenceMsg);

//Getting the placement of channels
for (let i = 0; i < myData.length; i++) {
  let myObj = myData[i];

  if (myObj.operation === "send") {
    let sendIndex = sequenceMsg.indexOf(myObj.origin);
    //console.log(myObj.origin, " is at index ", sendIndex);
    if (!sequenceMsg.includes(myObj.destination)) {
      sequenceMsg = insert(sequenceMsg, sendIndex, myObj.destination);
    }
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
let initialLength = -40;
let startHeight = 80;
let endHeight = 600;

//Coordinates of the goroutines and channels
const verticalCordinates = new Map();

//Drawing the vertical Lines i.e Goroutines and channels
for (let i = 0; i < numOfLines; i++) {
  initialLength = initialLength + divisions;
  let lineName = new PIXI.Text(sequenceMsg[i], textStyle);
  lineName.resolution = 2;
  lineName.position.set(initialLength - 22, 55);
  app.stage.addChild(lineName);

  let linetype = operationMap.get(sequenceMsg[i]);
  let lineColour = getLineColor(linetype);

  let goLine = new Graphics();
  goLine
    .lineStyle(3, lineColour, 1)
    .moveTo(initialLength, startHeight)
    .lineTo(initialLength, endHeight);
  app.stage.addChild(goLine);

  verticalCordinates.set(sequenceMsg[i], initialLength);
}

//console.log("Vertical Cordinates:", verticalCordinates);

//For the send and receive messages
let msgNums = sendRecArr.length;
let msgHeight = endHeight - startHeight;
let msgDivs = Math.floor(msgHeight / msgNums);
let startChanHeight = 78;

//Drawing the messages
for (let i = 0; i < msgNums; i++) {
  startChanHeight = startChanHeight + msgDivs - 1;

  let myMsg = sendRecArr[i];
  let msgColour = getLineColor(myMsg.operation);

  let start = verticalCordinates.get(myMsg.origin);
  let end = verticalCordinates.get(myMsg.destination);

  let chanLine = new Graphics();
  chanLine
    .lineStyle(3, msgColour, 1)
    .moveTo(start, startChanHeight)
    .lineTo(end, startChanHeight);

  let arrow = new Graphics();

  if (myMsg.operation === "send") {
    arrow
      .lineStyle(3, msgColour, 1)
      .moveTo(end + 16, startChanHeight - 10)
      .lineTo(end, startChanHeight)
      .lineTo(end + 16, startChanHeight + 10);
  } else if (myMsg.operation === "receive") {
    arrow
      .lineStyle(3, msgColour, 1)
      .moveTo(end - 16, startChanHeight - 10)
      .lineTo(end, startChanHeight)
      .lineTo(end - 16, startChanHeight + 10);
  }

  app.stage.addChild(chanLine);
  app.stage.addChild(arrow);
}
