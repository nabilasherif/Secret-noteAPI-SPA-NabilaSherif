<template>
  <v-container class="d-flex flex-column align-center justify-center fill-height">
    <v-row class="text-center">
      <v-col cols="12">
        <br><br><br><br>
        <img width=100px src="/root/Secret-noteAPI-SPA-NabilaSherif/frontend/src/assets/notelogo.png"/>
        <br>
        <h2 class="blue--text text--darken-2 display-2 font-weight-bold">
          Welcome to Secret Note
        </h2>
        <p>🤫🤫🤫🤫🤫🤫</p>
        <br><br>
        <v-btn class="mx-2" color="primary" @click="ShowCreateAccountDialog = !ShowCreateAccountDialog">Create Account</v-btn>
        <v-btn class="mx-2" color="primary" @click="ShowLogInDialog = !ShowLogInDialog">Log In</v-btn>
      </v-col>
    </v-row>

    <v-dialog v-model="ShowCreateAccountDialog" max-width="600px">
      <v-form @submit.prevent="postUser" >
        <v-card>
          <v-card-title>Create Account</v-card-title>
          <v-card-subtitle>
            <v-text-field v-model="user.username" label="Username" />
            <v-text-field v-model="user.password" label="Password" type="password" />
            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn @click="ShowCreateAccountDialog = !ShowCreateAccountDialog" type="submit" color="primary">Submit</v-btn>
              <v-btn @click="ShowCreateAccountDialog = !ShowCreateAccountDialog">Cancel</v-btn>
            </v-card-actions>
          </v-card-subtitle>
        </v-card>
      </v-form>
    </v-dialog>

    <v-dialog v-model="ShowLogInDialog" max-width="600px">
      <v-form @submit.prevent="logIn" >
        <v-card>
          <v-card-title>Log In</v-card-title>
          <v-card-subtitle>
            <v-text-field v-model="user.username" label="Username" />
            <v-text-field v-model="user.password" label="Password" type="password" />
            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn @click="ShowLogInDialog = !ShowLogInDialog" type="submit" color="primary">Login</v-btn>
              <v-btn @click="ShowLogInDialog = !ShowLogInDialog">Cancel</v-btn>
            </v-card-actions>
          </v-card-subtitle>
        </v-card>
      </v-form>
    </v-dialog>

  </v-container>
</template>

<script setup>
import { ref } from 'vue';
import userservice from '../services/userservice.ts';
import { useRouter } from 'vue-router';

const router=useRouter();


let ShowCreateAccountDialog= ref(false); 
let ShowLogInDialog= ref(false);

const user = ref({
  username: "",
  password: "",
});

function initForm() {
  user.value.username="";
  user.value.password="";
};

function postUser(){
  userservice.postUser(user.value);
  initForm();
}

function logIn() {
    userservice.logIn(user.value)
    .then(() => {
        router.push(`/profile`);
    })
    .catch(error => {
        toast.error("Login failed. Please check your username and password.");
        console.error('Error logging in:', error);
    })
    .finally(() => {
        initForm();
    });
}


</script>