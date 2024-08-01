import axios from "axios";
import { useToast } from "vue-toastification";

const toast = useToast();

const baseClient = () =>
    axios.create({
        baseURL: import.meta.env.VITE_SERVER_DOMAIN,
    });

const authClient=()=>
    axios.create({
        baseURL: import.meta.env.VITE_SERVER_DOMAIN,
        headers:{
            Authorization: "Bearer " + localStorage.getItem("token"),
            'x-access-token': localStorage.getItem("token") 
        }
    });

async function postUser(user:any ){
    console.log({user})
    const token= localStorage.getItem("token")
    console.log({token})
    await baseClient().post("/users", user)
    .then(response => {
        console.log(response.data);
        toast.success("Account created successfully");
    })
    .catch(err => {
        toast.error("couldn't create user");
        console.error('Error creating user:', err);
    });
  }

async function logIn(user: any) :Promise<void>{
    await baseClient().post("/login", user)
    .then(response => {
        let token =response.data.token;
        localStorage.setItem("token",token)
        console.log(response.data.user);
        return
    })
    .catch(err => {
        localStorage.removeItem("token");
        console.error('Error logging in:', err);
        toast.error("Login failed. Please check your username and password.");
        throw err;
        //return err
    });
}

async function postNote(note:any) {
    console.log(localStorage.getItem("token"));
    console.log(note);
    await authClient().post("/note", note)
    .then(response => {
        console.log(response.data);
        toast.success('Note created successfully');
        return
    })
    .catch(err => {
        toast.error("couldn't create note");
        return err
    });
}

function logout(){
    authClient().post("/logout");
    localStorage.removeItem("token");
}

export default{
    baseClient,
    authClient,
    postUser,
    logIn,
    postNote,
    logout,
}