import * as PIXI from "pixi.js";
import axios from "axios";
import { getLineColor, insert, capitalise } from "./utils";

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
//Send and receive operations
const sendRecArr = [];
const chanType = new Map();
const originMap = new Map();

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
    originMap.set(myObj.name, myObj.origin);
  }
  //check for channels
  if (myObj.operation === "send") {
    operationMap.set(myObj.destination, "channel");
    chanType.set(myObj.destination, myObj.value);
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
//console.log("Origin Map:", originMap);

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
  fontSize: 24, //Make this variable as the number of lines is unknown
});

const Graphics = PIXI.Graphics;

const numOfLines = sequenceMsg.length;
let divisions = Math.floor((innerWidth - 30) / numOfLines);
let initialLength = -40;
let startHeight = 90;
let endHeight = 510;

//Coordinates of the goroutines and channels
const verticalCordinates = new Map();

//Drawing the vertical Lines i.e Goroutines and channels
for (let i = 0; i < numOfLines; i++) {
  initialLength = initialLength + divisions;
  if (sequenceMsg.length > 8) {
    textStyle.fontSize = 21;
  }

  let lineName = new PIXI.Text(capitalise(sequenceMsg[i]), textStyle);
  lineName.resolution = 2;
  lineName.position.set(initialLength - 22, 65);
  app.stage.addChild(lineName);

  let linetype = operationMap.get(sequenceMsg[i]);
  let lineColour = getLineColor(linetype);

  let rect = new Graphics();
  rect
    .beginFill(lineColour, 1)
    .drawRect(initialLength, startHeight, 4, endHeight);
  rect.interactive = true;

  rect.hitArea = new PIXI.Rectangle(initialLength, startHeight, 5, endHeight);

  handleMouseOver(rect, linetype, sequenceMsg[i]);

  rect.mouseout = function (mouseData) {
    rect.removeChild(rect.message);
    delete rect.message;
  };

  app.stage.addChild(rect);
  verticalCordinates.set(sequenceMsg[i], initialLength);
}
//console.log("Vertical Cordinates:", verticalCordinates)
//console.log("Operations:", sendRecArr);

function handleMouseOver(rectangle, lineType, info) {
  const style = new PIXI.TextStyle({
    font: "5px Courier, monospace",
    fill: "#FFA500",
    background: 0x000000,
  });

  rectangle.mouseover = function (mouseData) {
    if (lineType === "goroutine") {
      let localOrigin = originMap.get(info);
      if (localOrigin !== "") {
        let originMessage = new PIXI.Text(`origin: ${localOrigin}`, style);
        originMessage.resolution = 2;
        originMessage.x = mouseData.data.global.x;
        originMessage.y = mouseData.data.global.y;

        //Checking for outbound message
        if (originMessage.x > 1200) {
          originMessage.x = 1140;
        }
        rectangle.message = originMessage;
        rectangle.addChild(originMessage);
      }
    }
  };
}

//For the send and receive messages
let msgNums = sendRecArr.length;
let msgHeight = endHeight - startHeight;
let msgDivs = Math.floor(msgHeight / msgNums);
let startChanHeight = 78;

const messageStyle = new PIXI.TextStyle({
  fontFamily: "Monteserrat",
  fontSize: 15, //Make this variable as the number of lines is unkown
  trim: true,
});

//Drawing the messages
for (let i = 0; i < msgNums; i++) {
  startChanHeight = startChanHeight + msgDivs - 1;

  let myMsg = sendRecArr[i];
  let msgColour = getLineColor(myMsg.operation);

  let start = verticalCordinates.get(myMsg.origin);
  let end = verticalCordinates.get(myMsg.destination);
  let msgInterval = Math.abs(end - start);

  let chanLine = new Graphics();
  chanLine
    .lineStyle(3, msgColour, 1)
    .moveTo(start, startChanHeight)
    .lineTo(end, startChanHeight);

  //The direction of the send and receive operations
  let arrow = new Graphics();
  let valPosition;

  if (end > start) {
    valPosition = end - msgInterval / 2 - 15;
    arrow
      .lineStyle(3, msgColour, 1)
      .moveTo(end - 16, startChanHeight - 10)
      .lineTo(end, startChanHeight)
      .lineTo(end - 16, startChanHeight + 10);
  } else {
    valPosition = start - msgInterval / 2 - 15;
    arrow
      .lineStyle(3, msgColour, 1)
      .moveTo(end + 16, startChanHeight - 10)
      .lineTo(end, startChanHeight)
      .lineTo(end + 16, startChanHeight + 10);
  }

  //Value being sent and received
  let sendRecVal;
  if (myMsg.operation === "send") {
    sendRecVal = chanType.get(myMsg.destination);
  } else if (myMsg.operation === "receive") {
    sendRecVal = chanType.get(myMsg.origin);
  }

  let valOperation = new PIXI.Text(sendRecVal, messageStyle);
  valOperation.resolution = 2;
  valOperation.position.set(valPosition, startChanHeight - 16);

  app.stage.addChild(chanLine);
  app.stage.addChild(arrow);
  app.stage.addChild(valOperation);
}

//Index for the colour representation
const infolist = ["Goroutine", "channel", "receive", "send"];
let infoContainer = new PIXI.Container();
infoContainer.x = 400;

function addIntro() {
  let initialX = 0;
  let spaceInterval = 50;

  for (let i = 0; i < infolist.length; i++) {
    drawInfoShape(infolist[i], initialX);
    initialX = initialX + spaceInterval * 3;
  }
}

function drawInfoShape(s, myX) {
  // console.log(s, getLineColor(s));
  let boxColor = getLineColor(s);
  let square = new Graphics();
  square.beginFill(boxColor).drawRect(myX, 20, 30, 30).endFill();

  let boxName = new PIXI.Text(capitalise(s), textStyle);
  boxName.resolution = 2;
  boxName.position.set(myX + 35, 23);

  infoContainer.addChild(square);
  infoContainer.addChild(boxName);
}
app.stage.addChild(infoContainer);
addIntro();
