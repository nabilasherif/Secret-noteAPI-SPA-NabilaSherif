import axios from "axios";

const baseClient = () =>
    axios.create({
      baseURL: 'http://localhost:8080/api',
    });

const authClient=()=>
    axios.create({
        baseURL: 'http://localhost:8080/api',
        headers:{
            Authorization: "Bearer " ,
            // + a way to get the token,
        }
    });

export default{
    
}