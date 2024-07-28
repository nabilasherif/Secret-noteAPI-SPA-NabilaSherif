import axios from "axios";

//import.meta.env.VITE_SERVER_DOMAIN

const baseClient = () =>
    axios.create({
         baseURL: 'http://localhost:8080/api',
    });

const authClient=()=>
    axios.create({
        baseURL: 'http://localhost:8080/api',
        headers:{
            Authorization: "Bearer " + localStorage.getItem("token"),
        }
    });

async function getToken(){
    await authClient().get("login")
    .then((response)=>{
        let token =response.data.token;
        localStorage.setItem("token",token)
    })
    .catch(()=>{
        localStorage.removeItem("token")
    })
};    
export default{
    baseClient,
    authClient,
    getToken,
}