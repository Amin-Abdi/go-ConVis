import axios from "axios";

const URL = "http://localhost:8000/operations";

//Fetching data from backend
const response = await axios(URL)
let myData = response.data

console.log("Here!:",myData)
