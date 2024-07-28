<template>
  <v-container class="d-flex flex-column align-center justify-center fill-height">
    <!-- <v-row class="text-center">
      <v-col>
        <img width=100px src="/root/Secret-noteAPI-SPA-NabilaSherif/frontend/src/assets/notelogo.png"></img>
      </v-col>
      <v-col cols="12">
        <h2 class="blue--text text--darken-2 display-2 font-weight-bold">
          Welcome to Secret Note
        </h2>
        <p>🤫🤫🤫🤫🤫🤫</p>
      </v-col>
      {{ ShowCreateAccountDialog }}
      <v-col cols="12" class="d-flex justify-center">
        <v-btn class="mx-2" color="primary" @click="ShowCreateAccountDialog = !ShowCreateAccountDialog">Create Account</v-btn>
        <v-btn class="mx-2" color="primary" @click="ShowLogInDialog == true">Log In</v-btn>
      </v-col>
    </v-row> -->

    <v-row>
      <!-- <v-dialog max-width="600px"> -->
      <v-card>
        <v-card-title>Create Account</v-card-title>
        <v-card-subtitle>
          <v-form @submit.prevent="onSubmitSignUp">
            <v-text-field v-model="user.username" label="Username"/>
            <v-text-field v-model="user.password" label="Password" type="password" />
            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn type="submit" color="primary">Submit</v-btn>
              <v-btn @click="ShowCreateAccountDialog = false">Cancel</v-btn>
            </v-card-actions>
          </v-form>
        </v-card-subtitle>
      </v-card>
    <!-- </v-dialog> -->
    </v-row>

    <!-- <div v-if="ShowCreateAccountDialog">
      <v-dialog max-width="600px">
      <v-card>
        <v-card-title>Create Account</v-card-title>
        <v-card-subtitle>
          <v-form @submit.prevent="onSubmitSignUp">
            <v-text-field v-model="username" label="Username"/>
            <v-text-field v-model="password" label="Password" type="password" />
            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn type="submit" color="primary">Submit</v-btn>
              <v-btn @click="ShowCreateAccountDialog = false">Cancel</v-btn>
            </v-card-actions>
          </v-form>
        </v-card-subtitle>
      </v-card>
    </v-dialog>
    </div> -->

    <v-dialog v-model="ShowLogInDialog" max-width="600px">
      <v-card>
        <v-card-title>Log In</v-card-title>
        <v-card-subtitle>
          <v-form @submit.prevent="onSubmitLogIn">
            <v-text-field v-model="user.username" label="Username" />
            <v-text-field v-model="user.password" label="Password" type="password" />
            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn type="submit" color="primary">Login</v-btn>
              <v-btn @click="ShowLogInDialog = false">Cancel</v-btn>
            </v-card-actions>
          </v-form>
        </v-card-subtitle>
      </v-card>
    </v-dialog>
  </v-container>
</template>

<!-- <script>
import axios from 'axios';
import { defineComponent } from 'vue';

const baseClient = () =>
      axios.create({
        baseURL: 'http://localhost:8080/api',
      });

export default defineComponent({
  name: 'HomePage',
  data() {
    return {
      addNewUser: {
        username: "",
        password: "",
      },
      logInData: {
        username: "",
        password: "",
      },
      ShowCreateAccountDialog: false, 
      ShowLogInDialog: false, 
    };
  },
  methods: {
    postUser() {
      // const path = 'http://localhost:8080/api/users';
      const userInfo = {
        username: this.addNewUser.username,
        password: this.addNewUser.password,
      };
      baseClient().post("/users", userInfo , {
      headers: {
      'Content-Type': 'application/json'
       // You can add other headers here as needed
    }
    })
        .then(response => {
          console.log(response.data);
          this.initForm();
          //show a success message here
        })
        .catch(err => {
          console.error('Error creating user:', err);
        });
    },
    logIn() {
      const path = 'http://localhost:8080/api/login';
      const userInfo = {
        username: this.logInData.username,
        password: this.logInData.password,
      };

      axios.post(path, userInfo)
        .then(response => {
          console.log(response.data);
          this.$router.push(`/profile}`);
        })
        .catch(err => {
          console.error('Error logging in:', err);
        });
    },
    initForm() {
      this.addNewUser.username = "";
      this.addNewUser.password = "";
      this.logInData.username = "";
      this.logInData.password = "";
    },
    onSubmitSignUp() {
      this.postUser();
      this.ShowCreateAccountDialog = false; 
    },
    onSubmitLogIn() {
      this.logIn();
      this.ShowLogInDialog = false;
    },
  },
});
</script> -->


<script setup>
import { ref } from 'vue';
//import { toastPlugin, useToast } from 'bootstrap-vue-next';

const ShowCreateAccountDialog= ref(false); 
const ShowLogInDialog= ref(false);

const username=ref("");
const password=ref("");

const user = ref({
  username: "",
  password: "",
});


// function initForm() {
//   addNewUser.value.username = "";
//   addNewUser.value.password = "";
//   logInData.username = "";
//   logInData.password = "";
// };

async function postUser(){
  await baseClient().post("/users", user.value)
  .then(response => {
      console.log(response.data);
      //toast.success('Account created')//errorrrrrr
      //this.initForm();
    })
    .catch(err => {
      to
      console.error('Error creating user:', err);
    });
}

async function logIn() {
      await baseClient().post("/login", user.value)
      .then(response => {
          console.log(response.data);
          this.$router.push(`/profile}`);
      })
      .catch(err => {
          console.error('Error logging in:', err);
      });
}
</script>