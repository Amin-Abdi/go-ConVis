
//function for ordering channels with their respective goroutines
export function insert(arr, index, newItem) {
    return [...arr.slice(0, index), newItem, ...arr.slice(index)]
}



// const insert = (arr, index, newItem) => [
//     // part of the array before the specified index
//     ...arr.slice(0, index),
//     // inserted items
//     newItem,
//     // part of the array after the specified index
//     ...arr.slice(index)
// ]