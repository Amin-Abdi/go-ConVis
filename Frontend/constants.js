export const BLACK = 0x000000;
export const BLUE = 0x0000ff;
export const GREEN = 0x00ff00;
export const RED = 0xff0000;

/*

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



//Value being sent and received
  let sendRecVal;
  if (myMsg.operation === "send") {
    sendRecVal = chanType.get(myMsg.destination);
  } else if (myMsg.operation === "receive") {
    sendRecVal = chanType.get(myMsg.origin);
  }

let valOperation = new PIXI.Text(sendRecVal, textStyle);
  valOperation.resolution = 2;
  let valPos = Math.abs((end - start) / 2) + 4;
  valOperation.position.set(valPos, startChanHeight - 50);
  console.log(`X:${valPos}`);


*/
