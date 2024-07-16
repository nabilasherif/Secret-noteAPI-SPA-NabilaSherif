<template>
    <div class="jumbotron vertical-center">
    <div class="container">
        <h1 class="text-center bg-primary text-grey" style="border-radius: 10px">Create new Profile</h1>
        <p>
            Create new Profile
        </p>
        <hr><br>
        <div class="form-group">
                <label for="username">Username:</label>
                <input type="text" class="form-control" id="username" v-model="username" required placeholder="enter your username">
            </div>
            <div class="form-group">
                <label for="password">Password:</label>
                <input type="password" class="form-control" id="password" v-model="password" required placeholder="enter your password">
            </div>
            <button type="button" class="btn btn-info btn-sm" @click="postUser">Create Profile</button>
            <!-- ehtmal fo2 b-button  & use outline instead of btn 3ashan y color when hover-->
            <br><br>
        <button type="button" class="btn btn-info btn-sm">Create Profile</button>
        <br><br>
        <footer class="bg-primary text-white text-center"></footer>
    <b-modal>

    </b-modal>
    </div>
    </div>
</template>

<script>
import axios from 'axios';
export default{
    data(){
        return{
            username:"",
            password:"",
        }
    },
    methods:{
        postUser(){
            const path ='';
            const userInfo={
                Name : this.username,
                Password :this.password
            }
            axios,post(path,userInfo)
            .then((res)=>{
                console.log(res.data);
            })
            .catch ((err)=>{
                console.err(err);
            })
        },
        onSubmit(e){
        }
    }
}
</script>