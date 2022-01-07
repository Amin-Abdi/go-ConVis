//function for ordering channels with their respective goroutines
import { BLACK, BLUE, GREEN, RED } from "./constants";

export function insert(arr, index, newItem) {
  return [...arr.slice(0, index), newItem, ...arr.slice(index)];
}

export function getLineColor(lnType) {
  if (lnType === "channel") {
    return BLUE;
  } else if (lnType === "send") {
    return RED;
  } else if (lnType === "receive") {
    return GREEN;
  }

  return BLACK;
}

// const insert = (arr, index, newItem) => [
//     // part of the array before the specified index
//     ...arr.slice(0, index),
//     // inserted items
//     newItem,
//     // part of the array after the specified index
//     ...arr.slice(index)
// ]

// for (let i = 0; i < numOfLines; i++) {
//   initialLength = initialLength + cors;
//   let lineName = new PIXI.Text(sequenceMsg[i], textStyle);
//   lineName.resolution = 1;
//   lineName.position.set(initialLength - 25, 55);

//   var goLine = new Graphics();

//   app.stage.addChild(lineName);
// }
