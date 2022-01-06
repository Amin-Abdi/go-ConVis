//function for ordering channels with their respective goroutines
export function insert(arr, index, newItem) {
  return [...arr.slice(0, index), newItem, ...arr.slice(index)];
}

export const BLACK = 0x000000;
export const BLUE = 0x0000ff;
export const GREEN = 0x00ff00;
export const RED = 0xff0000;

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
